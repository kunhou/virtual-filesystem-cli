package errors

const (
	ResourceNotFoundTempl      = "The %s doesn't exist."
	ResourceAlreadyExistsTempl = "The %s has already existed."
)

func ResourceNotFound(value ...string) error {
	return &Error{
		messageTemplate: ResourceNotFoundTempl,
		messageArgs:     value,
	}
}

func ResourceAlreadyExists(value ...string) error {
	return &Error{
		messageTemplate: ResourceAlreadyExistsTempl,
		messageArgs:     value,
	}
}
