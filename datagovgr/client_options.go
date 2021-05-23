package datagovgr

import (
	"errors"
	"net/http"
)

var (
	// ErrNilHTTPClient indicates an error when a user hasn't provide a valid http client.
	ErrNilHTTPClient = errors.New("http client can't be nil")
)

// ClientOption describes a functional parameter for the client constructor.
type ClientOption func(*clientOptions) error

type clientOptions struct {
	httpClient *http.Client
}

// WithHTTPClient allows the overriding of the http client.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(opts *clientOptions) error {
		if httpClient == nil {
			return ErrNilHTTPClient
		}

		opts.httpClient = httpClient

		return nil
	}
}
