package flyhttp

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	url2 "net/url"
	"strings"
)

func New(c *http.Client) Client {
	return Client{c}
}

func (c Client) Get(url string, args ...interface{}) (r Result) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Result{nil, err}
	}

	if len(args) > 1 {
		return Result{nil, errors.New("args's length must equal 0 or 1")}
	}

	if len(args) == 1 {
		switch v := args[0].(type) {
		case string:
			req.URL.RawQuery = v
		case map[string]string:
			res := url2.Values{}
			for key, value := range v {
				res.Set(key, value)
			}
			req.URL.RawQuery = res.Encode()
		case url2.Values:
			req.URL.RawQuery = v.Encode()
		}
	}

	return c.doRequest(req)
}

func (c Client) PostForm(url string, arg interface{}) Result {
	var body io.Reader
	if arg != nil {
		switch v := arg.(type) {
		case string:
			body = strings.NewReader(v)
		case url2.Values:
			body = strings.NewReader(v.Encode())
		case []byte:
			body = bytes.NewReader(v)
		default:
			return Result{nil, errors.New("wrong arg type, arg type must be string []byte url2.Values nil")}
		}
	}

	return c.Post(url, "application/x-www-form-urlencoded", body)
}

func (c Client) Post(url, contentType string, body io.Reader) Result {
	resp, err := c.inner.Post(url, contentType, body)
	return Result{resp, err}
}

func (c Client) doRequest(req *http.Request) Result {
	resp, err := c.inner.Do(req)
	return Result{resp, err}
}
