package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
}

func GetUser(name string, isStaticError bool, isMatchRequired bool) (*User, error) {
	if !isMatchRequired && isStaticError {
		return nil, errors.New("user not found")
	}
	if !isMatchRequired && !isStaticError {
		return nil, fmt.Errorf("user with name %q not found", name)
	}
	if isMatchRequired && isStaticError {
		return nil, ErrNotFound
	}
	if isMatchRequired && !isStaticError {
		return nil, &UserNotFoundError{Name: name}
	}
	return &User{Name: name}, nil
}
