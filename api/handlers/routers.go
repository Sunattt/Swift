package handlers

import (
	"github.com/gorilla/mux"
)

//описывает роуты
func InitRouter(h *Handler) *mux.Router{

	router := mux.NewRouter()
	router.HandleFunc("/sign-in", h.LoginSignIn).Methods("POST")
	router.HandleFunc("/sign-up", h.Registration).Methods("POST")

	router.HandleFunc("/hello", h.Hello).Methods("GET")
	authorized := router.NewRoute().Subrouter()
	authorized.Use(h.Authentication)
	authorized.Use(h.UserRole)	

	profile := router.PathPrefix("/profile").Subrouter()
	profile.Use(h.Authentication)
	profile.HandleFunc("/profile/{id}", h.GetProfileUser).Methods("GET")
	profile.HandleFunc("/profile/{id}", h.UpdateProfile).Methods("PUT")
	profile.HandleFunc("/profile/", h.CreateUser).Methods("POST")
	profile.HandleFunc("/profile/{id}", h.DeleteAccount).Methods("DELETE")
	// profile.HandleFunc("/image/{id}", h.)
	

	// posts := router.PathPrefix("/reels").Subrouter()
	// posts.HandleFunc("/view", h.ViewReels).Methods("GET")
	user := router.PathPrefix("/user").Subrouter() 
	// user.HandleFunc("/upload", h.UploadFile).Methods()
	user.HandleFunc("/download", h.DownloadFile).Methods()


	// post := authorized.PathPrefix("/post").Subrouter()
	// post.HandleFunc("/", handlers.Post)
	// post.HandleFunc("/", handlers.Post)
	// post.HandleFunc("/", handlers.Post)
	// post.HandleFunc("/", handlers.Post)


	//reels := authorized.PathPrefix("/reels").Subrouter()
	// reels.HandleFunc("/", handlers.Post)
	// reels.HandleFunc("/", handlers.Post)
	// reels.HandleFunc("/", handlers.Post)
	// reels.HandleFunc("/", handlers.Post)

	//message := authorized.PathPrefix("/reels").Subrouter()
	// message.HandleFunc("/", handlers.Post)
	// message.HandleFunc("/", handlers.Post)
	// message.HandleFunc("/", handlers.Post)
	// message.HandleFunc("/", handlers.Post)


	return router 

}
