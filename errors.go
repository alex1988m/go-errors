package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotFound = errors.New("user not found")
)

type UserNotFoundError struct {
	Name string
}

func (r *UserNotFoundError) Error() string {
	return fmt.Sprintf("user with name %q not found", r.Name)
}

func logError(err error, logger *log.Logger) {
	var notFound *UserNotFoundError
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			logger.Println("error is matched as ErrNotFound")
		} else if errors.As(err, &notFound) {
			logger.Println("error is matched as UserNotFoundError")
		} else {
			logger.Println("error is not matched as ErrNotFound or UserNotFoundError")
		}
		logger.Printf("get user: %v", err)
	}
}
