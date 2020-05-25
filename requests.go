package main

import (
	"log"

	"github.com/levigross/grequests"
)

// Client is the basic class for performing http requests
type Client struct {
	Endpoint string
}

var defaultClient *Client

// NewClient returns a new empty client
func NewClient() *Client {
	if defaultClient != nil {
		return defaultClient
	}
	defaultClient = &Client{}
	return defaultClient
}

// Post sends a simple Post requests to specific endpoint
func (c *Client) Post(endpoint string, params map[string]string) *grequests.Response {
	url := c.Endpoint + endpoint
	resp, err := grequests.Post(url, &grequests.RequestOptions{
		JSON:   params,
		IsAjax: true,
	})
	if err != nil {
		log.Println("Unable to make request", resp.Error)
	}

	if resp.Ok != true {
		log.Println(resp.String())
		log.Println("Request did not return OK")
	}
	return resp
}
