# flyhttp

### 初始化客户端

1. 设置公共header

```go
cli := flyhttp.New(
    flyhttp.WithHeader(http.Header{
        "Authorization": {"{token}"}, //将在后续请求中作为公共header
    }),
)
```

2. 设置 base url

```go
cli := flyhttp.New(
    flyhttp.WithBase("http://example.com"),
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
import "github.com/kuhufu/flyhttp"
```

方式一

```go
resp, err := flyhttp.Get("http://example.com?name=kuhufu&age=11")
```

方式二

```go
resp, err := flyhttp.Get("http://example.com",
    flyhttp.Query(url.Values{
        "name": {"kuhufu"},
        "age":  {"11"},
    }),
)
```

方式三

```go
resp, err := flyhttp.Get("http://example.com?name=kuhufu",
    flyhttp.Query(url.Values{
        "age": {"11"},
    }),
)
```


### Post请求
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



### response 的快捷操作

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

