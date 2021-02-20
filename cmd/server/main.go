package main

import (
	"archis/interceptor"
	pb "archis/pb"
	"archis/pkg/user"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("Started server on port %d", *port)

	db, err := sqlx.Connect("postgres", "user=postgres dbname=appointy password=1234")
	if err != nil {
		panic(err)
	}

	//schema := `
	//create table users (
	//	_id text primary key,
	//	email text,
	//	amount int
	//);`
	//db.MustExec(schema)

	userRepo := user.NewPostgresRepo(db)
	userSvc := user.NewUserService(userRepo)
	userInterceptor := interceptor.NewUserInterceptor(userSvc)

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userInterceptor)
	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Cannot start the server: ", err)
	}

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Cannot start the server: ", err)
	}
}
