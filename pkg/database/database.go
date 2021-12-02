package database

//go:generate statik -src=./ -include=*.sql

import (
	"database/sql"
	"fmt"
	"game-stats-api/pkg/datetime"
	"game-stats-api/pkg/game"
	"game-stats-api/pkg/player"
	"io/ioutil"
	"log"

	"github.com/jackskj/carta"
	"github.com/jmoiron/sqlx"
	"github.com/rakyll/statik/fs"

	_ "game-stats-api/pkg/database/statik"

	_ "github.com/lib/pq"
)

var db *sqlx.DB

const (
	allGamesQuery = `
select
	games.id           as  game_id,
	games.name         as  game_name,
	games.description  as  game_description,
	PG.player_id       as  player_id,
	P.name             as  player_name,
	S.date             as  player_stat
from games
	left outer join players_games PG on games.id = PG.game_id
	left outer join players P on PG.player_id = P.id
	left outer join stats S on PG.player_id = S.player_id and PG.game_id = S.game_id
`
	deleteOrphanPlayers = `
delete from players where id in (  
	select id from players
	left join players_games PG ON PG.player_id = players.id
	where PG.game_id is null
)
`
)

func Initialize(databasePath string) error {
	var err error
	connStr := "user=g603274 password=postgres dbname=postgres sslmode=require"
	db, err = sqlx.Open("postgres", connStr)
	if err != nil {
		return err
	}
	err = createTables()
	if err != nil {
		return err
	}
	return db.Ping()
}

func createTables() error {
	statikFS, err := fs.New()
	if err != nil {
		return err
	}
	r, err := statikFS.Open("/schema.sql")
	if err != nil {
		return err
	}
	defer r.Close()
	schema, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schema))
	return err
}

func getPlayerId(name string) (int64, error) {
	var id int64
	query, _ := db.Prepare("SELECT id FROM players WHERE name = ($1)")
	defer query.Close()
	err := query.QueryRow(name).Scan(&id)
	if err != nil {
		log.Println(err)
	}
	return id, err
}

func AddGamePlayerStat(gameId int, playerId int, date datetime.DateTime) error {
	tx := db.MustBegin()
	_, err := tx.Exec("INSERT INTO stats(player_id, game_id, date) VALUES ($1, $2, $3)", playerId, gameId, date.GetDate())
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func createGamePlayerTx(tx *sqlx.Tx, gameId int, player *player.Player) (int, error) {
	var playerId int64
	row := tx.QueryRow("INSERT INTO players(name) VALUES ($1) ON CONFLICT DO NOTHING RETURNING id", player.GetName())
	err := row.Scan(&playerId)
	if err != nil && err == sql.ErrNoRows {
		playerId, err = getPlayerId(player.GetName())
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}
	_, err = tx.Exec("INSERT INTO players_games(player_id, game_id) VALUES ($1, $2)", playerId, gameId)
	if err != nil {
		return 0, err
	}
	for _, date := range player.GetStats() {
		_, err = tx.Exec("INSERT INTO stats(player_id, game_id, date) VALUES ($1, $2, $3)", playerId, gameId, date.GetDate())
		if err != nil {
			return 0, err
		}
	}
	return int(playerId), err
}

func CreateGamePlayer(gameId int, player *player.Player) (int, error) {
	tx := db.MustBegin()
	id, err := createGamePlayerTx(tx, gameId, player)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	return id, err
}

func CreateGame(g *game.Game) (int, error) {
	var gameId int64
	tx := db.MustBegin()
	row := tx.QueryRow("INSERT INTO games(name, description) VALUES ($1, $2) RETURNING id", g.GetName(), g.GetDescription())
	err := row.Scan(&gameId)
	if err != nil {
		goto Err
	}
	for _, player := range g.GetPlayers() {
		_, err := createGamePlayerTx(tx, int(gameId), player)
		if err != nil {
			goto Err
		}
	}
	err = tx.Commit()
	return int(gameId), err
Err:
	log.Println(err)
	tx.Rollback()
	return 0, err
}

func DeleteGamePlayerLastStat(gameId int, playerId int) error {
	const query = `
delete from stats where date in (
select date from stats as S
where S.game_id = ($1) and S.player_id = ($2)
order by date desc limit 1
)
`
	tx := db.MustBegin()
	tx.Exec(query, gameId, playerId)
	return tx.Commit()
}

func DeleteGamePlayer(gameId int, playerId int) error {
	game, err := GetGame(gameId)
	if err != nil {
		return err
	}
	for _, player := range game.GetPlayers() {
		if player.GetId() == playerId {
			tx := db.MustBegin()
			tx.Exec("DELETE FROM players_games WHERE game_id = ($1) and player_id = ($2)", gameId, playerId)
			tx.Exec("DELETE FROM stats WHERE game_id = ($1) and player_id = ($2)", gameId, playerId)
			tx.Exec(deleteOrphanPlayers)
			err = tx.Commit()
			return err
		}
	}
	return fmt.Errorf("player with id '%v' not found", playerId)
}

func DeleteGame(id int) error {
	_, err := GetGame(id)
	if err != nil {
		return err
	}
	tx := db.MustBegin()
	tx.Exec("DELETE FROM games WHERE id = ($1)", id)
	tx.Exec("DELETE FROM players_games WHERE game_id = ($1)", id)
	tx.Exec("DELETE FROM stats WHERE game_id = ($1)", id)
	tx.Exec(deleteOrphanPlayers)
	err = tx.Commit()
	return err
}

func GetAllGames() []game.Game {
	rows, err := db.Query(allGamesQuery)
	if err != nil {
		log.Println(err)
	}

	allGames := []game.Game{}
	carta.Map(rows, &allGames)
	return allGames
}

func GetGame(id int) (game.Game, error) {
	var err error
	rows, err := db.Query(allGamesQuery+"WHERE games.id = ($1)", id)
	if err != nil {
		log.Println(err)
	}
	game := game.Game{}
	carta.Map(rows, &game)

	if game.GetId() == 0 {
		err = fmt.Errorf("game with id '%v' not found", id)
	}
	return game, err
}
