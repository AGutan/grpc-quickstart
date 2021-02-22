package service

import (
	"context"
	"github.com/alexgtn/grpc-quickstart/proto"
	"log"
)

type Service struct {
}

func (s *Service) Execute(_ context.Context, foo *foobar.Foo) (*foobar.Result, error) {
	log.Printf("Got Foo: %v", foo)
	res := &foobar.Result{
		Data: &foobar.Result_Bar{
			Bar: &foobar.Bar{
				Bar: "bar",
			},
		},
	}
	return res, nil
}

func NewService() *Service {
	return &Service{}
}
