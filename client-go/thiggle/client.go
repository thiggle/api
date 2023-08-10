package thiggle

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	Prompt       string   `json:"prompt"`
	Patterns     []string `json:"patterns"`
	MaxNewTokens int      `json:"max_new_tokens,omitempty"`
}

type RegexCompletionResponse struct {
	Completion      string `json:"completion"`
	TokensGenerated int    `json:"tokens_generated"`
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
	routeRegexCompletion = "/v1/completion/regex"
	routeCategorize      = "/v1/categorize"
)

func (c *Client) RegexCompletion(ctx context.Context, req *RegexCompletionRequest) (*RegexCompletionResponse, error) {
	return post[RegexCompletionRequest, RegexCompletionResponse](ctx, c, routeRegexCompletion, req)
}

func (c *Client) Categorize(ctx context.Context, req *CategorizeRequest) (*CategorizeResponse, error) {
	return post[CategorizeRequest, CategorizeResponse](ctx, c, routeCategorize, req)
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
