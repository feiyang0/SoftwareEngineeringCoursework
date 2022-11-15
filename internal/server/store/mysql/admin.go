package mysql

import (
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"gorm.io/gorm"
)

type admin struct {
	db *gorm.DB
}

func newAdmin(ds *datastore) *admin {
	return &admin{db: ds.db}
}

func (t *admin) Create(user *v1.User) error {

	return nil
}
func (t *admin) Update(user *v1.User) error {
	return nil
}
func (t *admin) Delete(username string) error {
	return nil
}
func (t *admin) Get(username string) (*v1.User, error) {

	return nil, nil
}
func (t *admin) List() ([]*v1.User, error) {
	return nil, nil
}
