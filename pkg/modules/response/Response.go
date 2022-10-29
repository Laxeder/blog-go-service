package response

import (
	"github.com/gofiber/fiber/v2"
)

func New(ctx *fiber.Ctx) *Response {
	return &Response{Ctx: ctx}
}

func Success(status int, data any) *Response {
	return &Response{Status: status, Data: data}
}

func Error(status int, code string, message string) *Response {
	return &Response{Status: status, Code: code, Message: message}
}

type Response struct {
	Ctx     *fiber.Ctx
	Message string
	Code    string
	Status  int
	Data    any
}

func (r *Response) Result(status int, data any) *Response {
	return &Response{Status: status, Data: data, Ctx: r.Ctx}
}

func (r *Response) Error(status int, code string, message string) *Response {
	return &Response{Status: status, Code: code, Message: message, Ctx: r.Ctx}
}

func (r *Response) ErrorDefault(code string) *Response {
	return r.Error(500, "BGS000", "Serviço indisponível no momento. Tente novamente mais tarde.")
}

func (r *Response) Json(res *Response) error {
	r.Status = r.HandleStatus(res.Status)
	r.Data = r.HandleMessage(res)

	return r.Ctx.Status(r.Status).JSON(r.Data)
}

func (r *Response) HandleStatus(status int) int {

	if status == 201 {
		return fiber.StatusCreated
	}

	if status == 204 {
		return fiber.StatusNoContent
	}

	if status == 401 {
		return fiber.StatusUnauthorized
	}

	if status == 0 || status < 100 {
		return fiber.StatusInternalServerError
	}

	if status < 200 {
		return fiber.StatusContinue
	}

	if status < 300 {
		return fiber.StatusOK
	}

	if status < 400 {
		return fiber.StatusMultipleChoices
	}

	if status < 500 {
		return fiber.StatusBadRequest
	}

	if status >= 500 {
		return fiber.StatusInternalServerError
	}

	return fiber.StatusInternalServerError
}

func (r *Response) HandleMessage(res *Response) any {
	if res.Status == fiber.StatusOK {
		if res.Data != nil {
			return res.Data
		}
	}

	if res.Status == fiber.StatusCreated || res.Status == fiber.StatusNoContent {
		return nil
	}

	if res.Status >= 400 {
		return fiber.Map{"code": res.Code, "message": res.Message}
	}

	return nil
}
