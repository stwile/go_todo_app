package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/stwile/go_todo_app/entity"
)

type RegisterUser struct {
	Service   RegisterUserService
	Validator *validator.Validate
}

func (ru *RegisterUser) ServeHttp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
		Role     string `json:"role" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJson(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := ru.Validator.Struct(b); err != nil {
		RespondJson(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	u, err := ru.Service.RegisterUser(ctx, b.Name, b.Password, b.Role)
	if err != nil {
		RespondJson(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	rsp := struct {
		ID entity.UserID `json:"id"`
	}{ID: u.ID}
	RespondJson(ctx, w, rsp, http.StatusOK)
}
