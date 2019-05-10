package flyhttp

import (
	"fmt"
	"net/http"
	"testing"
)

var obj = NewBase("https://starmicro.happyelements.cn/v1", http.DefaultClient)

func TestInstance_Get(t *testing.T) {
	s, e := obj.Get("/idol/idollist", map[string]string{
		"name": "jhon",
		"age":  "1",
	}).String()
	if e != nil {
		t.Error(e)
	}
	fmt.Println(s)
}
