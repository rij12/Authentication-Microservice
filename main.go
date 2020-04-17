package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rij12/Authentication-Microservice/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/rij12/Authentication-Microservice/controllers"
	"github.com/rij12/Authentication-Microservice/repository"
)

func main() {

	// Set a Gobal Var inside Repo Package
	db := repository.Database{}
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	mongoUrl := os.Getenv("MONGO_HOST")
	mongoPort := os.Getenv("MONGO_PORT")
	port, _ := strconv.Atoi(mongoPort)
	connection := db.ConnectDB(username, password, mongoUrl, port)
	defer connection.Disconnect(context.Background())

	// Routing
	router := mux.NewRouter()
	controller := controllers.UserController{}

	registerHandlers(router, &controller)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("GO API Running!")

	log.Fatal(srv.ListenAndServe())

}

func registerHandlers(router *mux.Router, controller *controllers.UserController) {
	router.HandleFunc("/api/login", controller.LoginController).Methods("POST")
	router.HandleFunc("/api/register", controller.RegisterController).Methods("POST")
	router.HandleFunc("/api/protected", utils.TokenVerifyMiddleWare(controller.ProtectedEndpointTest)).Methods("GET")
	router.HandleFunc("/api/user", controller.GetUserByEmailController).Methods("GET")
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")
}
