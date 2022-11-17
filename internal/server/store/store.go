package store

type Factory interface {
	Users() UserStore
	Problems() ProblemStore
	Close() error
}

//Students() StudentStore
//Teachers() TeacherStore
//Admins() AdminStore
