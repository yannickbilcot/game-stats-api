package main

import (
	"game-stats-api/pkg/database"
	"game-stats-api/pkg/server"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	godotenv.Load()

	address := os.Getenv("ADDRESS")
	if address == "" {
		address = "127.0.0.1"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	corsEnable := true
	if os.Getenv("CORS") == "" || os.Getenv("CORS") == "0" {
		corsEnable = false
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Println("database URL not defined")
	}

	err := database.Initialize(dbURL)
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

	spa := server.NewSpaHandler("ui/dist/spa", "index.html")
	router.PathPrefix("/").Handler(spa)

	var handler http.Handler
	if corsEnable {
		c := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedHeaders: []string{"X-Requested-With", "content-type"},
			AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		})
		handler = c.Handler(router)
	} else {
		handler = router
	}

	log.Printf("Listening on http://%v:%v", address, port)

	srv := &http.Server{
		Handler:      handler,
		Addr:         address + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
