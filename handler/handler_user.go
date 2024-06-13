package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/fajarcandraaa/dating_app_be/helpers"
	"github.com/fajarcandraaa/dating_app_be/internal/presentation"
	userPresentation "github.com/fajarcandraaa/dating_app_be/internal/presentation/user"
	userUsecase "github.com/fajarcandraaa/dating_app_be/internal/usecase/user"
)

type UserHandlerContract interface {
}

type UserHandler struct {
	authUsecase userUsecase.UserAuthUsecaseContract
}

func NewUserHandler(authUsecase userUsecase.UserAuthUsecaseContract) *UserHandler {
	return &UserHandler{authUsecase}
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var (
		payloadLogin userPresentation.LoginRequest
		responder    = helpers.NewHTTPResponse("loginUser")
		ctx          = context.Background()
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &payloadLogin)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	loginUser, err := h.authUsecase.LogIn(ctx, &payloadLogin)
	if err != nil {
		responder.FieldErrors(w, err, http.StatusNotFound, err.Error())
		return
	}

	resp := presentation.Response{
		Message: "success",
		Data: map[string]interface{}{
			"token": loginUser,
		},
		Error: "",
	}

	responder.SuccessJSON(w, resp, http.StatusOK, "login success")
}
