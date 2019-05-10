package flyhttp

import (
	"io"
	"net/http"
	url2 "net/url"
	p "path"
)

// 不要在 baseUrl 中包含 query params
func NewBase(baseUrl string, c *http.Client) BaseURLClient {
	return BaseURLClient{baseUrl: baseUrl, client:Client{c}}
}

func (b BaseURLClient) Get(path string, args ...interface{}) Result {
	url := buildUrl(b.baseUrl, path)
	return b.client.Get(url, args...)
}

func (b BaseURLClient) PostForm(path string, arg interface{}) Result {
	url := buildUrl(b.baseUrl, path)
	return b.client.PostForm(url, arg)
}

func (b BaseURLClient) Post(path, contentType string, body io.Reader) Result {
	url := buildUrl(b.baseUrl, path)
	return b.client.Post(url, contentType, body)
}

func buildUrl(baseUrl, path string) string {
	parse1, _ := url2.Parse(baseUrl)
	parse2, _ := url2.Parse(path)
	parse1.Path = p.Join(parse1.Path, parse2.Path)
	parse1.RawQuery = parse2.RawQuery
	return parse1.String()
}
