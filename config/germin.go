package config

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"log"
	"os"
)

type GeminiConfig struct {
	APIKey string
}

func Aimodel() {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	// Initialize the chat
	cs := model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Hello, I have 2 dogs in my house."),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Great to meet you. What would you like to know?"),
			},
			Role: "model",
		},
	}

	resp, err := cs.SendMessage(ctx, genai.Text("How many paws are in my house?"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
