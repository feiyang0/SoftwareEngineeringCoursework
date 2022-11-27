package store

type Factory interface {
	Users() UserStore
	Problems() ProblemStore
	Students() StudentStore
	Close() error
}

//Students() StudentStore
//Teachers() TeacherStore
//Admins() AdminStore
