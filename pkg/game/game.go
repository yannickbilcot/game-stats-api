package game

import (
	"fmt"
	"game-stats-api/pkg/player"
)

type Game struct {
	Id          int              `db:"game_id" json:"id"`
	Name        string           `db:"game_name" json:"name"`
	Description string           `db:"game_description" json:"description"`
	Players     []*player.Player `json:"players,omitempty"`
}

func New(name string) *Game {
	g := Game{Name: name}
	return &g
}

func (g *Game) GetName() string {
	return g.Name
}

func (g *Game) GetId() int {
	return g.Id
}

func (g *Game) GetDescription() string {
	return g.Description
}

func (g *Game) SetDescription(desc string) {
	g.Description = desc
}

func (g *Game) AddPlayer(name string) (*player.Player, error) {
	for _, player := range g.GetPlayers() {
		if player.GetName() == name {
			return nil, fmt.Errorf("player with name '%v' already exist", name)
		}
	}
	p := player.New(name)
	g.Players = append(g.Players, p)
	return p, nil
}

func (g *Game) DeletePlayer(name string) error {
	for i, player := range g.GetPlayers() {
		if player.GetName() == name {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("player with name '%v' not found", name)
}

func (g *Game) GetPlayers() []*player.Player {
	return g.Players
}
func (g *Game) GetPlayer(id int) (*player.Player, error) {
	for _, player := range g.GetPlayers() {
		if player.GetId() == id {
			return player, nil
		}
	}
	return nil, fmt.Errorf("player with id '%v' not found", id)
}
