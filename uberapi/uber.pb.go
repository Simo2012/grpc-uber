// Code generated by protoc-gen-go.
// source: uber.proto
// DO NOT EDIT!

/*
Package uberapi is a generated protocol buffer package.

It is generated from these files:
	uber.proto

It has these top-level messages:
	GetEstimatesPriceRequest
	GetEstimatesPriceResponse
	GetEstimatesTimeRequest
	GetEstimatesTimeResponse
	GetHistoryRequest
	GetProductsRequest
	GetProductsResponse
	Activities
	Activity
	Error
	PriceEstimate
	Product
	Profile
*/
package uberapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "google/protobuf"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetEstimatesPriceRequest struct {
	EndLatitude    float64 `protobuf:"fixed64,1,opt,name=end_latitude,json=endLatitude" json:"end_latitude,omitempty"`
	EndLongitude   float64 `protobuf:"fixed64,2,opt,name=end_longitude,json=endLongitude" json:"end_longitude,omitempty"`
	StartLatitude  float64 `protobuf:"fixed64,3,opt,name=start_latitude,json=startLatitude" json:"start_latitude,omitempty"`
	StartLongitude float64 `protobuf:"fixed64,4,opt,name=start_longitude,json=startLongitude" json:"start_longitude,omitempty"`
}

func (m *GetEstimatesPriceRequest) Reset()                    { *m = GetEstimatesPriceRequest{} }
func (m *GetEstimatesPriceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEstimatesPriceRequest) ProtoMessage()               {}
func (*GetEstimatesPriceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetEstimatesPriceResponse struct {
	Items []*PriceEstimate `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *GetEstimatesPriceResponse) Reset()                    { *m = GetEstimatesPriceResponse{} }
func (m *GetEstimatesPriceResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEstimatesPriceResponse) ProtoMessage()               {}
func (*GetEstimatesPriceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetEstimatesPriceResponse) GetItems() []*PriceEstimate {
	if m != nil {
		return m.Items
	}
	return nil
}

type GetEstimatesTimeRequest struct {
	CustomerUuid   string  `protobuf:"bytes,1,opt,name=customer_uuid,json=customerUuid" json:"customer_uuid,omitempty"`
	ProductId      string  `protobuf:"bytes,2,opt,name=product_id,json=productId" json:"product_id,omitempty"`
	StartLatitude  float64 `protobuf:"fixed64,3,opt,name=start_latitude,json=startLatitude" json:"start_latitude,omitempty"`
	StartLongitude float64 `protobuf:"fixed64,4,opt,name=start_longitude,json=startLongitude" json:"start_longitude,omitempty"`
}

func (m *GetEstimatesTimeRequest) Reset()                    { *m = GetEstimatesTimeRequest{} }
func (m *GetEstimatesTimeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEstimatesTimeRequest) ProtoMessage()               {}
func (*GetEstimatesTimeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GetEstimatesTimeResponse struct {
	Items []*Product `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *GetEstimatesTimeResponse) Reset()                    { *m = GetEstimatesTimeResponse{} }
func (m *GetEstimatesTimeResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEstimatesTimeResponse) ProtoMessage()               {}
func (*GetEstimatesTimeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetEstimatesTimeResponse) GetItems() []*Product {
	if m != nil {
		return m.Items
	}
	return nil
}

type GetHistoryRequest struct {
	Limit  int32 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *GetHistoryRequest) Reset()                    { *m = GetHistoryRequest{} }
func (m *GetHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetHistoryRequest) ProtoMessage()               {}
func (*GetHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type GetProductsRequest struct {
	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *GetProductsRequest) Reset()                    { *m = GetProductsRequest{} }
func (m *GetProductsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetProductsRequest) ProtoMessage()               {}
func (*GetProductsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type GetProductsResponse struct {
	Items []*Product `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *GetProductsResponse) Reset()                    { *m = GetProductsResponse{} }
func (m *GetProductsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetProductsResponse) ProtoMessage()               {}
func (*GetProductsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetProductsResponse) GetItems() []*Product {
	if m != nil {
		return m.Items
	}
	return nil
}

type Activities struct {
	Count   int32       `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	History []*Activity `protobuf:"bytes,2,rep,name=history" json:"history,omitempty"`
	Limit   int32       `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset  int32       `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *Activities) Reset()                    { *m = Activities{} }
func (m *Activities) String() string            { return proto.CompactTextString(m) }
func (*Activities) ProtoMessage()               {}
func (*Activities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Activities) GetHistory() []*Activity {
	if m != nil {
		return m.History
	}
	return nil
}

type Activity struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
}

func (m *Activity) Reset()                    { *m = Activity{} }
func (m *Activity) String() string            { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()               {}
func (*Activity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type Error struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Fields  string `protobuf:"bytes,2,opt,name=fields" json:"fields,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type PriceEstimate struct {
	CurrencyCode    string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode" json:"currency_code,omitempty"`
	DisplayName     string `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	Estimate        string `protobuf:"bytes,3,opt,name=estimate" json:"estimate,omitempty"`
	HighEstimate    int32  `protobuf:"varint,4,opt,name=high_estimate,json=highEstimate" json:"high_estimate,omitempty"`
	LowEstimate     int32  `protobuf:"varint,5,opt,name=low_estimate,json=lowEstimate" json:"low_estimate,omitempty"`
	ProductId       string `protobuf:"bytes,6,opt,name=product_id,json=productId" json:"product_id,omitempty"`
	SurgeMultiplier int32  `protobuf:"varint,7,opt,name=surge_multiplier,json=surgeMultiplier" json:"surge_multiplier,omitempty"`
}

func (m *PriceEstimate) Reset()                    { *m = PriceEstimate{} }
func (m *PriceEstimate) String() string            { return proto.CompactTextString(m) }
func (*PriceEstimate) ProtoMessage()               {}
func (*PriceEstimate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type Product struct {
	Capacity    string `protobuf:"bytes,1,opt,name=capacity" json:"capacity,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	Image       string `protobuf:"bytes,4,opt,name=image" json:"image,omitempty"`
	ProductId   string `protobuf:"bytes,5,opt,name=product_id,json=productId" json:"product_id,omitempty"`
}

func (m *Product) Reset()                    { *m = Product{} }
func (m *Product) String() string            { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()               {}
func (*Product) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type Profile struct {
	Email     string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Picture   string `protobuf:"bytes,4,opt,name=picture" json:"picture,omitempty"`
	PromoCode string `protobuf:"bytes,5,opt,name=promo_code,json=promoCode" json:"promo_code,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func init() {
	proto.RegisterType((*GetEstimatesPriceRequest)(nil), "uberapi.GetEstimatesPriceRequest")
	proto.RegisterType((*GetEstimatesPriceResponse)(nil), "uberapi.GetEstimatesPriceResponse")
	proto.RegisterType((*GetEstimatesTimeRequest)(nil), "uberapi.GetEstimatesTimeRequest")
	proto.RegisterType((*GetEstimatesTimeResponse)(nil), "uberapi.GetEstimatesTimeResponse")
	proto.RegisterType((*GetHistoryRequest)(nil), "uberapi.GetHistoryRequest")
	proto.RegisterType((*GetProductsRequest)(nil), "uberapi.GetProductsRequest")
	proto.RegisterType((*GetProductsResponse)(nil), "uberapi.GetProductsResponse")
	proto.RegisterType((*Activities)(nil), "uberapi.Activities")
	proto.RegisterType((*Activity)(nil), "uberapi.Activity")
	proto.RegisterType((*Error)(nil), "uberapi.Error")
	proto.RegisterType((*PriceEstimate)(nil), "uberapi.PriceEstimate")
	proto.RegisterType((*Product)(nil), "uberapi.Product")
	proto.RegisterType((*Profile)(nil), "uberapi.Profile")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for UberAPIService service

type UberAPIServiceClient interface {
	GetEstimatesPrice(ctx context.Context, in *GetEstimatesPriceRequest, opts ...grpc.CallOption) (*GetEstimatesPriceResponse, error)
	GetEstimatesTime(ctx context.Context, in *GetEstimatesTimeRequest, opts ...grpc.CallOption) (*GetEstimatesTimeResponse, error)
	GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*Activities, error)
	GetMe(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*Profile, error)
	GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
}

type uberAPIServiceClient struct {
	cc *grpc.ClientConn
}

func NewUberAPIServiceClient(cc *grpc.ClientConn) UberAPIServiceClient {
	return &uberAPIServiceClient{cc}
}

func (c *uberAPIServiceClient) GetEstimatesPrice(ctx context.Context, in *GetEstimatesPriceRequest, opts ...grpc.CallOption) (*GetEstimatesPriceResponse, error) {
	out := new(GetEstimatesPriceResponse)
	err := grpc.Invoke(ctx, "/uberapi.UberAPIService/GetEstimatesPrice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uberAPIServiceClient) GetEstimatesTime(ctx context.Context, in *GetEstimatesTimeRequest, opts ...grpc.CallOption) (*GetEstimatesTimeResponse, error) {
	out := new(GetEstimatesTimeResponse)
	err := grpc.Invoke(ctx, "/uberapi.UberAPIService/GetEstimatesTime", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uberAPIServiceClient) GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*Activities, error) {
	out := new(Activities)
	err := grpc.Invoke(ctx, "/uberapi.UberAPIService/GetHistory", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uberAPIServiceClient) GetMe(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := grpc.Invoke(ctx, "/uberapi.UberAPIService/GetMe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uberAPIServiceClient) GetProducts(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := grpc.Invoke(ctx, "/uberapi.UberAPIService/GetProducts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UberAPIService service

type UberAPIServiceServer interface {
	GetEstimatesPrice(context.Context, *GetEstimatesPriceRequest) (*GetEstimatesPriceResponse, error)
	GetEstimatesTime(context.Context, *GetEstimatesTimeRequest) (*GetEstimatesTimeResponse, error)
	GetHistory(context.Context, *GetHistoryRequest) (*Activities, error)
	GetMe(context.Context, *google_protobuf.Empty) (*Profile, error)
	GetProducts(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
}

func RegisterUberAPIServiceServer(s *grpc.Server, srv UberAPIServiceServer) {
	s.RegisterService(&_UberAPIService_serviceDesc, srv)
}

func _UberAPIService_GetEstimatesPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEstimatesPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UberAPIServiceServer).GetEstimatesPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uberapi.UberAPIService/GetEstimatesPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UberAPIServiceServer).GetEstimatesPrice(ctx, req.(*GetEstimatesPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UberAPIService_GetEstimatesTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEstimatesTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UberAPIServiceServer).GetEstimatesTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uberapi.UberAPIService/GetEstimatesTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UberAPIServiceServer).GetEstimatesTime(ctx, req.(*GetEstimatesTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UberAPIService_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UberAPIServiceServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uberapi.UberAPIService/GetHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UberAPIServiceServer).GetHistory(ctx, req.(*GetHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UberAPIService_GetMe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UberAPIServiceServer).GetMe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uberapi.UberAPIService/GetMe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UberAPIServiceServer).GetMe(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UberAPIService_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UberAPIServiceServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uberapi.UberAPIService/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UberAPIServiceServer).GetProducts(ctx, req.(*GetProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UberAPIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "uberapi.UberAPIService",
	HandlerType: (*UberAPIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEstimatesPrice",
			Handler:    _UberAPIService_GetEstimatesPrice_Handler,
		},
		{
			MethodName: "GetEstimatesTime",
			Handler:    _UberAPIService_GetEstimatesTime_Handler,
		},
		{
			MethodName: "GetHistory",
			Handler:    _UberAPIService_GetHistory_Handler,
		},
		{
			MethodName: "GetMe",
			Handler:    _UberAPIService_GetMe_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _UberAPIService_GetProducts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func init() { proto.RegisterFile("uber.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 873 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xeb, 0x44,
	0x14, 0x56, 0x92, 0x3a, 0x69, 0x4e, 0x7e, 0x6e, 0xef, 0xa4, 0x5c, 0x82, 0x53, 0x50, 0x6b, 0x04,
	0x5c, 0x04, 0x4a, 0xc4, 0x65, 0xcd, 0x22, 0xa0, 0x0a, 0x2a, 0xb5, 0x55, 0x65, 0xe8, 0x96, 0xc8,
	0xb5, 0x27, 0xe9, 0x48, 0xfe, 0x63, 0x66, 0xdc, 0x2a, 0x0b, 0x36, 0x48, 0x3c, 0x01, 0xec, 0x79,
	0x02, 0x36, 0xbc, 0x0a, 0xaf, 0xc0, 0x2b, 0xb0, 0x67, 0xfe, 0x3c, 0xb1, 0xdd, 0x50, 0x09, 0x89,
	0x9d, 0xcf, 0x37, 0xdf, 0x7c, 0xe7, 0x7c, 0xe3, 0x33, 0x67, 0x00, 0x8a, 0x3b, 0x4c, 0xe7, 0x39,
	0xcd, 0x78, 0x86, 0x7a, 0xf2, 0x3b, 0xc8, 0x89, 0x3b, 0xdb, 0x64, 0xd9, 0x26, 0xc6, 0x0b, 0x05,
	0xdf, 0x15, 0xeb, 0x05, 0x4e, 0x72, 0xbe, 0xd5, 0x2c, 0xf7, 0xc4, 0x2c, 0x0a, 0xe2, 0x22, 0x48,
	0xd3, 0x8c, 0x07, 0x9c, 0x64, 0x29, 0xd3, 0xab, 0xde, 0x1f, 0x2d, 0x98, 0x7e, 0x8d, 0xf9, 0x39,
	0xe3, 0x24, 0x09, 0x38, 0x66, 0x37, 0x94, 0x84, 0xd8, 0xc7, 0x3f, 0x14, 0x98, 0x71, 0x74, 0x06,
	0x43, 0x9c, 0x46, 0xab, 0x58, 0xec, 0xe0, 0x45, 0x84, 0xa7, 0xad, 0xd3, 0xd6, 0xeb, 0x96, 0x3f,
	0x10, 0xd8, 0xa5, 0x81, 0xd0, 0xfb, 0x30, 0x52, 0x94, 0x2c, 0xdd, 0x68, 0x4e, 0x5b, 0x71, 0xe4,
	0xbe, 0xcb, 0x12, 0x43, 0x1f, 0xc0, 0x98, 0xf1, 0x80, 0xf2, 0x9d, 0x52, 0x47, 0xb1, 0x46, 0x0a,
	0xb5, 0x5a, 0x1f, 0xc1, 0x0b, 0x43, 0xb3, 0x6a, 0x07, 0x8a, 0xa7, 0x77, 0x5b, 0x3d, 0xef, 0x02,
	0xde, 0xd9, 0x53, 0x33, 0xcb, 0x85, 0x2d, 0x8c, 0x3e, 0x05, 0x87, 0x70, 0x9c, 0x30, 0x51, 0x6d,
	0xe7, 0xf5, 0xe0, 0xcd, 0xab, 0xb9, 0x39, 0xa5, 0xb9, 0xa2, 0x95, 0x9b, 0x7c, 0x4d, 0xf2, 0x7e,
	0x6f, 0xc1, 0xdb, 0x55, 0xad, 0xef, 0x48, 0x62, 0xed, 0x0b, 0x6f, 0x61, 0xc1, 0x78, 0x96, 0x60,
	0xba, 0x2a, 0x0a, 0x12, 0x29, 0xff, 0x7d, 0x7f, 0x58, 0x82, 0xb7, 0x02, 0x43, 0xef, 0x02, 0x88,
	0x93, 0x8c, 0x8a, 0x90, 0xaf, 0x04, 0xa3, 0xad, 0x18, 0x7d, 0x83, 0x5c, 0x44, 0xff, 0xbb, 0xf5,
	0x2f, 0xeb, 0xbf, 0x4b, 0x97, 0x6b, 0x9c, 0x7f, 0x58, 0x77, 0x7e, 0x54, 0x71, 0xae, 0xca, 0x29,
	0x3d, 0x2f, 0xe1, 0xa5, 0xd0, 0xf8, 0x86, 0x08, 0x17, 0x74, 0x5b, 0x9a, 0x3d, 0x06, 0x27, 0x26,
	0x09, 0xe1, 0xca, 0xa4, 0xe3, 0xeb, 0x00, 0xbd, 0x82, 0x6e, 0xb6, 0x5e, 0x33, 0xcc, 0x95, 0x33,
	0xc7, 0x37, 0x91, 0x77, 0x0d, 0x48, 0x48, 0x18, 0x5d, 0x56, 0x6a, 0xb8, 0x70, 0xd8, 0xe8, 0x15,
	0x1b, 0xa3, 0x13, 0xe8, 0x37, 0x9b, 0x64, 0x07, 0x78, 0x5f, 0xc0, 0xa4, 0xa6, 0xf7, 0x1f, 0x1d,
	0xfd, 0x08, 0xb0, 0x0c, 0x39, 0x79, 0x20, 0x9c, 0x60, 0x26, 0xad, 0x84, 0x59, 0x91, 0x5a, 0x2b,
	0x2a, 0x40, 0x9f, 0x40, 0xef, 0x5e, 0x5b, 0x16, 0xe9, 0xa5, 0xda, 0x4b, 0xab, 0x66, 0xf6, 0x6e,
	0xfd, 0x92, 0xb1, 0x3b, 0x8d, 0xce, 0xfe, 0xd3, 0x38, 0xa8, 0x9d, 0xc6, 0x7b, 0x70, 0x58, 0x4a,
	0x20, 0x04, 0x07, 0x95, 0x5e, 0x51, 0xdf, 0xde, 0x15, 0x38, 0xe7, 0x94, 0x66, 0x54, 0x2e, 0x86,
	0x99, 0x39, 0x1c, 0xc7, 0x57, 0xdf, 0x52, 0x74, 0x4d, 0x70, 0x1c, 0x31, 0xd3, 0x3c, 0x26, 0x42,
	0x53, 0xe8, 0x25, 0x98, 0xb1, 0x60, 0xa3, 0x5b, 0xa6, 0xef, 0x97, 0xa1, 0xf7, 0x73, 0x1b, 0x46,
	0xb5, 0x66, 0xd6, 0x9d, 0x4a, 0x29, 0x4e, 0xc3, 0xed, 0xca, 0x26, 0x50, 0x9d, 0xaa, 0xc1, 0xaf,
	0x64, 0x22, 0x71, 0x9b, 0x23, 0xc2, 0xf2, 0x38, 0xd8, 0xae, 0xd2, 0x20, 0xc1, 0x26, 0xdd, 0xc0,
	0x60, 0xd7, 0x02, 0x92, 0x3f, 0x10, 0x1b, 0x4d, 0x93, 0xd4, 0xc6, 0x32, 0xc7, 0x3d, 0xd9, 0xdc,
	0xaf, 0x2c, 0x41, 0x9f, 0xc1, 0x50, 0x82, 0xb6, 0x10, 0x91, 0x23, 0xce, 0x1e, 0x77, 0x1c, 0x47,
	0x71, 0x06, 0x02, 0xb3, 0x94, 0xfa, 0x85, 0xe9, 0x36, 0x2f, 0xcc, 0xc7, 0x70, 0xc4, 0x0a, 0xba,
	0xc1, 0xab, 0xa4, 0x88, 0x39, 0xc9, 0x63, 0x82, 0xe9, 0xb4, 0xa7, 0x54, 0x5e, 0x28, 0xfc, 0xca,
	0xc2, 0xde, 0x6f, 0x2d, 0xe8, 0x99, 0x46, 0x90, 0x95, 0x87, 0x41, 0x1e, 0x84, 0xe2, 0x17, 0x18,
	0xf3, 0x36, 0x46, 0xa7, 0x30, 0x88, 0x30, 0x0b, 0x29, 0xc9, 0xe5, 0xe4, 0xb3, 0xbe, 0x77, 0xd0,
	0x93, 0xa3, 0xe9, 0x3c, 0x3d, 0x1a, 0xd1, 0x11, 0xa2, 0xfe, 0x8d, 0xb6, 0xdd, 0xf7, 0x75, 0xd0,
	0x30, 0xe3, 0x34, 0xcc, 0x78, 0xbf, 0xea, 0x0a, 0xd7, 0x24, 0x56, 0x02, 0x38, 0x09, 0x48, 0x6c,
	0xca, 0xd3, 0x81, 0x14, 0x58, 0x13, 0xca, 0x78, 0xf5, 0x97, 0xf4, 0x15, 0xa2, 0xb2, 0xce, 0xc4,
	0xad, 0x09, 0xca, 0x55, 0xf3, 0x47, 0x24, 0xa0, 0x16, 0x45, 0x87, 0xe4, 0x24, 0xe4, 0x05, 0x2d,
	0x8b, 0x2a, 0x43, 0x53, 0x56, 0x92, 0xe9, 0x66, 0xd8, 0x95, 0x95, 0x64, 0xb2, 0x13, 0xde, 0xfc,
	0xdd, 0x81, 0xf1, 0xad, 0xe8, 0xfd, 0xe5, 0xcd, 0xc5, 0xb7, 0x98, 0x3e, 0x88, 0x4e, 0x42, 0x8f,
	0x6a, 0x26, 0xd4, 0x47, 0x2a, 0x3a, 0xb3, 0x37, 0xe4, 0xdf, 0x9e, 0x08, 0xd7, 0x7b, 0x8e, 0xa2,
	0x6f, 0xb1, 0x37, 0xfb, 0xe9, 0xcf, 0xbf, 0x7e, 0x69, 0xbf, 0x85, 0x26, 0x8b, 0x87, 0xcf, 0x16,
	0x65, 0x6b, 0x30, 0xf1, 0x5a, 0xc9, 0x1c, 0x0c, 0x8e, 0x9a, 0x03, 0x0d, 0x9d, 0xee, 0x15, 0xad,
	0x8c, 0x66, 0xf7, 0xec, 0x19, 0x86, 0xc9, 0xea, 0xaa, 0xac, 0xc7, 0x08, 0xd5, 0xb3, 0x72, 0x99,
	0xc0, 0x07, 0xd8, 0x4d, 0x40, 0xe4, 0x56, 0xc5, 0xea, 0x63, 0xd1, 0x9d, 0x34, 0x87, 0x84, 0x18,
	0x30, 0xde, 0x44, 0x49, 0x8f, 0xd0, 0x40, 0x4a, 0x97, 0x23, 0x63, 0x09, 0x8e, 0xd8, 0x7e, 0x25,
	0x2e, 0xf4, 0x5c, 0xbf, 0xb8, 0xf3, 0xf2, 0x39, 0x9e, 0x9f, 0xcb, 0xe7, 0xd8, 0xad, 0x4d, 0x2f,
	0xd9, 0x12, 0xde, 0x58, 0xe9, 0x1c, 0xa2, 0xae, 0xd4, 0x11, 0x65, 0x7d, 0x0f, 0x83, 0xca, 0x14,
	0x44, 0xb3, 0x6a, 0x5d, 0x8d, 0x59, 0xeb, 0x9e, 0xec, 0x5f, 0x34, 0xe6, 0x8f, 0x95, 0xf2, 0x18,
	0x0d, 0xa5, 0xb2, 0xe9, 0x47, 0x76, 0xd7, 0x55, 0x15, 0x7d, 0xfe, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9e, 0x4e, 0x8a, 0x1d, 0x45, 0x08, 0x00, 0x00,
}
