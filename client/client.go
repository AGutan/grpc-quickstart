package main

import (
	"context"
	foobar "github.com/alexgtn/grpc-quickstart/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := foobar.NewFooBarServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	foo := &foobar.Foo{
		Foo: "foo",
	}
	r, err := c.Execute(ctx, foo)
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Got response %v", r)
}
