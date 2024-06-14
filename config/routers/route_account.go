package routers

import (
	"github.com/fajarcandraaa/dating_app_be/handler"
	"github.com/fajarcandraaa/dating_app_be/internal/usecase"
	"github.com/fajarcandraaa/dating_app_be/middleware"
)

func accountRoute(se *Serve, s *usecase.UsecaseContract) {
	var (
		accountHandler = handler.NewAccountHandler(s.AccountSubscribtionContract)
	)

	se.Router.HandleFunc("/package/{userId}", middleware.Authentication(accountHandler.BuyPackage)).Methods("PATCH")

}
