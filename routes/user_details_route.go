package routes

import (
	"sport-venue-management/controllers"

	"github.com/gorilla/mux"
)

var UserDetailRoutes = func(router *mux.Router) {
	router.HandleFunc("/register/user", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/login/user", controllers.LoginUser).Methods("POST")
}
