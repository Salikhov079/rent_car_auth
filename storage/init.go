package postgres

import (
	pb "github.com/Salikhov079/rent_car/genprotos"
)

type InitRoot interface {
	User() User
}
type User interface {
	Create(user *pb.User) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.User, error)
	GetAll(_ *pb.User) (*pb.GetAllUsers, error)
	Update(user *pb.User) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
	Login(user *pb.User ) (*pb.User, error)
}
