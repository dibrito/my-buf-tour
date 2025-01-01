package main

import (
	"context"
	"log"
	"net/http"

	"github.com/dibrito/my-buf-tour/gen/boking/v1/bookingv1connect"

	bookingv1 "github.com/dibrito/my-buf-tour/gen/booking/v1"

	connect "connectrpc.com/connect"
)

func main() {
	client := bookingv1connect.NewBookingServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	_, err := client.CreateBooking(
		context.Background(),
		connect.NewRequest(&bookingv1.CreateBookingRequest{
			// Pet: &bookingv1.Pet{},
		}),
	)
	if err != nil {
		log.Println(err)
		return
	}
	// log.Println(res.Msg, "Successfully PutPet")
}
