package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/naveeshkumar24/internals/models"
	"github.com/naveeshkumar24/pkg/database"
)

type LeaveFormRepository struct {
	db   *sql.DB
	Time *time.Location
}

func NewLeaveFormRepository(db *sql.DB) *LeaveFormRepository {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Fatalf("Failed to load time zone: %v", err)
	}
	return &LeaveFormRepository{
		db:   db,
		Time: loc,
	}
}
func (lfr *LeaveFormRepository) SubmitLeaveForm(form []models.LeaveForm) error {
	query := database.NewLeaveQuery(lfr.db)
	err := query.CreateLeaveFormTable()
	if err != nil {
		log.Printf("Failed to create leave form table: %v", err)
		return err
	}
	err = query.SubmitLeaveForm(form[0])
	if err != nil {
		log.Printf("Failed to submit leave form: %v", err)
		return err
	}
	return nil
}
func (lfr *LeaveFormRepository) UpdateLeaveStatus(studentusn string, status string, remark string) error {
	query := database.NewLeaveQuery(lfr.db)
	err := query.UpdateLeaveStatus(studentusn, status, remark)
	if err != nil {
		log.Printf("Failed to update leave status: %v", err)
		return err
	}
	return nil
}
func (lfr *LeaveFormRepository) GetPendingLeaves() ([]models.LeaveForm, error) {
	query := database.NewLeaveQuery(lfr.db)
	leaves, err := query.GetPendingLeaves()
	if err != nil {
		log.Printf("Failed to get pending leaves: %v", err)
		return nil, err
	}
	return leaves, nil
}
func (lfr *LeaveFormRepository) GetSanctionedLeaves() ([]models.LeaveForm, error) {
	query := database.NewLeaveQuery(lfr.db)
	leaves, err := query.GetSanctionedLeaves()
	if err != nil {
		log.Printf("Failed to get sanctioned leaves: %v", err)
		return nil, err
	}
	return leaves, nil
}
