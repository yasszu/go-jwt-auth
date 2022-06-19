package usecase

func newUnexpectedError() *ErrorUnexpected {
	return &ErrorUnexpected{
		Message: "unexpected error",
	}
}

type ErrorUnexpected struct {
	Message string
}

func (e *ErrorUnexpected) Error() string {
	return e.Message
}

func newNotFoundError() *ErrorNotFound {
	return &ErrorNotFound{
		Message: "not found resource",
	}
}

type ErrorNotFound struct {
	Message string
}

func (e *ErrorNotFound) Error() string {
	return e.Message
}

func newErrorUnauthorized() *ErrorUnauthorized {
	return &ErrorUnauthorized{
		Massage: "unauthorized",
	}
}

type ErrorUnauthorized struct {
	Massage string
}

func (e *ErrorUnauthorized) Error() string {
	return e.Massage
}
