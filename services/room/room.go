package room

import (
	"context"

	r "github.com/dimitarsi/go-chat-app/grpc/room"
)

type RoomService struct {
	r.UnimplementedRoomsServer
}


func (r * RoomService)	CreateRoom(context.Context, *r.CreateRoomRequest) (*r.NewRoomResponse, error) {
	nr := r.NewRoomResponse{}
}
func (r * RoomService)	JoinRoom(context.Context, *r.JoinRequest) (*r.JoinResponse, error) {}
func (r * RoomService)	LeaveRoom(context.Context, *r.LeaveRequest) (*r.LeaveResponse, error) {}
func (r * RoomService)	FindRoom(context.Context, *r.SearchRequest) (*r.RoomsFoundResponse, error) {}
func (r * RoomService)	RoomDetails(*r.RoomDetailsRequest, *r.Rooms_RoomDetailsServer) error {}