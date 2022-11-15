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

func (s *students) Create(user *v1.User) error {

	return nil
}
func (s *students) Update(user *v1.User) error {
	return nil
}
func (s *students) Delete(username string) error {
	return nil
}
func (s *students) Get(username string) (*v1.User, error) {

	return nil, nil
}
func (s *students) List() ([]*v1.User, error) {
	return nil, nil
}
