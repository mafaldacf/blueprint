package coupons_app

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type StudentService interface {
	CreateStudent(ctx context.Context, studentID int, name string) (Student, error)
	AddToBalance(ctx context.Context, studentID int, value int) (Student, error)
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
		StudentID: studentID,
		Name:      name,
	}
	collection, err := s.studentsDB.GetCollection(ctx, "students", "students")
	if err != nil {
		return coupon, err
	}
	err = collection.InsertOne(ctx, coupon)
	return coupon, err
}

func (s *StudentServiceImpl) AddToBalance(ctx context.Context, studentID int, value int) (Student, error) {
	var student Student

	collection, err := s.studentsDB.GetCollection(ctx, "students", "students")
	if err != nil {
		return student, err
	}

	query := bson.D{{Key: "studentID", Value: student}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return student, nil
	}
	found, err := result.One(ctx, student)
	if err != nil || !found {
		return student, err
	}

	student.Balance += value
	err = collection.InsertOne(ctx, student)
	return student, err
}
