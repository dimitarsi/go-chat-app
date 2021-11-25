package room

import (
	"context"

	r "github.com/dimitarsi/go-chat-app/grpc/room"
)

type RoomService struct {
	r.UnimplementedRoomsServer
}


func (rs * RoomService)	CreateRoom(context.Context, *r.CreateRoomRequest) (*r.NewRoomResponse, error) {
	nr := &r.NewRoomResponse{}

	return nr, nil
}
func (rs * RoomService)	JoinRoom(context.Context, *r.JoinRequest) (*r.JoinResponse, error) {
	return &r.JoinResponse{}, nil
}

func (rs * RoomService)	LeaveRoom(ctx context.Context,in *r.LeaveRequest) (*r.LeaveResponse, error) {
	return &r.LeaveResponse{}, nil
}
func (rs * RoomService)	FindRoom(ctx context.Context,in *r.SearchRequest) (*r.RoomsFoundResponse, error) {
	return &r.RoomsFoundResponse{}, nil
}
func (rs * RoomService)	RoomDetails(in *r.RoomDetailsRequest, s r.Rooms_RoomDetailsServer) error {

	return nil
}