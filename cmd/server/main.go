package main

import (
	"archis/interceptor"
	pb "archis/pb"
	"archis/pkg/user"
	"context"
	"flag"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()

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

	//address := fmt.Sprintf("0.0.0.0:%d", *port)
	//lis, err := net.Listen("tcp", address)
	//if err != nil {
	//	log.Fatal("Cannot start the server: ", err)
	//}

	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:8081",
		grpc.WithInsecure(),
		//grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Failure dialing the server", err)
	}

	router := runtime.NewServeMux()
	if err = pb.RegisterUserServiceHandler(context.Background(), router, conn); err != nil {
		log.Fatal("Failed to register gateway", err)
	}

	portStr := ":" + strconv.Itoa(*port)
	log.Println(portStr)
	err = http.ListenAndServe(portStr, httpGrpcRouter(grpcServer, router))
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	} else {
		log.Println("Serving server on port ", port)
	}

	//err = grpcServer.Serve(lis)
	//if err != nil {
	//	log.Fatal("Cannot start the server: ", err)
	//}
}

func httpGrpcRouter(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			log.Println("Serving through grpc")
			grpcServer.ServeHTTP(w, r)
		} else {
			log.Println("Serving through http")
			httpHandler.ServeHTTP(w, r)
		}
	})
}
