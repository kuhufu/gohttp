package test

import (
	"net/http"
	"net/url"
	"testing"
)
import "github.com/kuhufu/gohttp"

func Test_Get(t *testing.T) {
	cli := gohttp.New(
		gohttp.WithBase("http://example.com"),
		gohttp.WithHeader(http.Header{
			"Authorization": {"{token}"},
		}),
	)

	foo := cli.Group("/foo")
	_, err := foo.Get("/bar",
		gohttp.Query(url.Values{
			"name": {"kuhufu"},
			"age":  {"11"},
		}),
	)
	if err != nil {
		t.Error(err)
	}

	// GET http://example.com/foo?name=kuhufu&age=11
	_, err = cli.Get("/foo/bar",
		gohttp.Query(url.Values{
			"name": {"kuhufu"},
			"age":  {"11"},
		}),
	)
	if err != nil {
		t.Error(err)
	}
}
