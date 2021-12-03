package server

import (
	"encoding/json"
	"game-stats-api/pkg/database"
	"game-stats-api/pkg/datetime"
	"game-stats-api/pkg/game"
	"game-stats-api/pkg/player"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

type SpaHandler struct {
	staticPath string
	indexPath  string
}

func NewSpaHandler(staticPath string, indexPath string) SpaHandler {
	s := SpaHandler{staticPath, indexPath}
	return s
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (h SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
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
