package handlers

import (
	"archis/api/views"
	"archis/pkg"
	error "archis/pkg"
	"archis/pkg/user"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func createUser(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := make(map[string]interface{})

		req := user.AuthRequest{}

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&req)
		if err != nil {
			log.Println(err)
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusBadRequest, error.ErrWrongFormat.Error(), message)
			return
		}

		defer r.Body.Close()

		err = userSvc.CreateUser(req)
		if err == pkg.ErrInvalidEmailAndPass {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusUnauthorized, err.Error(), message)
			return
		} else if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusInternalServerError, err.Error(), message)
			return
		}

		message["message"] = "Successfully created the user!"
		views.SendResponse(w, http.StatusCreated, "", message)
	}
}

func updateUser(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := make(map[string]interface{})

		req := user.UpdateRequest{}

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&req)
		if err != nil {
			log.Println(err)
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusBadRequest, error.ErrWrongFormat.Error(), message)
			return
		}

		defer r.Body.Close()

		user, err := userSvc.UpdateUser(req)
		if err == pkg.ErrInvalidEmailAndPass {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusUnauthorized, err.Error(), message)
			return
		} else if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusInternalServerError, err.Error(), message)
			return
		}

		message["message"] = "Successfully updated the user!"
		message["user"] = user
		views.SendResponse(w, http.StatusCreated, "", message)
	}
}

func getUser(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := make(map[string]interface{})

		v := r.URL.Query()
		id := v.Get("id")

		user, err := userSvc.GetUser(id)
		if err == pkg.ErrInvalidEmailAndPass {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusUnauthorized, err.Error(), message)
			return
		} else if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusInternalServerError, err.Error(), message)
			return
		}

		message["message"] = "Successfully fetched the user!"
		message["user"] = user
		views.SendResponse(w, http.StatusCreated, "", message)
	}
}

func deleteUser(userSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := make(map[string]interface{})

		v := r.URL.Query()
		id := v.Get("id")

		err := userSvc.DeleteUser(id)
		if err == pkg.ErrInvalidEmailAndPass {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusUnauthorized, err.Error(), message)
			return
		} else if err != nil {
			message["message"] = "An error occurred"
			views.SendResponse(w, http.StatusInternalServerError, err.Error(), message)
			return
		}

		message["message"] = "Successfully deleted the user!"
		views.SendResponse(w, http.StatusCreated, "", message)
	}
}

//MakeUserHandler defines the route handlers for user
func MakeUserHandler(r *mux.Router, userSvc user.Service) {
	r.HandleFunc("", (createUser(userSvc))).Methods("POST")
	r.HandleFunc("", (updateUser(userSvc))).Methods("PUT")
	r.HandleFunc("", (deleteUser(userSvc))).Queries("id", "{id}").Methods("DELETE")
	r.HandleFunc("", (getUser(userSvc))).Queries("id", "{id}").Methods("GET")
}
