package handler

import (
	"encoding/json"
	"github.com/vaidehiadhi/threeLayerArc/models"
	"github.com/vaidehiadhi/threeLayerArc/service"
	"net/http"

	"github.com/gorilla/mux"
)

type userHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	user, err := h.service.GetUser(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
	}
}

func (h *userHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.AddUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "error encoding", http.StatusInternalServerError)
	}
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid body response", http.StatusBadRequest)
		return
	}

	if err := user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateUser(name, &user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	if err := h.service.DeleteUser(name); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
