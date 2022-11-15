package store

type Factory interface {
	Users() UserStore
	Students() StudentStore
	Teachers() TeacherStore
	Admins() AdminStore
	Close() error
}
