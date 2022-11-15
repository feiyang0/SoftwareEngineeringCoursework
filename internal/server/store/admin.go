package store

import v1 "SoftwareEngine/internal/pkg/model/server/v1"

type AdminStore interface {
	Create(user *v1.User) error
	Update(user *v1.User) error
	Delete(username string) error
	Get(username string) (*v1.User, error)
	List() ([]*v1.User, error)
}
