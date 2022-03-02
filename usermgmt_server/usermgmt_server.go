package main

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/tjoe1985/hello_grpc.git/usermgmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	user_list *pb.UserList
}

func (s *UserManagementServer) CreateNewUser(c context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Println("received: ", in.GetName(), in.GetAge())
	user_uuid := uuid.NewString()
	user := &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Uuid: user_uuid,
	}
	s.user_list.Users = append(s.user_list.Users, user)
	return user, nil
}
func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	return s.user_list, nil
}
func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{
		user_list: &pb.UserList{},
	}
}
func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("error failed to listen: ", err)
	}
	s := grpc.NewServer()
	//register server as new grpc service
	pb.RegisterUserManagementServer(s, server)
	log.Println("server listening on :", lis.Addr())
	//start the server
	return s.Serve(lis)
}
func main() {
	user_mgmt_server := NewUserManagementServer()
	if err := user_mgmt_server.Run(); err != nil {
		log.Println("failed to serve: ", err)
	}

}
