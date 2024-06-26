package service

import (
	"context"
	"log"

	pb "github.com/Salikhov079/rent_car/genprotos"
	s "github.com/Salikhov079/rent_car/storage"
)

type UserService struct {
	stg s.InitRoot
	pb.UnimplementedUserServiceServer
}

func NewUserService(stg s.InitRoot) *UserService {
	return &UserService{stg: stg}
}

func (c *UserService) Create(ctx context.Context, User *pb.User) (*pb.Void, error) {
	pb, err := c.stg.User().Create(User)
	if err != nil {
		log.Print(err)
	}
	return pb, err
}

func (c *UserService) GetAll(ctx context.Context, pb *pb.User) (*pb.GetAllUsers, error) {
	Users, err := c.stg.User().GetAll(pb)
	if err != nil {
		log.Print(err)
	}

	return Users, err
}

func (c *UserService) GetById(ctx context.Context, id *pb.ById) (*pb.User, error) {
	prod, err := c.stg.User().GetById(id)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}

func (c *UserService) Update(ctx context.Context, User *pb.User) (*pb.Void, error) {
	pb, err := c.stg.User().Update(User)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *UserService) Delete(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	pb, err := c.stg.User().Delete(id)
	if err != nil {
		log.Print(err)
	}

	return pb, err
}

func (c *UserService) Login(ctx context.Context, username *pb.User) (*pb.User, error) {
	prod, err := c.stg.User().Login(username)
	if err != nil {
		log.Print(err)
	}

	return prod, err
}
