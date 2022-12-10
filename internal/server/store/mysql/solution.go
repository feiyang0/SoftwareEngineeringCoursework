package mysql

import (
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/pkg/errno"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type solutions struct {
	db *gorm.DB
}

func newSolutions(ds *datastore) *solutions {
	return &solutions{ds.db}
}

func (s *solutions) Create(pid uint64, solution *v1.Solution) error {
	s.db.Create(&solution)
	s.db.Save(&v1.ProblemSolution{
		ProblemId: pid, SolutionId: solution.ID,
	})
	return nil
}

func (s *solutions) Update(solution *v1.Solution) error {
	tempS, _ := s.GetSolution(solution.ID)

	solution.CreatedAt = tempS.CreatedAt

	return s.db.Save(&solution).Error
}

func (s *solutions) Delete(pid, sid uint64) error {
	s.db.Where("problem_id = ?, solution_id = ?", pid, sid).Delete(&v1.ProblemSolution{})
	return s.db.Where("id = ?", sid).Delete(&v1.Solution{}).Error
}

func (s *solutions) GetSolutionList(opts *v1.SolutionListOption) ([]*v1.Solution, int64, error) {
	var ss []*v1.Solution
	var solutionNumber int64

	tx := s.db.Joins("left join problem_solution on problem_solution.solution_id = solution.id").
		Where("problem_solution.problem_id = ?", opts.Pid).
		Find(&ss).Count(&solutionNumber)

	tx.Offset(opts.Offset).Limit(opts.Limit).Find(&ss)

	return ss, solutionNumber, nil
}

func (s *solutions) GetSolution(sid uint64) (*v1.Solution, error) {
	solution := &v1.Solution{}
	err := s.db.Where("id = ?", sid).First(&solution).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrSolutionNotFound
		}
	}
	err = s.db.Model(&solution).Association("Comments").Find(&solution.Comments)
	return solution, nil
}

func (s *solutions) AddComment(sid uint64, comment *v1.Comment) error {
	s.db.Create(&comment)
	fmt.Println("comment.ID:", comment.ID)
	return s.db.Save(&v1.SolutionComment{
		SolutionId: sid, CommentId: comment.ID,
	}).Error
}

func (s *solutions) DelComment(sid, cid uint64) error {
	s.db.Where("solution_id = ? and comment_id = ?", sid, cid).Delete(&v1.SolutionComment{})
	s.db.Where("id = ?", cid).Delete(&v1.Comment{})
	return nil
}
