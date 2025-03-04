package repository

import (
	"database/sql"
	"log"

	"github.com/naveeshkumar24/internals/models"
	"github.com/naveeshkumar24/pkg/database"
)

type FacultyLoginRepository struct {
	db *sql.DB
}

func NewFacultyLoginRepository(db *sql.DB) *FacultyLoginRepository {
	return &FacultyLoginRepository{
		db: db,
	}
}
func (f *FacultyLoginRepository) SubmitFacultyRegisterForm(data []models.FacultyRegister) error {
	query := database.NewFacultyQuery(f.db)
	err := query.CreateFacultyTable()
	if err != nil {
		log.Printf("Failed to create faculty table: %v", err)
		return err
	}
	err = query.SubmitFacultyRegisterForm(data[0])
	if err != nil {
		log.Printf("Failed to submit faculty register form : %v", err)
		return err
	}
	return nil
}
func (f *FacultyLoginRepository) SubmitFacultyLoginForm(data []models.FacultyLogin) error {
	query := database.NewFacultyQuery(f.db)
	_, err := query.SubmitFacultyLoginForm(data[0].Email, data[0].Password)
	if err != nil {
		log.Printf("Failed to Login: %v", err)
		return err
	}
	return nil
}
