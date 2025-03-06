package models

type StudentLogin struct {
	USN      string `json:"usn"`
	Password string `json:"password"`
}

type Studentregister struct {
	ID              int    `json:"id"`
	UserId          string `json:"userid"`
	UserName        string `json:"name"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
	Email           string `json:"email"`
	PhoneNumber     int64  ` json:"phonenumber"`
	Section         string `json:"section"`
	Sem             string `json:"sem"`
	USN             string `json:"usn"`
}

type StudentregisterInterface interface {
	SubmitStudentRegisterForm(Studentregister) error
	Submitstudentloginform(StudentLogin) error
}
