package domain

import "github.com/google/uuid"

const (
	ADMIN  = "ADMIN"
	WORKER = "WORKER"
)

type User struct {
	Id         uuid.UUID
	Username   string
	Role       string
	DutyAmount int
}
