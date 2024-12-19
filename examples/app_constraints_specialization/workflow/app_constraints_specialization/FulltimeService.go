package app_constraints_specialization

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type FulltimeService interface {
	CreateFulltime(ctx context.Context, employeeID string, freelancerID string, rate string, terms string) (Fulltime, error)
	//GetFulltimeByEmployeeID(ctx context.Context, employeeID string) (Fulltime, error)
	GetFulltime(ctx context.Context, fulltimeID string) (Fulltime, error)
}

type FulltimeServiceImpl struct {
	fulltimesDB backend.NoSQLDatabase
}

func NewFulltimeServiceImpl(ctx context.Context, fulltimesDB backend.NoSQLDatabase) (FulltimeService, error) {
	s := &FulltimeServiceImpl{fulltimesDB: fulltimesDB}
	return s, nil
}

func (s *FulltimeServiceImpl) CreateFulltime(ctx context.Context, employeeID string, freelancerID string, salary string, position string) (Fulltime, error) {
	employee := Fulltime{
		EmployeeID: employeeID,
		FulltimeID: freelancerID,
		Salary:     salary,
		Position:   position,
	}
	collection, err := s.fulltimesDB.GetCollection(ctx, "fulltimes", "fulltimes")
	if err != nil {
		return employee, err
	}
	err = collection.InsertOne(ctx, employee)
	return employee, err
}

/* func (s *FulltimeServiceImpl) GetFulltimeByEmployeeID(ctx context.Context, employeeID string) (Fulltime, error) {
	var fulltime Fulltime
	collection, err := s.fulltimesDB.GetCollection(ctx, "fulltimes", "fulltimes")
	if err != nil {
		return fulltime, err
	}
	query := bson.D{{Key: "employeeID", Value: employeeID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return Fulltime{}, err
	}
	res, err := result.One(ctx, &fulltime)
	if !res || err != nil {
		return Fulltime{}, err
	}
	return fulltime, err
} */

func (s *FulltimeServiceImpl) GetFulltime(ctx context.Context, fulltimeID string) (Fulltime, error) {
	var fulltime Fulltime
	collection, err := s.fulltimesDB.GetCollection(ctx, "fulltimes", "fulltimes")
	if err != nil {
		return fulltime, err
	}
	query := bson.D{{Key: "fulltimeID", Value: fulltimeID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return Fulltime{}, err
	}
	res, err := result.One(ctx, &fulltime)
	if !res || err != nil {
		return Fulltime{}, err
	}
	return fulltime, err
}
