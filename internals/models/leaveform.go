package models

type LeaveForm struct {
	ID            int    `json:"id"`
	LeaveUUID     string `json:"leave_uuid"`
	StudentUSN    string `json:"student_usn"`
	Reason        string `json:"reason"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	Status        string `json:"status"`
	FacultyRemark string `json:"faculty_remark"`
}

type LeaveFormInterface interface {
	SubmitLeaveForm(form LeaveForm) error
	UpdateLeaveStatus(leaveID int, status string, remark string) error
	GetPendingLeaves() ([]LeaveForm, error)
	GetSanctionedLeaves() ([]LeaveForm, error)
}
