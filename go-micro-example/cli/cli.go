package main

import (
	proto "git"
	micro "github.com/micro/go-micro"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	client := proto.NewGreeterService(
		"greeter",
		service.Client(),
	)
}
