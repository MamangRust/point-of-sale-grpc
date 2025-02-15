// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.2
// source: order_item.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FindAllOrderItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Search        string                 `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindAllOrderItemRequest) Reset() {
	*x = FindAllOrderItemRequest{}
	mi := &file_order_item_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindAllOrderItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllOrderItemRequest) ProtoMessage() {}

func (x *FindAllOrderItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_item_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllOrderItemRequest.ProtoReflect.Descriptor instead.
func (*FindAllOrderItemRequest) Descriptor() ([]byte, []int) {
	return file_order_item_proto_rawDescGZIP(), []int{0}
}

func (x *FindAllOrderItemRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindAllOrderItemRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FindAllOrderItemRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type FindByIdOrderItemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindByIdOrderItemRequest) Reset() {
	*x = FindByIdOrderItemRequest{}
	mi := &file_order_item_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindByIdOrderItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIdOrderItemRequest) ProtoMessage() {}

func (x *FindByIdOrderItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_order_item_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIdOrderItemRequest.ProtoReflect.Descriptor instead.
func (*FindByIdOrderItemRequest) Descriptor() ([]byte, []int) {
	return file_order_item_proto_rawDescGZIP(), []int{1}
}

func (x *FindByIdOrderItemRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrderItemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       int32                  `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ProductId     int32                  `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price         int32                  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItemResponse) Reset() {
	*x = OrderItemResponse{}
	mi := &file_order_item_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemResponse) ProtoMessage() {}

func (x *OrderItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_item_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemResponse.ProtoReflect.Descriptor instead.
func (*OrderItemResponse) Descriptor() ([]byte, []int) {
	return file_order_item_proto_rawDescGZIP(), []int{2}
}

func (x *OrderItemResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderItemResponse) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderItemResponse) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *OrderItemResponse) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItemResponse) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderItemResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *OrderItemResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type OrderItemResponseDeleteAt struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       int32                  `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ProductId     int32                  `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price         int32                  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt     string                 `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderItemResponseDeleteAt) Reset() {
	*x = OrderItemResponseDeleteAt{}
	mi := &file_order_item_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderItemResponseDeleteAt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItemResponseDeleteAt) ProtoMessage() {}

func (x *OrderItemResponseDeleteAt) ProtoReflect() protoreflect.Message {
	mi := &file_order_item_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItemResponseDeleteAt.ProtoReflect.Descriptor instead.
func (*OrderItemResponseDeleteAt) Descriptor() ([]byte, []int) {
	return file_order_item_proto_rawDescGZIP(), []int{3}
}

func (x *OrderItemResponseDeleteAt) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderItemResponseDeleteAt) GetOrderId() int32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderItemResponseDeleteAt) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *OrderItemResponseDeleteAt) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItemResponseDeleteAt) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderItemResponseDeleteAt) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *OrderItemResponseDeleteAt) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *OrderItemResponseDeleteAt) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type ApiResponseOrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          *OrderItemResponse     `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiResponseOrderItem) Reset() {
	*x = ApiResponseOrderItem{}
	mi := &file_order_item_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiResponseOrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponseOrderItem) ProtoMessage() {}

func (x *ApiResponseOrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_item_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponseOrderItem.ProtoReflect.Descriptor instead.
func (*ApiResponseOrderItem) Descriptor() ([]byte, []int) {
	return file_order_item_proto_rawDescGZIP(), []int{4}
}

func (x *ApiResponseOrderItem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ApiResponseOrderItem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApiResponseOrderItem) GetData() *OrderItemResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type ApiResponsesOrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          []*OrderItemResponse   `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiResponsesOrderItem) Reset() {
	*x = ApiResponsesOrderItem{}
	mi := &file_order_item_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiResponsesOrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponsesOrderItem) ProtoMessage() {}

func (x *ApiResponsesOrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_item_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponsesOrderItem.ProtoReflect.Descriptor instead.
func (*ApiResponsesOrderItem) Descriptor() ([]byte, []int) {
	return file_order_item_proto_rawDescGZIP(), []int{5}
}

func (x *ApiResponsesOrderItem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ApiResponsesOrderItem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApiResponsesOrderItem) GetData() []*OrderItemResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type ApiResponseOrderItemDelete struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiResponseOrderItemDelete) Reset() {
	*x = ApiResponseOrderItemDelete{}
	mi := &file_order_item_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiResponseOrderItemDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponseOrderItemDelete) ProtoMessage() {}

func (x *ApiResponseOrderItemDelete) ProtoReflect() protoreflect.Message {
	mi := &file_order_item_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponseOrderItemDelete.ProtoReflect.Descriptor instead.
func (*ApiResponseOrderItemDelete) Descriptor() ([]byte, []int) {
	return file_order_item_proto_rawDescGZIP(), []int{6}
}

func (x *ApiResponseOrderItemDelete) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ApiResponseOrderItemDelete) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ApiResponseOrderItemAll struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiResponseOrderItemAll) Reset() {
	*x = ApiResponseOrderItemAll{}
	mi := &file_order_item_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiResponseOrderItemAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponseOrderItemAll) ProtoMessage() {}

func (x *ApiResponseOrderItemAll) ProtoReflect() protoreflect.Message {
	mi := &file_order_item_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponseOrderItemAll.ProtoReflect.Descriptor instead.
func (*ApiResponseOrderItemAll) Descriptor() ([]byte, []int) {
	return file_order_item_proto_rawDescGZIP(), []int{7}
}

func (x *ApiResponseOrderItemAll) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ApiResponseOrderItemAll) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ApiResponsePaginationOrderItemDeleteAt struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Status        string                       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          []*OrderItemResponseDeleteAt `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Pagination    *PaginationMeta              `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiResponsePaginationOrderItemDeleteAt) Reset() {
	*x = ApiResponsePaginationOrderItemDeleteAt{}
	mi := &file_order_item_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiResponsePaginationOrderItemDeleteAt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponsePaginationOrderItemDeleteAt) ProtoMessage() {}

func (x *ApiResponsePaginationOrderItemDeleteAt) ProtoReflect() protoreflect.Message {
	mi := &file_order_item_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponsePaginationOrderItemDeleteAt.ProtoReflect.Descriptor instead.
func (*ApiResponsePaginationOrderItemDeleteAt) Descriptor() ([]byte, []int) {
	return file_order_item_proto_rawDescGZIP(), []int{8}
}

func (x *ApiResponsePaginationOrderItemDeleteAt) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ApiResponsePaginationOrderItemDeleteAt) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApiResponsePaginationOrderItemDeleteAt) GetData() []*OrderItemResponseDeleteAt {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ApiResponsePaginationOrderItemDeleteAt) GetPagination() *PaginationMeta {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ApiResponsePaginationOrderItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data          []*OrderItemResponse   `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Pagination    *PaginationMeta        `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiResponsePaginationOrderItem) Reset() {
	*x = ApiResponsePaginationOrderItem{}
	mi := &file_order_item_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiResponsePaginationOrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiResponsePaginationOrderItem) ProtoMessage() {}

func (x *ApiResponsePaginationOrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_item_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiResponsePaginationOrderItem.ProtoReflect.Descriptor instead.
func (*ApiResponsePaginationOrderItem) Descriptor() ([]byte, []int) {
	return file_order_item_proto_rawDescGZIP(), []int{9}
}

func (x *ApiResponsePaginationOrderItem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ApiResponsePaginationOrderItem) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ApiResponsePaginationOrderItem) GetData() []*OrderItemResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ApiResponsePaginationOrderItem) GetPagination() *PaginationMeta {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_order_item_proto protoreflect.FileDescriptor

var file_order_item_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x62, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x2a, 0x0a, 0x18, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49,
	0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xcd, 0x01, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xf4, 0x01, 0x0a, 0x19, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x73, 0x0a, 0x14, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x74, 0x0a,
	0x15, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x1a, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x4b, 0x0a, 0x17, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x6c, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0xc1, 0x01, 0x0a, 0x26, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x62,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x32, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb1, 0x01, 0x0a, 0x1e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xe2, 0x02, 0x0a, 0x10, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a,
	0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x57, 0x0a, 0x0c, 0x46, 0x69, 0x6e,
	0x64, 0x42, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x69, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x74, 0x12, 0x58, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x54, 0x72, 0x61, 0x73,
	0x68, 0x65, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x12, 0x4f, 0x0a, 0x14,
	0x46, 0x69, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x49, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x19, 0x5a,
	0x17, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6f, 0x66, 0x73, 0x61, 0x6c, 0x65, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_order_item_proto_rawDescOnce sync.Once
	file_order_item_proto_rawDescData []byte
)

func file_order_item_proto_rawDescGZIP() []byte {
	file_order_item_proto_rawDescOnce.Do(func() {
		file_order_item_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_order_item_proto_rawDesc), len(file_order_item_proto_rawDesc)))
	})
	return file_order_item_proto_rawDescData
}

var file_order_item_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_order_item_proto_goTypes = []any{
	(*FindAllOrderItemRequest)(nil),                // 0: pb.FindAllOrderItemRequest
	(*FindByIdOrderItemRequest)(nil),               // 1: pb.FindByIdOrderItemRequest
	(*OrderItemResponse)(nil),                      // 2: pb.OrderItemResponse
	(*OrderItemResponseDeleteAt)(nil),              // 3: pb.OrderItemResponseDeleteAt
	(*ApiResponseOrderItem)(nil),                   // 4: pb.ApiResponseOrderItem
	(*ApiResponsesOrderItem)(nil),                  // 5: pb.ApiResponsesOrderItem
	(*ApiResponseOrderItemDelete)(nil),             // 6: pb.ApiResponseOrderItemDelete
	(*ApiResponseOrderItemAll)(nil),                // 7: pb.ApiResponseOrderItemAll
	(*ApiResponsePaginationOrderItemDeleteAt)(nil), // 8: pb.ApiResponsePaginationOrderItemDeleteAt
	(*ApiResponsePaginationOrderItem)(nil),         // 9: pb.ApiResponsePaginationOrderItem
	(*PaginationMeta)(nil),                         // 10: pb.PaginationMeta
}
var file_order_item_proto_depIdxs = []int32{
	2,  // 0: pb.ApiResponseOrderItem.data:type_name -> pb.OrderItemResponse
	2,  // 1: pb.ApiResponsesOrderItem.data:type_name -> pb.OrderItemResponse
	3,  // 2: pb.ApiResponsePaginationOrderItemDeleteAt.data:type_name -> pb.OrderItemResponseDeleteAt
	10, // 3: pb.ApiResponsePaginationOrderItemDeleteAt.pagination:type_name -> pb.PaginationMeta
	2,  // 4: pb.ApiResponsePaginationOrderItem.data:type_name -> pb.OrderItemResponse
	10, // 5: pb.ApiResponsePaginationOrderItem.pagination:type_name -> pb.PaginationMeta
	0,  // 6: pb.OrderItemService.FindAll:input_type -> pb.FindAllOrderItemRequest
	0,  // 7: pb.OrderItemService.FindByActive:input_type -> pb.FindAllOrderItemRequest
	0,  // 8: pb.OrderItemService.FindByTrashed:input_type -> pb.FindAllOrderItemRequest
	1,  // 9: pb.OrderItemService.FindOrderItemByOrder:input_type -> pb.FindByIdOrderItemRequest
	9,  // 10: pb.OrderItemService.FindAll:output_type -> pb.ApiResponsePaginationOrderItem
	8,  // 11: pb.OrderItemService.FindByActive:output_type -> pb.ApiResponsePaginationOrderItemDeleteAt
	8,  // 12: pb.OrderItemService.FindByTrashed:output_type -> pb.ApiResponsePaginationOrderItemDeleteAt
	5,  // 13: pb.OrderItemService.FindOrderItemByOrder:output_type -> pb.ApiResponsesOrderItem
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_order_item_proto_init() }
func file_order_item_proto_init() {
	if File_order_item_proto != nil {
		return
	}
	file_api_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_order_item_proto_rawDesc), len(file_order_item_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_item_proto_goTypes,
		DependencyIndexes: file_order_item_proto_depIdxs,
		MessageInfos:      file_order_item_proto_msgTypes,
	}.Build()
	File_order_item_proto = out.File
	file_order_item_proto_goTypes = nil
	file_order_item_proto_depIdxs = nil
}
