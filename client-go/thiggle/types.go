package thiggle

import (
	"time"

	"github.com/google/uuid"
)

type CompletionRequest struct {
	Prompt           string         `json:"prompt"`
	Model            *StringOrSlice `json:"model"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	Temperature      float32        `json:"temperature,omitempty"`
	TopP             float32        `json:"top_p,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	PresencePenalty  float32        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32        `json:"frequency_penalty,omitempty"`
}

type CompletionResponse struct {
	ID      uuid.UUID          `json:"id"`
	Object  string             `json:"object"`
	Created time.Time          `json:"created"`
	Choices []CompletionChoice `json:"choices"`
	Usage   *Usage             `json:"usage"`
}

type Usage struct {
	PromptTokens     int   `json:"prompt_tokens"`
	PredictionTimeMs int64 `json:"prediction_time_ms"`
	CompletionTokens int   `json:"completion_tokens"`
	TotalTokens      int   `json:"total_tokens"`
}

type CompletionChoice struct {
	Text         string `json:"text"`
	Created      int    `json:"created"`
	Model        string `json:"model"`
	FinishReason string `json:"finish_reason"`
	Usage        *Usage `json:"usage"`
}
