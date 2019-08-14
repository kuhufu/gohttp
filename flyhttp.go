package flyhttp

import (
	"net/http"
)

var defaultClient = Client{http.DefaultClient}

type Interface interface {
	Get(url string, args ...interface{}) Result
	Post(url string, args ...interface{}) Result
	PostForm(url string, args interface{}) Result
	PostJson(url string, args interface{}) Result
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

func PostJson(url string, arg interface{}) Result {
	return defaultClient.PostJson(url, arg)
}
