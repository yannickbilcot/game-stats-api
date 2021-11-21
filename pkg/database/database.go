package database

import (
	"fmt"
	"game-stats-api/pkg/datetime"
	"game-stats-api/pkg/game"
	"game-stats-api/pkg/player"
	"log"
	"os"

	"github.com/jackskj/carta"
	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

const schemaPath = "pkg/database/schema.sql"

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
	db, err = sqlx.Open("sqlite3", databasePath)
	if err != nil {
		return err
	}
	createTables()
	return db.Ping()
}

func createTables() {
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Println(err)
	}
	db.MustExec(string(schema))
}

func getPlayerId(name string) (int64, error) {
	var id int64
	query, _ := db.Prepare("SELECT id FROM players WHERE name = (?)")
	defer query.Close()
	err := query.QueryRow(name).Scan(&id)
	if err != nil {
		log.Println(err)
	}
	return id, err
}

func AddGamePlayerStat(gameId int, playerId int, date datetime.DateTime) error {
	tx := db.MustBegin()
	_, err := tx.Exec("INSERT INTO stats(player_id, game_id, date) VALUES (?, ?, ?)", playerId, gameId, date.GetDate())
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	return err
}

func createGamePlayerTx(tx *sqlx.Tx, gameId int, player *player.Player) (int, error) {
	var playerId int64
	res, err := tx.Exec("INSERT INTO players(name) VALUES (?)", player.GetName())
	if err != nil {
		if sqlite3.ErrNoExtended(err.(sqlite3.Error).ExtendedCode) == sqlite3.ErrConstraintUnique {
			playerId, err = getPlayerId(player.GetName())
			if err != nil {
				log.Println(err)
				return 0, fmt.Errorf("failed to create a new player")
			}
		}
	} else {
		playerId, _ = res.LastInsertId()
	}
	tx.MustExec("INSERT INTO players_games(player_id, game_id) VALUES (?, ?)", playerId, gameId)
	for _, date := range player.GetStats() {
		tx.MustExec("INSERT INTO stats(player_id, game_id, date) VALUES (?, ?, ?)", playerId, gameId, date.GetDate())
	}
	return int(playerId), err
}

func CreateGamePlayer(gameId int, player *player.Player) (int, error) {
	tx := db.MustBegin()
	id, err := createGamePlayerTx(tx, gameId, player)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	return id, err
}

func CreateGame(g *game.Game) (int, error) {
	tx := db.MustBegin()
	res := tx.MustExec("INSERT INTO games(name, description) VALUES (?, ?)", g.GetName(), g.GetDescription())
	gameId, _ := res.LastInsertId()
	for _, player := range g.GetPlayers() {
		_, err := createGamePlayerTx(tx, int(gameId), player)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	err := tx.Commit()
	return int(gameId), err
}


func DeleteGamePlayerLastStat(gameId int, playerId int) error {
	const query = `
delete from stats where date in (
select date from stats as S
where S.game_id = (?) and S.player_id = (?)
order by date desc limit 1
)
`
	tx := db.MustBegin()
	tx.MustExec(query, gameId, playerId)
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
			tx.MustExec("DELETE FROM players_games WHERE game_id = (?) and player_id = (?)", gameId, playerId)
			tx.MustExec("DELETE FROM stats WHERE game_id = (?) and player_id = (?)", gameId, playerId)
			tx.MustExec(deleteOrphanPlayers)
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
	tx.MustExec("DELETE FROM games WHERE id = (?)", id)
	tx.MustExec("DELETE FROM players_games WHERE game_id = (?)", id)
	tx.MustExec("DELETE FROM stats WHERE game_id = (?)", id)
	tx.MustExec(deleteOrphanPlayers)
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
	rows, err := db.Query(allGamesQuery+"WHERE games.id = (?)", id)
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
