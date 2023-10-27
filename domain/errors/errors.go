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
	Err  error
	Type string
}

func NewAppErrorImpl(err error, errType string) AppError {
	return &AppErrorImpl{
		Err:  err,
		Type: errType,
	}
}

func (a *AppErrorImpl) Error() string {
	return a.Err.Error()
}

func NewAppErrorWithType(errType string) AppError {
	var err error
	switch errType {
	case NotFound:
		err = goErrors.New(NotFoundMessage)
	case ValidationError:
		err = goErrors.New(ValidationErrorMessage)
	case ResourceAlreadyExists:
		err = goErrors.New(AlreadyExistsErrorMessage)
	case RepositoryError:
		err = goErrors.New(RepositoryErrorMessage)
	case NotAuthenticated:
		err = goErrors.New(NotAuthenticatedErrorMessage)
	case NotAuthorized:
		err = goErrors.New(NotAuthorizedErrorMessage)
	case TokenGeneratorError:
		err = goErrors.New(TokenGeneratorErrorMessage)
	default:
		err = goErrors.New(UnknownErrorMessage)
	}
	return NewAppErrorImpl(err, errType)
}
