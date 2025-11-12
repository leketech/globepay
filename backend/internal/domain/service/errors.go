package service

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// ConflictError represents a conflict error
type ConflictError struct {
	Message string
}

func (e *ConflictError) Error() string {
	return e.Message
}

// AuthenticationError represents an authentication error
type AuthenticationError struct {
	Message string
}

func (e *AuthenticationError) Error() string {
	return e.Message
}

// NotFoundError represents a not found error
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

// InsufficientFundsError represents an insufficient funds error
type InsufficientFundsError struct {
	Message string
}

func (e *InsufficientFundsError) Error() string {
	return e.Message
}
