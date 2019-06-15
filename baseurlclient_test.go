package flyhttp

import (
	"net/http"
	"net/url"
	"testing"
)

var bc = NewBase("https://starmicro.happyelements.cn/v1", http.DefaultClient)
var bc2 = Base("https://starmicro.happyelements.cn/v1")

func TestInstance_Get(t *testing.T) {
	tests := []struct {
		args []interface{}
		want bool
	}{
		{want: true, args: nil},
		{want: false, args: []interface{}{nil}},
		{want: false, args: []interface{}{nil, nil}},
		{want: true, args: []interface{}{"page=1"}},
		{want: true, args: []interface{}{map[string]string{"name": "jhon", "age": "1"}}},
		{want: true, args: []interface{}{url.Values{"name": {"jhon"}, "age": {"1"}}}},
	}

	for i, v := range tests {
		_, err := bc.Get("/idol/idollist", v.args...).String()
		if (err == nil) != v.want {
			t.Error(i, err)
		}
	}
}
