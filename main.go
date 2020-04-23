package main

import (
	"encoding/json"
	"github.com/alexcesaro/log/stdlog"
	"github.com/gorilla/mux"
	"github.com/rij12/Authentication-Microservice/controllers"
	"github.com/rij12/Authentication-Microservice/repository"
	"github.com/rij12/Authentication-Microservice/service"
	"log"
	"net/http"
	"time"
)

var logger = stdlog.GetFromFlags()

func main() {

	// Dependency Injection
	// See README for more details about DI.

	// Config
	configurationService := service.ConfigurationService{}
	config := configurationService.GetConfig()

	// Repo
	mongoRepository := repository.MongoRepository{}
	mongoRepository.Init(config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port)
	//defer mongoRepository.GetDb().Disconnect(context.Background())
	userRepository := repository.UserRepositoryImpl{MongoRepository: &mongoRepository}

	// Services
	cryptoService := service.CryptoService{}
	cryptoService.Init(config)
	userService := service.UserService{ConfigurationService: configurationService,
									   CryptoService: cryptoService,
									   UserRepository: userRepository}

	// Routing and Controllers
	router := mux.NewRouter()
	controller := controllers.UserController{UserService: &userService}

	registerHandlers(router, &controller, cryptoService)

	srv := &http.Server{
		Handler:      router,
		Addr:         config.Server.Host + ":" + config.Server.Port,
		WriteTimeout: config.Server.TimeoutInSeconds * time.Second,
		ReadTimeout:  config.Server.TimeoutInSeconds * time.Second,
	}

	logger.Info("Starting server on %s", config.Server.Host+":"+config.Server.Port)
	//TODO Fix SSL!
	//err := srv.ListenAndServeTLS(config.Crypto.SSLCert, config.Crypto.SSLKey)
	err := srv.ListenAndServe()
	log.Fatal("Server failed with error: ", err)

}

func registerHandlers(router *mux.Router, controller *controllers.UserController, cryptoService service.CryptoService) {
	router.HandleFunc("/api/login", controller.LoginController).Methods("POST")
	router.HandleFunc("/api/register", controller.RegisterController).Methods("POST")
	router.HandleFunc("/api/protected", cryptoService.TokenVerifyMiddleWare(controller.ProtectedEndpointTest)).Methods("GET")
	router.HandleFunc("/api/user", cryptoService.TokenVerifyMiddleWare(controller.GetUserByEmailController)).Methods("GET")
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")
}
