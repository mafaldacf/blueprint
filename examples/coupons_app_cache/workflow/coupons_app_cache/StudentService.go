package coupons_app_cache

import (
	"context"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type StudentService interface {
	CreateStudent(ctx context.Context, studentID string, name string) error
	AddToBalance(ctx context.Context, studentID string, couponValue int) error
}

type StudentServiceImpl struct {
	studentsCache backend.Cache
}

func NewStudentServiceImpl(ctx context.Context, studentsCache backend.Cache) (StudentService, error) {
	s := &StudentServiceImpl{studentsCache: studentsCache}
	return s, nil
}

func (s *StudentServiceImpl) CreateStudent(ctx context.Context, studentID string, name string) error {
	student := Student{
		StudentID: studentID,
		Name: name,
	}
	err := s.studentsCache.Put(ctx, studentID, student)
	return err
}

func (s *StudentServiceImpl) AddToBalance(ctx context.Context, studentID string, couponValue int) error {
	_, err := s.studentsCache.Incr(ctx, studentID)
	return err
}
