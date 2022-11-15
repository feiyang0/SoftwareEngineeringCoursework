package errno

import "fmt"

var (
	// OK represents a successful request.
	OK = &Errno{Code: 0, Message: "OK"}

	// InternalServerError represents all unknown server-side errors.
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}

	// ErrBind represents a failed parameter binding.
	ErrBind = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// ErrPageNotFound represents a route not matched error.
	ErrPageNotFound = &Errno{Code: 10003, Message: "Page not found."}

	// ErrValidation represents all validation failed errors.
	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}

	// ErrDatabase represents a database error.
	ErrDatabase = &Errno{Code: 20002, Message: "Database error."}

	// ErrToken represents a error when signing JWT token.
	ErrToken = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// ErrEncrypt represents a encrypting error.
	ErrEncrypt = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}

	// ErrUserNotFound represents the user not found.
	ErrUserNotFound = &Errno{Code: 20102, Message: "User was not found."}

	// ErrUserAlreadyExist represents the user already exist.
	ErrUserAlreadyExist = &Errno{Code: 20103, Message: "User already exist."}

	// ErrTokenInvalid represents the token format is wrong.
	ErrTokenInvalid = &Errno{Code: 20104, Message: "Token was invalid."}

	// ErrPasswordIncorrect represents the password is incorrect.
	ErrPasswordIncorrect = &Errno{Code: 20105, Message: "Password was incorrect."}

	// ErrCaptchaIncorrect represents the Captcha is incorrect.
	ErrCaptchaIncorrect = &Errno{Code: 20106, Message: "Captcha was incorrect."}

	// ErrPostNotFound represents  the post not found.
	ErrPostNotFound = &Errno{Code: 20201, Message: "Post was not found."}

	// ErrPostAlreadyExist represents the post already exist.
	ErrPostAlreadyExist = &Errno{Code: 20202, Message: "Post already exist."}
)

// Errno defines a new error type used by goserver.
type Errno struct {
	Code    int
	Message string
}

// Error implement the `Error` method in error interface.
func (err Errno) Error() string {
	return err.Message
}

// Err represents an error.
type Err struct {
	Code    int
	Message string
	Err     error
}

// New create a new `Err` error.
func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

// Add add message to `Err` error.
func (err *Err) Add(message string) error {
	err.Message += " " + message

	return err
}

// Addf add a formated message to `Err` error.
func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)

	return err
}

// Error return error message in string format.
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// IsErrUserNotFound return true if the `err` is a `ErrUserNotFound` type error.
func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)

	return code == ErrUserNotFound.Code
}

// DecodeErr decode an err message.
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	//nolint: errorlint
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
