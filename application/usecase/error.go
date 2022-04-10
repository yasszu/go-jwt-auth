package usecase

type UnexpectedError struct {
	Err error
}

func (e *UnexpectedError) Error() string {
	return e.Err.Error()
}

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string {
	return "not found " + e.Name
}

type UnauthorizedError struct {
	Massage string
}

func (e *UnauthorizedError) Error() string {
	return e.Massage
}
