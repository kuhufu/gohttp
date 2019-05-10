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
flyhttp.Post("http://example.com", "application/json", strings.NewReader("data"))
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