package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/fajarcandraaa/dating_app_be/helpers"
	"github.com/fajarcandraaa/dating_app_be/internal/presentation"
	userPresentation "github.com/fajarcandraaa/dating_app_be/internal/presentation/user"
	userUsecase "github.com/fajarcandraaa/dating_app_be/internal/usecase/user"
)

type UserHandler struct {
	authUsecase userUsecase.UserAuthUsecaseContract
	infoUsecase userUsecase.UserInfoUsecaseContract
}

func NewUserHandler(authUsecase userUsecase.UserAuthUsecaseContract, infoUsecase userUsecase.UserInfoUsecaseContract) *UserHandler {
	return &UserHandler{authUsecase, infoUsecase}
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

	responder.SuccessJSONV2(w, resp.Data, http.StatusOK, "login success")
}

func (h *UserHandler) Registration(w http.ResponseWriter, r *http.Request) {
	var (
		payloadRegistration userPresentation.RegistrationRequest
		responder           = helpers.NewHTTPResponse("registrationUser")
		ctx                 = context.Background()
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &payloadRegistration)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = h.authUsecase.SignUp(ctx, &payloadRegistration)
	if err != nil {
		responder.FieldErrors(w, err, http.StatusUnprocessableEntity, err.Error())
		return
	}

	responder.SuccessWithoutData(w, http.StatusCreated, "registration success")
}

func (h *UserHandler) RandomProfile(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("getRandomProfile")
		ctx       = context.Background()
		authToken = r.Header.Get("Authorization")
	)

	if len(strings.Split(authToken, " ")) != 2 {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, "please check your token authorization")
		return
	}
	token := strings.Split(authToken, " ")[1]
	parseToken, err := helpers.ParseToken(token)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.infoUsecase.GetRandomProfile(ctx, parseToken.UserCode)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	responder.SuccessJSONV2(w, resp, http.StatusOK, "get profile")
}
