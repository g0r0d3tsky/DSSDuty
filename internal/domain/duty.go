package domain

import "time"

type Duty struct {
	Id     int
	Date   time.Time
	UserId int
}
