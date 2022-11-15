package mysql

import (
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"gorm.io/gorm"
)

type teachers struct {
	db *gorm.DB
}

func newTeachers(ds *datastore) *teachers {
	return &teachers{db: ds.db}
}

func (t *teachers) Create(user *v1.User) error {

	return nil
}
func (t *teachers) Update(user *v1.User) error {
	return nil
}
func (t *teachers) Delete(username string) error {
	return nil
}
func (t *teachers) Get(username string) (*v1.User, error) {

	return nil, nil
}
func (t *teachers) List() ([]*v1.User, error) {
	return nil, nil
}
