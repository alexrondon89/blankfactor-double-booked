package server

import "github.com/gofiber/fiber/v2"

type Handler interface {
	InsertEvent(c *fiber.Ctx) error
	GetCollisionInEvents(c *fiber.Ctx) error
}

type Server struct {
	app     *fiber.App
	handler Handler
}

func New(handler Handler) *Server {
	return &Server{
		app:     fiber.New(),
		handler: handler,
	}
}

func (s *Server) AddPostEvent() *Server {
	s.app.Post("/addEvent", s.handler.InsertEvent)
	return s
}

func (s *Server) AddGetCollisionInEvents() *Server {
	s.app.Get("/collisions", s.handler.GetCollisionInEvents)
	return s
}

func (s *Server) Start() {
	defer s.app.Shutdown()
	s.app.Listen(":3000")
}
