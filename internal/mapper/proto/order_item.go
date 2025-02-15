package protomapper

import (
	"pointofsale/internal/domain/response"
	"pointofsale/internal/pb"
)

type orderItemProtoMapper struct{}

func NewOrderItemProtoMapper() *orderItemProtoMapper {
	return &orderItemProtoMapper{}
}

func (o *orderItemProtoMapper) ToProtoResponseOrderItem(status string, message string, pbResponse *response.OrderItemResponse) *pb.ApiResponseOrderItem {
	return &pb.ApiResponseOrderItem{
		Status:  status,
		Message: message,
		Data:    o.mapResponseOrderItem(pbResponse),
	}
}

func (o *orderItemProtoMapper) ToProtoResponsesOrderItem(status string, message string, pbResponse []*response.OrderItemResponse) *pb.ApiResponsesOrderItem {
	return &pb.ApiResponsesOrderItem{
		Status:  status,
		Message: message,
		Data:    o.mapResponsesOrderItem(pbResponse),
	}
}

func (o *orderItemProtoMapper) ToProtoResponseOrderItemDelete(status string, message string) *pb.ApiResponseOrderItemDelete {
	return &pb.ApiResponseOrderItemDelete{
		Status:  status,
		Message: message,
	}
}

func (o *orderItemProtoMapper) ToProtoResponseOrderItemAll(status string, message string) *pb.ApiResponseOrderItemAll {
	return &pb.ApiResponseOrderItemAll{
		Status:  status,
		Message: message,
	}
}

func (o *orderItemProtoMapper) ToProtoResponsePaginationOrderItemDeleteAt(pagination *pb.PaginationMeta, status string, message string, orderItems []*response.OrderItemResponseDeleteAt) *pb.ApiResponsePaginationOrderItemDeleteAt {
	return &pb.ApiResponsePaginationOrderItemDeleteAt{
		Status:     status,
		Message:    message,
		Data:       o.mapResponsesOrderItemDeleteAt(orderItems),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (o *orderItemProtoMapper) ToProtoResponsePaginationOrderItem(pagination *pb.PaginationMeta, status string, message string, orderItems []*response.OrderItemResponse) *pb.ApiResponsePaginationOrderItem {
	return &pb.ApiResponsePaginationOrderItem{
		Status:     status,
		Message:    message,
		Data:       o.mapResponsesOrderItem(orderItems),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (o *orderItemProtoMapper) mapResponseOrderItem(orderItem *response.OrderItemResponse) *pb.OrderItemResponse {
	return &pb.OrderItemResponse{
		Id:        int32(orderItem.ID),
		OrderId:   int32(orderItem.OrderID),
		ProductId: int32(orderItem.ProductID),
		Quantity:  int32(orderItem.Quantity),
		Price:     int32(orderItem.Price),
		CreatedAt: orderItem.CreatedAt,
		UpdatedAt: orderItem.UpdatedAt,
	}
}

func (o *orderItemProtoMapper) mapResponsesOrderItem(orderItems []*response.OrderItemResponse) []*pb.OrderItemResponse {
	var mappedOrderItems []*pb.OrderItemResponse

	for _, orderItem := range orderItems {
		mappedOrderItems = append(mappedOrderItems, o.mapResponseOrderItem(orderItem))
	}

	return mappedOrderItems
}

func (o *orderItemProtoMapper) mapResponseOrderItemDelete(orderItem *response.OrderItemResponseDeleteAt) *pb.OrderItemResponseDeleteAt {
	return &pb.OrderItemResponseDeleteAt{
		Id:        int32(orderItem.ID),
		OrderId:   int32(orderItem.OrderID),
		ProductId: int32(orderItem.ProductID),
		Quantity:  int32(orderItem.Quantity),
		Price:     int32(orderItem.Price),
		CreatedAt: orderItem.CreatedAt,
		UpdatedAt: orderItem.UpdatedAt,
		DeletedAt: orderItem.DeleteAt,
	}
}

func (o *orderItemProtoMapper) mapResponsesOrderItemDeleteAt(orderItems []*response.OrderItemResponseDeleteAt) []*pb.OrderItemResponseDeleteAt {
	var mappedOrderItems []*pb.OrderItemResponseDeleteAt

	for _, orderItem := range orderItems {
		mappedOrderItems = append(mappedOrderItems, o.mapResponseOrderItemDelete(orderItem))
	}

	return mappedOrderItems
}
