package mux

import "errors"

var (
	// ErrMethodMismatch is returned when the method in the request does not match with
	// the method defined against the route.
	ErrMethodMismatch = errors.New("Method is not allowed.")

	// ErrNotFound -> if no route matches.
	ErrNotFound = errors.New("No matching route was found.")
)
