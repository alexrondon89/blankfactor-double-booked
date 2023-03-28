package main

import (
	"blankfactor/cmd/server"
	handler2 "blankfactor/cmd/server/handler"
	"blankfactor/internal"
	"blankfactor/internal/Repository"
	validator2 "blankfactor/internal/validator"
	checker2 "blankfactor/internal/validator/checker"
)

func main() {
	checker := checker2.New()
	validator := validator2.New(checker)
	tempDb := Repository.New()
	service := internal.New(validator, tempDb)
	handler := handler2.New(service)
	server.New(handler).
		AddGetCollisionInEvents().
		AddPostEvent().
		Start()
}
