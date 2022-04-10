package app

import (
	"fmt"
	"log"
	"net/http"
	"testmod/app/controller"

	"github.com/gorilla/mux"
)

type app struct {
	mc   controller.MovieController
	port string
}

type App interface {
	StartApp() error
}

func NewApp(mc controller.MovieController, port int) App {

	return &app{mc: mc, port: fmt.Sprintf("%d", port)}
}
func (a *app) StartApp() error {
	// Register the API routes

	router := mux.NewRouter()

	router.HandleFunc("/api/find/{title}", a.mc.Find).Methods("GET")
	router.HandleFunc("/api/update", a.mc.Update).Methods("POST")
	router.HandleFunc("/ping", a.mc.Ping).Methods("GET")

	log.Println("Starting the server....")
	http.ListenAndServe(":"+a.port, router)
	return nil
}
