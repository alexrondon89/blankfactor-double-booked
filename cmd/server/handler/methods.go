package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	AnalyzeEventCalendar() []map[string]interface{}
	InsertEventInCalendar([]float64) []float64
}
type Handler struct {
	service Service
}

type event struct {
	Event []float64 `json:"event"`
}

func New(service Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) InsertEvent(ctx *fiber.Ctx) error {
	obj := new(event)
	if err := ctx.BodyParser(obj); err != nil {
		return ctx.Status(500).JSON(err)
	}
	resp := h.service.InsertEventInCalendar(obj.Event)
	return ctx.Status(201).JSON(resp)

}

func (h Handler) GetCollisionInEvents(ctx *fiber.Ctx) error {
	resp := h.service.AnalyzeEventCalendar()
	return ctx.Status(200).JSON(resp)

}
