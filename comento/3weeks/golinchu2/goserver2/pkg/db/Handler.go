package db

import (
	"fmt"

	"github.com/pkg/errors"
	"main.go/pkg/model"
)

func (h *DBHandler) CreateStudent(student *model.Students) (*model.Students, error) {
	result := h.gDB.Create(student)
	return student, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) GetStudentlist() ([]*model.Students, error) {
	studentlist := []*model.Students{}
	result := h.gDB.Find(&studentlist)
	fmt.Println(studentlist)

	return studentlist, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) GetStudentlistByrollID(rollno string) (*model.Students, error) {
	students := &model.Students{}
	result := h.gDB.First(students, rollno)

	return students, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) UpdateStudentList(rollno string, newstudent *model.Students) (*model.Students, error) {
	oldStudent := &model.Students{}
	result := h.gDB.Model(oldStudent).Where(rollno).Updates(newstudent)
	return oldStudent, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) DeleteStudent(rollno string) error {
	result := h.gDB.Delete(&model.Students{}, rollno)
	return errors.Wrap(result.Error, "db handler error")
}
