package models

//User new User
type User struct {
	Username    string    
	Password    string    
	Firstname   string   
	Lastname    string   
	PhoneNo     int      
	Address     string   
	Email       string   
	DateofBirth string   
	Description string   
	Type        string    
	Investor    Investor  
	Student     Student   
}

//Student  new Student
type Student struct {
	ID         int
	UserID     int
	TeamID     int
	Profession int
	CV         string
	TeamRole   string
}

//Investor new investor
type Investor struct {
	ID       int
	UserID   int
	Linkedin string
	Company  string
	Projects []int
}

//Project new host of project
type Project struct {
	ID             int
	Title          string
	Description    string
	CreatedDate    string
	ClosedDate     string
	InvestorTeamID int
	StudentTeamIDs []int  // Team id is used to get the info
	Category       int //Agriculture, IT,
}

//StudentTeam is used to make a team by students
type StudentTeam struct {
	ID        int
	StudenIDs []int //Students currently working
}

//InvestorTeam is used make a team of investor by investors
type InvestorTeam struct {
	ID            int
	InvestorIDs   []int
	StudentTeamID int //can have multiple student teams but they'll be added to the existing team
	Project       int
}

//Profession to get the professions
type Profession struct {
	ID   int
	Name string
}

//ProjectCategory catagories of the project
type ProjectCategory struct {
	ID   int
	Name string
}

// type date struct {
// 	Date  int
// 	Month int
// 	Year  int
// }
