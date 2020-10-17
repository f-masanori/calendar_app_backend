# net/httpパッケージのメモ

- http.Handlerとは?

```go
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}

ServeHTTP関数を持つだけのインターフェイスのことで、httpリクエストを受けて、レスポンスすることが責務
```

- http.Handleとは？

```go
func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }

URLに対応するhttp.HandlerをDefaultServeMuxに登録する関数
(DefaultServeMux)
```

- DefaultServeMuxとは？

ServeMuxのインスタンス。ServeMuxはServeHTTPと言う名前のメソッドを持っている



- http.HandlerFuncとは？

```go
type HandlerFunc func(ResponseWriter, *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
  f(w, r)
}

```

- http.HandleFuncとは？

```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}
```



____

構造体Serverのhandlerフィールドに値を指定してしまうと、きたリクエスト全てそのハンドラーに流れてしまう。