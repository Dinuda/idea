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
	Ideas []int
}

//Idea new idea
type Idea struct{
	ID int
	Name string
	Description string
	Team int // Team id is used to get the info
	Category string //Agriculture, IT, 
}

//Team is used to make a team
type Team struct{
	ID int
	InvestorIDs []int //can have multiple investors in one idea
	LookingRoles []string //Remaining role to be found
	AppliedStudents []Student //Students who are applied
	CurrentStudents []Student //Students currently working
}

type dateofBirth struct{
	Date int
	Month int
	Year int 
}