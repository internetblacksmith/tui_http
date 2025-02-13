package models

import (
	"encoding/json"
	"time"
)

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
	PATCH  HTTPMethod = "PATCH"
	HEAD   HTTPMethod = "HEAD"
)

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Request struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Method    HTTPMethod        `json:"method"`
	URL       string            `json:"url"`
	Headers   []Header          `json:"headers"`
	Body      string            `json:"body"`
	Params    map[string]string `json:"params"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type Response struct {
	StatusCode int               `json:"status_code"`
	Status     string            `json:"status"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
	Size       int64             `json:"size"`
	Duration   time.Duration     `json:"duration"`
	Timestamp  time.Time         `json:"timestamp"`
}

func (r *Request) AddHeader(key, value string) {
	r.Headers = append(r.Headers, Header{Key: key, Value: value})
}

func (r *Request) ToJSON() (string, error) {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (r *Response) ToJSON() (string, error) {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
