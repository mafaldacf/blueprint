package app_constraints_specialization

type Employee struct {
	EmployeeID     string
	Name           string
	IBAN           string
	Specialization string
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
