package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const baseURL = "https://api.telegram.org/bot"

type Client struct {
	token  string
	http   *http.Client
	prefix string
}

func New(token string) *Client {
	return &Client{
		token:  token,
		http:   &http.Client{},
		prefix: baseURL + token + "/",
	}
}

// Call makes a GET request with url-encoded params and returns the raw result.
func (c *Client) Call(method string, params url.Values) (json.RawMessage, error) {
	u := c.prefix + method
	if len(params) > 0 {
		u += "?" + params.Encode()
	}
	resp, err := c.http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("http get: %w", err)
	}
	defer resp.Body.Close()
	return parseResponse(resp.Body)
}

// CallJSON makes a POST request with a JSON body and returns the raw result.
func (c *Client) CallJSON(method string, payload any) (json.RawMessage, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	resp, err := c.http.Post(c.prefix+method, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("http post: %w", err)
	}
	defer resp.Body.Close()
	return parseResponse(resp.Body)
}

// Upload sends a multipart form request (used by upload.go).
func (c *Client) Upload(method string, contentType string, body io.Reader) (json.RawMessage, error) {
	resp, err := c.http.Post(c.prefix+method, contentType, body)
	if err != nil {
		return nil, fmt.Errorf("http post: %w", err)
	}
	defer resp.Body.Close()
	return parseResponse(resp.Body)
}

func parseResponse(r io.Reader) (json.RawMessage, error) {
	var apiResp APIResponse
	if err := json.NewDecoder(r).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	if !apiResp.OK {
		return nil, fmt.Errorf("api error %d: %s", apiResp.ErrorCode, apiResp.Description)
	}
	return apiResp.Result, nil
}
