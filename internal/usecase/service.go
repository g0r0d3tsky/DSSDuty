package usecase

import (
	"github.com/g0r0d3tsky/DSSDutyBot/internal/repository"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/usecase/auth"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/usecase/duty"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/usecase/duty/impl"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/usecase/event"
	impl3 "github.com/g0r0d3tsky/DSSDutyBot/internal/usecase/event/impl"
	"github.com/g0r0d3tsky/DSSDutyBot/internal/usecase/user"
	impl2 "github.com/g0r0d3tsky/DSSDutyBot/internal/usecase/user/impl"
)

type Service struct {
	Duty  duty.DutyUseCase
	User  user.UserUseCase
	Event event.EventUseCase
	Auth  auth.AuthUseCase
}
type UC struct {
	Repo repository.ServiceRepository
}

func New(repo repository.ServiceRepository) *Service {
	duty := impl.DutyUseCase{
		Repo: repo,
	}
	user := impl2.UserUseCase{
		Repo: repo,
	}
	event := impl3.EventUseCase{
		Repo: repo,
	}
	return &Service{
		Duty:  &duty,
		User:  &user,
		Event: &event,
	}
}

//реализация всех методов, что в api/rest/service.go, по сути просто будешь дергать duty.SomeMethod()
