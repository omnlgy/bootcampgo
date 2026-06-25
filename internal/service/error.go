package service

import "fmt"

type NotFoundError struct {
	Resource string
	ID       string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with id %v not found", e.Resource, e.ID)
}

type ConflictError struct {
	Resource string
	Field    string
	value    any
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("%s with %s already exists", e.Resource, e.Field)
}

type ValidationError struct {
	Field   string
	Rule    string
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

// Bundle multiple validation errors
type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string {
	return "validation failed"
}
