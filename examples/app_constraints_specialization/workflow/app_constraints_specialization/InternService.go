package app_constraints_specialization

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type InternService interface {
	CreateIntern(ctx context.Context, employeeID string, internID string, mentorID string, stipend string, duration string) (Intern, error)
	//GetInternByEmployeeID(ctx context.Context, employeeID string) (Intern, error)
	GetIntern(ctx context.Context, internID string) (Intern, error)
}

type InternServiceImpl struct {
	internsDB backend.NoSQLDatabase
}

func NewInternServiceImpl(ctx context.Context, internsDB backend.NoSQLDatabase) (InternService, error) {
	s := &InternServiceImpl{internsDB: internsDB}
	return s, nil
}

func (s *InternServiceImpl) CreateIntern(ctx context.Context, employeeID string, internID string, mentorID string, stipend string, duration string) (Intern, error) {
	employee := Intern{
		EmployeeID: employeeID,
		InternID:   internID,
		Mentor:     mentorID,
		Stipend:    stipend,
		Duration:   duration,
	}
	collection, err := s.internsDB.GetCollection(ctx, "interns", "interns")
	if err != nil {
		return employee, err
	}
	err = collection.InsertOne(ctx, employee)
	return employee, err
}

/* func (s *InternServiceImpl) GetInternByEmployeeID(ctx context.Context, employeeID string) (Intern, error) {
	var intern Intern
	collection, err := s.internsDB.GetCollection(ctx, "interns", "interns")
	if err != nil {
		return intern, err
	}
	query := bson.D{{Key: "employeeID", Value: employeeID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return Intern{}, err
	}
	res, err := result.One(ctx, &intern)
	if !res || err != nil {
		return Intern{}, err
	}
	return intern, err
} */

func (s *InternServiceImpl) GetIntern(ctx context.Context, internID string) (Intern, error) {
	var intern Intern
	collection, err := s.internsDB.GetCollection(ctx, "interns", "interns")
	if err != nil {
		return intern, err
	}
	query := bson.D{{Key: "internID", Value: internID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return Intern{}, err
	}
	res, err := result.One(ctx, &intern)
	if !res || err != nil {
		return Intern{}, err
	}
	return intern, err
}
