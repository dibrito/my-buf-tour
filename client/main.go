package main

import (
	"context"
	"log"
	"net/http"

	bookingv1 "github.com/dibrito/my-buf-tour/gen/booking/v1"
	"github.com/dibrito/my-buf-tour/gen/booking/v1/bookingv1connect"

	connect "connectrpc.com/connect"
	"google.golang.org/genproto/googleapis/type/datetime"
)

func main() {
	client := bookingv1connect.NewBookingServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.CreateBooking(
		context.Background(),
		connect.NewRequest(&bookingv1.CreateBookingRequest{
			UserId:  "1",
			HotelId: "1",
			CheckInDate: &datetime.DateTime{
				Year:  2025,
				Month: 1,
				Day:   1,
			},
			CheckOutDate: &datetime.DateTime{
				Year:  2025,
				Month: 1,
				Day:   5,
			},
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Confirmation Id:", res.Msg.ConfirmationId)
	log.Println("Status:", res.Msg.Status)
}
