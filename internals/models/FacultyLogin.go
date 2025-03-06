package models

type FacultyLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type FacultyRegister struct {
	ID            int    `json:"id"`
	FacultyUuId   string `json:"facultyuuid"`
	FacultyId     string `json:"facultyid"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Department    string `json:"department"`
	Phone         string `json:"phone"`
	Designation   string `json:"designation"`
	Qualification string `json:"qualification"`
	Experience    string `json:"experience"`
}
type FacultyRegisterInterface interface {
	SubmitFacultyRegisterForm(FacultyRegister) error
	SubmitFacultyLoginForm(FacultyLogin) error
}
