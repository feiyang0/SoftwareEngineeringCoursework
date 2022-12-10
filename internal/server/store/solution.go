package store

import v1 "SoftwareEngine/internal/pkg/model/server/v1"

type SolutionStore interface {
	Create(pid uint64, solution *v1.Solution) error
	Update(solution *v1.Solution) error
	Delete(pid, sid uint64) error
	GetSolutionList(opts *v1.SolutionListOption) ([]*v1.Solution, int64, error)
	GetSolution(sid uint64) (*v1.Solution, error)
	AddComment(sid uint64, comment *v1.Comment) error
	DelComment(sid, cid uint64) error
}
