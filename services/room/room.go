package room

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/dimitarsi/go-chat-app/database"
	r "github.com/dimitarsi/go-chat-app/grpc/room"
	"github.com/dimitarsi/go-chat-app/grpc/shared"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type RoomService struct {
	r.UnimplementedRoomsServer

	DB *redis.Client
}


func (rs * RoomService)	CreateRoom(ctx context.Context, in *r.CreateRoomRequest) (*r.NewRoomResponse, error) {

	roomId, err := uuid.NewRandom();

	if err != nil {
		return nil, errors.New("unable to create room")
	}

	roomIdName :=  fmt.Sprintf("%s:%s", roomId, in.RoomName)

	rs.DB.LPush("rooms", roomIdName)

	nr := &r.NewRoomResponse{
		Room: &shared.Room {
			RoomId: roomId.String(),
			RoomName: in.RoomName,
		},
	}

	return nr, nil
}
func (rs * RoomService)	JoinRoom(ctx context.Context,in *r.JoinRequest) (*r.JoinResponse, error) {


	return &r.JoinResponse{}, nil
}

func (rs * RoomService)	LeaveRoom(ctx context.Context,in *r.LeaveRequest) (*r.LeaveResponse, error) {
	return &r.LeaveResponse{}, nil
}
func (rs * RoomService)	FindRoom(ctx context.Context,in *r.SearchRequest) (*r.RoomsFoundResponse, error) {
	totalRooms := rs.DB.LLen("rooms")
	
	resp := &r.RoomsFoundResponse{
		RoomsFound: make([]*shared.Room, totalRooms.Val()),
	}

	rooms, _ := rs.DB.LRange("rooms", 0, totalRooms.Val()).Result()

	for i, room := range rooms {
		roomDataPair := strings.Split(room, ":")
		if len(roomDataPair) == 2 {
			resp.RoomsFound[i] = &shared.Room{
				RoomId: roomDataPair[0],
				RoomName: roomDataPair[1],
			}
		}
	}

	return resp, nil
}
func (rs * RoomService)	RoomDetails(in *r.RoomDetailsRequest, s r.Rooms_RoomDetailsServer) error {

	return nil
}

func NewRoomService() *RoomService {
	
	return &RoomService{
		DB: database.NewRedisClient(),
	}
}