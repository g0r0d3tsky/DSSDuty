package domain

import (
	"github.com/google/uuid"
	"time"
)

const (
	ADMIN  = "ADMIN"
	WORKER = "WORKER"
)

type Stimulation struct {
	Id          uuid.UUID
	Stimulation bool
	Info        string
	DateTime    time.Time
}

type User struct {
	Id          uuid.UUID
	Username    string
	Role        string
	DutyAmount  int
	FullName    string
	Course      int
	Stimulation []Stimulation
}
