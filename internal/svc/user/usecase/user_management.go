package usecase

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/captain-corgi/grpc-go-example/api/user"
)

//UserManagement implement UserServiceServer
type UserManagement struct {
	pb.UnimplementedUserServiceServer
}

//NewUserManagement return new user management service.
func NewUserManagement() *UserManagement {
	return &UserManagement{}
}

//ListUsers return list of users
func (r *UserManagement) ListUsers(context.Context, *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users := pb.ListUsersResponse{}
	users.Users = []*pb.User{
		{
			Name:     "Anh1",
			Email:    "anh1@mail.com",
			Password: "******",
		},
		{
			Name:     "Anh2",
			Email:    "anh2@mail.com",
			Password: "******",
		},
	}
	users.NextPageToken = "END"
	return &users, nil
}

//GetUser return an user details
func (r *UserManagement) GetUser(context.Context, *pb.GetUserRequest) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

//CreateUser add an user to database
func (r *UserManagement) CreateUser(context.Context, *pb.CreateUserRequest) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

//UpdateUser change an user info
func (r *UserManagement) UpdateUser(context.Context, *pb.UpdateUserRequest) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

//DeleteUser remove an user from database
func (r *UserManagement) DeleteUser(context.Context, *pb.DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
