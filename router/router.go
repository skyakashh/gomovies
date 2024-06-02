package router

import (
	"github.com/gorilla/mux"
	"github.com/skyakashh/mongo/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/create", controller.CreateMovie).Methods("PUT", "POST")
	r.HandleFunc("/delete/{id}", controller.DeleteAMovie).Methods("DELETE", "POST")
	r.HandleFunc("/deleteall", controller.DeleteAllMovies).Methods("DELETE", "POST")
	r.HandleFunc("/watched/{id}", controller.MarkAsWatched).Methods("POST")

	return r
}
