package routers

import (
	"github.com/fajarcandraaa/dating_app_be/handler"
	"github.com/fajarcandraaa/dating_app_be/internal/usecase"
)

func userRouter(se *Serve, s *usecase.UsecaseContract) {
	var (
		authHandler = handler.NewUserHandler(s.UserAuthContract)
	)

	se.Router.HandleFunc("/login", authHandler.LoginUser).Methods("POST")
	se.Router.HandleFunc("/registration", authHandler.Registration).Methods("POST")
}
