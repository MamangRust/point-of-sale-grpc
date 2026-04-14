package gapi_test

import (
	"context"
	"pointofsale/internal/cache"
	order_cache "pointofsale/internal/cache/order"
	orderitem_cache_srv "pointofsale/internal/cache/order_item"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/handler/gapi"
	"pointofsale/internal/pb"
	"pointofsale/internal/repository"
	"pointofsale/internal/service"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/observability"
	"pointofsale/tests"
	"net"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderItemGapiTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *pgxpool.Pool
	redisClient *redis.Client
	grpcServer  *grpc.Server
	client      pb.OrderItemServiceClient
	orderSvc    service.OrderService
	conn        *grpc.ClientConn
	merchantID  int
	userID      int
	categoryID  int
	productID   int
	orderID     int
	cashierID   int
}

func (s *OrderItemGapiTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := pgxpool.New(s.ts.Ctx, s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	opts, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	s.redisClient = redis.NewClient(opts)

	queries := db.New(pool)
	repos := repository.NewRepositories(queries)

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test-gapi-oi", lp)
	obs, _ := observability.NewObservability("test-gapi-oi", log)
	
	cacheMetrics, _ := observability.NewCacheMetrics("test-gapi-oi")
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)
	orderCache := order_cache.NewOrderMencache(cacheStore)
	orderItemCache := orderitem_cache_srv.NewOrderItemCache(cacheStore)

	orderItemService := service.NewOrderItemService(service.OrderItemServiceDeps{
		OrderItemRepo: repos.OrderItem,
		Logger:        log,
		Observability: obs,
		Cache:         orderItemCache,
	})
	s.orderSvc = service.NewOrderService(service.OrderServiceDeps{
		OrderRepo:     repos.Order,
		OrderItemRepo: repos.OrderItem,
		ProductRepo:   repos.Product,
		CashierRepo:   repos.Cashier,
		MerchantRepo:  repos.Merchant,
		Logger:        log,
		Observability: obs,
		Cache:         orderCache,
	})

	// Start gRPC Server
	orderItemHandler := gapi.NewOrderItemHandleGrpc(orderItemService)
	server := grpc.NewServer()
	pb.RegisterOrderItemServiceServer(server, orderItemHandler)
	s.grpcServer = server

	lis, err := net.Listen("tcp", ":0")
	s.Require().NoError(err)

	go func() {
		_ = server.Serve(lis)
	}()

	// Create Client
	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	s.conn = conn
	s.client = pb.NewOrderItemServiceClient(conn)

	ctx := context.Background()

	// Setup Dependencies
	user, err := repos.User.CreateUser(ctx, &requests.CreateUserRequest{
		FirstName: "GapiOI", LastName: "User", Email: "gapioi.user@example.com", Password: "password123",
	})
	s.Require().NoError(err)
	s.userID = int(user.UserID)

	merchant, err := repos.Merchant.CreateMerchant(ctx, &requests.CreateMerchantRequest{
		UserID: s.userID, Name: "GapiOI Merchant",
	})
	s.Require().NoError(err)
	s.merchantID = int(merchant.MerchantID)

	cashier, err := repos.Cashier.CreateCashier(ctx, &requests.CreateCashierRequest{
		MerchantID: s.merchantID,
		UserID:     s.userID,
		Name:       "GapiOI Cashier",
	})
	s.Require().NoError(err)
	s.cashierID = int(cashier.CashierID)

	slugCat := "gapioi-cat"
	category, err := repos.Category.CreateCategory(ctx, &requests.CreateCategoryRequest{
		Name: "GapiOI Cat", SlugCategory: &slugCat,
	})
	s.Require().NoError(err)
	s.categoryID = int(category.CategoryID)

	slugProd := "gapioi-prod"
	product, err := repos.Product.CreateProduct(ctx, &requests.CreateProductRequest{
		MerchantID: s.merchantID, CategoryID: s.categoryID, Name: "GapiOI Prod", Price: 100, CountInStock: 100, SlugProduct: &slugProd,
	})
	s.Require().NoError(err)
	s.productID = int(product.ProductID)

	// Create Order
	order, err := s.orderSvc.CreateOrder(ctx, &requests.CreateOrderRequest{
		MerchantID: s.merchantID,
		CashierID:  s.cashierID,
		Items: []requests.CreateOrderItemRequest{
			{
				ProductID: s.productID,
				Quantity:  1,
			},
		},
	})
	s.Require().NoError(err)
	s.orderID = int(order.OrderID)
}

func (s *OrderItemGapiTestSuite) TearDownSuite() {
	if s.conn != nil {
		s.conn.Close()
	}
	if s.grpcServer != nil {
		s.grpcServer.Stop()
	}
	if s.redisClient != nil {
		s.redisClient.Close()
	}
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	if s.ts != nil {
		s.ts.Teardown()
	}
}

func (s *OrderItemGapiTestSuite) TestOrderItemGapiLifecycle() {
	ctx := context.Background()

	// 1. Find By Order
	res, err := s.client.FindOrderItemByOrder(ctx, &pb.FindByIdOrderItemRequest{Id: int32(s.orderID)})
	s.NoError(err)
	s.NotEmpty(res.Data)

	// 2. Find All
	allRes, err := s.client.FindAll(ctx, &pb.FindAllOrderItemRequest{Page: 1, PageSize: 10})
	s.NoError(err)
	s.NotEmpty(allRes.Data)

	// 3. Find By Active
	activeRes, err := s.client.FindByActive(ctx, &pb.FindAllOrderItemRequest{Page: 1, PageSize: 10})
	s.NoError(err)
	s.NotEmpty(activeRes.Data)
}

func TestOrderItemGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(OrderItemGapiTestSuite))
}
