package flyhttp

import (
	"net/http"
	"net/url"
	"testing"
)

var bc = NewBase("https://starmicro.happyelements.cn/v1", http.DefaultClient)
var bc2 = Base("https://starmicro.happyelements.cn/v1")

func TestInstance_Get(t *testing.T) {
	values := []struct {
		args   []interface{}
		expect bool
	}{
		{
			args:   nil,
			expect: true,
		},
		{
			args:   []interface{}{nil},
			expect: false,
		},
		{
			args:   []interface{}{nil, nil},
			expect: false,
		},
		{
			args:   []interface{}{"page=1"},
			expect: true,
		},
		{
			args:   []interface{}{map[string]string{"name": "jhon", "age": "1"}},
			expect: true,
		},
		{
			args:   []interface{}{url.Values{"name": {"jhon"}, "age": {"1"}}},
			expect: true,
		},
	}

	for i, v := range values {
		_, err := bc.Get("/idol/idollist", v.args...).String()
		if (err == nil) != v.expect {
			t.Error(i, err)
		}
	}
}
