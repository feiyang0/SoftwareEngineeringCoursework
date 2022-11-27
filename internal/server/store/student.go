package store

import v1 "SoftwareEngine/internal/pkg/model/server/v1"

type StudentStore interface {
	Get(uid, pid uint64) (*v1.StudentProblem, error)
	Commit(uid, pid uint64) error
	CancelCommit(uid, pid uint64) error
	Collect(uid, pid uint64) error
	CancelCollect(uid, pid uint64) error
}
