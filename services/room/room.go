package room

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dimitarsi/go-chat-app/database"
	r "github.com/dimitarsi/go-chat-app/grpc/room"
	"github.com/dimitarsi/go-chat-app/grpc/shared"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
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

	res := rs.DB.LPush(in.RoomId, in.UserId)

	if res.Err() != nil {
		return &r.JoinResponse{
			Status: r.JoinResponse_DENIED,
		}, errors.New("unable to join room")
	}

	usersInRoom := rs.DB.LLen(in.RoomId).Val()

	fmt.Printf("Total users in this room - %d", usersInRoom)


	return &r.JoinResponse{
		Status: r.JoinResponse_ALLOWED,
	}, nil
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

	roomUsers := rs.DB.LLen(in.RoomId).Val()
	errChan := make(chan error)

	dline, ok := s.Context().Deadline()

	s.SendHeader(metadata.MD{
		"hello": {"world", dline.String(), fmt.Sprintf("%v", ok)},
	})

	if roomUsers == 0 {
		return errors.New("empty room")
	}

	go func(errChan chan error) {
		for {
			roomUsers = rs.DB.LLen(in.RoomId).Val()
			allUsers := rs.DB.LRange(in.RoomId, 0 , roomUsers)

			fmt.Printf("Users found %d", allUsers)

			err := s.Send(&r.RoomDetailsResponse{
				Participants: &r.RoomDetailsResponse_Participants{
					Count: 0,
					Users: getUserFromIds(allUsers.Val()),
				},
			})

			ctx := s.Context()
			deadline, ok := ctx.Deadline()

			fmt.Printf("Deadline by %v, is ok - %v", deadline, ok)

			if err != nil {
				fmt.Printf("error sending the request %v", err)
				errChan <- err;
				break;
			} else {
				fmt.Printf("Message send")
			}
		
			time.Sleep(time.Second*10)
			
		}
	}(errChan)
	// var sendErr error
	sendErr := <- errChan

	return sendErr
}

func getUserFromIds(userIds []string) []*shared.User {
	users := make([]*shared.User, len(userIds))

	for ind, userId := range userIds {
		users[ind] = &shared.User {
			UserId: userId,
			DisplayName: fmt.Sprintf("User-%s", userId),
		}
	}

	return users;
}

func NewRoomService() *RoomService {
	
	return &RoomService{
		DB: database.NewRedisClient(),
	}
}