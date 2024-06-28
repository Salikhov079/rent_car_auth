package models

type Users struct {
	UserName    string
	Email       string
	Password    string
	PhoneNumber string
	Role        string
}

type Login struct {
	UserName string
}

type Filter struct {
	Name  string
	Email string
}
