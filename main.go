package main

import (
	"flag"
	"fmt"
	"net"

	"google.golang.org/grpc"

	room "github.com/dimitarsi/go-chat-app/grpc/room"
	"github.com/dimitarsi/go-chat-app/grpc/user_signup"
	service "github.com/dimitarsi/go-chat-app/services/room"
	"github.com/dimitarsi/go-chat-app/services/user"
)

func main() {

	runAsClient := flag.Bool("client", false, "Run application in client mode")
	roomId := flag.String("r", "", "Room to monitor")

	flag.Parse()

	fmt.Printf("run as client %v", *runAsClient)

	if *runAsClient {
		client(*roomId)
		return;
	}

	lis, err := net.Listen("tcp", ":5000")

	if err != nil {
		fmt.Printf("Unable to listen on port %s\n", "5000")
		panic(err)
	}

	s := grpc.NewServer()

	room.RegisterRoomsServer(s, service.NewRoomService())
	user_signup.RegisterSignupServiceServer(s, user.NewUsersService())

	err = s.Serve(lis)

	if err == nil {
		fmt.Printf("Chat service started on port :5000")
	}
}
