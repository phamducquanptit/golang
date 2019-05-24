package errno

var (
	// Common errors
	OK = &Errno{Code: 0, Message: "OK"}

	// Error connect to Server
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBadRequest       = &Errno{Code: 400, Message: "Bad request"}
	ErrNotFound         = &Errno{Code: 404, Message: "Error not found"}
	ErrInternalServer   = &Errno{Code: 500, Message: "Error Internal server"}

	ErrBind       = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrValidation = &Errno{Code: 20001, Message: "Validation failed"}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error during registration JSON web token"}

	// User errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "User not found"}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "Token invalid"}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "Password incorrect"}

	// Todo errors
	ErrUpdateTodo = &Errno{Code: 20200, Message: "Error within update to do"}
)
