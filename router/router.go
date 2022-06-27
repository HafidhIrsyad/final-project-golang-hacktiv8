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
	r.Handle("/user", middleware.Authentication(http.HandlerFunc(ah.Delete))).Methods("DELETE")
}

func PhotoPath(r *mux.Router, ph *handler.PhotoHandler) {
	r.Handle("/photo", middleware.Authentication(http.HandlerFunc(ph.CreatePhoto))).Methods("POST")
	r.Handle("/photos", middleware.Authentication(http.HandlerFunc(ph.GetPhotos))).Methods("GET")
	r.Handle("/photo/{id}", middleware.Authentication(http.HandlerFunc(ph.GetPhotoById))).Methods("GET")
	r.Handle("/photo/{id}", middleware.Authentication(http.HandlerFunc(ph.UpdatePhoto))).Methods("PUT")
	r.Handle("/photo/{id}", middleware.Authentication(http.HandlerFunc(ph.DeletePhoto))).Methods("DELETE")
}

func CommentPath(r *mux.Router, ch *handler.CommentHandler) {
	r.Handle("/comment", middleware.Authentication(http.HandlerFunc(ch.CreateComment))).Methods("POST")
	r.Handle("/comments", middleware.Authentication(http.HandlerFunc(ch.GetAllComment))).Methods("GET")
	r.Handle("/comment/{id}", middleware.Authentication(http.HandlerFunc(ch.GetCommentById))).Methods("GET")
	r.Handle("/comment/{id}", middleware.Authentication(http.HandlerFunc(ch.UpdateComment))).Methods("PUT")
	r.Handle("/comment/{id}", middleware.Authentication(http.HandlerFunc(ch.DeleteComment))).Methods("DELETE")
}

func SocialMediaPath(r *mux.Router, sh *handler.SocialMediaHandler) {
	r.Handle("/social-media", middleware.Authentication(http.HandlerFunc(sh.CreateSocialMedia))).Methods("POST")
	r.Handle("/social-medias", middleware.Authentication(http.HandlerFunc(sh.GetAllSocialMedia))).Methods("GET")
	r.Handle("/social-media/{id}", middleware.Authentication(http.HandlerFunc(sh.GetAllSocialMedia))).Methods("GET")
	r.Handle("/social-media/{id}", middleware.Authentication(http.HandlerFunc(sh.UpdateSocialMedia))).Methods("PUT")
	r.Handle("/social-media/{id}", middleware.Authentication(http.HandlerFunc(sh.DeleteSocialMedia))).Methods("DELETE")
}
