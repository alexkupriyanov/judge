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
	api.HandleFunc("/games", Controllers.GameTypeList).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/game/{Id:[0-9]+}", Controllers.GetCurrentGameType).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/team/{Id:[0-9]+}", Controllers.GetTeam).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/player/{Id:[0-9]+}", Controllers.GetPlayer).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/match/{Id:[0-9]+}", Controllers.GetMatch).Methods(http.MethodGet, http.MethodOptions)
	return
}

func initInternal(api *mux.Router) {
	api.HandleFunc("/game/create", Controllers.CreateGameType).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/game/create", Controllers.CreateGameTypePost).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/game/{Id:[0-9]+}/event/create", Controllers.CreateEventType).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/game/{Id:[0-9]+}/event/create", Controllers.CreateEventTypePost).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/team/create/{Id:[0-9]+}", Controllers.CreateTeam).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/team/create/{Id:[0-9]+}", Controllers.CreateTeamPost).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/player/create/{Id:[0-9]+}", Controllers.CreatePlayer).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/player/create/{Id:[0-9]+}", Controllers.CreatePlayerPost).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/match/create/{Id:[0-9]+}", Controllers.CreateMatch).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/match/create/{Id:[0-9]+}", Controllers.CreateMatchPost).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/match/{Id:[0-9]+}/event", Controllers.CreateEvent).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/match/{Id:[0-9]+}/event", Controllers.CreateEventPost).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/match/{MatchId:[0-9]+}/event/{EventId:[0-9]}", Controllers.CreateCurrentEvent).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/match/{MatchId:[0-9]+}/event/{EventId:[0-9]}", Controllers.CreateCurrentEventPost).Methods(http.MethodPost, http.MethodOptions)
	return
}
