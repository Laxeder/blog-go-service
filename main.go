package main

import (
	"blog/pkg/configs"
	"blog/pkg/routes"
	"blog/pkg/services"
	"fmt"
)

func main() {
	configs.Load()

	srv := &services.Server{}
	srv.New()
	srv.Routes(routes.ApiV1)

	err := srv.Listen(configs.GetAPIPort())

	if err != nil {
		fmt.Printf("Server down. Error: %v", err)
	}
}
