package test

import (
	"net/http"
	"net/url"
	"testing"
)
import "github.com/kuhufu/flyhttp"

func Test_Get(t *testing.T) {
	cli := flyhttp.New(
		flyhttp.WithBase("http://example.com"),
		flyhttp.WithHeader(http.Header{
			"Authorization": {"{token}"},
		}),
	)

	foo := cli.Group("/foo")
	_, err := foo.Get("/bar",
		flyhttp.Query(url.Values{
			"name": {"kuhufu"},
			"age":  {"11"},
		}),
	)
	if err != nil {
		t.Error(err)
	}

	// GET http://example.com/foo?name=kuhufu&age=11
	_, err = cli.Get("/foo/bar",
		flyhttp.Query(url.Values{
			"name": {"kuhufu"},
			"age":  {"11"},
		}),
	)
	if err != nil {
		t.Error(err)
	}
}
