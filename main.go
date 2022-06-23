package main

import (
	"final-project/config"
	"final-project/handler"
	"final-project/middleware"
	"final-project/repository"
	"final-project/router"
	"final-project/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	configs, errConf := config.Config()
	if errConf != nil {
		return
	}

	userRepo := repository.NewUserRepository(configs)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	router.UserAuthPath(r, userHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:4000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Listening on port -> 127.0.0.1:4000")

	log.Fatal(srv.ListenAndServe())

}
