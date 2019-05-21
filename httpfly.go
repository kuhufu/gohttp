package flyhttp

import (
	"net/http"
)

var defaultClient Client

func init() {
	defaultClient = Client{http.DefaultClient}
}

func New(c *http.Client) Client {
	return Client{c}
}

// 不要在 baseUrl 中包含 query params
func NewBase(baseUrl string, c *http.Client) BaseURLClient {
	return BaseURLClient{Client{c}, baseUrl}
}

func Base(baseUrl string) BaseURLClient {
	return BaseURLClient{defaultClient, baseUrl}
}

func Get(url string, args ...interface{}) (r Result) {
	return defaultClient.Get(url, args...)
}

func Post(url string, args ...interface{}) (r Result) {
	return defaultClient.Post(url, args...)
}

func PostForm(url string, arg interface{}) Result {
	return defaultClient.PostForm(url, arg)
}
