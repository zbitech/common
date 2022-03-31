package entity

import "time"

type Events struct {
	Name      string
	UserId    string
	Timestamp time.Time
}
