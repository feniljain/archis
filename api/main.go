package main

import (
	"archis/api/handlers"
	"archis/pkg/user"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	sqlx "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	mR := mux.NewRouter()
	authR := mR.PathPrefix("/api/user").Subrouter()

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

	handlers.MakeUserHandler(authR, userSvc)

	srv := http.Server{
		Handler:      mR,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening on port 8080")
	log.Fatal(srv.ListenAndServe())
}
