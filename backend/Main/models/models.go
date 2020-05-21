package models

//User new User
type User struct {
	Username    string   `json:"username,omitempty"`
	Password    string   `json:"password,omitempty"`
	Name        string   `json:"name,omitempty"`
	PhoneNo     int      `json:"phone_no,omitempty"`
	Address     string   `json:"address,omitempty"`
	Email       string   `json:"email,omitempty"`
	Description string   `json:"description,omitempty"`
	Type        string   `json:"type,omitempty"`
	Investor    Investor `json:"investor,omitempty"`
	Student     Student  `json:"student,omitempty"`
}

//Student  new Student
type Student struct {
	ID         int    `json:"id,omitempty"`
	UserID     int    `json:"user_id,omitempty"`
	TeamID     int    `json:"team_id,omitempty"`
	Profession int    `json:"profession,omitempty"`
	CV         string `json:"cv,omitempty"`
	TeamRole   string `json:"team_role,omitempty"`
}

//Investor new investor
type Investor struct {
	ID       int    `json:"id,omitempty"`
	UserID   int    `json:"user_id,omitempty"`
	Linkedin string `json:"linkedin,omitempty"`
	Company  string `json:"company,omitempty"`
	Projects []int  `json:"projects,omitempty"`
}

//Project new host of project
type Project struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedDate string `json:"created_date,omitempty"`
	ClosedDate  string `json:"closed_date,omitempty"`
	Category    int    `json:"category,omitempty"` //Agriculture, IT,
	Host        int    `json:"host,omitempty"`
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
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

//ProjectCategory catagories of the project
type ProjectCategory struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// 	Date  int
// 	Month int
// 	Year  int
// }
