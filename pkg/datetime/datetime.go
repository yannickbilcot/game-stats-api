package datetime

import "time"

type DateTime struct {
	Date *time.Time `db:"player_stat" json:"date,omitempty" example:"2021-08-12T16:28:56Z"`
}

func New(date time.Time) DateTime {
	d := DateTime{&date}
	return d
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	if t.Date == nil {
		return []byte("null"), nil
	}
	if t.Date.IsZero() {
		return []byte("null"), nil
	} else {
		return t.Date.MarshalJSON()
	}
}

func (d *DateTime) GetDate() time.Time {
	if d.Date == nil {
		return time.Time{}
	}
	return *d.Date
}
