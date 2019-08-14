package flyhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"
	"strings"
)

type Client struct {
	inner *http.Client
}

var _ Interface = Client{}

func New(c *http.Client) Client {
	return Client{c}
}

func (c Client) Get(url string, args ...interface{}) (r Result) {
	length := len(args)
	if length > 2 {
		return Result{nil, errors.New("the length of args must <= 2")}
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
		default:
			return Result{
				resp: nil,
				err:  errors.New(fmt.Sprintf("wrong args[0] type: %T, type must be map[string]string or url.Values", args[0])),
			}
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
	if length != 2 {
		return Result{nil, errors.New("the length of args must == 2, use the format: Post(url, contentType|headr, data")}
	}

	header := make(http.Header)
	switch v := args[0].(type) {
	case string:
		header.Set("Content-Type", v)
	case http.Header:
		header = v
	default:
		return Result{nil, errors.New(fmt.Sprintf("wrong args[0] type: %T, wrong args[0] type, type must be http.Header or string", args[0]))}
	}

	var body io.Reader
	switch v := args[1].(type) {
	case []byte:
		body = bytes.NewReader(v)
	case string:
		body = strings.NewReader(v)
	case io.Reader:
		body = v
	default:
		return Result{nil, errors.New(fmt.Sprintf("wrong args[1] type: %T, wrong args[1] type, arg type must be string or []byte or io.Reader", args[1]))}
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return Result{nil, err}
	}
	req.Header = header
	return c.Do(req)
}

func (c Client) PostForm(url string, data interface{}) Result {
	var body io.Reader
	switch v := data.(type) {
	case map[string]string:
		res := url2.Values{}
		for key, value := range v {
			res.Set(key, value)
		}
		body = strings.NewReader(res.Encode())
	case url2.Values:
		body = strings.NewReader(v.Encode())
	default:
		return Result{resp: nil, err: errors.New(fmt.Sprintf("wrong data type: %T, data type must be map[string]string or url.Values", data))}
	}
	return c.Post(url, "application/x-www-form-urlencoded", body)
}

func (c Client) PostJson(url string, data interface{}) Result {
	res, _ := json.Marshal(data)
	body := bytes.NewReader(res)
	return c.Post(url, "application/json", body)
}

func (c Client) Do(req *http.Request) Result {
	resp, err := c.inner.Do(req)
	return Result{resp, err}
}

//Base 生成一个与当前client关联的 BaseURLClient
func (c Client) Base(baseUrl string) BaseURLClient {
	return BaseURLClient{baseUrl: baseUrl, client: c}
}
