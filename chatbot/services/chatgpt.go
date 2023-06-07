package service

import (
	"fmt"
	"main/chatbot/api"
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
func (s *MessageService) GetChatResponse(message string) string {
	response, err := s.apiClient.SendCompletionRequest(message)
	if err != nil {

		fmt.Println("API Error:", err)
		return "Error: " + err.Error()
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content
	}

	return "Error: No response"
}
