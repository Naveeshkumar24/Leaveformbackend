package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/naveeshkumar24/internals/models"

	"github.com/naveeshkumar24/repository"
)

type wardenLoginHandler struct {
	repo *repository.WardenLoginRepository
}

func NewWardenLoginHandler(repo repository.WardenLoginRepository) *wardenLoginHandler {
	return &wardenLoginHandler{
		repo: &repo,
	}
}
func (wl *wardenLoginHandler) SubmitWardenRegisterForm(w http.ResponseWriter, r *http.Request) {
	var data models.WardenRegister
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	err = wl.repo.SubmitWardenRegisterForm(data)
	if err != nil {
		http.Error(w, "Failed to submit warden register form", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Data submitted successfully"})
}
func (wl *wardenLoginHandler) SubmitWardenLoginForm(w http.ResponseWriter, r *http.Request) {
	var data models.WardenLogin
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	err = wl.repo.SubmitWardenLoginForm(data)
	if err != nil {
		http.Error(w, "Failed to submit warden login form", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Data submitted successfully"})
}
