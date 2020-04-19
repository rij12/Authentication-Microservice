package main

import (
	"context"
	"encoding/json"
	"github.com/alexcesaro/log/stdlog"
	"github.com/gorilla/mux"
	"github.com/rij12/Authentication-Microservice/controllers"
	"github.com/rij12/Authentication-Microservice/repository"
	"github.com/rij12/Authentication-Microservice/service"
	"github.com/rij12/Authentication-Microservice/utils"
	"log"
	"net/http"
	"time"
)

var logger = stdlog.GetFromFlags()

func main() {

	db := repository.Database{}

	// Populate Config Struct
	configurationService := service.ConfigurationService{}
	config := configurationService.GetConfig()

	connection := db.ConnectDB(config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port)
	defer connection.Disconnect(context.Background())

	// Routing
	router := mux.NewRouter()
	controller := controllers.UserController{}

	registerHandlers(router, &controller)

	srv := &http.Server{
		Handler: router,
		Addr:    config.Server.Host + ":" + config.Server.Port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("Starting server on %s", config.Server.Host+":"+config.Server.Port)
	err := srv.ListenAndServeTLS("./crypto/cert.pem", "./crypto/key.pem")
	log.Fatal("Server failed with error: ", err)

}

func registerHandlers(router *mux.Router, controller *controllers.UserController) {
	router.HandleFunc("/api/login", controller.LoginController).Methods("POST")
	router.HandleFunc("/api/register", controller.RegisterController).Methods("POST")
	router.HandleFunc("/api/protected", utils.TokenVerifyMiddleWare(controller.ProtectedEndpointTest)).Methods("GET")
	router.HandleFunc("/api/user", utils.TokenVerifyMiddleWare(controller.GetUserByEmailController)).Methods("GET")
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")
}
