package controller

import (
	service "main/chatbot/services"
	"main/chatbot/structures"

	"github.com/gofiber/fiber/v2"
)

type ChatController struct {
	messageService *service.MessageService
}

func NewChatController(messageService *service.MessageService) *ChatController {
	return &ChatController{
		messageService: messageService,
	}
}

// ApiMessage, "/api/message" yoluna yapılan POST isteğini yönetir
func (c *ChatController) ApiMessage(ctx *fiber.Ctx) error {

	var req structures.Request
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	response := c.messageService.GetChatResponse(req.Message)

	return ctx.JSON(fiber.Map{"response": response})
}
