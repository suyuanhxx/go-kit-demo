package main

import (
	"fmt"
)

type UserService interface {
	GetUserByUserName(string) string
}

type userService struct{}

func (userService) GetUserByUserName(s string) string {
	fmt.Println("GetUserByUserName param is", s)
	return s + " demo"
}

// ServiceMiddleware is a chainable behavior modifier for StringService.
type ServiceMiddleware func(UserService) UserService
