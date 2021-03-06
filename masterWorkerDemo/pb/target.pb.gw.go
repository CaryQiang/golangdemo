// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: target.proto

/*
Package target is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package target

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_TargetService_TargetLibList_0(ctx context.Context, marshaler runtime.Marshaler, client TargetServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TargetLibListRequest
	var metadata runtime.ServerMetadata

	msg, err := client.TargetLibList(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_TargetService_TargetAdd_0(ctx context.Context, marshaler runtime.Marshaler, client TargetServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TargetAddRequest
	var metadata runtime.ServerMetadata

	if req.ContentLength > 0 {
		if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	msg, err := client.TargetAdd(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_TargetService_TargetDelete_0(ctx context.Context, marshaler runtime.Marshaler, client TargetServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq TargetDelRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["target_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "target_id")
	}

	protoReq.TargetId, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "target_id", err)
	}

	msg, err := client.TargetDelete(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterTargetServiceHandlerFromEndpoint is same as RegisterTargetServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterTargetServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterTargetServiceHandler(ctx, mux, conn)
}

// RegisterTargetServiceHandler registers the http handlers for service TargetService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterTargetServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterTargetServiceHandlerClient(ctx, mux, NewTargetServiceClient(conn))
}

// RegisterTargetServiceHandler registers the http handlers for service TargetService to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "TargetServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "TargetServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "TargetServiceClient" to call the correct interceptors.
func RegisterTargetServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client TargetServiceClient) error {

	mux.Handle("GET", pattern_TargetService_TargetLibList_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
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
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_TargetService_TargetLibList_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TargetService_TargetLibList_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_TargetService_TargetAdd_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
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
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_TargetService_TargetAdd_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TargetService_TargetAdd_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_TargetService_TargetDelete_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
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
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_TargetService_TargetDelete_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_TargetService_TargetDelete_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_TargetService_TargetLibList_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"target_libs"}, ""))

	pattern_TargetService_TargetAdd_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"targets"}, ""))

	pattern_TargetService_TargetDelete_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1}, []string{"targets", "target_id"}, ""))
)

var (
	forward_TargetService_TargetLibList_0 = runtime.ForwardResponseMessage

	forward_TargetService_TargetAdd_0 = runtime.ForwardResponseMessage

	forward_TargetService_TargetDelete_0 = runtime.ForwardResponseMessage
)
