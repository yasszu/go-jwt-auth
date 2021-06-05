package entity

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
	return e.Name + ": not found"
}

type UnauthorizedError struct {
	Massage string
}

func (e *UnauthorizedError) Error() string {
	return e.Massage
}
