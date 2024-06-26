package main

import (
	"fmt"
	"log"

	"github.com/Salikhov079/rent_car/api"
	"github.com/Salikhov079/rent_car/api/handler"
	pb "github.com/Salikhov079/rent_car/genprotos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	UserConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8085"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UserConn.Close()
	us := pb.NewUserServiceClient(UserConn)
	h := handler.NewHandler(us)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8081")

	err = r.Run(":8081")
	if err != nil {
		log.Fatal("Error while Run: ", err.Error())
	}
}
