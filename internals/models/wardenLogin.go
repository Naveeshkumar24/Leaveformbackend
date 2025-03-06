package models

type WardenLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type WardenRegister struct {
	ID            int    `json:"id"`
	WardenId      string `json:"wardenid"`
	WardenUuId    string `json:"wardenuuid"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Phone         string `json:"phone"`
	Designation   string `json:"designation"`
	Qualification string `json:"qualification"`
}
type WardenRegisterInterface interface {
	SubmitWardenRegiterForm(WardenRegister) error
	SubmitWardenLoginForm(WardenLogin) error
}
