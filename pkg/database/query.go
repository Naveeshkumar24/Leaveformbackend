package database

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/naveeshkumar24/internals/models"
)

type StudentQuery struct {
	db   *sql.DB
	Time *time.Location
}

func NewStudentQuery(db *sql.DB) *StudentQuery {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Fatalf("Failed to load time zone: %v", err)
	}

	return &StudentQuery{
		db:   db,
		Time: loc,
	}
}

// CreateStudentTables creates the tables for student login and registration.
func (sq *StudentQuery) CreateStudentTables() error {
	tx, err := sq.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queries := []string{
		`CREATE TABLE IF NOT EXISTS studentregister (
			id SERIAL ,
			userid VARCHAR(255) PRIMARY KEY ,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			confirmpassword VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			phonenumber BIGINT,
			section VARCHAR(50),
			sem VARCHAR(50),
			usn VARCHAR(50) UNIQUE NOT NULL
		)`,
	}

	for _, query := range queries {
		if _, err := tx.Exec(query); err != nil {
			log.Printf("Failed to execute query: %s", query)
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	log.Println("Student tables created successfully.")
	return nil
}

func (sq *StudentQuery) SubmitStudentRegister(data models.Studentregister) error {
	newUserId := uuid.New()
	data.UserId = newUserId.String()
	_, err := sq.db.Exec(`
		INSERT INTO studentregister (userid, username, password, confirmpassword, email, phonenumber, section, sem, usn)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		data.UserId, data.UserName, data.Password, data.ConfirmPassword, data.Email,
		data.PhoneNumber, data.Section, data.Sem, data.USN,
	)
	if err != nil {
		log.Printf("Failed to insert student register data: %v", err)
		return err
	}

	log.Println("Student register data submitted successfully.")
	log.Println("New ID:", data.ID)
	return nil
}

func (s *StudentQuery) SubmitStudentLogin(usn, password string) (bool, error) {
	var storedPassword string

	err := s.db.QueryRow(`
		SELECT password FROM studentregister WHERE usn = $1
	`, usn).Scan(&storedPassword)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("Login failed: USN not found")
			return false, errors.New("invalid USN or password")
		}
		log.Printf("Database error: %v", err)
		return false, err
	}

	if storedPassword != password {
		log.Println("Login failed: Incorrect password")
		return false, errors.New("invalid USN or password")
	}

	log.Println("Login successful for USN:", usn)
	return true, nil
}
