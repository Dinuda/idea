package models

//Student  new Student
type Student struct{
	ID string
	Firstname string
	Lastname string
	PhoneNo int
	Address string
	Email string
	DateofBirth date
	Description string
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
	DateofBirth date
	Linkedin string
	Description string
	Company string
	Ideas []int
}

//Idea new idea
type Idea struct{
	ID int 
	Title string
	Description string
	CreatedDay date
	ClosedDay date
	TeamID  int // Team id is used to get the info
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

type date struct{
	Date int
	Month int
	Year int 
}