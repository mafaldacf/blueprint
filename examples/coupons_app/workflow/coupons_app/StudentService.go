package coupons_app

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type StudentService interface {
	CreateStudent(ctx context.Context, studentID int, name string) (Student, error)
	AddToBalance(ctx context.Context, studentID int, value int) error
}

type StudentServiceImpl struct {
	studentsDB backend.NoSQLDatabase
}

func NewStudentServiceImpl(ctx context.Context, studentsDB backend.NoSQLDatabase) (StudentService, error) {
	s := &StudentServiceImpl{studentsDB: studentsDB}
	return s, nil
}

func (s *StudentServiceImpl) CreateStudent(ctx context.Context, studentID int, name string) (Student, error) {
	coupon := Student{
		//StudentID:         studentID,
		Name:              name,
		NumClaimedCoupons: 0,
	}
	collection, err := s.studentsDB.GetCollection(ctx, "students", "Student")
	if err != nil {
		return coupon, err
	}
	err = collection.InsertOne(ctx, coupon)
	return coupon, err
}

func (s *StudentServiceImpl) AddToBalance(ctx context.Context, studentID int, value int) error {
	collection, err := s.studentsDB.GetCollection(ctx, "students", "Student")
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "StudentID", Value: studentID}}
	update := bson.D{
		{Key: "$inc", Value: bson.D{{Key: "Balance", Value: 1}}},
		{Key: "$inc", Value: bson.D{{Key: "ClaimedCoupons", Value: 1}}}}

	res, err := collection.UpdateOne(ctx, filter, update)
	if res != 1 || err != nil {
		return err
	}
	return nil
}
