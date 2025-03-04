package repository

import (
	"database/sql"
	"log"

	"github.com/naveeshkumar24/internals/models"
	"github.com/naveeshkumar24/pkg/database"
)

type WardenLoginRepository struct {
	db *sql.DB
}

func NewWardenLoginRepository(db *sql.DB) *WardenLoginRepository {
	return &WardenLoginRepository{
		db: db,
	}
}
func (w *WardenLoginRepository) SubmitWardenRegisterForm(data []models.WardenRegister) error {
	query := database.NewWardenQuery(w.db)
	err := query.CreateWardenTable()
	if err != nil {
		log.Printf("Failed to create warden table: %v", err)
		return err
	}
	err = query.SubmitWardenRegisterForm(data[0])
	if err != nil {
		log.Printf("Failed to submit warden register form : %v", err)
		return err
	}
	return nil
}

func (w *WardenLoginRepository) SubmitWardenLoginForm(data []models.WardenLogin) error {
	query := database.NewWardenQuery(w.db)
	_, err := query.SubmitWardenLoginForm(data[0].Email, data[0].Password)
	if err != nil {
		log.Printf("Failed to Login: %v", err)
		return err
	}
	return nil
}
