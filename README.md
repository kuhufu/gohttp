##### GET
```go
flyhttp.Get("http://example.com?id=1&page=2")

flyhttp.Get("http://example.com", "id=1&page=2")

flyhttp.Get("http://example.com", map[string][string]{
	"name": "tom",
	"age":"11",
})

flyhttp.Get("http://example.com", url.Values{
	"id": {"1"},
	"page":{"2"},
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
	"age":{"11"},
})

flyhttp.PostForum("http://example.com", map[string][string]{
	"name": "tom",
	"age":"11",
})

```