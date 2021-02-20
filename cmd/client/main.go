package main

import (
	user "archis/client"
	"flag"
	"log"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("Dial server: %s", *serverAddress)

	cc1, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Cannot Dial server: ", err)
	}

	userClient := user.NewUserClient(cc1)
	//resp, err := userClient.CreateUser("someone@email.com", 10)
	//resp, err := userClient.UpdateUser("db83cb07-2758-4d6a-9833-6de58ef7f874", "someone@email.com", 20)
	//resp, err := userClient.GetUser("db83cb07-2758-4d6a-9833-6de58ef7f874")
	resp, err := userClient.DeleteUser("db83cb07-2758-4d6a-9833-6de58ef7f874")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
}
