package main

type Service interface {
	Do()
}

type MyService struct{}

func (s *MyService) Do() {}

func NewService() Service {
	return &MyService{}
}
