package orderitem_cache

import (
	"context"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
)

type OrderItemQueryCache interface {
	GetCachedOrderItemsAll(ctx context.Context, req *requests.FindAllOrderItems) (*response.ApiResponsePaginationOrderItem, bool)
	SetCachedOrderItemsAll(ctx context.Context, req *requests.FindAllOrderItems, res *response.ApiResponsePaginationOrderItem)

	GetCachedOrderItemActive(ctx context.Context, req *requests.FindAllOrderItems) (*response.ApiResponsePaginationOrderItemDeleteAt, bool)
	SetCachedOrderItemActive(ctx context.Context, req *requests.FindAllOrderItems, res *response.ApiResponsePaginationOrderItemDeleteAt)

	GetCachedOrderItemTrashed(ctx context.Context, req *requests.FindAllOrderItems) (*response.ApiResponsePaginationOrderItemDeleteAt, bool)
	SetCachedOrderItemTrashed(ctx context.Context, req *requests.FindAllOrderItems, res *response.ApiResponsePaginationOrderItemDeleteAt)

	GetCachedOrderItems(ctx context.Context, orderID int) (*response.ApiResponsesOrderItem, bool)
	SetCachedOrderItems(ctx context.Context, res *response.ApiResponsesOrderItem)
}
