package app_constraints_specialization

import (
	"context"
	"sync"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type EmployeeService interface {
	//CreateEmployee(ctx context.Context, employeeID string, name string, IBAN string) (Employee, error)
	CreateEmployeeFreelancer(ctx context.Context, employeeID string, name string, IBAN string, freelancerID string, rate string, terms string) (Employee, Freelancer, error)
	CreateEmployeeFulltime(ctx context.Context, employeeID string, name string, IBAN string, fulltimeID string, salary string, position string) (Employee, Fulltime, error)
	CreateEmployeeIntern(ctx context.Context, employeeID string, name string, IBAN string, internID string, mentorID string, stipend string, duration string) (Employee, Intern, error)
	PromoteFreelancerToFulltime(ctx context.Context, employeeID string, name string, IBAN string, fulltimeID string, salary string, position string) (Fulltime, error)
	DeleteEmployee(ctx context.Context, employeeID string) error
	/* GetEmployee(ctx context.Context, employeeID string) (Employee, error)
	GetEmployeeFreelancer(ctx context.Context, freelancerID string) (Employee, Freelancer, error)
	GetEmployeeFulltime(ctx context.Context, fulltimeID string) (Employee, Fulltime, error) */
}

type EmployeeServiceImpl struct {
	freelancerService FreelancerService
	fulltimeService   FulltimeService
	internService     InternService
	employeesDB       backend.NoSQLDatabase
}

func NewEmployeeServiceImpl(ctx context.Context, freelancerService FreelancerService, fulltimeService FulltimeService, internService InternService, employeesDB backend.NoSQLDatabase) (EmployeeService, error) {
	s := &EmployeeServiceImpl{freelancerService: freelancerService, fulltimeService: fulltimeService, internService: internService, employeesDB: employeesDB}
	return s, nil
}

/* func (s *EmployeeServiceImpl) CreateEmployee(ctx context.Context, employeeID string, name string, IBAN string) (Employee, error) {
	employee := Employee{
		EmployeeID: employeeID,
		Name:       name,
		IBAN:       IBAN,
	}
	collection, err := s.employeesDB.GetCollection(ctx, "employees", "employees")
	if err != nil {
		return employee, err
	}
	err = collection.InsertOne(ctx, employee)
	return employee, err
} */

func (s *EmployeeServiceImpl) CreateEmployeeFreelancer(ctx context.Context, employeeID string, name string, IBAN string, freelancerID string, rate string, terms string) (Employee, Freelancer, error) {
	employee := Employee{
		EmployeeID: employeeID,
		Name:       name,
		IBAN:       IBAN,
		SpecFlag:   "freelancer",
	}
	collection, err := s.employeesDB.GetCollection(ctx, "employees", "employees")
	if err != nil {
		return Employee{}, Freelancer{}, err
	}
	err = collection.InsertOne(ctx, employee)
	if err != nil {
		return Employee{}, Freelancer{}, err
	}
	freelancer, err := s.freelancerService.CreateFreelancer(ctx, employeeID, freelancerID, rate, terms)
	if err != nil {
		return Employee{}, Freelancer{}, err
	}
	return employee, freelancer, nil
}

func (s *EmployeeServiceImpl) CreateEmployeeFulltime(ctx context.Context, employeeID string, name string, IBAN string, fulltimeID string, salary string, position string) (Employee, Fulltime, error) {
	employee := Employee{
		EmployeeID: employeeID,
		Name:       name,
		IBAN:       IBAN,
		SpecFlag:   "fulltime",
	}
	collection, err := s.employeesDB.GetCollection(ctx, "employees", "employees")
	if err != nil {
		return Employee{}, Fulltime{}, err
	}
	err = collection.InsertOne(ctx, employee)
	if err != nil {
		return Employee{}, Fulltime{}, err
	}
	fulltime, err := s.fulltimeService.CreateFulltime(ctx, employeeID, fulltimeID, salary, position)
	if err != nil {
		return Employee{}, Fulltime{}, err
	}
	return employee, fulltime, nil
}

func (s *EmployeeServiceImpl) CreateEmployeeIntern(ctx context.Context, employeeID string, name string, IBAN string, internID string, mentorID string, stipend string, duration string) (Employee, Intern, error) {
	employee := Employee{
		EmployeeID: employeeID,
		Name:       name,
		IBAN:       IBAN,
		SpecFlag:   "intern",
	}
	collection, err := s.employeesDB.GetCollection(ctx, "employees", "employees")
	if err != nil {
		return Employee{}, Intern{}, err
	}
	err = collection.InsertOne(ctx, employee)
	if err != nil {
		return Employee{}, Intern{}, err
	}
	intern, err := s.internService.CreateIntern(ctx, employeeID, internID, mentorID, stipend, duration)
	if err != nil {
		return Employee{}, Intern{}, err
	}
	return employee, intern, nil
}

func (s *EmployeeServiceImpl) PromoteFreelancerToFulltime(ctx context.Context, employeeID string, name string, IBAN string, fulltimeID string, salary string, position string) (Fulltime, error) {
	collection, err := s.employeesDB.GetCollection(ctx, "employees", "employees")
	if err != nil {
		return Fulltime{}, err
	}

	var fulltime Fulltime
	var err1, err2, err3 error
	var wg sync.WaitGroup
	wg.Add(3)

	go func() { // update employee specialization flag
		update := bson.D{{Key: "$set", Value: bson.D{{Key: "SpecFlag", Value: "fulltime"}}}}
		filter := bson.D{{Key: "EmployeeID", Value: employeeID}}
		_, err1 = collection.UpdateOne(ctx, filter, update)
		wg.Done()
	}()

	go func() { // delete freelancer
		err2 = s.freelancerService.DeleteFreelancer(ctx, employeeID)
		wg.Done()
	}()

	go func() { // create fulltime
		fulltime, err3 = s.fulltimeService.CreateFulltime(ctx, employeeID, fulltimeID, salary, position)
		wg.Done()
	}()

	wg.Wait()
	if err1 != nil || err2 != nil || err3 != nil {
		return Fulltime{}, err
	}

	return fulltime, nil
}

func (s *EmployeeServiceImpl) DeleteEmployee(ctx context.Context, employeeID string) error {
	collection, err := s.employeesDB.GetCollection(ctx, "employees", "employees")
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "freelancerID", Value: employeeID}}
	err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

/* func (s *EmployeeServiceImpl) GetEmployee(ctx context.Context, employeeID string) (Employee, error) {
	var employee Employee
	collection, err := s.employeesDB.GetCollection(ctx, "employees", "employees")
	if err != nil {
		return employee, nil
	}
	query := bson.D{{Key: "employeeID", Value: employeeID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return Employee{}, nil
	}
	res, err := result.One(ctx, &employee)
	if !res || err != nil {
		return Employee{}, err
	}
	return employee, nil
} */

/* func (s *EmployeeServiceImpl) GetEmployeeSpecializationByName(ctx context.Context, name string) (Employee, interface{}, error) {
	var employee Employee
	collection, err := s.employeesDB.GetCollection(ctx, "employees", "employees")
	if err != nil {
		return employee, nil, err
	}
	query := bson.D{{Key: "name", Value: name}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return Employee{}, nil, err
	}
	res, err := result.One(ctx, &employee)
	if !res || err != nil {
		return Employee{}, nil, err
	}
	if employee.Specialization == "freelancer" {
		freelancer, err := s.freelancerService.GetFreelancerByEmployeeID(ctx, employee.EmployeeID)
		return employee, freelancer, err
	}
	fulltime, err := s.fulltimeService.GetFulltimeByEmployeeID(ctx, employee.EmployeeID)
	return employee, fulltime, err
} */

/* func (s *EmployeeServiceImpl) GetEmployeeFreelancer(ctx context.Context, freelancerID string) (Employee, Freelancer, error) {
	freelancer, err := s.freelancerService.GetFreelancer(ctx, freelancerID)
	if err != nil {
		return Employee{}, Freelancer{}, err
	}

	var employee Employee
	collection, err := s.employeesDB.GetCollection(ctx, "employees", "employees")
	if err != nil {
		return Employee{}, Freelancer{}, err
	}
	query := bson.D{{Key: "employeeID", Value: freelancer.EmployeeID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return Employee{}, Freelancer{}, err
	}
	res, err := result.One(ctx, &employee)
	if !res || err != nil {
		return Employee{}, Freelancer{}, err
	}

	return employee, freelancer, nil
} */

/* func (s *EmployeeServiceImpl) GetEmployeeFulltime(ctx context.Context, fulltimeID string) (Employee, Fulltime, error) {
	fulltime, err := s.fulltimeService.GetFulltime(ctx, fulltimeID)
	if err != nil {
		return Employee{}, Fulltime{}, err
	}

	var employee Employee
	collection, err := s.employeesDB.GetCollection(ctx, "employees", "employees")
	if err != nil {
		return Employee{}, Fulltime{}, err
	}
	query := bson.D{{Key: "employeeID", Value: fulltime.EmployeeID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return Employee{}, Fulltime{}, err
	}
	res, err := result.One(ctx, &employee)
	if !res || err != nil {
		return Employee{}, Fulltime{}, err
	}

	return employee, fulltime, nil
} */
