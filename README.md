# flyhttp

#### 初始化客户端

1. 设置公共header

```go
cli := flyhttp.New(
    flyhttp.WithHost("http://example.com"),
    flyhttp.WithHeader(http.Header{
        "Authorization": {"{token}"}, //将在后续请求中作为公共header
    }),
)
```

2. 设置host

```go
cli := flyhttp.New(
    flyhttp.WithHost("http://example.com"),
)

// GET http://example.com/foo/bar?name=kuhufu&age=11
_, err := cli.Get("/foo/bar",
    flyhttp.QueryParams(url.Values{
        "name": {"kuhufu"},
        "age":  {"11"},
    }),
)
```

3. 创建子分组

```go
foo := cli.Group("/foo")
_, err := foo.Get("/bar",
    flyhttp.QueryParams(url.Values{
        "name": {"kuhufu"},
        "age":  {"11"},
    }),
)
```



#### Get

*GET http://example.com?name=kuhufu&age=11*

```go
import "github.com/kuhufu/flyhttp"
```

方式一

```go
resp, err := flyhttp.Get("http://example.com?name=kuhufu&age=11")
```

方式二

```go
resp, err := flyhttp.Get("http://example.com",
    flyhttp.QueryParams(url.Values{
        "name": {"kuhufu"},
        "age":  {"11"},
    }),
)
```

方式三

```go
resp, err := flyhttp.Get("http://example.com?name=kuhufu",
    flyhttp.QueryParams(url.Values{
        "age": {"11"},
    }),
)
```


#### Post
```go
resp, err := flyhttp.Post("http://example.com",
    flyhttp.Header("Content-Type", "application/json"),
    flyhttp.Body([]byte(`{"name":"kuhufu","age":11}`)),
)
```

通过 `JSONBody` 将 `Content-Type` 设置 为 `application/json`，并将对象序列化为json字符串后作为body

```go
resp, err := flyhttp.Post("http://example.com",
    flyhttp.JSONBody(map[string]interface{}{
        "name": "kuhufu",
        "age":  11,
    }),
)
```

通过 `FormBody` 将 `Content-Type` 设置为 `application/x-www-form-urlencoded`

```go
resp, err := flyhttp.Post("http://example.com",
    flyhttp.FormBody(url.Values{
        "name": {"kuhufu"},
        "age":  {"11"},
    }),
)
```



#### response 的快捷操作

```go
res := flyhttp.Wrap(flyhttp.Get("http://example.com"))

//获取body字节
bytes, err := res.Bytes()

//获取body并转为string
str, err := res.String()

//JSON反序列化
var v map[string]interface{}
err := res.JSON(&v)
```

