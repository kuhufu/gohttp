package flyhttp

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

var client = New(&http.Client{})

func TestClient_Get(t *testing.T) {
	baseURL := "https://starmicro.happyelements.cn/v1"
	testSet := []struct {
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
			args:   []interface{}{nil, nil, nil},
			expect: false,
		},
		{
			args:   []interface{}{"page=1"},
			expect: true,
		},
		{
			args:   []interface{}{"page=1", http.Header{}},
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

	for i, v := range testSet {
		_, err := bc.Get(baseURL+"/idol/idollist", v.args...).String()
		if (err == nil) != v.expect {
			t.Error(i, err)
		}
	}
}

func TestPost(t *testing.T) {
	header := http.Header{
		"content-type": {"application/x-www-form-urlencoded"},
	}
	testSet := []struct {
		args   []interface{}
		expect bool
	}{
		{
			args:   nil,
			expect: false,
		},
		{
			args:   []interface{}{nil},
			expect: false,
		},
		{
			args:   []interface{}{"id=1"},
			expect: false,
		},
		{
			args:   []interface{}{nil, nil},
			expect: false,
		},
		{
			args:   []interface{}{[]byte("id=1"), header},
			expect: false,
		},
		{
			args:   []interface{}{header, []byte("id=1")},
			expect: true,
		},
		{
			args:   []interface{}{header, nil},
			expect: true,
		},
		{
			args:   []interface{}{"application/x-www-form-urlencoded", "id=1"},
			expect: true,
		},
		{
			args:   []interface{}{header, strings.NewReader("id=1")},
			expect: true,
		},
	}

	for i, v := range testSet {
		_, err := client.Post(
			"https://starmicro.happyelements.cn/v1/comment/comment",
			v.args...).String()
		if (err == nil) != v.expect {
			t.Error(i, err)
		}
	}
}

func TestPostForm(t *testing.T) {
	data := url.Values{
		"content": {"打卡"},  //评论内容
		"id":      {"622"}, //视频id
		"pid":     {"0"},   //unknown
		"type":    {"1"},   //unknown
		"idol_id": {"1"},   //
	}
	s, err := client.PostForm("https://starmicro.happyelements.cn/v1/comment/comment", data).String()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(s)
}
