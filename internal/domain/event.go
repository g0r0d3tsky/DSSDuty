package domain

import "time"

type Event struct {
	Id             int
	UserId         int
	DutyId         int
	EventType      string
	EventTimeStamp time.Time
}
