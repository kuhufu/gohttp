package flyhttp

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"testing"
)

var obj = NewBase("https://starmicro.happyelements.cn/v1", http.DefaultClient)

func TestInstance_Get(t *testing.T) {
	s, e := obj.Get("/idol/idollist", map[string]string{
		"name": "jhon",
		"age":  "1",
	}).String()
	if e != nil {
		t.Error()
	}
	fmt.Println(s)
}

func Test_path(t *testing.T) {
	p1, _ := url.Parse("https://starmicro.happyelements.cn/v1")
	p := "idol?id=2&name=jhon"
	p2, _ := url.Parse(p)
	p1.Path = path.Join(p1.Path, p2.Path)
	p1.RawQuery = p2.RawQuery
	fmt.Println(p1.String())
}
