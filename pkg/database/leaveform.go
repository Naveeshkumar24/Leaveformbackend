package database

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/naveeshkumar24/internals/models"
)

type LeaveQuery struct {
	db   *sql.DB
	Time *time.Location
}

func NewLeaveQuery(db *sql.DB) *LeaveQuery {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Fatalf("Failed to load time zone: %v", err)
	}
	return &LeaveQuery{
		db:   db,
		Time: loc,
	}
}

func (lq *LeaveQuery) CreateLeaveFormTable() error {
	tx, err := lq.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
	CREATE TABLE IF NOT EXISTS leaveforms (
		id SERIAL PRIMARY KEY,
		leave_uuid VARCHAR(255)  NOT NULL,
		student_usn VARCHAR(50) UNIQUE NOT NULL,
		reason TEXT NOT NULL,
		start_date VARCHAR(50) NOT NULL,
		end_date VARCHAR(50) NOT NULL,
		status VARCHAR(50) NOT NULL,
		faculty_remark TEXT
	)
	`

	if _, err := tx.Exec(query); err != nil {
		log.Printf("Failed to execute query: %s", query)
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	log.Println("LeaveForm Table Created")
	return nil
}

func (lq *LeaveQuery) SubmitLeaveForm(form models.LeaveForm) error {
	if form.Status == "" || form.FacultyRemark == "" {
		form.Status = "Pending"
		form.FacultyRemark = "Pending"
	}
	NewLeaveId := uuid.New().String()
	form.LeaveUUID = NewLeaveId
	_, err := lq.db.Exec(`
		INSERT INTO leaveforms (leave_uuid,student_usn, reason, start_date, end_date, status, faculty_remark)
		VALUES ($1, $2, $3, $4, $5, $6,$7)
	`, form.LeaveUUID, form.StudentUSN, form.Reason, form.StartDate, form.EndDate, form.Status, form.FacultyRemark)
	if err != nil {
		log.Printf("Failed to insert leave form data: %v", err)
		return err
	}
	log.Println("Leave form submitted successfully for student:", form.StudentUSN)
	return nil
}

func (lq *LeaveQuery) UpdateLeaveStatus(StudentUSN string, status string, remark string) error {
	result, err := lq.db.Exec(`
		UPDATE leaveforms SET status = $1, faculty_remark = $2 WHERE  student_usn= $3
	`, status, remark, StudentUSN)
	if err != nil {
		log.Printf("Failed to update leave status: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no leave form found with the provided ID")
	}

	log.Printf("Leave form (ID: %d) updated to status: %s", StudentUSN, status)
	return nil
}

func (lq *LeaveQuery) GetSanctionedLeaves() ([]models.LeaveForm, error) {
	rows, err := lq.db.Query(`
		SELECT id, student_usn, reason, start_date, end_date, status, faculty_remark
		FROM leaveforms
		WHERE status = 'Approved'
	`)
	if err != nil {
		log.Printf("Failed to query sanctioned leaves: %v", err)
		return nil, err
	}
	defer rows.Close()

	var leaves []models.LeaveForm
	for rows.Next() {
		var lf models.LeaveForm
		if err := rows.Scan(&lf.ID, &lf.StudentUSN, &lf.Reason, &lf.StartDate, &lf.EndDate, &lf.Status, &lf.FacultyRemark); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		leaves = append(leaves, lf)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	log.Printf("Retrieved %d sanctioned leave forms", len(leaves))
	return leaves, nil
}

func (lq *LeaveQuery) GetPendingLeaves() ([]models.LeaveForm, error) {
	rows, err := lq.db.Query(`
		SELECT id, student_usn, reason, start_date, end_date, status, faculty_remark
		FROM leaveforms
		WHERE status = 'Pending'
	`)
	if err != nil {
		log.Printf("Failed to query pending leaves: %v", err)
		return nil, err
	}
	defer rows.Close()

	var leaves []models.LeaveForm
	for rows.Next() {
		var lf models.LeaveForm
		if err := rows.Scan(&lf.ID, &lf.StudentUSN, &lf.Reason, &lf.StartDate, &lf.EndDate, &lf.Status, &lf.FacultyRemark); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		leaves = append(leaves, lf)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	log.Printf("Retrieved %d pending leave forms", len(leaves))
	return leaves, nil
}
