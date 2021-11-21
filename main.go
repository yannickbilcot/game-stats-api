package main

import (
	"game-stats-api/pkg/database"
	"game-stats-api/pkg/server"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func main() {
	err := database.Initialize("data.db")
	checkErr(err)

	//now := time.Now()
	//g := game.New("Monopoly")
	//g.SetDescription("Monopoly Deal game at Sagemcom")
	//p, _ := g.AddPlayer("Yannick")
	//p.AddStat(datetime.New(now))
	//p.AddStat(datetime.New(now.Add(time.Second)))
	//p, _ = g.AddPlayer("John")
	//p.AddStat(datetime.New(now.Add(2*time.Second)))
	//p, _ = g.AddPlayer("Davy")
	//p.AddStat(datetime.New(now.Add(3*time.Second)))
	//_, err = database.CreateGame(g)
	//checkErr(err)

	//g2 := game.New("Tarot")
	//p, _ = g2.AddPlayer("Patrice")
	//g2.AddPlayer("Yannick")
	//g2.AddPlayer("John")
	//g2.AddPlayer("Davy")
	//p.AddStat(datetime.New(now.Add(4*time.Second)))
	//_, err = database.CreateGame(g2)
	//checkErr(err)

	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/games/", server.CreateGameHandler).Methods("POST")
	router.HandleFunc("/games/", server.GetAllGamesHandler).Methods("GET")
	router.HandleFunc("/games/{id:[0-9]+}/", server.GetGameHandler).Methods("GET")
	router.HandleFunc("/games/{id:[0-9]+}/", server.DeleteGameHandler).Methods("DELETE")
	router.HandleFunc("/games/{id:[0-9]+}/players/", server.CreateGamePlayerHandler).Methods("POST")
	router.HandleFunc("/games/{gid:[0-9]+}/players/{pid:[0-9]+}/", server.DeleteGamePlayerHandler).Methods("DELETE")
	router.HandleFunc("/games/{gid:[0-9]+}/players/{pid:[0-9]+}/stats/", server.AddGamePlayerStatHandler).Methods("POST")
	router.HandleFunc("/games/{gid:[0-9]+}/players/{pid:[0-9]+}/laststat/", server.DeleteGamePlayerLastStatHandler).Methods("DELETE")
	router.Use(middleware.Logger)
	router.Use(middleware.AllowContentEncoding("application/json"))
	log.Println("Listening...")

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
