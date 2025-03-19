package examples

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"google.golang.org/genai"
)

func TextGenTextOnlyPrompt() (*genai.GenerateContentResponse, error) {
	// [START text_gen_text_only_prompt]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	contents := []*genai.Content{
		genai.NewUserContentFromText("Write a story about a magic backpack."),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	printResponse(response)
	// [END text_gen_text_only_prompt]
	return response, err
}

func TextGenTextOnlyPromptStreaming() error {
	// [START text_gen_text_only_prompt_streaming]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	contents := []*genai.Content{
		genai.NewUserContentFromText("Write a story about a magic backpack."),
	}
	for response, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(response.Candidates[0].Content.Parts[0].Text)
	}
	// [END text_gen_text_only_prompt_streaming]
	return err
}

func TextGenMultimodalOneImagePrompt() (*genai.GenerateContentResponse, error) {
	// [START text_gen_multimodal_one_image_prompt]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	imagePath := filepath.Join(getMedia(), "organ.jpg")
	f, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	imagePart := &genai.Part{
		InlineData: &genai.Blob{
			Data:     data,
			MIMEType: "image/jpeg",
		},
	}
	contents := []*genai.Content{
		genai.NewUserContentFromText("Tell me about this instrument"),
		genai.NewUserContentFromParts([]*genai.Part{imagePart}),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	text, err := response.Text()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
	// [END text_gen_multimodal_one_image_prompt]
	return response, err
}

func TextGenMultimodalOneImagePromptStreaming() error {
	// [START text_gen_multimodal_one_image_prompt_streaming]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	imagePath := filepath.Join(getMedia(), "organ.jpg")
	f, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	imagePart := &genai.Part{
		InlineData: &genai.Blob{
			Data:     data,
			MIMEType: "image/jpeg",
		},
	}
	contents := []*genai.Content{
		genai.NewUserContentFromText("Tell me about this instrument"),
		genai.NewUserContentFromParts([]*genai.Part{imagePart}),
	}
	for response, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(response.Candidates[0].Content.Parts[0].Text)
	}
	// [END text_gen_multimodal_one_image_prompt_streaming]
	return err
}

func TextGenMultimodalMultiImagePrompt() (*genai.GenerateContentResponse, error) {
	// [START text_gen_multimodal_multi_image_prompt]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	organPath := filepath.Join(getMedia(), "organ.jpg")
	cajunPath := filepath.Join(getMedia(), "Cajun_instruments.jpg")
	fOrg, err := os.Open(organPath)
	if err != nil {
		log.Fatal(err)
	}
	dataOrg, err := io.ReadAll(fOrg)
	fOrg.Close()
	if err != nil {
		log.Fatal(err)
	}
	fCajun, err := os.Open(cajunPath)
	if err != nil {
		log.Fatal(err)
	}
	dataCajun, err := io.ReadAll(fCajun)
	fCajun.Close()
	if err != nil {
		log.Fatal(err)
	}
	parts := []*genai.Part{
		{InlineData: &genai.Blob{Data: dataOrg, MIMEType: "image/jpeg"}},
		{InlineData: &genai.Blob{Data: dataCajun, MIMEType: "image/jpeg"}},
	}
	contents := []*genai.Content{
		genai.NewUserContentFromText(
			"What is the difference between both of these instruments?",
		),
		genai.NewUserContentFromParts(parts),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	text, err := response.Text()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
	// [END text_gen_multimodal_multi_image_prompt]
	return response, err
}

func TextGenMultimodalMultiImagePromptStreaming() error {
	// [START text_gen_multimodal_multi_image_prompt_streaming]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	organPath := filepath.Join(getMedia(), "organ.jpg")
	cajunPath := filepath.Join(getMedia(), "Cajun_instruments.jpg")
	fOrg, err := os.Open(organPath)
	if err != nil {
		log.Fatal(err)
	}
	dataOrg, err := io.ReadAll(fOrg)
	fOrg.Close()
	if err != nil {
		log.Fatal(err)
	}
	fCajun, err := os.Open(cajunPath)
	if err != nil {
		log.Fatal(err)
	}
	dataCajun, err := io.ReadAll(fCajun)
	fCajun.Close()
	if err != nil {
		log.Fatal(err)
	}
	parts := []*genai.Part{
		{InlineData: &genai.Blob{Data: dataOrg, MIMEType: "image/jpeg"}},
		{InlineData: &genai.Blob{Data: dataCajun, MIMEType: "image/jpeg"}},
	}
	contents := []*genai.Content{
		genai.NewUserContentFromText(
			"What is the difference between both of these instruments?",
		),
		genai.NewUserContentFromParts(parts),
	}
	for result, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Candidates[0].Content.Parts[0].Text)
	}
	// [END text_gen_multimodal_multi_image_prompt_streaming]
	return err
}

func TextGenMultimodalAudio() (*genai.GenerateContentResponse, error) {
	// [START text_gen_multimodal_audio]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	prompt := "Give me a summary of this audio file."
	audioPath := filepath.Join(getMedia(), "sample.mp3")
	f, err := os.Open(audioPath)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	audioPart := &genai.Part{
		InlineData: &genai.Blob{Data: data, MIMEType: "audio/mpeg"},
	}
	contents := []*genai.Content{
		genai.NewUserContentFromText(prompt),
		genai.NewUserContentFromParts([]*genai.Part{audioPart}),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	text, err := response.Text()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
	// [END text_gen_multimodal_audio]
	return response, err
}

func TextGenMultimodalAudioStreaming() error {
	// [START text_gen_multimodal_audio_streaming]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	prompt := "Give me a summary of this audio file."
	audioPath := filepath.Join(getMedia(), "sample.mp3")
	f, err := os.Open(audioPath)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	audioPart := &genai.Part{
		InlineData: &genai.Blob{Data: data, MIMEType: "audio/mpeg"},
	}
	contents := []*genai.Content{
		genai.NewUserContentFromText(prompt),
		genai.NewUserContentFromParts([]*genai.Part{audioPart}),
	}
	for result, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Candidates[0].Content.Parts[0].Text)
	}
	// [END text_gen_multimodal_audio_streaming]
	return err
}

func TextGenMultimodalVideoPrompt() (*genai.GenerateContentResponse, error) {
	// [START text_gen_multimodal_video_prompt]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	videoPath := filepath.Join(getMedia(), "Big_Buck_Bunny.mp4")
	f, err := os.Open(videoPath)
	if err != nil {
		log.Fatal(fmt.Errorf("error opening file: %v", err))
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(fmt.Errorf("error reading file: %v", err))
	}
	parts := []*genai.Part{
		{Text: "Describe this video clip"},
		{InlineData: &genai.Blob{Data: data, MIMEType: "video/mp4"}},
	}
	contents := []*genai.Content{
		genai.NewUserContentFromParts(parts),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	text, err := response.Text()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response.text=%s\n", text)
	// [END text_gen_multimodal_video_prompt]
	return response, err
}

func TextGenMultimodalVideoPromptStreaming() error {
	// [START text_gen_multimodal_video_prompt_streaming]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	videoPath := filepath.Join(getMedia(), "Big_Buck_Bunny.mp4")
	f, err := os.Open(videoPath)
	if err != nil {
		log.Fatal(fmt.Errorf("error opening file: %v", err))
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(fmt.Errorf("error reading file: %v", err))
	}
	parts := []*genai.Part{
		{Text: "Describe this video clip"},
		{InlineData: &genai.Blob{Data: data, MIMEType: "video/mp4"}},
	}
	contents := []*genai.Content{
		genai.NewUserContentFromParts(parts),
	}
	for result, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Candidates[0].Content.Parts[0].Text)
	}
	// [END text_gen_multimodal_video_prompt_streaming]
	return err
}

func TextGenMultimodalPdf() (*genai.GenerateContentResponse, error) {
	// [START text_gen_multimodal_pdf]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	pdfPath := filepath.Join(getMedia(), "test.pdf")
	f, err := os.Open(pdfPath)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	pdfPart := &genai.Part{
		InlineData: &genai.Blob{Data: data, MIMEType: "application/pdf"},
	}
	contents := []*genai.Content{
		genai.NewUserContentFromText("Give me a summary of this document:"),
		genai.NewUserContentFromParts([]*genai.Part{pdfPart}),
	}
	response, err := client.Models.GenerateContent(ctx, "gemini-2.0-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}
	text, err := response.Text()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response.text=%s\n", text)
	// [END text_gen_multimodal_pdf]
	return response, err
}

func TextGenMultimodalPdfStreaming() error {
	// [START text_gen_multimodal_pdf_streaming]
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}
	pdfPath := filepath.Join(getMedia(), "test.pdf")
	f, err := os.Open(pdfPath)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(f)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	pdfPart := &genai.Part{
		InlineData: &genai.Blob{Data: data, MIMEType: "application/pdf"},
	}
	contents := []*genai.Content{
		genai.NewUserContentFromText("Give me a summary of this document:"),
		genai.NewUserContentFromParts([]*genai.Part{pdfPart}),
	}
	for result, err := range client.Models.GenerateContentStream(
		ctx,
		"gemini-2.0-flash",
		contents,
		nil,
	) {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(result.Candidates[0].Content.Parts[0].Text)
	}
	// [END text_gen_multimodal_pdf_streaming]
	return err
}
