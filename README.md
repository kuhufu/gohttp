# flyhttp
### Example
##### GET
`http://example.com?id=1&page=2`
```go
flyhttp.Get("http://example.com?id=1&page=2")

flyhttp.Get("http://example.com", "id=1&page=2")

flyhttp.Get("http://example.com", map[string][string]{
	"id":   "1",
	"page": "2",
})

flyhttp.Get("http://example.com", url.Values{
	"id":   {"1"},
	"page": {"2"},
})
```
##### POST
```go
flyhttp.Post("http://example.com", strings.NewReader(`{"name":"jhon", "age":11}`), http.Header{
	"content-type":{"application/json"}
})
equal
flyhttp.Post("http://example.com", `{"name":"jhon", "age":11}`, "application/json")
flyhttp.Post("http://example.com", []byte(`{"name":"jhon", "age":11}`), "application/json")
flyhttp.Post("http://example.com", strings.NewReader(`{"name":"jhon", "age":11}`), "application/json")
```

##### POST Forum
```go
flyhttp.PostForum("http://example.com", url.Values{
	"name": {"tom"},
	"age":  {"11"},
})

flyhttp.PostForum("http://example.com", map[string][string]{
	"name": "tom",
	"age":  "11",
})

```
### Client
```go
client := flyhttp.New(&http.Client)
//just like above
```


### BaseURLClient
`http://example.com/path/path`
```go
client := flyhttp.NewBase("http://example.com", &http.Client{})
client.Get("/path/path")
```

### 注意
##### xxx.Get
虽然使用了可变长参数，
但`Get(url string, args ...interface())`至多三个实参。

三个实参按 `(url, query_params, header)` 排列

以下为实参允许的类型

|名称|类型|
|-----|----|
|url|`string`|
|query_params|`string`, `map[string][string]`, `url.Value`|
|header|`http.Header`|
>*query_params 会覆盖掉 url 中的查询参数*

#### xxx.Post

虽然使用了可变长参数，

但`Post(url string, args ...interface())`至多三个实参。

三个实参按 `(url, data, header|contentType)` 排列

以下为实参允许的类型

|名称|类型|
|-----|----|
|url|`string`|
|data|`[]byte`, `string`, `io.Reader`|
|header|`http.Header`|
|contentType|`string`|
-------

更多请见测试文件