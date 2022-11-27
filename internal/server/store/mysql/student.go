package mysql

import (
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"gorm.io/gorm"
)

type students struct {
	db *gorm.DB
}

func newStudents(ds *datastore) *students {
	return &students{ds.db}
}
func (s *students) Get(uid, pid uint64) (*v1.StudentProblem, error) {
	sp := &v1.StudentProblem{}
	err := s.db.Where("user_id = ? and problem_id = ?", uid, pid).First(&sp).Error
	return sp, err
}
func (s *students) calCnt(pid uint64, num int) {
	problem := &v1.Problem{}
	s.db.Where("id = ?", pid).First(&problem)
	problem.Cnt += num
	s.db.Save(&problem)
}
func (s *students) Commit(uid, pid uint64) error {
	sp, err := s.Get(uid, pid)
	if err != nil {
		sp = &v1.StudentProblem{
			UserId:    uid,
			ProblemId: pid,
			Pass:      true,
		}
	} else {
		sp.Pass = true
	}
	s.calCnt(pid, 1)
	return s.db.Save(sp).Error
}

func (s *students) CancelCommit(uid, pid uint64) error {
	sp, err := s.Get(uid, pid)
	if err != nil {
		return err
	}
	sp.Pass = false
	s.calCnt(pid, -1)
	return s.db.Save(sp).Error
}

func (s *students) Collect(uid, pid uint64) error {
	sp, err := s.Get(uid, pid)
	if err != nil {
		sp = &v1.StudentProblem{
			UserId:    uid,
			ProblemId: pid,
			Favour:    true,
		}
	} else {
		sp.Favour = true
	}
	return s.db.Save(sp).Error
}

func (s *students) CancelCollect(uid, pid uint64) error {
	sp, err := s.Get(uid, pid)
	if err != nil {
		return err
	}
	sp.Favour = false
	return s.db.Save(sp).Error
}
