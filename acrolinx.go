package acrolinx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	userAgent       = "go-acrolinx"
	headerSignature = "X-Acrolinx-Client"
	headerToken     = "X-Acrolinx-Auth"
	headerLocale    = "X-Acrolinx-Client-Locale"
)

type Client struct {
	// Signature identifies this client
	signature string

	platformURL *url.URL

	accessToken string

	client *http.Client

	// Services for different parts of the API
	Checking *CheckingService
}

func NewClient(signature string, urlStr string, options ...ClientOptionFunc) (*Client, error) {
	platformURL, err := makePlatformURL(urlStr)
	if err != nil {
		return nil, fmt.Errorf("Error creating new client: %w", err)
	}

	// build HTTP NewClient

	client := &Client{
		signature:   signature,
		platformURL: platformURL,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	for _, fn := range options {
		if fn == nil {
			continue
		}

		if err := fn(client); err != nil {
			return nil, err
		}
	}

	client.Checking = &CheckingService{client}

	return client, nil
}

func (c *Client) SignIn(username string, password string) error {
	creds := Credentials{username, password}
	path := "dashboard/api/signin/authenticate"

	req, err := c.newRequest(http.MethodPost, path, creds)
	if err != nil {
		return fmt.Errorf("Error signing in, could not prepare request: %w", err)
	}

	var token accessToken
	err = c.do(req, token)
	if err != nil {
		return fmt.Errorf("Error signing in, could not prepare request: %w", err)
	}
	c.accessToken = token.AccessToken

	return nil
}

func (c *Client) newRequest(method, path string, creds interface{}) (*http.Request, error) {
	u := *c.platformURL
	u.Path = c.platformURL.Path + path

	jsonBody, err := json.Marshal(creds)
	if err != nil {
		return nil, fmt.Errorf("Error encoding JSON: %w", err)
	}
	body := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, fmt.Errorf("Error creating new request: %w", err)
	}
	req.Header.Set(headerSignature, c.signature)
	if c.accessToken != "" {
		req.Header.Set(headerToken, c.accessToken)
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) error {
	res, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("Error submitting request: %w", err)
	}

	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return fmt.Errorf("Error decoding JSON response: %w", err)
	}
	return nil
}

func makePlatformURL(urlStr string) (*url.URL, error) {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	platformURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("Error parsing platform URL: %w", err)
	}

	return platformURL, nil
}

func (c *Client) setToken(token string) {
	c.accessToken = token
}

type Links = map[string]string

type Response struct {
	Data     interface{}   `json:"data,omitempty"`
	Links    Links         `json:"links,omitempty"`
	Progress *Progress     `json:"progress,omitempty"`
	Error    *RequestError `json:"error,omitempty"`
}

type Progress struct {
	Percent    int    `json:"percent"`
	Message    string `json:"message"`
	RetryAfter int    `json:"retryAfter"`
}

type RequestError struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

func (r *RequestError) Error() string {
	return r.Detail
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type accessToken struct {
	AccessToken string `json:"accessToken"`
}
