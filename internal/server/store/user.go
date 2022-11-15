package store

import v1 "SoftwareEngine/internal/pkg/model/server/v1"

type UserStore interface {
	Create(user *v1.User) error
	Update(user *v1.User) error
	Delete(username string) error
	GetById(schoolId uint64) (*v1.User, error)
	GetByEmail(email string) (*v1.User, error)
	List() ([]*v1.User, error)
}
