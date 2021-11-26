package datetime

import "time"

type DateTime struct {
	Date *time.Time `db:"player_stat" json:"date,omitempty"`
}

func New(date time.Time) DateTime {
	d := DateTime{&date}
	return d
}

func (d *DateTime) GetDate() time.Time {
	if d.Date == nil {
		return time.Time{}
	}
	return *d.Date
}
