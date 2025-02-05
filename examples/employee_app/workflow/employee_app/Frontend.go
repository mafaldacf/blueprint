package employee_app

import (
	"context"
)

type Frontend interface {
	//CreateEmployee(ctx context.Context, employeeID string, name string, IBAN string) (Employee, error)
	CreateEmployeeFreelancer(ctx context.Context, employeeID string, name string, IBAN string, freelancerID string, rate string, terms string) (Employee, Freelancer, error)
	CreateEmployeeFulltime(ctx context.Context, employeeID string, name string, IBAN string, fulltimeID string, salary string, position string) (Employee, Fulltime, error)
	CreateEmployeeIntern(ctx context.Context, employeeID string, name string, IBAN string, internID string, mentorID string, stipend string, duration string) (Employee, Intern, error)
	PromoteFreelancerToFulltime(ctx context.Context, employeeID string, name string, IBAN string, fulltimeID string, salary string, position string) (Fulltime, error)
	DeleteEmployee(ctx context.Context, employeeID string) error
	/* GetEmployeeFreelancer(ctx context.Context, freelancerID string) (Employee, Freelancer, error)
	GetEmployeeFulltime(ctx context.Context, fulltimeID string) (Employee, Fulltime, error) */
}

type FrontendImpl struct {
	employeeService   EmployeeService
	freelancerService FreelancerService
	fulltimeService   FulltimeService
	internService     InternService
}

func NewFrontendImpl(ctx context.Context, employeeService EmployeeService, freelancerService FreelancerService, fulltimeService FulltimeService, internService InternService) (Frontend, error) {
	return &FrontendImpl{employeeService: employeeService, freelancerService: freelancerService, fulltimeService: fulltimeService, internService: internService}, nil
}

/* func (u *FrontendImpl) CreateEmployee(ctx context.Context, employeeID string, name string, IBAN string) (Employee, error) {
	employee, err := u.employeeService.CreateEmployee(ctx, employeeID, name, IBAN)
	return employee, err
} */

func (u *FrontendImpl) CreateEmployeeFreelancer(ctx context.Context, employeeID string, name string, IBAN string, freelancerID string, rate string, terms string) (Employee, Freelancer, error) {
	employee, freelancer, err := u.employeeService.CreateEmployeeFreelancer(ctx, employeeID, name, IBAN, freelancerID, rate, terms)
	return employee, freelancer, err
}

func (u *FrontendImpl) CreateEmployeeFulltime(ctx context.Context, employeeID string, name string, IBAN string, fulltimeID string, salary string, position string) (Employee, Fulltime, error) {
	employee, fulltime, err := u.employeeService.CreateEmployeeFulltime(ctx, employeeID, name, IBAN, fulltimeID, salary, position)
	return employee, fulltime, err
}

func (u *FrontendImpl) CreateEmployeeIntern(ctx context.Context, employeeID string, name string, IBAN string, internID string, mentorID string, stipend string, duration string) (Employee, Intern, error) {
	employee, intern, err := u.employeeService.CreateEmployeeIntern(ctx, employeeID, name, IBAN, internID, mentorID, stipend, duration)
	return employee, intern, err
}

func (u *FrontendImpl) PromoteFreelancerToFulltime(ctx context.Context, employeeID string, name string, IBAN string, fulltimeID string, salary string, position string) (Fulltime, error) {
	fulltime, err := u.employeeService.PromoteFreelancerToFulltime(ctx, employeeID, name, IBAN, fulltimeID, salary, position)
	return fulltime, err
}

func (u *FrontendImpl) DeleteEmployee(ctx context.Context, employeeID string) error {
	err := u.employeeService.DeleteEmployee(ctx, employeeID)
	return err
}

/* func (u *FrontendImpl) GetEmployeeFreelancer(ctx context.Context, freelancerID string) (Employee, Freelancer, error) { */
/* freelancer, err := u.freelancerService.GetFreelancer(ctx, freelancerID)
if err != nil {
	return Employee{}, Freelancer{}, err
}
employee, err := u.employeeService.GetEmployee(ctx, freelancer.EmployeeID)
if err != nil {
	return Employee{}, Freelancer{}, err
}
return employee, freelancer, nil */
/* return u.employeeService.GetEmployeeFreelancer(ctx, freelancerID)
} */

/* func (u *FrontendImpl) GetEmployeeFulltime(ctx context.Context, fulltimeID string) (Employee, Fulltime, error) { */
/* fulltime, err := u.fulltimeService.GetFulltime(ctx, fulltimeID)
if err != nil {
	return Employee{}, Fulltime{}, err
}
employee, err := u.employeeService.GetEmployee(ctx, fulltime.EmployeeID)
if err != nil {
	return Employee{}, Fulltime{}, err
}
return employee, fulltime, nil */
/* return u.employeeService.GetEmployeeFulltime(ctx, fulltimeID)
} */
