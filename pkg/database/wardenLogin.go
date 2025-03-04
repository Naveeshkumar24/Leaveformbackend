package database

import (
	"database/sql"
	"errors"

	"log"
	"time"

	"github.com/google/uuid"
	"github.com/naveeshkumar24/internals/models"
)

type WardenQuery struct {
	db   *sql.DB
	Time *time.Location
}

func NewWardenQuery(db *sql.DB) *WardenQuery {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Fatalf("Failed to load time zone: %v", err)
	}
	return &WardenQuery{
		db:   db,
		Time: loc,
	}
}
func (w *WardenQuery) CreateWardenTable() error {
	tx, err := w.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	queries := []string{
		`CREATE TABLE IF NOT EXISTS wardenregister (
			id SERIAL ,
			wardenid VARCHAR(255) UNIQUE NOT NULL,
			wardenuuid VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			phone VARCHAR(255) NOT NULL,
			designation VARCHAR(255) NOT NULL,
			qualification VARCHAR(255) NOT NULL
		)`,
	}
	for _, query := range queries {
		if _, err := tx.Exec(query); err != nil {
			log.Printf("Failed to execute query: %s", query)
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	log.Println("Warden Table Created")
	return nil
}

func (w *WardenQuery) SubmitWardenRegisterForm(data models.WardenRegister) error {
	newWardenId := uuid.New().String()
	data.WardenUuId = newWardenId
	_, err := w.db.Exec(`INSERT INTO wardenregister (wardenuuid,wardenid,name,email,password,phone,designation,qualification) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`, data.WardenUuId, data.WardenId, data.Name, data.Email, data.Password, data.Phone, data.Designation, data.Qualification)
	if err != nil {
		log.Printf("Failed to submit warden register form : %v", err)
		return err
	}
	log.Println("Warden Register Form Submitted")
	return nil
}
func (w *WardenQuery) SubmitWardenLoginForm(email, password string) (bool, error) {
	var storedPassword string
	err := w.db.QueryRow(`SELECT password FROM wardenregister WHERE email = $1`, email).Scan(&storedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("No warden found with the given email")
			return false, err
		}
		log.Printf("Failed to query warden password: %v", err)
		return false, err
	}
	if storedPassword != password {
		log.Println("Invalid password")
		return false, errors.New("Invalid password")
	}
	log.Println("Login successful")
	return true, nil
}
