package mysql

import (
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/pkg/errno"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"strconv"
)

type problems struct {
	db *gorm.DB
}

func newProblems(ds *datastore) *problems {
	return &problems{db: ds.db}
}

func (p *problems) Create(problem *v1.Problem) error {
	// 存problem
	problem.Cnt = 0
	if err := p.db.Create(&problem).Error; err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'problem.title'", err.Error()); match {
			return errno.ErrProblemAlreadyExist
		}
		return err
	}
	return nil
}
func (p *problems) getProblem(id uint64) (*v1.Problem, error) {
	problem := &v1.Problem{}
	err := p.db.Where("id = ?", id).First(&problem).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrProblemNotFound
		}
	}
	err = p.db.Model(&problem).Association("Tags").Find(&problem.Tags)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrProblemNotFound
		}
	}
	return problem, nil
}
func (p *problems) Update(problem *v1.Problem) error {
	tempP, _ := p.getProblem(problem.ID)

	pt := &v1.ProblemTag{}
	p.db.Where("problem_id = ?", problem.ID).Delete(&pt)

	problem.CreatedAt = tempP.CreatedAt
	return p.db.Save(&problem).Error
}

func (p *problems) Delete(id uint64) error {
	err := p.db.Where("id = ?", id).Delete(&v1.Problem{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errno.ErrProblemNotFound
	}
	return nil
}

func (p *problems) GetTags() ([]*v1.Tag, error) {
	var tags []*v1.Tag
	if err := p.db.Find(&tags).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrTagNotFound
		}
		return nil, err
	}
	return tags, nil
}
func (p *problems) GetAll(opts *v1.ProblemListOption) ([]*v1.Problem, error) {
	var ps []*v1.Problem
	p.db.Offset(opts.Offset).Limit(opts.Limit).
		Order("id").Find(&ps)

	for _, problem := range ps {
		p.db.Model(&problem).Association("Tags").Find(&problem.Tags)
	}

	return ps, nil
}
func (p *problems) GetAllWithTag(uid uint64, opts *v1.ProblemListOption) ([]*v1.Problem, int64, error) {
	var ps []*v1.Problem
	//orders := fmt.Sprintf("%s %s", opts.OrderBy, opts.SortOrder)
	var orders string

	if opts.Orders != nil {
		for _, o := range opts.Orders {
			orders += fmt.Sprintf("%s %s,", o.OrderBy, o.SortOrder)
		}
		orders = orders[0 : len(orders)-1]
	}

	tx := p.db.Where("").Find(&ps) // 这里需要先初始化
	// 先关键词搜索
	if opts.SearchKeyWords != "" {
		keyId, err := strconv.ParseInt(opts.SearchKeyWords, 10, 64)
		if err == nil {
			//fmt.Println("--------", keyId, "------------")
			tx.Where("id = ?", uint64(keyId)).Find(&ps)
		} else {
			tx.Where(fmt.Sprintf(" title like '%%%s%%' ", opts.SearchKeyWords)).Find(&ps)
		}
	}
	if opts.CourseName != "" {
		tx.Where("courseName = ?", opts.CourseName).Find(&ps)
	}
	if opts.Category != "" {
		//fmt.Println("--------", opts.Category, "------------")
		tx.Where("category = ?", opts.Category).Find(&ps)
	}
	if opts.Difficulty != 0 {
		tx.Where("difficulty = ?", opts.Difficulty).Find(&ps)
	}
	if opts.Tag != "" {
		//fmt.Println("--------", opts.Tag, "------------")
		tx.Joins("left join problem_tags on problem_tags.problem_id = problem.id").
			Where("problem_tags.tag_name = ?", opts.Tag).Find(&ps)
	}
	var problemNumber int64
	tx.Count(&problemNumber)
	// 取数据
	tx.Offset(opts.Offset).Limit(opts.Limit).Order(orders).Find(&ps)
	// 同步题目状态
	for _, problem := range ps {
		p.db.Model(&problem).Association("Tags").Find(&problem.Tags)

		stuP := &v1.StudentProblem{}
		p.db.Where("user_id = ? and problem_id = ?", uid, problem.ID).First(&stuP)
		problem.Pass, problem.Favour = stuP.Pass, stuP.Favour
	}

	return ps, problemNumber, nil
}
func (p *problems) GetProblem(uid, pid uint64) (*v1.Problem, error) {
	problem := &v1.Problem{}
	err := p.db.Where("id = ?", pid).First(&problem).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrProblemNotFound
		}
	}
	err = p.db.Model(&problem).Association("Tags").Find(&problem.Tags)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrProblemNotFound
		}
	}
	sp := &v1.StudentProblem{}
	p.db.Where("user_id = ? and problem_id = ?", uid, pid).First(&sp)
	problem.Pass = sp.Pass
	problem.Favour = sp.Favour
	return problem, nil
}
