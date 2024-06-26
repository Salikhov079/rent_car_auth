package handler

import pb "github.com/Salikhov079/rent_car/genprotos"

type Handler struct {
	User   pb.UserServiceClient
}

func NewHandler(us pb.UserServiceClient) *Handler {
	return &Handler{us}
}
