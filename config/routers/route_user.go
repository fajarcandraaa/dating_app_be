package routers

import (
	"github.com/fajarcandraaa/dating_app_be/handler"
	"github.com/fajarcandraaa/dating_app_be/internal/usecase"
	"github.com/fajarcandraaa/dating_app_be/middleware"
)

func userRouter(se *Serve, s *usecase.UsecaseContract) {
	var (
		authHandler = handler.NewUserHandler(s.UserAuthContract, s.UserInfoContract)
	)

	se.Router.HandleFunc("/login", authHandler.LoginUser).Methods("POST")
	se.Router.HandleFunc("/registration", authHandler.Registration).Methods("POST")
	se.Router.HandleFunc("/view", middleware.Authentication(authHandler.RandomProfile)).Methods("GET")
}
