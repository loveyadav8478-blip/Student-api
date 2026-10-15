package storage

import "github.com/loveyadav8478-blip/student-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)

	GetStudentById(id int64) (types.Student,error)

	// DeleteByID(id int64) (types.Student,error)
	// UpdateByID(id int64) (types.Student,error)
}