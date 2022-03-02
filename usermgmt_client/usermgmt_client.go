package main

import (
	"context"
	pb "github.com/tjoe1985/hello_grpc.git/usermgmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Println("error connecting : ", err)
	}
	defer conn.Close()

	client := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// create a map with users to create
	var new_users = make(map[string]int32)
	new_users["Joel"] = 36
	new_users["Baba"] = 37
	new_users["Titi"] = 38

	for name, age := range new_users {
		r, err := client.CreateNewUser(ctx, &pb.NewUser{
			Name: name,
			Age:  age,
		})
		if err != nil {
			log.Println("Error creating user: ", err)
		}
		log.Println(" Created user is: ", r.GetName(), " Age: ", r.GetAge(), " Uuid: ", r.GetUuid())
	}
	params := &pb.GetUsersParams{}
	response, err := client.GetUsers(ctx, params)
	if err != nil {
		println("Error getting users: ", err)
	}
	u := response.GetUsers()
	log.Printf("Users from get users : %v ", u[0])
}
