package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"restless/pkg/models"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) ExecuteRequest(req *models.Request) (*models.Response, error) {
	start := time.Now()

	fullURL, err := c.buildURL(req.URL, req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to build URL: %w", err)
	}

	var body io.Reader
	if req.Body != "" {
		body = strings.NewReader(req.Body)
	}

	httpReq, err := http.NewRequest(string(req.Method), fullURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for _, header := range req.Headers {
		if header.Key != "" && header.Value != "" {
			httpReq.Header.Set(header.Key, header.Value)
		}
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	headers := make(map[string]string)
	for key, values := range resp.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}

	duration := time.Since(start)

	return &models.Response{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Headers:    headers,
		Body:       string(respBody),
		Size:       int64(len(respBody)),
		Duration:   duration,
		Timestamp:  time.Now(),
	}, nil
}

func (c *Client) buildURL(baseURL string, params map[string]string) (string, error) {
	if len(params) == 0 {
		return baseURL, nil
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	query := u.Query()
	for key, value := range params {
		if key != "" && value != "" {
			query.Set(key, value)
		}
	}

	u.RawQuery = query.Encode()
	return u.String(), nil
}

func (c *Client) SetTimeout(timeout time.Duration) {
	c.httpClient.Timeout = timeout
}
