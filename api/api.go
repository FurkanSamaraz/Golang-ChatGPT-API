package api

import (
	"bytes"
	"chat-bot/structures"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type OpenAIAPI struct {
	apiKey string
	model  string
}

func NewOpenAIAPI(apiKey, model string) *OpenAIAPI {
	return &OpenAIAPI{
		apiKey: apiKey,
		model:  model,
	}
}

func (a *OpenAIAPI) SendCompletionRequest(userInput string) (*structures.APIResponse, error) {

	conversation := []map[string]interface{}{
		{"role": "system", "content": "You are a helpful assistant that generates questions."},
	}
	conversation = append(conversation, map[string]interface{}{"role": "user", "content": userInput})
	conversation = append(conversation, map[string]interface{}{"role": "system", "content": "Generate 5 questions based on the articles."})

	data := map[string]interface{}{
		"model":    a.model,
		"messages": conversation,
	}

	requestBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	completion_url := os.Getenv("COMPLETION_URL")
	req, err := http.NewRequest("POST", completion_url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.apiKey))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response structures.APIResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if len(response.Choices) == 0 {
		return nil, &structures.CustomError{
			Type:    "Invalid Data",
			Message: "len(response.Choices) == 0",
		}
	}

	return &response, nil
}
