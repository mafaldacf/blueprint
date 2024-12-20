package app_constraints_specialization

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type FreelancerService interface {
	CreateFreelancer(ctx context.Context, employeeID string, freelancerID string, rate string, terms string) (Freelancer, error)
	DeleteFreelancer(ctx context.Context, employeeID string) error
	//GetFreelancerByEmployeeID(ctx context.Context, employeeID string) (Freelancer, error)
	GetFreelancer(ctx context.Context, freelancerID string) (Freelancer, error)
}

type FreelancerServiceImpl struct {
	freelancersDB backend.NoSQLDatabase
}

func NewFreelancerServiceImpl(ctx context.Context, freelancersDB backend.NoSQLDatabase) (FreelancerService, error) {
	s := &FreelancerServiceImpl{freelancersDB: freelancersDB}
	return s, nil
}

func (s *FreelancerServiceImpl) CreateFreelancer(ctx context.Context, employeeID string, freelancerID string, rate string, terms string) (Freelancer, error) {
	employee := Freelancer{
		EmployeeID:   employeeID,
		FreelancerID: freelancerID,
		Rate:         rate,
		Terms:        terms,
	}
	collection, err := s.freelancersDB.GetCollection(ctx, "freelancers", "freelancers")
	if err != nil {
		return employee, err
	}
	err = collection.InsertOne(ctx, employee)
	return employee, err
}

/* func (s *FreelancerServiceImpl) GetFreelancerByEmployeeID(ctx context.Context, employeeID string) (Freelancer, error) {
	var freelancer Freelancer
	collection, err := s.freelancersDB.GetCollection(ctx, "freelancers", "freelancers")
	if err != nil {
		return freelancer, err
	}
	query := bson.D{{Key: "employeeID", Value: employeeID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return Freelancer{}, err
	}
	res, err := result.One(ctx, &freelancer)
	if !res || err != nil {
		return Freelancer{}, err
	}
	return freelancer, err
} */

func (s *FreelancerServiceImpl) GetFreelancer(ctx context.Context, freelancerID string) (Freelancer, error) {
	var freelancer Freelancer
	collection, err := s.freelancersDB.GetCollection(ctx, "freelancers", "freelancers")
	if err != nil {
		return freelancer, err
	}
	query := bson.D{{Key: "freelancerID", Value: freelancerID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return Freelancer{}, err
	}
	res, err := result.One(ctx, &freelancer)
	if !res || err != nil {
		return Freelancer{}, err
	}
	return freelancer, err
}

func (s *FreelancerServiceImpl) DeleteFreelancer(ctx context.Context, employeeID string) error {
	collection, err := s.freelancersDB.GetCollection(ctx, "freelancers", "freelancers")
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "employeeID", Value: employeeID}}
	err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
