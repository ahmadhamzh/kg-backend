package models

type Teacher struct {
	FirstName   string `json:"firstName" required:"true"`
	LastName    string `json:"lastName" required:"true"`
	Email       string `json:"email" required:"true"`
	Password    string `json:"password" required:"true"`
	PhoneNumber string `json:"phoneNumber" required:"true"`
}
