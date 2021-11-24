package player

import (
	"fmt"
	"game-stats-api/pkg/datetime"
)

type Player struct {
	Id    *int                 `db:"player_id" json:"id"`
	Name  *string              `db:"player_name" json:"name"`
	Stats []*datetime.DateTime `json:"stats"`
}

func New(name string) *Player {
	g := &Player{Name: &name}
	return g
}

func (p *Player) GetId() int {
	return *p.Id
}

func (p *Player) GetName() string {
	return *p.Name
}

func (p *Player) GetStats() []*datetime.DateTime {
	return p.Stats
}

func (p *Player) AddStat(date datetime.DateTime) error {
	for _, stat := range p.Stats {
		if stat.GetDate().Equal(date.GetDate()) {
			return fmt.Errorf("stat with date '%v' already exist", date.GetDate())
		}
	}
	p.Stats = append(p.Stats, &date)
	return nil
}
