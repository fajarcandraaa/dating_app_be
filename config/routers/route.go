package routers

import (
	"fmt"
	"log"

	"github.com/fajarcandraaa/dating_app_be/internal/repository"
	"github.com/fajarcandraaa/dating_app_be/internal/usecase"
	"github.com/fajarcandraaa/dating_app_be/pkg/storage/redis"
)

func (se *Serve) initializeRoutes() {
	rds, err := redis.NewRedisConfig()
	if err != nil {
		fmt.Printf("Cannot connect to redis")
		log.Fatal("This is the error:", err)
	}

	r := repository.NewRepositories(se.DB)
	s := usecase.NewUsecaseService(r, rds)

	userRouter(se, s)
}
