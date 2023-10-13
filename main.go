package main

import (
	"chat-bot/api"
	cache "chat-bot/cache"
	controller "chat-bot/controllers"
	"chat-bot/server"
	service "chat-bot/services"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api_key := os.Getenv("API_KEY")
	api_model := os.Getenv("MODEL")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	// Bağımlılıkları oluşturma
	apiClient := api.NewOpenAIAPI(api_key, api_model)
	messageService := service.NewMessageService(apiClient)
	cah := cache.NewCache()
	chatController := controller.NewChatController(messageService, cah)

	// Sunucuyu başlatma
	server.Start(app, chatController)
	log.Fatal(app.Listen(":3000"))
}
