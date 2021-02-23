package interceptor

import (
	"archis/pb"
	"archis/pkg/user"
	"context"
)

//Intercept holds the userSvc instance and implements gRPC generated interface
type Intercept struct {
	userSvc user.Service
}

//NewUserInterceptor returns a new instance of interceptor
func NewUserInterceptor(userSvc user.Service) pb.UserServiceServer {
	return Intercept{userSvc}
}

//CreateUser creates a user
func (i Intercept) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.Response, error) {
	err := i.userSvc.CreateUser(user.AuthRequest{
		Amount: int(req.Amount),
		Email:  req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Message: "User created successfully!",
		Error:   "",
	}, nil
}

//UpdateUser updates a user
func (i Intercept) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.Response, error) {
	intAmount := int(req.Amount)

	_, err := i.userSvc.UpdateUser(user.UpdateRequest{
		ID:     req.Id,
		Amount: &intAmount,
		Email:  &req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Message: "User updated successfully!",
		Error:   "",
		//User: &pb.User{
		//	Amount: uint64(updatedUser.Amount),
		//	Id:     updatedUser.ID,
		//	Email:  updatedUser.Email,
		//},
	}, nil
}

//GetUser fetches a user
func (i Intercept) GetUser(ctx context.Context, req *pb.IDRequest) (*pb.Response, error) {

	_, err := i.userSvc.GetUser(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Message: "User fetched successfully!",
		Error:   "",
		//User: &pb.User{
		//	Amount: uint64(fetchedUser.Amount),
		//	Id:     fetchedUser.ID,
		//	Email:  fetchedUser.Email,
		//},
	}, nil
}

//DeleteUser deletes a user
func (i Intercept) DeleteUser(ctx context.Context, req *pb.IDRequest) (*pb.Response, error) {

	err := i.userSvc.DeleteUser(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Message: "User deleted successfully!",
		Error:   "",
	}, nil
}
