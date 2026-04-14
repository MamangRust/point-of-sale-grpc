package gapi_test

import (
	"context"
	"pointofsale/internal/cache"
	order_cache "pointofsale/internal/cache/order"
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

type OrderGapiTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *pgxpool.Pool
	redisClient *redis.Client
	grpcServer  *grpc.Server
	client      pb.OrderServiceClient
	repos       *repository.Repositories
	conn        *grpc.ClientConn
	merchantID  int
	userID      int
	categoryID  int
	productID   int
	cashierID   int
}

func (s *OrderGapiTestSuite) SetupSuite() {
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
	s.repos = repository.NewRepositories(queries)

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test-gapi", lp)
	obs, _ := observability.NewObservability("test-gapi", log)
	
	cacheMetrics, _ := observability.NewCacheMetrics("test-gapi")
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)
	orderCache := order_cache.NewOrderMencache(cacheStore)

	orderService := service.NewOrderService(service.OrderServiceDeps{
		OrderRepo:     s.repos.Order,
		OrderItemRepo: s.repos.OrderItem,
		ProductRepo:   s.repos.Product,
		CashierRepo:   s.repos.Cashier,
		MerchantRepo:  s.repos.Merchant,
		Logger:        log,
		Observability: obs,
		Cache:         orderCache,
	})

	// Start gRPC Server
	orderHandler := gapi.NewOrderHandleGrpc(orderService)
	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, orderHandler)
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
	s.client = pb.NewOrderServiceClient(conn)

	ctx := context.Background()

	// Setup Dependencies
	user, err := s.repos.User.CreateUser(ctx, &requests.CreateUserRequest{
		FirstName: "Gapi", LastName: "User", Email: "gapi.user@example.com", Password: "password123",
	})
	s.Require().NoError(err)
	s.userID = int(user.UserID)

	merchant, err := s.repos.Merchant.CreateMerchant(ctx, &requests.CreateMerchantRequest{
		UserID: s.userID, Name: "Gapi Merchant",
	})
	s.Require().NoError(err)
	s.merchantID = int(merchant.MerchantID)

	slugCat := "gapi-cat"
	category, err := s.repos.Category.CreateCategory(ctx, &requests.CreateCategoryRequest{
		Name: "Gapi Cat", SlugCategory: &slugCat,
	})
	s.Require().NoError(err)
	s.categoryID = int(category.CategoryID)

	slugProd := "gapi-prod"
	product, err := s.repos.Product.CreateProduct(ctx, &requests.CreateProductRequest{
		MerchantID: s.merchantID, CategoryID: s.categoryID, Name: "Gapi Prod", Price: 100, CountInStock: 100, SlugProduct: &slugProd,
	})
	s.Require().NoError(err)
	s.productID = int(product.ProductID)

	cashier, err := s.repos.Cashier.CreateCashier(ctx, &requests.CreateCashierRequest{
		MerchantID: s.merchantID,
		UserID:     s.userID,
		Name:       "Gapi Cashier",
	})
	s.Require().NoError(err)
	s.cashierID = int(cashier.CashierID)
}

func (s *OrderGapiTestSuite) TearDownSuite() {
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

func (s *OrderGapiTestSuite) TestOrderGapiLifecycle() {
	ctx := context.Background()

	// 1. Create
	createReq := &pb.CreateOrderRequest{
		MerchantId: int32(s.merchantID),
		CashierId:  int32(s.cashierID),
		Items: []*pb.CreateOrderItemRequest{
			{
				ProductId: int32(s.productID),
				Quantity:  2,
			},
		},
	}

	res, err := s.client.Create(ctx, createReq)
	s.NoError(err)
	s.NotNil(res)
	orderID := res.Data.Id

	// 2. Find By ID
	found, err := s.client.FindById(ctx, &pb.FindByIdOrderRequest{Id: orderID})
	s.NoError(err)
	s.Equal(orderID, found.Data.Id)

	// 3. Find All
	orders, err := s.client.FindAll(ctx, &pb.FindAllOrderRequest{Page: 1, PageSize: 10})
	s.NoError(err)
	s.NotEmpty(orders.Data)

	// Fetch Order Items to get IDs
	orderItems, err := s.repos.OrderItem.FindOrderItemByOrder(ctx, int(orderID))
	s.Require().NoError(err)
	s.NotEmpty(orderItems)
	orderItemID := orderItems[0].OrderItemID

	// 4. Update
	updateReq := &pb.UpdateOrderRequest{
		OrderId:    orderID,
		Items: []*pb.UpdateOrderItemRequest{
			{
				OrderItemId: int32(orderItemID),
				ProductId:   int32(s.productID),
				Quantity:    3,
			},
		},
	}
	updated, err := s.client.Update(ctx, updateReq)
	s.NoError(err)
	s.NotNil(updated)

	// 5. Trash
	_, err = s.client.TrashedOrder(ctx, &pb.FindByIdOrderRequest{Id: orderID})
	s.NoError(err)

	// 6. Restore
	_, err = s.client.RestoreOrder(ctx, &pb.FindByIdOrderRequest{Id: orderID})
	s.NoError(err)

	// 7. Delete Permanent
	_, err = s.client.TrashedOrder(ctx, &pb.FindByIdOrderRequest{Id: orderID})
	s.NoError(err)

	delRes, err := s.client.DeleteOrderPermanent(ctx, &pb.FindByIdOrderRequest{Id: orderID})
	s.NoError(err)
	s.Equal("success", delRes.Status)
}

func TestOrderGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(OrderGapiTestSuite))
}
