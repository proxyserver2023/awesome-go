package main

import (
	"context"
	"fmt"
	proto "github.com/alamin-mahamud/awesome-go/go-micro-example/proto"
	"github.com/micro/cli"
	micro "github.com/micro/go-micro"
)

// Greeter - ...
type Greeter struct{}

// Hello - Method
func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {

	// 1. create a new service. Optionally include some options here.
	service := micro.NewService(

		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "HelloWorld",
		}),

		micro.Flags(cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)

	// 2. init will parse the command line flags
	service.Init()

	// 3. register the handler
	proto.RegisterGreeterHandler(service.Server(), &Greeter{})

	// 4. Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
