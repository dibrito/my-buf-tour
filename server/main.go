package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	bookingv1 "github.com/dibrito/my-buf-tour/gen/booking/v1"
	bookingv1connect "github.com/dibrito/my-buf-tour/gen/booking/v1/bookingv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const address = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	path, handler := bookingv1connect.NewBookingServiceHandler(&bookingServiceServer{})
	mux.Handle(path, handler)
	fmt.Println("... Listen on:", address)
	http.ListenAndServe(address,
		h2c.NewHandler(mux, &http2.Server{}),
	)

}

// bookingServiceServer implementes the BookingService API.
type bookingServiceServer struct {
}

// CreateBooking implements bookingv1connect.BookingServiceHandler.
func (b *bookingServiceServer) CreateBooking(
	ctx context.Context,
	req *connect.Request[bookingv1.CreateBookingRequest],
) (*connect.Response[bookingv1.CreateBookingResponse], error) {
	userId := req.Msg.GetUserId()
	hotelId := req.Msg.GetHotelId()
	checkInDate := req.Msg.GetCheckInDate()
	checkOutDate := req.Msg.GetCheckOutDate()

	fmt.Println("Got a request to create a booking")
	fmt.Println("User", userId, "is booking hotel", hotelId, "from", checkInDate, "to", checkOutDate)

	return connect.NewResponse(&bookingv1.CreateBookingResponse{
		ConfirmationId: "123",
		Status:         bookingv1.BookingStatus_BOOKING_STATUS_CONFIRMED,
	}), nil
}

// GetBooking implements bookingv1connect.BookingServiceHandler.
func (b *bookingServiceServer) GetBooking(
	ctx context.Context,
	req *connect.Request[bookingv1.GetBookingRequest],
) (*connect.Response[bookingv1.GetBookingResponse], error) {
	userId := req.Msg.GetBookingId()
	bookingId := req.Msg.GetBookingId()

	fmt.Println("Got a request to get a booking")
	fmt.Println("Booking ID:", bookingId)

	return connect.NewResponse(&bookingv1.GetBookingResponse{
		UserId:    userId,
		HotelId:   "1",
		BookingId: "1",
	}), nil
}
