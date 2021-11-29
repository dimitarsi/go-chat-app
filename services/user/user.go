package user

import (
	"context"
	"errors"

	"github.com/dimitarsi/go-chat-app/grpc/shared"
	"github.com/dimitarsi/go-chat-app/grpc/user_signup"
	"github.com/google/uuid"
)

type UserService struct {
	user_signup.UnimplementedSignupServiceServer
}

func (us *UserService) Signup(ctx context.Context, in *user_signup.UserDetailsRequest) (*user_signup.UserLogginResponse, error) {
	random_user_id, err := uuid.NewRandom()

	if err != nil {
		return nil, errors.New("unable to create user")
	}

	return &user_signup.UserLogginResponse{
		User: &shared.User{
			UserId: random_user_id.String(),
			DisplayName: in.DisplayName,
		},
	}, nil
}

func NewUsersService() *UserService {
	return &UserService{}
}