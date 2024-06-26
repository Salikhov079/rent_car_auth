package models

type Users struct {
	UserName    string
	Email       string
	Password    string
	PhoneNumber string
	Role        string
}

type UsersFilter struct {
	Name  string
	Email string
}
