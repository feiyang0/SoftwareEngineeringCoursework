package store

import v1 "SoftwareEngine/internal/pkg/model/server/v1"

type ProblemStore interface {
	Create(problem *v1.Problem) error
	GetTags() ([]*v1.Tag, error)
	GetAll(opts *v1.ListOption) ([]*v1.Problem, error)
	GetProblem(title string) (*v1.Problem, error)
}
