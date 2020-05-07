package models

//User new User
type User struct{
	Firstname string
	Lastname string
	Email string
	Occupation string
	DateofBirth dateofBirth
}


type dateofBirth struct{
	Date int
	Month int
	Year int 
}