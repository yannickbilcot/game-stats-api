package main

import (
	"game-stats-api/pkg/database"
	"game-stats-api/pkg/server"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	err := database.Initialize("data.db")
	if err != nil {
		log.Fatalln(err)
	}

	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/api/v1/games/", server.CreateGameHandler).Methods("POST")
	router.HandleFunc("/api/v1/games/", server.GetAllGamesHandler).Methods("GET")
	router.HandleFunc("/api/v1/games/{id:[0-9]+}/", server.GetGameHandler).Methods("GET")
	router.HandleFunc("/api/v1/games/{id:[0-9]+}/", server.DeleteGameHandler).Methods("DELETE")
	router.HandleFunc("/api/v1/games/{id:[0-9]+}/players/", server.CreateGamePlayerHandler).Methods("POST")
	router.HandleFunc("/api/v1/games/{gid:[0-9]+}/players/{pid:[0-9]+}/", server.DeleteGamePlayerHandler).Methods("DELETE")
	router.HandleFunc("/api/v1/games/{gid:[0-9]+}/players/{pid:[0-9]+}/stats/", server.AddGamePlayerStatHandler).Methods("POST")
	router.HandleFunc("/api/v1/games/{gid:[0-9]+}/players/{pid:[0-9]+}/laststat/", server.DeleteGamePlayerLastStatHandler).Methods("DELETE")
	router.Use(middleware.Logger)
	router.Use(middleware.AllowContentEncoding("application/json"))
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "content-type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"})
	log.Println("Listening...")

	srv := &http.Server{
		Handler:      handlers.CORS(allowedOrigins, allowedHeaders, allowedMethods)(router),
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
