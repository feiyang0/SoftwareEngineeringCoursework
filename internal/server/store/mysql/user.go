package mysql

import (
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/pkg/errno"
	"errors"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *users {
	return &users{ds.db}
}

func (u *users) Create(user *v1.User) error {
	return u.db.Create(&user).Error
}
func (u *users) Update(user *v1.User) error {
	return u.db.Save(user).Error
}
func (u *users) Delete(username string) error {
	return nil
}
func (u *users) GetById(schoolId uint64) (*v1.User, error) {
	user := &v1.User{}
	err := u.db.Where("id = ?", schoolId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (u *users) GetByEmail(email string) (*v1.User, error) {
	user := &v1.User{}
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (u *users) List() ([]*v1.User, error) {
	return nil, nil
}
