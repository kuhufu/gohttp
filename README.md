# gohttp

### 初始化客户端

1. 设置公共header

```go
cli := gohttp.New(
    gohttp.WithHeader(http.Header{
        "Authorization": {"{token}"}, //将在后续请求中作为公共header
    }),
)
```

2. 设置 base url

```go
cli := gohttp.New(
    gohttp.WithBase("http://example.com"),
)

// GET http://example.com/foo/bar
_, err := cli.Get("/foo/bar")
```

3. 创建子分组

```go
// GET http://example.com/foo/bar
foo := cli.Group("http://example.com")

_, err := foo.Get("/foo/bar")
```



### Get请求

```go
//GET http://example.com?name=kuhufu&age=11
import "github.com/kuhufu/gohttp"
```

方式一

```go
resp, err := gohttp.Get("http://example.com?name=kuhufu&age=11")
```

方式二

```go
resp, err := gohttp.Get("http://example.com",
    gohttp.Query(url.Values{
        "name": {"kuhufu"},
        "age":  {"11"},
    }),
)
```

方式三

```go
resp, err := gohttp.Get("http://example.com?name=kuhufu",
    gohttp.Query(url.Values{
        "age": {"11"},
    }),
)
```


### Post请求
```go
resp, err := gohttp.Post("http://example.com",
    gohttp.Header("Content-Type", "application/json"),
    gohttp.Body([]byte(`{"name":"kuhufu","age":11}`)),
)
```

通过 `JSONBody` 将 `Content-Type` 设置 为 `application/json`，并将对象序列化为json字符串后作为body

```go
resp, err := gohttp.Post("http://example.com",
    gohttp.JSONBody(map[string]interface{}{
        "name": "kuhufu",
        "age":  11,
    }),
)
```

通过 `FormBody` 将 `Content-Type` 设置为 `application/x-www-form-urlencoded`

```go
resp, err := gohttp.Post("http://example.com",
    gohttp.FormBody(url.Values{
        "name": {"kuhufu"},
        "age":  {"11"},
    }),
)
```



### 响应的快捷操作

```go
res := gohttp.Wrap(gohttp.Get("http://example.com"))

//获取body字节
bytes, err := res.Bytes()

//获取body并转为string
str, err := res.String()

//JSON反序列化
var v map[string]interface{}
err := res.JSON(&v)
```

