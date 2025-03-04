package database

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/naveeshkumar24/internals/models"
)

type FacultyQuery struct {
	db   *sql.DB
	Time *time.Location
}

func NewFacultyQuery(db *sql.DB) *FacultyQuery {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Fatalf("Failed to load time zone: %v", err)
	}
	return &FacultyQuery{
		db:   db,
		Time: loc,
	}
}
func (f *FacultyQuery) CreateFacultyTable() error {
	tx, err := f.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	queries := []string{
		`CREATE TABLE IF NOT EXISTS facultyregister (
				id SERIAL ,
				facultyuuid VARCHAR(255) PRIMARY KEY,
				facultyid VARCHAR(255) UNIQUE NOT NULL,
				name VARCHAR(255) NOT NULL,
				email VARCHAR(255) NOT NULL,
				password VARCHAR(255) NOT NULL,
				department VARCHAR(255) NOT NULL,
				phone VARCHAR(255) NOT NULL,
				designation VARCHAR(255) NOT NULL,
				qualification VARCHAR(255) NOT NULL,
				experience VARCHAR(255) NOT NULL
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
	log.Println("Faculty Table Created")
	return nil
}

func (f *FacultyQuery) SubmitFacultyRegisterForm(facultyRegister models.FacultyRegister) error {
	NewFacultyId := uuid.New().String()
	facultyRegister.FacultyId = NewFacultyId
	_, err := f.db.Exec("INSERT INTO facultyregister (facultyuuid,facultyid,name,email,password,department,phone,designation,qualification,experience) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)",
		facultyRegister.FacultyUuId, facultyRegister.FacultyId, facultyRegister.Name, facultyRegister.Email, facultyRegister.Password, facultyRegister.Department, facultyRegister.Phone, facultyRegister.Designation, facultyRegister.Qualification, facultyRegister.Experience)
	if err != nil {
		log.Printf("Failed to insert faculty register data: %v", err)
		return err
	}
	log.Println("Faculty Register Data Inserted Successfully")
	return nil
}
func (f *FacultyQuery) SubmitFacultyLoginForm(email, password string) (bool, error) {
	var storedPassword string
	err := f.db.QueryRow("SELECT password FROM facultyregister WHERE email = $1", email).Scan(&storedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("Login failed: Email not found")
			return false, errors.New("Invalid email or password")
		}
		log.Printf("Failed to query database: %v", err)
		return false, err
	}
	if storedPassword != password {
		log.Println("Login failed: Incorrect password")
		return false, errors.New("Invalid email or password")
	}
	log.Println("Login successful", email)
	return true, nil
}
