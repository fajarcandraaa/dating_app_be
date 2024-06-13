package routers

import (
	"github.com/fajarcandraaa/dating_app_be/internal/repository"
	"github.com/fajarcandraaa/dating_app_be/internal/usecase"
)

func (se *Serve) initializeRoutes() {
	r := repository.NewRepositories(se.DB)
	s := usecase.NewUsecaseService(r)

	userRouter(se, s)
}
