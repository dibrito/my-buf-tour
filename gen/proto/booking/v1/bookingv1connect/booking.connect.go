// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/booking/v1/booking.proto

package bookingv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/dibrito/my-buf-tour/gen/proto/booking/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// BookingServiceName is the fully-qualified name of the BookingService service.
	BookingServiceName = "proto.booking.v1.BookingService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BookingServiceCreateBookingProcedure is the fully-qualified name of the BookingService's
	// CreateBooking RPC.
	BookingServiceCreateBookingProcedure = "/proto.booking.v1.BookingService/CreateBooking"
	// BookingServiceGetBookingProcedure is the fully-qualified name of the BookingService's GetBooking
	// RPC.
	BookingServiceGetBookingProcedure = "/proto.booking.v1.BookingService/GetBooking"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	bookingServiceServiceDescriptor             = v1.File_proto_booking_v1_booking_proto.Services().ByName("BookingService")
	bookingServiceCreateBookingMethodDescriptor = bookingServiceServiceDescriptor.Methods().ByName("CreateBooking")
	bookingServiceGetBookingMethodDescriptor    = bookingServiceServiceDescriptor.Methods().ByName("GetBooking")
)

// BookingServiceClient is a client for the proto.booking.v1.BookingService service.
type BookingServiceClient interface {
	// Sends a booking request.
	CreateBooking(context.Context, *connect.Request[v1.CreateBookingRequest]) (*connect.Response[v1.CreateBookingResponse], error)
	// Get booking details.
	GetBooking(context.Context, *connect.Request[v1.GetBookingRequest]) (*connect.Response[v1.GetBookingResponse], error)
}

// NewBookingServiceClient constructs a client for the proto.booking.v1.BookingService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBookingServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BookingServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &bookingServiceClient{
		createBooking: connect.NewClient[v1.CreateBookingRequest, v1.CreateBookingResponse](
			httpClient,
			baseURL+BookingServiceCreateBookingProcedure,
			connect.WithSchema(bookingServiceCreateBookingMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getBooking: connect.NewClient[v1.GetBookingRequest, v1.GetBookingResponse](
			httpClient,
			baseURL+BookingServiceGetBookingProcedure,
			connect.WithSchema(bookingServiceGetBookingMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// bookingServiceClient implements BookingServiceClient.
type bookingServiceClient struct {
	createBooking *connect.Client[v1.CreateBookingRequest, v1.CreateBookingResponse]
	getBooking    *connect.Client[v1.GetBookingRequest, v1.GetBookingResponse]
}

// CreateBooking calls proto.booking.v1.BookingService.CreateBooking.
func (c *bookingServiceClient) CreateBooking(ctx context.Context, req *connect.Request[v1.CreateBookingRequest]) (*connect.Response[v1.CreateBookingResponse], error) {
	return c.createBooking.CallUnary(ctx, req)
}

// GetBooking calls proto.booking.v1.BookingService.GetBooking.
func (c *bookingServiceClient) GetBooking(ctx context.Context, req *connect.Request[v1.GetBookingRequest]) (*connect.Response[v1.GetBookingResponse], error) {
	return c.getBooking.CallUnary(ctx, req)
}

// BookingServiceHandler is an implementation of the proto.booking.v1.BookingService service.
type BookingServiceHandler interface {
	// Sends a booking request.
	CreateBooking(context.Context, *connect.Request[v1.CreateBookingRequest]) (*connect.Response[v1.CreateBookingResponse], error)
	// Get booking details.
	GetBooking(context.Context, *connect.Request[v1.GetBookingRequest]) (*connect.Response[v1.GetBookingResponse], error)
}

// NewBookingServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBookingServiceHandler(svc BookingServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	bookingServiceCreateBookingHandler := connect.NewUnaryHandler(
		BookingServiceCreateBookingProcedure,
		svc.CreateBooking,
		connect.WithSchema(bookingServiceCreateBookingMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	bookingServiceGetBookingHandler := connect.NewUnaryHandler(
		BookingServiceGetBookingProcedure,
		svc.GetBooking,
		connect.WithSchema(bookingServiceGetBookingMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/proto.booking.v1.BookingService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BookingServiceCreateBookingProcedure:
			bookingServiceCreateBookingHandler.ServeHTTP(w, r)
		case BookingServiceGetBookingProcedure:
			bookingServiceGetBookingHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBookingServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBookingServiceHandler struct{}

func (UnimplementedBookingServiceHandler) CreateBooking(context.Context, *connect.Request[v1.CreateBookingRequest]) (*connect.Response[v1.CreateBookingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.booking.v1.BookingService.CreateBooking is not implemented"))
}

func (UnimplementedBookingServiceHandler) GetBooking(context.Context, *connect.Request[v1.GetBookingRequest]) (*connect.Response[v1.GetBookingResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.booking.v1.BookingService.GetBooking is not implemented"))
}
