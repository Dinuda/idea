package models


//Entrepreneur  new Entrepreneur
type Entrepreneur struct{
	Firstname string
	Lastname string
	Email string
	Occupation string
	DateofBirth dateofBirth
	Description string
	Startup startup
}

type startup struct{
	Name string
	Description string
	PitchDeck string
}

//Investor new investor
type Investor struct{
	Firstname string
	Lastname string
	Email string
	Occupation string
	DateofBirth dateofBirth
	Description string
	Company string
}



type dateofBirth struct{
	Date int
	Month int
	Year int 
}