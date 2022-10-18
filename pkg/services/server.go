package services

import (
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func (srv *Server) New() *fiber.App {
	srv.app = fiber.New()
	return srv.app
}

func (srv *Server) Listen(address string) error {
	return srv.app.Listen(address)
}

func (srv *Server) Routes(routes func(*fiber.App)) *fiber.App {
	defer routes(srv.app)
	return srv.app
}

func (srv *Server) Middlewares(middlewares func(*fiber.App)) *fiber.App {
	defer middlewares(srv.app)
	return srv.app
}
