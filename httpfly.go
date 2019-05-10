package flyhttp

import (
	"net/http"
)

var defaultClient Client

func init() {
	defaultClient = Client{http.DefaultClient}
}

type Client struct {
	inner *http.Client
}

type BaseURLClient struct {
	client Client
	baseUrl string
}

type Result struct {
	resp *http.Response
	err  error
}


func Get(url string, args ...interface{}) (r Result) {
	return defaultClient.Get(url, args...)
}

func PostForm(url string, arg interface{}) Result {
	return defaultClient.PostForm(url, arg)
}