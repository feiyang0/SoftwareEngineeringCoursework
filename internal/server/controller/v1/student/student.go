package student

import "SoftwareEngine/internal/server/store"

type StudentController struct {
	studentS store.StudentStore
}

func NewStudentController(store store.Factory) *StudentController {
	return &StudentController{
		studentS: store.Students(),
	}
}
