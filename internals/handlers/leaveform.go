package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/naveeshkumar24/internals/models"
	"github.com/naveeshkumar24/repository"
)

type LeaveFormHandler struct {
	repo *repository.LeaveFormRepository
}

func NewLeaveFormHandler(repo repository.LeaveFormRepository) *LeaveFormHandler {
	return &LeaveFormHandler{
		repo: &repo,
	}
}
func (lfr *LeaveFormHandler) SubmitLeaveForm(w http.ResponseWriter, r *http.Request) {
	var data []models.LeaveForm
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Failed to decode  request body", http.StatusBadRequest)
		return
	}
	err = lfr.repo.SubmitLeaveForm(data)
	if err != nil {
		http.Error(w, "Failed to submit leave form", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Data submitted successfully"})
}
func (lfr *LeaveFormHandler) UpdateLeaveStatus(w http.ResponseWriter, r *http.Request) {
	var data []models.LeaveForm
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	err = lfr.repo.UpdateLeaveStatus(data[0].StudentUSN, data[0].Status, data[0].FacultyRemark)
	if err != nil {
		http.Error(w, "Failed to update leave status", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Data updated successfully"})
}

func (lfr *LeaveFormHandler) GetPendingLeaves(w http.ResponseWriter, r *http.Request) {
	leaves, err := lfr.repo.GetPendingLeaves()
	if err != nil {
		http.Error(w, "Failed to get pending leaves", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(leaves)
}
func (lfr *LeaveFormHandler) GetSanctionedLeaves(w http.ResponseWriter, r *http.Request) {
	leaves, err := lfr.repo.GetSanctionedLeaves()
	if err != nil {
		http.Error(w, "Failed to get approved leaves", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(leaves)
}
