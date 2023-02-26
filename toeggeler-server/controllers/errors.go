package controllers

const (
	// 4xx
	ErrAuthenticate   = "authen-failed"
	ErrUserNotFound   = "user-not-found"
	ErrUserExists     = "user-already-exists"
	ErrMailExists     = "email-already-exists"
	ErrInvalidEmail   = "email-invalid"
	ErrInvalidPayload = "payload-invalid"
	ErrInvalidEvent   = "event-invalid"

	// 5xx
	ErrGenericError = "internal-error"
)
