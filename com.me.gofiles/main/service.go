package main

import (
	"errors"
	"strings"
)

type StringService interface {
	Status() string
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct {}
var ErrEmpty = errors.New("Empty error")
var Status = "Ok"

func (stringService)Uppercase(s string) (string,error)  {
	if s == "" {
		return "",ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int  {
	return len(s)
}

func (stringService) Status() string {
	return Status
}

type ServiceMiddleware func(StringService) StringService

