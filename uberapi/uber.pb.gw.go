// Code generated by protoc-gen-grpc-gateway
// source: uber.proto
// DO NOT EDIT!

/*
Package uberapi is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package uberapi

import (
	"google/protobuf"
	"io"
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/gengo/grpc-gateway/utilities"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = utilities.NewDoubleArray

var (
	filter_UberAPIService_GetEstimatesPrice_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_UberAPIService_GetEstimatesPrice_0(ctx context.Context, marshaler runtime.Marshaler, client UberAPIServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetEstimatesPriceRequest
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_UberAPIService_GetEstimatesPrice_0); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetEstimatesPrice(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_UberAPIService_GetEstimatesTime_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_UberAPIService_GetEstimatesTime_0(ctx context.Context, marshaler runtime.Marshaler, client UberAPIServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetEstimatesTimeRequest
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_UberAPIService_GetEstimatesTime_0); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetEstimatesTime(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_UberAPIService_GetHistory_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_UberAPIService_GetHistory_0(ctx context.Context, marshaler runtime.Marshaler, client UberAPIServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetHistoryRequest
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_UberAPIService_GetHistory_0); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetHistory(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_UberAPIService_GetMe_0(ctx context.Context, marshaler runtime.Marshaler, client UberAPIServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq google_protobuf.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.GetMe(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_UberAPIService_GetProducts_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_UberAPIService_GetProducts_0(ctx context.Context, marshaler runtime.Marshaler, client UberAPIServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetProductsRequest
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_UberAPIService_GetProducts_0); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetProducts(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterUberAPIServiceHandlerFromEndpoint is same as RegisterUberAPIServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterUberAPIServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterUberAPIServiceHandler(ctx, mux, conn)
}

// RegisterUberAPIServiceHandler registers the http handlers for service UberAPIService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterUberAPIServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewUberAPIServiceClient(conn)

	mux.Handle("GET", pattern_UberAPIService_GetEstimatesPrice_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, req)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
		}
		resp, md, err := request_UberAPIService_GetEstimatesPrice_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_UberAPIService_GetEstimatesPrice_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_UberAPIService_GetEstimatesTime_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, req)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
		}
		resp, md, err := request_UberAPIService_GetEstimatesTime_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_UberAPIService_GetEstimatesTime_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_UberAPIService_GetHistory_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, req)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
		}
		resp, md, err := request_UberAPIService_GetHistory_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_UberAPIService_GetHistory_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_UberAPIService_GetMe_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, req)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
		}
		resp, md, err := request_UberAPIService_GetMe_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_UberAPIService_GetMe_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_UberAPIService_GetProducts_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, req)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
		}
		resp, md, err := request_UberAPIService_GetProducts_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_UberAPIService_GetProducts_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_UberAPIService_GetEstimatesPrice_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "estimates", "price"}, ""))

	pattern_UberAPIService_GetEstimatesTime_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "estimates", "time"}, ""))

	pattern_UberAPIService_GetHistory_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "history"}, ""))

	pattern_UberAPIService_GetMe_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "me"}, ""))

	pattern_UberAPIService_GetProducts_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "products"}, ""))
)

var (
	forward_UberAPIService_GetEstimatesPrice_0 = runtime.ForwardResponseMessage

	forward_UberAPIService_GetEstimatesTime_0 = runtime.ForwardResponseMessage

	forward_UberAPIService_GetHistory_0 = runtime.ForwardResponseMessage

	forward_UberAPIService_GetMe_0 = runtime.ForwardResponseMessage

	forward_UberAPIService_GetProducts_0 = runtime.ForwardResponseMessage
)
