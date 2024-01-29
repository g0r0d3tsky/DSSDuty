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
	Id          uuid.UUID `json:"id"`
	Stimulation bool      `json:"stimulation"`
	Info        string    `json:"info"`
	DateTime    time.Time `json:"date_time"`
}

type User struct {
	Id          uuid.UUID     `json:"id"`
	Username    string        `json:"username"`
	Role        string        `json:"role"`
	DutyAmount  int           `json:"duty_amount"`
	FullName    string        `json:"full_name"`
	Course      int           `json:"course"`
	Stimulation []Stimulation `json:"stimulation"`
}
