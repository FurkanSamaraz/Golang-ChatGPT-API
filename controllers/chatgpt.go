package controller

import (
	cache "chat-bot/cache"
	service "chat-bot/services"
	"chat-bot/structures"

	"github.com/gofiber/fiber/v2"
)

type ChatController struct {
	messageService *service.MessageService
	cache          *cache.Cache
}

func NewChatController(messageService *service.MessageService, cache *cache.Cache) *ChatController {
	return &ChatController{
		messageService: messageService,
		cache:          cache,
	}
}

// ApiMessage, "/api/message" yoluna yapılan POST isteğini yönetir
func (c *ChatController) ApiMessage(ctx *fiber.Ctx) error {

	var req structures.Request
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}
	response := c.messageService.GetChatResponse(req.Message, c.cache)

	return ctx.JSON(fiber.Map{"response": response})
}
