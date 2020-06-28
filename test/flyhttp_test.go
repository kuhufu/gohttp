package test

import (
	"fmt"
	"github.com/kuhufu/flyhttp"
	"testing"
)

func TestGet(t *testing.T) {
	res := flyhttp.Wrap(flyhttp.Get("http://example.com"))

	fmt.Println(res.String())
}

func TestGroup(t *testing.T) {
	cli := flyhttp.Group("http://example.com")
	res := flyhttp.Wrap(cli.Get("/foo"))
	fmt.Println(res.String())
}
