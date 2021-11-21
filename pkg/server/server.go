package server

import (
	"encoding/json"
	"game-stats-api/pkg/database"
	"game-stats-api/pkg/datetime"
	"game-stats-api/pkg/game"
	"game-stats-api/pkg/player"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetAllGamesHandler(w http.ResponseWriter, req *http.Request) {
	allGames := database.GetAllGames()
	renderJSON(w, allGames)
}

func GetGameHandler(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	game, err := database.GetGame(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, game)
}

func CreateGameHandler(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	g := game.New("")
	err := dec.Decode(&g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := database.CreateGame(g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, map[string]int{"id": id})
}

func DeleteGameHandler(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	err := database.DeleteGame(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func CreateGamePlayerHandler(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	game, err := database.GetGame(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	var p player.Player
	err = dec.Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = game.AddPlayer(p.GetName())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err = database.CreateGamePlayer(id, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, map[string]int{"id": id})
}

func DeleteGamePlayerHandler(w http.ResponseWriter, req *http.Request) {
	gid, _ := strconv.Atoi(mux.Vars(req)["gid"])
	game, err := database.GetGame(gid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	pid, _ := strconv.Atoi(mux.Vars(req)["pid"])
	_, err = game.GetPlayer(pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	err = database.DeleteGamePlayer(gid, pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func AddGamePlayerStatHandler(w http.ResponseWriter, req *http.Request) {
	gid, _ := strconv.Atoi(mux.Vars(req)["gid"])
	game, err := database.GetGame(gid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	pid, _ := strconv.Atoi(mux.Vars(req)["pid"])
	player, err := game.GetPlayer(pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	var date datetime.DateTime
	err = dec.Decode(&date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = player.AddStat(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.AddGamePlayerStat(gid, pid, date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeleteGamePlayerLastStatHandler(w http.ResponseWriter, req *http.Request) {
	gid, _ := strconv.Atoi(mux.Vars(req)["gid"])
	game, err := database.GetGame(gid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	pid, _ := strconv.Atoi(mux.Vars(req)["pid"])
	_, err = game.GetPlayer(pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = database.DeleteGamePlayerLastStat(gid, pid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
