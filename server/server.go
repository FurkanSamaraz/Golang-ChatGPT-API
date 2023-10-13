package server

import (
	controller "chat-bot/controllers"

	"github.com/gofiber/fiber/v2"
)

// Start, sunucuyu başlatır ve yolları yönlendirir
func Start(app *fiber.App, chatController *controller.ChatController) {
	app.Post("/api/message", chatController.ApiMessage)
}
