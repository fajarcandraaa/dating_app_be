package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/fajarcandraaa/dating_app_be/helpers"
	"github.com/fajarcandraaa/dating_app_be/internal/presentation/subscription"
	"github.com/fajarcandraaa/dating_app_be/internal/usecase/account"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	acountUsecase account.AccountSubscribtionUsecaseContract
}

func NewAccountHandler(acountUsecase account.AccountSubscribtionUsecaseContract) *AccountHandler {
	return &AccountHandler{acountUsecase: acountUsecase}
}

func (h *AccountHandler) BuyPackage(w http.ResponseWriter, r *http.Request) {
	var (
		payload   subscription.UpdatePackageRequest
		userId    = mux.Vars(r)["userId"]
		responder = helpers.NewHTTPResponse("updatePackage")
		ctx       = context.Background()
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusBadRequest, err.Error())
		return
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	payload.UserId = userId

	err = h.acountUsecase.Subscribe(ctx, payload)
	if err != nil {
		responder.FieldErrors(w, err, http.StatusNotFound, err.Error())
		return
	}

	responder.SuccessWithoutData(w, http.StatusOK, "buy packages success")
}
