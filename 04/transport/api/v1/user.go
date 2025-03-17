package v1

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"

	"user-management-api/service/model"
	"user-management-api/transport/util"
)

var validate = validator.New()

func getEmailFromURL(r *http.Request) string {
	email := chi.URLParam(r, "email")
	return email
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var user model.User
	if err := json.Unmarshal(b, &user); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	if err := validate.Struct(user); err != nil {
		//validace tady???
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return

	}

	if err := h.service.CreateUser(r.Context(), user); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	util.WriteResponse(w, http.StatusCreated, user)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := h.service.GetUser(r.Context(), getEmailFromURL(r))
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, user)
}

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.ListUsers(r.Context())
	util.WriteResponse(w, http.StatusOK, users)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	var user model.User
	if err := json.Unmarshal(b, &user); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	email := user.Email
	userUpdated, err := h.service.UpdateUser(r.Context(), email, user)
	if err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
	util.WriteResponse(w, http.StatusOK, userUpdated)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if err := h.service.DeleteUser(r.Context(), getEmailFromURL(r)); err != nil {
		util.WriteErrResponse(w, http.StatusBadRequest, err)
		return
	}
}
