package test

import (
	"fmt"
	"github.com/kuhufu/gohttp"
	"testing"
)

func TestGet(t *testing.T) {
	res := gohttp.Wrap(gohttp.Get("http://example.com"))

	fmt.Println(res.String())
}

func TestGroup(t *testing.T) {
	cli := gohttp.Group("http://example.com")
	res := gohttp.Wrap(cli.Get("/foo"))
	fmt.Println(res.String())
}
