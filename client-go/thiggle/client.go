package thiggle

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const (
	thiggleVersion = "0.0.2"
	defaultBaseURL = "https://api.thiggle.com"
)

type CategorizeRequest struct {
	Prompt               string   `json:"prompt"`
	Categories           []string `json:"categories"`
	AllowNullCategory    Bool     `json:"allow_null_category,omitempty"`
	AllowMultipleClasses Bool     `json:"allow_multiple_classes,omitempty"`
}

type CategorizeResponse struct {
	Choices []string `json:"choices"`
}

type RegexCompletionRequest struct {
	Prompt            string  `json:"prompt"`
	Pattern           string  `json:"pattern"`
	MaxNewTokens      int     `json:"max_new_tokens,omitempty"`
	TopP              float64 `json:"top_p,omitempty"`
	TopK              int     `json:"top_k,omitempty"`
	Temperature       float64 `json:"temperature,omitempty"`
	StopAfterMatch    Bool    `json:"stop_after_match,omitempty"`
	RepetitionPenalty float64 `json:"repetition_penalty,omitempty"`
}

type RegexCompletionResponse struct {
	Completion      string `json:"completion"`
	TokensGenerated int    `json:"tokens_generated"`
	StopReason      string `json:"stop_reason"`
}

type ContextFreeCompletionRequest struct {
	Prompt            string  `json:"prompt"`
	Grammar           string  `json:"grammar"`
	MaxNewTokens      int     `json:"max_new_tokens,omitempty"`
	TopP              float64 `json:"top_p,omitempty"`
	TopK              int     `json:"top_k,omitempty"`
	Temperature       float64 `json:"temperature,omitempty"`
	StopAfterMatch    Bool    `json:"stop_after_match,omitempty"`
	RepetitionPenalty float64 `json:"repetition_penalty,omitempty"`
}

type ContextFreeCompletionResponse struct {
	Completion      string `json:"completion"`
	TokensGenerated int    `json:"tokens_generated"`
	StopReason      string `json:"stop_reason"`
}

type TypedCompletionRequest struct {
	UserID           uuid.UUID `json:"user_id"`
	Prompt           string    `json:"prompt"`
	MaxTokens        *uint     `json:"max_tokens,omitempty"`
	Temperature      *float64  `json:"temperature,omitempty"`
	TopP             *float64  `json:"top_p,omitempty"`
	Stop             []string  `json:"stop,omitempty"`
	PresencePenalty  *float64  `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64  `json:"frequency_penalty,omitempty"`

	Schema   string `json:"schema"`
	TypeName string `json:"type_name"`
}

type TypedCompletionResponse struct {
	ID      uuid.UUID           `json:"id"`
	Object  string              `json:"object"`
	Created time.Time           `json:"created"`
	Choices []*CompletionChoice `json:"choices"`
	Usage   *Usage              `json:"usage"`
}

type Client struct {
	BaseURL string
	APIKey  string
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: defaultBaseURL,
		APIKey:  apiKey,
	}
}

const (
	routeRegexCompletion       = "/v1/completion/regex"
	routeTypedCompletion       = "/v1/completion/typed"
	routeContextFreeCompletion = "/v1/completion/cfg"
	routeCompletion            = "/v1/completion"
	routeCategorize            = "/v1/categorize"
)

func (c *Client) Completion(ctx context.Context, req *CompletionRequest) (*CompletionResponse, error) {
	return post[CompletionRequest, CompletionResponse](ctx, c, routeCompletion, req)
}

func (c *Client) RegexCompletion(ctx context.Context, req *RegexCompletionRequest) (*RegexCompletionResponse, error) {
	return post[RegexCompletionRequest, RegexCompletionResponse](ctx, c, routeRegexCompletion, req)
}

func (c *Client) ContextFreeCompletion(ctx context.Context, req *ContextFreeCompletionRequest) (*ContextFreeCompletionResponse, error) {
	return post[ContextFreeCompletionRequest, ContextFreeCompletionResponse](ctx, c, routeContextFreeCompletion, req)
}

func (c *Client) Categorize(ctx context.Context, req *CategorizeRequest) (*CategorizeResponse, error) {
	return post[CategorizeRequest, CategorizeResponse](ctx, c, routeCategorize, req)
}

func (c *Client) TypedCompletion(ctx context.Context, req *TypedCompletionRequest) (*TypedCompletionResponse, error) {
	return post[TypedCompletionRequest, TypedCompletionResponse](ctx, c, routeTypedCompletion, req)
}

func post[Request any, Response any](ctx context.Context, c *Client, url string, req *Request) (*Response, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", c.BaseURL+url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	httpReq = httpReq.WithContext(ctx)

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)
	httpReq.Header.Set("User-Agent", "thiggle-client-go/"+thiggleVersion)

	client := &http.Client{}

	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	body, err = io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	var resp Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
