package gohttp

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type RequestOption func(req *http.Request)

func Query(params url.Values) RequestOption {
	return func(req *http.Request) {
		if len(req.URL.RawQuery) != 0 {
			req.URL.RawQuery += "&" + params.Encode()
		} else {
			req.URL.RawQuery = params.Encode()
		}
	}
}

func Cookie(c *http.Cookie) RequestOption {
	return func(req *http.Request) {
		req.AddCookie(c)
	}
}

func Header(key, val string) RequestOption {
	return func(req *http.Request) {
		if req.Header == nil {
			req.Header = http.Header{}
		}
		req.Header.Set(key, val)
	}
}

func ContentType(contentType string) RequestOption {
	return Header("Content-Type", contentType)
}

func Body(body []byte) RequestOption {
	return func(req *http.Request) {
		req.Body = ioutil.NopCloser(bytes.NewReader(body))
	}
}

func ReaderBody(r io.Reader) RequestOption {
	return func(req *http.Request) {
		if readCloser, ok := r.(io.ReadCloser); ok {
			req.Body = readCloser
		} else {
			req.Body = ioutil.NopCloser(r)
		}
	}
}

func FormBody(form url.Values) RequestOption {
	return func(req *http.Request) {
		Header("Content-Type", "application/x-www-form-urlencoded")(req)
		req.Body = ioutil.NopCloser(strings.NewReader(form.Encode()))
	}
}

func JSONBody(obj interface{}) RequestOption {
	return func(req *http.Request) {
		marshal, err := json.Marshal(obj)
		if err != nil {
			panic(err)
		}

		Header("Content-Type", "application/json")(req)
		req.Body = ioutil.NopCloser(bytes.NewReader(marshal))
	}
}
