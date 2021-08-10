package test

import (
	"fmt"
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

type reader string

func (r reader) Read(p []byte) (n int, err error) {
	return 0, nil
}

func Test(t *testing.T) {
	request, err := http.NewRequest("Post", "https://www.baidu.com", reader("sdfdf"))
	if err != nil {
		t.Error(err)
	}

	fmt.Println(request.ContentLength)
}
