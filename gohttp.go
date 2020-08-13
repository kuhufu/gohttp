package gohttp

import (
	"net/http"
)

var defaultClient *GroupClient

func init() {
	defaultClient = New()
}

func Wrap(response *http.Response, err error) Result {
	return Result{
		resp: response,
		err:  err,
	}
}

func Group(relativePath string, opts ...ClientOption) *GroupClient {
	return defaultClient.Group(relativePath, opts...)
}

func Do(method string, url string, opts ...RequestOption) (*http.Response, error) {
	return defaultClient.Do(method, url, opts...)
}

func Get(url string, opts ...RequestOption) (*http.Response, error) {
	return defaultClient.Get(url, opts...)
}

func Post(url string, opts ...RequestOption) (*http.Response, error) {
	return defaultClient.Post(url, opts...)
}
