package mysql

import (
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/pkg/errno"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"regexp"
)

type problems struct {
	db *gorm.DB
}

func newProblems(ds *datastore) *problems {
	return &problems{db: ds.db}
}

func (p *problems) Create(problem *v1.Problem) error {
	// 存problem
	if err := p.db.Create(&problem).Error; err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'problem.title'", err.Error()); match {
			return errno.ErrProblemAlreadyExist
		}
		return err
	}
	//// 存tag
	//var tags []v1.Tag
	//var pTags []v1.ProblemTag
	//// 存problem-tag
	//for _, t := range problem.Tags {
	//	pTags = append(pTags, v1.ProblemTag{ProblemId: problem.ID, TagName: t.Name})
	//	// tag存在就插入
	//	tag := &v1.Tag{}
	//	err := u.db.Where("name = ?", t).First(&tag).Error
	//	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
	//		tags = append(tags, *tag)
	//	}
	//}
	//if err := u.db.Create(&tags).Error; err != nil {
	//	return err
	//}
	//if err := u.db.Create(&pTags).Error; err != nil {
	//	return err
	//}
	return nil
}

func (p *problems) Update(problem *v1.Problem) error {
	return p.db.Save(problem).Error
}
func (p *problems) Delete(title string) error {
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
func (p *problems) GetAll(opts *v1.ListOption) ([]*v1.Problem, error) {
	var ps []*v1.Problem
	p.db.Offset(opts.Offset).Limit(opts.Limit).
		Order("id").Find(&ps)

	for _, problem := range ps {
		p.db.Model(&problem).Association("Tags").Find(&problem.Tags)
	}
	return ps, nil
}
func (p *problems) GetAllWithTag(opts *v1.ListOption) ([]*v1.Problem, error) {
	var ps []*v1.Problem
	orders := fmt.Sprintf("%s %s", opts.OrderBy, opts.SortOrder)

	p.db.Order(orders)
	return ps, nil
}
func (p *problems) GetProblem(title string) (*v1.Problem, error) {
	problem := &v1.Problem{}
	err := p.db.Where("title = ?", title).First(&problem).Error
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

//func (u *problems) Get(problem *v1.Problem) error {
//	user := &v1.User{}
//	err := u.db.Where("id = ?", schoolId).First(&user).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, errno.ErrUserNotFound
//		}
//		return nil, err
//	}
//	return user, nil
//}

//func (u *problems) GetByName(name string) (*v1.User, error) {
//	user := &v1.User{}
//	err := u.db.Where("email = ?", name).First(&user).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, errno.ErrUserNotFound
//		}
//		return nil, err
//	}
//	return user, nil
//}

func (p *problems) List() ([]*v1.User, error) {
	return nil, nil
}
