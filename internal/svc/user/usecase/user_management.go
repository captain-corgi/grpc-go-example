package usecase

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/captain-corgi/grpc-go-example/api/user"
)

type UserManagement struct {
	pb.UnimplementedUserServiceServer
}

//NewUserManagement return new user management service.
func NewUserManagement() *UserManagement {
	return &UserManagement{}
}

func (r *UserManagement) ListUsers(context.Context, *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (r *UserManagement) GetUser(context.Context, *pb.GetUserRequest) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (r *UserManagement) CreateUser(context.Context, *pb.CreateUserRequest) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (r *UserManagement) UpdateUser(context.Context, *pb.UpdateUserRequest) (*pb.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (r *UserManagement) DeleteUser(context.Context, *pb.DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
