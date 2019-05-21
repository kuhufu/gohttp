package flyhttp

import (
	"fmt"
	"net/http"
	"testing"
)

var bc = NewBase("https://starmicro.happyelements.cn/v1", http.DefaultClient)
var bc2 = Base("https://starmicro.happyelements.cn/v1")

func TestInstance_Get(t *testing.T) {
	s, e := bc.Get("/idol/idollist", map[string]string{
		"name": "jhon",
		"age":  "1",
	}).String()
	if e != nil {
		t.Error(e)
	}
	fmt.Println(s)
}
