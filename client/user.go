package user

import (
	"archis/pb"
	"context"
	"time"

	"google.golang.org/grpc"
)

//Client implements UserServiceClient
type Client struct {
	service pb.UserServiceClient
}

//NewUserClient returns a new instance of user client
func NewUserClient(cc *grpc.ClientConn) *Client {
	service := pb.NewUserServiceClient(cc)
	return &Client{service}
}

//CreateUser makes req. for creating a user
func (client *Client) CreateUser(email string, amount int) (*pb.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.service.CreateUser(ctx, &pb.CreateUserRequest{
		Email:  email,
		Amount: uint64(amount),
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//UpdateUser makes req. for updating a user
func (client *Client) UpdateUser(id, email string, amount int) (*pb.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.service.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:     id,
		Email:  email,
		Amount: uint64(amount),
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//GetUser makes req. for getting a user
func (client *Client) GetUser(id string) (*pb.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.service.GetUser(ctx, &pb.IDRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//DeleteUser makes req. for deleting a user
func (client *Client) DeleteUser(id string) (*pb.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.service.DeleteUser(ctx, &pb.IDRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
