package main

import (
	"context"
	"log"

	"github.com/qazaqpyn/grpc-app/proto/notification"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	c := notification.NewNotificationServiceClient(conn)

	response, err := c.Notify(context.Background(), &notification.NotificationRequest{Message: "WE are from CLIENT"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Status: ", response.Status)
}
