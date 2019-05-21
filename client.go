package flyhttp

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	url2 "net/url"
	"strings"
)

type Client struct {
	inner *http.Client
}

func (c Client) Get(url string, args ...interface{}) (r Result) {
	length := len(args)
	if length > 2 {
		return Result{nil, errors.New("args's length must <= 2")}
	}

	var query string
	if length >= 1 {
		switch v := args[0].(type) {
		case string:
			query = v
		case map[string]string:
			res := url2.Values{}
			for key, value := range v {
				res.Set(key, value)
			}
			query = res.Encode()
		case url2.Values:
			query = v.Encode()
		}
	}

	header := make(http.Header)
	if length == 2 && args[1] != nil {
		header = args[1].(http.Header)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Result{nil, err}
	}

	if query != "" {
		req.URL.RawQuery = query
	}
	req.Header = header
	return c.Do(req)
}

func (c Client) Post(url string, args ...interface{}) Result {
	length := len(args)
	if length > 2 {
		return Result{nil, errors.New("args's length must <= 2")}
	}

	var body io.Reader
	if length >= 1 {
		switch v := args[0].(type) {
		case []byte:
			body = bytes.NewReader(v)
		case string:
			body = strings.NewReader(v)
		case io.Reader:
			body = v
		default:
			return Result{nil, errors.New("wrong arg type, arg type must be string []byte or io.Reader")}
		}
	}

	header := make(http.Header)
	if length == 2 && args[1] != nil {
		switch v := args[1].(type) {
		case http.Header:
			header = v
		case string:
			header.Set("Content-Type", v)
		}

	}

	req, err := http.NewRequest("POST", url, body)
	req.Header = header
	if err != nil {
		return Result{nil, err}
	}
	return c.Do(req)
}

func (c Client) PostForm(url string, arg interface{}) Result {
	body := arg
	switch v := arg.(type) {
	case map[string]string:
		res := url2.Values{}
		for key, value := range v {
			res.Set(key, value)
		}
		body = strings.NewReader(res.Encode())
	case url2.Values:
		body = strings.NewReader(v.Encode())
	}
	return c.Post(url, body, "application/x-www-form-urlencoded")
}

func (c Client) Do(req *http.Request) Result {
	resp, err := c.inner.Do(req)
	return Result{resp, err}
}
