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
		fmt.Println(errConf.Error())
		return
	}

	userRepo := repository.NewUserRepository(configs)
	commentRepo := repository.NewCommentRepository(configs)
	photoRepo := repository.NewPhotoRepo(configs)
	smRepo := repository.NewSocialMediaRepo(configs)

	userService := service.NewUserService(userRepo)
	photoService := service.NewPhotoService(photoRepo, commentRepo)
	commentService := service.NewCommentService(commentRepo)
	smService := service.NewSocialMediaService(smRepo)

	userHandler := handler.NewUserHandler(userService)
	photoHandler := handler.NewPhotoHandler(photoService)
	commentHandler := handler.NewCommentHandler(commentService)
	smHandler := handler.NewSocialMediaHandler(smService)

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	router.UserAuthPath(r, userHandler)
	router.PhotoPath(r, photoHandler)
	router.CommentPath(r, commentHandler)
	router.SocialMediaPath(r, smHandler)

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
