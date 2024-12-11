package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	v1 "user/api/user/v1"
)

var userClient v1.UserClient
var conn *grpc.ClientConn

func main() {
	Init()
	TestCreateUser()
	conn.Close()
}

func Init() {
	var err error
	conn, err = grpc.NewClient("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("grpc link err:" + err.Error())
	}
	userClient = v1.NewUserClient(conn)
}

// TestCreateUser 创建用户
func TestCreateUser() {
	rsp, err := userClient.CreateUser(context.TODO(), &v1.CreateUserRequest{
		Mobile:   fmt.Sprintf("1380013800%d", 1),
		Password: "<PASSWORD>",
		NickName: fmt.Sprintf("test%d", 1),
	})
	if err != nil {
		panic("grpc create user err:" + err.Error())
	}
	fmt.Println(rsp)
}
