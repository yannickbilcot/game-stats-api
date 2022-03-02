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

// @Summary Get all games
// @Description Retrieve the data for all the games.
// @Tags List
// @ID get-all-games
// @Produce json
// @Success 200 {array} game.Game
// @Router /games [get]
func GetAllGamesHandler(w http.ResponseWriter, req *http.Request) {
	allGames := database.GetAllGames()
	renderJSON(w, allGames)
}

// @Summary Get a game by ID
// @Description Return the data for a single game.
// @Tags List
// @ID get-game-by-id
// @Produce json
// @Param id path int true "game ID"
// @Success 200 {object} game.Game
// @Failure 404 {string} string
// @Router /games/{id} [get]
func GetGameHandler(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	game, err := database.GetGame(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, game)
}

// @Summary Create a new game
// @Description Create a new game with a specific Name and Description.
// @Description The id parameter is not needed (automatically created).
// @Tags Create
// @ID create-game
// @Accept json
// @Param body-json-param body game.Game true "game"
// @Produce json
// @Success 200 {object} object
// @Failure 400 {object} string
// @Router /games/ [post]
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

// @Summary Delete a game by ID
// @Description Delete a game with a specific ID.
// @Tags Delete
// @ID delete-game
// @Produce json
// @Param id path int true "game ID"
// @Success 200
// @Failure 404 {string} string
// @Router /games/{id} [delete]
func DeleteGameHandler(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	err := database.DeleteGame(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

// @Summary Create a game player
// @Description Add a new player for a specific game.
// @Description The player name should be unique.
// @Description The id parameter is not needed (automatically created).
// @Tags Create
// @ID create-game-player
// @Param id path int true "game ID"
// @Accept json
// @Param body-json-param body player.Player true "player"
// @Produce json
// @Success 200 {object} object
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /games/{id}/players/ [post]
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

// @Summary Delete a game player
// @Description Delete a player for a specific game.
// @Tags Delete
// @ID delete-game-player
// @Param gid path int true "game ID"
// @Param pid path int true "player ID"
// @Accept json
// @Produce json
// @Success 200
// @Failure 404 {object} string
// @Router /games/{gid}/players/{pid} [delete]
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

// @Summary Add a game win to a player
// @Description Add a game win to a player at the chosen date and time.
// @Tags Create
// @ID add-game-player-win
// @Param gid path int true "game ID"
// @Param pid path int true "player ID"
// @Accept json
// @Param body-json-param body datetime.DateTime true "date and time"
// @Produce json
// @Success 200
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /games/{gid}/players/{pid}/stats/ [post]
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

// @Summary Delete the latest game win to a player
// @Description The most recent "game win" for a specific player will be deleted.
// @Tags Delete
// @ID delete-game-player-latest-win
// @Param gid path int true "game ID"
// @Param pid path int true "player ID"
// @Produce json
// @Success 200
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /games/{gid}/players/{pid}/laststat/ [delete]
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
