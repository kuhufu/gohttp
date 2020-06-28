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
