package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/dimitarsi/go-chat-app/grpc/room"
	"google.golang.org/grpc"
)


func client(roomId string) {
	cl, err := grpc.Dial("localhost:5000", grpc.WithInsecure())


	if err != nil {
		fmt.Println("error connecting to server")
		return
	}


	roomClient := room.NewRoomsClient(cl)

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Minute))

	stream, err := roomClient.RoomDetails(ctx, &room.RoomDetailsRequest{
		RoomId: roomId,
	})

	for {
		resp, err := stream.Recv()

		if err != io.EOF {
			fmt.Printf("Receive %v", *resp.Participants)
		}
	}
}