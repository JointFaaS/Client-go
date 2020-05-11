package client

import (
	"errors"
	"net/http"
	"net/url"
)
// A Client implements the base client request and response handling
// used by all service clients.
type Client struct {
	hc *http.Client

	host string
}

// NewClient 
func NewClient(config *Config) (*Client, error) {
	if config == nil {
		return nil, errors.New("config is nil")
	}
	url, err := url.Parse(config.Host)
	if err != nil {
		return nil, err
	}
	client := Client{
		hc: http.DefaultClient,
		host: url.Host,
	}
	return &client, nil
}