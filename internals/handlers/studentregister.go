package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/naveeshkumar24/internals/models"
	"github.com/naveeshkumar24/pkg/utils"
	"github.com/naveeshkumar24/repository"
)

type StudentLoginHandler struct {
	StudentLogin *repository.StudentLoginRepository
}

func NewStudentLoginHandler(studentLogin repository.StudentLoginRepository) *StudentLoginHandler {
	return &StudentLoginHandler{
		StudentLogin: &studentLogin,
	}
}
func (b *StudentLoginHandler) SubmitStudentRegisterForm(w http.ResponseWriter, r *http.Request) {
	var data models.Studentregister
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("Failed to decode student register form: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		utils.Encode(w, map[string]string{"message": "invalid request"})
		return
	}
	err = b.StudentLogin.SubmitStudentRegisterForm(data)
	if err != nil {
		log.Printf("Failed to submit student register form: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.Encode(w, map[string]string{"message": "failed to submit student register form"})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "student register form submitted"})
}
func (b *StudentLoginHandler) SubmitStudentLogin(w http.ResponseWriter, r *http.Request) {
	var data models.StudentLogin
	err := utils.Decode(r, &data)
	if err != nil {
		log.Printf("Failed to decode student login form: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		utils.Encode(w, map[string]string{"message": "invalid request"})
		return
	}
	err = b.StudentLogin.SubmitStudentLogin(data)
	if err != nil {
		log.Printf("Failed to submit student login form: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		utils.Encode(w, map[string]string{"message": "failed to submit student login form"})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Data submitted successfully"})
}
