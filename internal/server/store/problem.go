package store

import v1 "SoftwareEngine/internal/pkg/model/server/v1"

type ProblemStore interface {
	Create(problem *v1.Problem) error
	GetTags() ([]*v1.Tag, error)
	GetAll(opts *v1.ProblemListOption) ([]*v1.Problem, error)
	GetAllWithTag(uid uint64, opts *v1.ProblemListOption) ([]*v1.Problem, error)
	GetProblem(title string) (*v1.Problem, error)
	Update(problem *v1.Problem) error
	Delete(ID uint64) error
}
