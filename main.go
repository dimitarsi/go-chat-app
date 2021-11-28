package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	room "github.com/dimitarsi/go-chat-app/grpc/room"
	service "github.com/dimitarsi/go-chat-app/services/room"
)

func main() {
	lis, err := net.Listen("tcp", ":5000")

	if err != nil {
		fmt.Printf("Unable to listen on port %s\n", "5000")
		panic(err)
	}

	s := grpc.NewServer()

	room.RegisterRoomsServer(s, service.NewRoomService())

	err = s.Serve(lis)

	if err == nil {
		fmt.Printf("Chat service started on port :5000")
	}
}
