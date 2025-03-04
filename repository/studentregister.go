package repository

import (
	"database/sql"
	"log"

	"github.com/naveeshkumar24/internals/models"
	"github.com/naveeshkumar24/pkg/database"
)

type StudentLoginRepository struct {
	db *sql.DB
}

func NewStudentLoginRepository(db *sql.DB) *StudentLoginRepository {
	return &StudentLoginRepository{
		db: db,
	}
}
func (s *StudentLoginRepository) SubmitStudentRegisterForm(data []models.Studentregister) error {
	query := database.NewStudentQuery(s.db)
	err := query.SubmitStudentRegister(data[0])
	if err != nil {
		log.Printf("Failed to submit student register form : %v", err)
		return err
	}
	return nil
}
func (s *StudentLoginRepository) SubmitStudentLogin(data []models.StudentLogin) error {
	query := database.NewStudentQuery((s.db))
	_, err := query.SubmitStudentLogin(data[0].USN, data[0].Password)
	if err != nil {
		log.Printf("Failed to Login: %v", err)
		return err
	}
	return nil
}
