package acrolinx

import (
	"bytes"
	"encoding/json"
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

func NewClient(signature string, urlStr string) (*Client, error) {
	platformURL, err := makePlatformURL(urlStr)
	if err != nil {
		return nil, err
	}

	// build HTTP NewClient

	client := &Client{
		signature:   signature,
		platformURL: platformURL,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	client.Checking = &CheckingService{client}

	return client, nil
}

func (c *Client) SignIn(username string, password string) error {
	creds := Credentials{username, password}
	path := "dashboard/api/signin/authenticate"

	req, err := c.newRequest(path, creds)
	if err != nil {
		return err
	}

	var token accessToken
	err = c.do(req, token)
	if err != nil {
		return err
	}
	c.accessToken = token.AccessToken

	return nil
}

func (c *Client) newRequest(path string, creds interface{}) (*http.Request, error) {
	u := *c.platformURL
	u.Path = c.platformURL.Path + path

	jsonBody, err := json.Marshal(creds)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, u.String(), body)
	if err != nil {
		return nil, err
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
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return err
	}
	return nil
}

func makePlatformURL(urlStr string) (*url.URL, error) {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}

	platformURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	return platformURL, nil
}

type Links = map[string]string

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Links Links       `json:"links,omitempty"`
	Error Error       `json:"error,omitempty"`
}

type Error struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type accessToken struct {
	AccessToken string `json:"accessToken"`
}
