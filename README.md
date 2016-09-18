# jsonlog (beta)

Pretty Print your Go JSON logs!

If you log json formatted logs using the stdlib, you may have output that looks similar to the following:

```
$ go run site/main.go
2016/09/18 08:13:39 starting on :9126
2016/09/18 08:13:42 {"code":200,"content_length":0,"event":"request","method":"GET","remote_addr":"[::1]:51731","request_id":"09068621","request_time":1474211622,"tts_ns":0,"url":"/dash"}
```

Piping your log output to jsonlog gives the following:
```
$ go run site/main.go 2>&1 | jsonlog
2016/09/18 08:12:36 starting on :9126
2016/09/18 08:12:42 {
     "code": 200,
     "content_length": 0,
     "event": "request",
     "method": "GET",
     "remote_addr": "[::1]:51689",
     "request_id": "152d4501",
     "request_time": 1.474211562e+09,
     "tts_ns": 0,
     "url": "/dash"
 }
```


### Motivation

They key motivation in making this package is to easily parse panics/traces that are json values in some log statements. For example:

```
$ go run main.go
2016/09/18 07:59:46 starting on :9126
2016/09/18 07:59:49 {"event":"panic","message":"interface conversion: interface {} is nil, not *main.logWriter goroutine 20 [running]:\nruntime/debug.Stack(0xc42003a8e0, 0x25bb60, 0xc420076740)\n\t/usr/local/go/src/runtime/debug/stack.go:24 +0x79\nmain.mwPanic.func1.1(0xc4200e21e0)\n\t/Users/sethammons/workspace/go/src/github.com/sethgrid/cloudlock/main.go:73 +0x76\npanic(0x25bb60, 0xc420076740)\n\t/usr/local/go/src/runtime/panic.go:458 +0x243\nmain.mwLog.func1(0x3b8d40, 0xc42007d5f0, 0xc4200e21e0)\n\t/Users/sethammons/workspace/go/src/github.com/sethgrid/cloudlock/main.go:126 +0x8fc\nnet/http.HandlerFunc.ServeHTTP(0xc4200de820, 0x3b8d40, 0xc42007d5f0, 0xc4200e21e0)\n\t/usr/local/go/src/net/http/server.go:1726 +0x44\nmain.mwPanic.func1(0x3b8d40, 0xc42007d5f0, 0xc4200e21e0)\n\t/Users/sethammons/workspace/go/src/github.com/sethgrid/cloudlock/main.go:76 +0x8b\nnet/http.HandlerFunc.ServeHTTP(0xc4200de840, 0x3b8d40, 0xc42007d5f0, 0xc4200e21e0)\n\t/usr/local/go/src/net/http/server.go:1726 +0x44\nnet/http.(*ServeMux).ServeHTTP(0x3d3720, 0x3b8d40, 0xc42007d5f0, 0xc4200e21e0)\n\t/usr/local/go/src/net/http/server.go:2022 +0x7f\nnet/http.serverHandler.ServeHTTP(0xc4200a4500, 0x3b8d40, 0xc42007d5f0, 0xc4200e21e0)\n\t/usr/local/go/src/net/http/server.go:2202 +0x7d\nnet/http.(*conn).serve(0xc4200a4580, 0x3b92c0, 0xc420076640)\n\t/usr/local/go/src/net/http/server.go:1579 +0x4b7\ncreated by net/http.(*Server).Serve\n\t/usr/local/go/src/net/http/server.go:2293 +0x44d\n"}
```
becomes
```
$ go run main.go 2>&1 | jsonlog
2016/09/18 07:57:43 starting on :9126
2016/09/18 07:57:46 {
     "event": "panic",
     "message": "interface conversion: interface {} is nil, not *main.logWriter goroutine 5 [running]:
runtime/debug.Stack(0xc42003f8e0, 0x25bb60, 0xc420014840)
    /usr/local/go/src/runtime/debug/stack.go:24 +0x79
main.mwPanic.func1.1(0xc4200d01e0)
    /Users/sethammons/workspace/go/src/github.com/sethgrid/cloudlock/main.go:73 +0x76
panic(0x25bb60, 0xc420014840)
    /usr/local/go/src/runtime/panic.go:458 +0x243
main.mwLog.func1(0x3b8d40, 0xc420075860, 0xc4200d01e0)
    /Users/sethammons/workspace/go/src/github.com/sethgrid/cloudlock/main.go:126 +0x8fc
net/http.HandlerFunc.ServeHTTP(0xc4200cca60, 0x3b8d40, 0xc420075860, 0xc4200d01e0)
    /usr/local/go/src/net/http/server.go:1726 +0x44
main.mwPanic.func1(0x3b8d40, 0xc420075860, 0xc4200d01e0)
    /Users/sethammons/workspace/go/src/github.com/sethgrid/cloudlock/main.go:76 +0x8b
net/http.HandlerFunc.ServeHTTP(0xc4200cca80, 0x3b8d40, 0xc420075860, 0xc4200d01e0)
    /usr/local/go/src/net/http/server.go:1726 +0x44
net/http.(*ServeMux).ServeHTTP(0x3d3720, 0x3b8d40, 0xc420075860, 0xc4200d01e0)
    /usr/local/go/src/net/http/server.go:2022 +0x7f
net/http.serverHandler.ServeHTTP(0xc420092500, 0x3b8d40, 0xc420075860, 0xc4200d01e0)
    /usr/local/go/src/net/http/server.go:2202 +0x7d
net/http.(*conn).serve(0xc420092580, 0x3b92c0, 0xc420014740)
    /usr/local/go/src/net/http/server.go:1579 +0x4b7
created by net/http.(*Server).Serve
    /usr/local/go/src/net/http/server.go:2293 +0x44d
"
 }
```