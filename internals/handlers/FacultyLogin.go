package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/naveeshkumar24/internals/models"
	"github.com/naveeshkumar24/pkg/utils"
	"github.com/naveeshkumar24/repository"
)

type FacultyLoginHandler struct {
	repo *repository.FacultyLoginRepository
}

func NewFacultyLoginHandler(repo repository.FacultyLoginRepository) *FacultyLoginHandler {
	return &FacultyLoginHandler{
		repo: &repo,
	}
}
func (b *FacultyLoginHandler) SubmitFacultyRegisterForm(w http.ResponseWriter, r *http.Request) {
	var data models.FacultyRegister
	err := utils.Decode(r, &data)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	err = b.repo.SubmitFacultyRegisterForm(data)
	if err != nil {
		http.Error(w, "Failed to submit faculty register form", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.Encode(w, map[string]string{"message": "Data submitted successfully"})
}
func (b *FacultyLoginHandler) SubmitFacultyLoginForm(w http.ResponseWriter, r *http.Request) {
	var data models.FacultyLogin
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	err = b.repo.SubmitFacultyLoginForm(data)
	if err != nil {
		http.Error(w, "Failed to submit faculty login form", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Data submitted successfully"})
}
