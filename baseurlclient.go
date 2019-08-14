package flyhttp

import (
	"net/http"
	url2 "net/url"
	p "path"
)

type BaseURLClient struct {
	client  Client
	baseUrl string
}

var _ BaseURLClient = BaseURLClient{}

// 不要在 baseUrl 中包含 query params
func NewBase(baseUrl string, c *http.Client) BaseURLClient {
	return BaseURLClient{Client{c}, baseUrl}
}

func Base(baseUrl string) BaseURLClient {
	return BaseURLClient{defaultClient, baseUrl}
}

func (b BaseURLClient) Get(path string, args ...interface{}) Result {
	url := buildUrl(b.baseUrl, path)
	return b.client.Get(url, args...)
}

func (b BaseURLClient) PostForm(path string, data interface{}) Result {
	url := buildUrl(b.baseUrl, path)
	return b.client.PostForm(url, data)
}

//Post Post(url, contentType|header, data)
func (b BaseURLClient) Post(path string, args ...interface{}) Result {
	url := buildUrl(b.baseUrl, path)
	return b.client.Post(url, args...)
}

func buildUrl(baseUrl, path string) string {
	baseURL, _ := url2.Parse(baseUrl)
	pathURL, _ := url2.Parse(path)
	baseURL.Path = p.Join(baseURL.Path, pathURL.Path)
	baseURL.RawQuery = pathURL.RawQuery
	return baseURL.String()
}
