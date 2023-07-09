package thiggle

import (
	"os"
	"testing"
)

func TestIntegrationCategorize(t *testing.T) {
	client := NewClient(os.Getenv("THIGGLE_API_KEY"))
	client.BaseURL = "http://localhost:8080"

	req := CategorizeRequest{
		Prompt: "What animals have four legs?",
		Categories: []string{
			"dog",
			"fish",
			"snake",
			"cow",
		},
		AllowMultipleClasses: NewBool(true),
	}

	resp, err := client.Categorize(req)
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if len(resp.Choices) != 2 {
		t.Fatal("expected 2 choices")
	}

}
