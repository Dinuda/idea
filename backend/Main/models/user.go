package models

//Student  new Student
type Student struct{
	ID string
	Firstname string
	Lastname string
	PhoneNo int
	StudentType string
	Address string
	Linkedin string
	Website string
	DateofBirth dateofBirth
	Description string
	Qualifications qualifications
}

type qualifications struct{
	StudentType string // University, Bachelor
	Role string
	University string
	CV string
}


//Investor new investor
type Investor struct{
	ID int
	Firstname string
	Lastname string
	Email string
	Occupation string
	Gender string
	PhoneNo int
	DateofBirth dateofBirth
	Description string
	Company string
	Ideas []idea
}

type idea struct{
	ID int
	InvestorID int
	Name string
	Description string
	TeamNo int
	Category string //Agriculture, IT, 
}

type dateofBirth struct{
	Date int
	Month int
	Year int 
}