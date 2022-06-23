package router

import (
	"final-project/handler"
	"final-project/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func UserAuthPath(r *mux.Router, ah *handler.UserHandler) {
	r.HandleFunc("/user", ah.Register).Methods("POST")
	r.HandleFunc("/user/login", ah.Login).Methods("POST")
	r.Handle("/user", middleware.Authentication(http.HandlerFunc(ah.Update))).Methods("PUT")
	r.Handle("/user/profile", middleware.Authentication(http.HandlerFunc(ah.GetById))).Methods("GET")
}
