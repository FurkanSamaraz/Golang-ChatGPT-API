package service

import (
	"chat-bot/api"
	cache "chat-bot/cache"
	"fmt"
)

type MessageService struct {
	apiClient *api.OpenAIAPI
}

func NewMessageService(apiClient *api.OpenAIAPI) *MessageService {
	return &MessageService{
		apiClient: apiClient,
	}
}

// GetChatResponse, ChatGPT'den bir yanıt almak için API isteği yapar
func (s *MessageService) GetChatResponse(message string, cache *cache.Cache) string {
	// Önbellekte yanıtı ara
	cachedResponse := cache.Get(message)
	if cachedResponse != "" {
		return cachedResponse
	}

	response, err := s.apiClient.SendCompletionRequest(message)
	if err != nil {
		fmt.Println("API Error:", err)
		return "Error: " + err.Error()
	}

	if len(response.Choices) > 0 {
		result := response.Choices[0].Message.Content
		// Yanıtı önbelleğe al
		cache.Set(message, result)
		return result
	}

	return "Error: No response"
}
