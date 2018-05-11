package shapeshift

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const (
	// BaseURL is an api endpoint
	BaseURL = "https://shapeshift.io/%s"
)

// Client is an interface for a client.
type Client interface {
	GetAPIClient
	PostAPIClient
}

// New creates Client
func New() Client {
	return &client{}
}

type client struct{}

func (c *client) do(ctx context.Context, method, path string, body io.Reader) (json.RawMessage, error) {
	url := fmt.Sprintf(BaseURL, path)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		const format = "failed to HTTP request with status code %d: %s"
		msg, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			msg = []byte("no error message returned from shapeshift server")
		}
		return nil, errors.Errorf(format, resp.StatusCode, msg)
	}

	var rawMsg json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&rawMsg); err != nil {
		return nil, err
	}
	return rawMsg, nil
}
