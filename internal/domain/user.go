package domain

import (
	"github.com/google/uuid"
	"time"
)

const (
	ADMIN  = "ADMIN"
	WORKER = "WORKER"
)

type User struct {
	Id          uuid.UUID
	Username    string
	Role        string
	DutyAmount  int
	FullName    string
	Course      int
	Stimulation []struct {
		Id        uuid.UUID
		Rewards   int
		Sanctions int
		Info      string
		DateTime  time.Time
	}
}
