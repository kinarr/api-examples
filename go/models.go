package examples

import (
	"context"
	"fmt"
	"os"
	"log"

	"google.golang.org/genai"
)

func ModelsList() error {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}


	// Retrieve the list of models.
	models, err := client.Models.List(ctx, &genai.ListModelsConfig{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("List of models that support generateContent:")
	for _, m := range models.Items {
		for _, action := range m.SupportedActions {
			if action == "generateContent" {
				fmt.Println(m.Name)
				break
			}
		}
	}

	fmt.Println("\nList of models that support embedContent:")
	for _, m := range models.Items {
		for _, action := range m.SupportedActions {
			if action == "embedContent" {
				fmt.Println(m.Name)
				break
			}
		}
	}
	return err
}

func ModelsGet() error {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	modelInfo, err := client.Models.Get(ctx, "gemini-2.0-flash", nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(modelInfo)
	return err
}
