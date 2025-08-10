package service

import (
	"github.com/sthitasahu/tasker/internal/lib/job"
	"github.com/sthitasahu/tasker/internal/repository"
	"github.com/sthitasahu/tasker/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
