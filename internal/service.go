package internal

import (
	"github.com/g0r0d3tsky/DSSDutyBot/internal/usecase/duty"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/usecase/event"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/usecase/user"
)

type Service struct {
	Duty  duty.DutyUseCase
	User  user.UserUseCase
	Event event.EventUseCase
}

//реализация всех методов, что в api/rest/service.go, по сути просто будешь дергать duty.SomeMethod()
