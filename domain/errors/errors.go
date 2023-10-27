package errors

import goErrors "errors"

const (
	NotFound        = "NotFound"
	NotFoundMessage = "Record Not Found"

	ValidationError        = "ValidationError"
	ValidationErrorMessage = "Validation error"

	ResourceAlreadyExists     = "ResourceAlreadyExists"
	AlreadyExistsErrorMessage = "Resource already exists"

	RepositoryError        = "RepositoryError"
	RepositoryErrorMessage = "Error in repository operation"

	NotAuthenticated             = "NotAuthenticated"
	NotAuthenticatedErrorMessage = "Not Authenticated"

	TokenGeneratorError        = "TokenGeneratorError"
	TokenGeneratorErrorMessage = "Error in token generation"

	NotAuthorized             = "NotAuthorized"
	NotAuthorizedErrorMessage = "Not authorized"

	UnknownError        = "UnknownError"
	UnknownErrorMessage = "Something went wrong"
)

type AppError interface {
	Error() string
}

type AppErrorImpl struct {
	Err    error
	Type   string
	Status int
}

func NewAppErrorImpl(err error, errType string, status int) AppError {
	return &AppErrorImpl{
		Err:    err,
		Type:   errType,
		Status: status,
	}
}

func (a *AppErrorImpl) Error() string {
	return a.Err.Error()
}

func NewAppErrorWithType(errType string) AppError {
	var err error
	var status int
	switch errType {
	case NotFound:
		err = goErrors.New(NotFoundMessage)
		status = 404
	case ValidationError:
		err = goErrors.New(ValidationErrorMessage)
		status = 400
	case ResourceAlreadyExists:
		err = goErrors.New(AlreadyExistsErrorMessage)
		status = 500
	case RepositoryError:
		err = goErrors.New(RepositoryErrorMessage)
		status = 500
	case NotAuthenticated:
		err = goErrors.New(NotAuthenticatedErrorMessage)
		status = 401
	case NotAuthorized:
		err = goErrors.New(NotAuthorizedErrorMessage)
		status = 403
	case TokenGeneratorError:
		err = goErrors.New(TokenGeneratorErrorMessage)
		status = 500
	default:
		err = goErrors.New(UnknownErrorMessage)
		status = 500
	}
	return NewAppErrorImpl(err, errType, status)
}
