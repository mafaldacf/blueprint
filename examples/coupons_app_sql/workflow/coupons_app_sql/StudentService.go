package coupons_app_sql

import (
	"context"
	"os"
	"strings"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type StudentService interface {
	CreateStudent(ctx context.Context, name string) error
	AddToBalance(ctx context.Context, studentID int, couponValue int) error
}

type StudentServiceImpl struct {
	studentsDB backend.RelationalDB
}

func NewStudentServiceImpl(ctx context.Context, studentsDB backend.RelationalDB) (StudentService, error) {
	s := &StudentServiceImpl{studentsDB: studentsDB}
	s.createTables(ctx)
	return s, nil
}

func (s *StudentServiceImpl) CreateStudent(ctx context.Context, name string) error {
	_, err := s.studentsDB.Exec(ctx, "INSERT INTO students(name) VALUES (?);", name)
	return err
}

func (s *StudentServiceImpl) AddToBalance(ctx context.Context, studentID int, couponValue int) error {
	var student Student
	err := s.studentsDB.Select(ctx, &student, "SELECT * FROM students WHERE student_id = ?", studentID)
	if err != nil {
		return err
	}

	newBalance := student.Balance + couponValue
	_, err = s.studentsDB.Exec(ctx, "UPDATE students SET balance = ? WHERE student_id = ?", newBalance, studentID)
	return err
}

func (s *StudentServiceImpl) createTables(ctx context.Context) error {
	sqlBytes, err := os.ReadFile("database/students.sql")
	if err != nil {
		return err
	}
	sqlStatements := strings.Split(string(sqlBytes), ";")
	for _, stmt := range sqlStatements {
		_, err := s.studentsDB.Exec(ctx, stmt)
		if err != nil {
			return err
		}
	}
	return nil
}
