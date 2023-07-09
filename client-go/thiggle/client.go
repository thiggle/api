package thiggle

import (
	"bytes"
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

func (c *Client) Categorize(req CategorizeRequest) (*CategorizeResponse, error) {
	url := fmt.Sprintf("%s/v1/categorize", c.BaseURL)

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)
	httpReq.Header.Set("User-Agent", "thiggle-client-go/"+thiggleVersion)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	var categorizeResp CategorizeResponse
	err = json.Unmarshal(body, &categorizeResp)
	if err != nil {
		return nil, err
	}

	return &categorizeResp, nil
}
