package employee_app

type Employee struct {
	EmployeeID string
	Name       string
	IBAN       string
	SpecFlag   string
}

type Freelancer struct {
	EmployeeID   string
	FreelancerID string
	Rate         string
	Terms        string
}

type Fulltime struct {
	EmployeeID string
	FulltimeID string
	Salary     string
	Position   string
}

type Intern struct {
	EmployeeID string
	InternID   string
	Stipend    string
	Duration   string
	Mentor     string
}
