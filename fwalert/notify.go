package fwalert

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
	"net/url"
)

const (
	ApiURL = "https://fwalert.com"
)

// Options allows full configuration of the message sent to the Pushover API
type Options struct {
	Token string `json:"token"`
	// User may be either a user key or a group key.
	Message string `json:"message"`
}

type client struct {
	opt Options
}

func New(opt Options) *client {
	return &client{opt: opt}
}

type Resp struct {
	Status int      `json:"status"`
	Errors []string `json:"errors"`
}

func (c *client) Send(message string) error {
	if c.opt.Token == "" {
		return errors.New("missing token")
	}

	if message == "" {
		return errors.New("missing message")
	}

	encodedMessage := url.QueryEscape(message)
	_url := fmt.Sprintf("%s/%s?message=%s", ApiURL, c.opt.Token, encodedMessage)
	resp, err := req.Get(_url)
	if err != nil {
		return nil
	}
	r := &Resp{}
	err = resp.ToJSON(r)
	if err != nil {
		return fmt.Errorf("fwalert send error: %w", err)
	}
	if r.Status != 1 {
		return errors.New(r.Errors[0])
	}
	return nil
}
