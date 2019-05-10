package flyhttp

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

var client = New(&http.Client{})

func TestClient_Get(t *testing.T) {
	data, err := client.Get("https://starmicro.happyelements.cn/v1/media/media-detail?id=497&idol_id=4").String()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(data)
}

func TestGet2(t *testing.T) {
	data, err := client.Get("https://starmicro.happyelements.cn/v1/media/media-detail?id=496&idol_id=4", "id=497&idol_id=4").String()
	if err != nil {
		t.Error()
		fmt.Println(err)
	}
	fmt.Println(data)
}

func TestGet3(t *testing.T) {
	data, err := client.Get("https://starmicro.happyelements.cn/v1/media/media-detail?id=496&idol_id=4",
		map[string]string{
			"id":      "497",
			"idol_id": "4",
		}).String()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

func TestGet4(t *testing.T) {
	data, err := client.Get("https://starmicro.happyelements.cn/v1/media/media-detail",
		url.Values{
			"id":      {"497"},
			"idol_id": {"4"},
		}).String()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}

func TestPost(t *testing.T) {
	header := http.Header{
		"authorization": {"Miinno o44kJAQ3S_nAcl3uSe9GsD8_u7dV6mvT_1548283248"},
		"referer":       {"https://servicewechat.com/wx0a48c468391428fd/19/page-frame.html"},
	}
	data := url.Values{
		"content": {"打卡"},  //评论内容
		"id":      {"621"}, //视频id
		"pid":     {"0"},   //unknown
		"type":    {"1"},   //unknown
		"idol_id": {"1"},   //
	}

	s, err := client.Post("https://starmicro.happyelements.cn/v1/comment/comment", data, header).String()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(s)
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
