package main

import (
	"github.com/gorilla/mux"
	"judge/Controllers"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(Logger)
	external := router.PathPrefix("/").Subrouter()
	initExternal(external)
	internal := external.PathPrefix("/").Subrouter()
	initInternal(internal)
	internal.Use(Authentication)
	return router
}

func initExternal(api *mux.Router) {
	api.HandleFunc("/", Controllers.Home).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/register", Controllers.Register).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/register", Controllers.RegisterPost).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/login", Controllers.Login).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/login", Controllers.LoginPost).Methods(http.MethodPost, http.MethodOptions)
	//api.HandleFunc("/queue", controllers.Queue).Methods(http.MethodPost, http.MethodOptions)
	return
}

func initInternal(api *mux.Router) {
	api.HandleFunc("/games", Controllers.GameTypeList).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/game/create", Controllers.CreateGameType).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/game/create", Controllers.CreateGameTypePost).Methods(http.MethodPost, http.MethodOptions)
	//api.HandleFunc("/logout", controllers.Logout).Methods(http.MethodPost, http.MethodOptions)
	//api.HandleFunc("/token", controllers.CheckToken).Methods(http.MethodPost, http.MethodOptions)
	//user := api.PathPrefix("/user").Subrouter()
	//initUser(user)
	//sites := api.PathPrefix("/site").Subrouter()
	//initSite(sites)
	//	api.HandleFunc("/test", controllers.GetUserInfo).Methods(http.MethodPost, http.MethodOptions)
	return
}