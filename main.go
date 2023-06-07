package main

import (
	"log"

	"main/chatbot/api"
	controller "main/chatbot/controllers"
	"main/chatbot/server"
	service "main/chatbot/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	// Bağımlılıkları oluşturma
	apiClient := api.NewOpenAIAPI("API-KEY", "MODEL")
	messageService := service.NewMessageService(apiClient)
	chatController := controller.NewChatController(messageService)

	// Sunucuyu başlatma
	server.Start(app, chatController)
	log.Fatal(app.Listen(":3000"))
}
