package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestFormatter(t *testing.T) {
	input := `2016/09/18 06:26:10 starting on :9126
2016/09/18 06:26:13 {"event":"panic","message":"interface conversion: interface {} is nil, not *main.logWriter\ngoroutine 5 [running]:\nruntime/debug.Stack(0xc42003f8e0, 0x25bb60, 0xc420014840)\n\t/usr/local/go/src/runtime/debug/stack.go:24 +0x79\nmain.mwPanic.func1.1(0xc4200d01e0)\n\t/Users/sethammons/workspace/go/src/github.com/sethgrid/cloudlock/main.go:73 +0x76\npanic(0x25bb60, 0xc420014840)\n\t/usr/local/go/src/runtime/panic.go:458 +0x243\nmain.mwLog.func1(0x3b8d40, 0xc420075930, 0xc4200d01e0)\n\t/Users/sethammons/workspace/go/src/github.com/sethgrid/cloudlock/main.go:126 +0x8fc\nnet/http.HandlerFunc.ServeHTTP(0xc4200cca60, 0x3b8d40, 0xc420075930, 0xc4200d01e0)\n\t/usr/local/go/src/net/http/server.go:1726 +0x44\nmain.mwPanic.func1(0x3b8d40, 0xc420075930, 0xc4200d01e0)\n\t/Users/sethammons/workspace/go/src/github.com/sethgrid/cloudlock/main.go:76 +0x8b\nnet/http.HandlerFunc.ServeHTTP(0xc4200cca80, 0x3b8d40, 0xc420075930, 0xc4200d01e0)\n\t/usr/local/go/src/net/http/server.go:1726 +0x44\nnet/http.(*ServeMux).ServeHTTP(0x3d3720, 0x3b8d40, 0xc420075930, 0xc4200d01e0)\n\t/usr/local/go/src/net/http/server.go:2022 +0x7f\nnet/http.serverHandler.ServeHTTP(0xc420092500, 0x3b8d40, 0xc420075930, 0xc4200d01e0)\n\t/usr/local/go/src/net/http/server.go:2202 +0x7d\nnet/http.(*conn).serve(0xc420092580, 0x3b92c0, 0xc420014740)\n\t/usr/local/go/src/net/http/server.go:1579 +0x4b7\ncreated by net/http.(*Server).Serve\n\t/usr/local/go/src/net/http/server.go:2293 +0x44d\n"}`

	want := `2016/09/18 06:26:10 starting on :9126
2016/09/18 06:26:13 {
     "event": "panic",
     "message": "interface conversion: interface {} is nil, not *main.logWriter
goroutine 5 [running]:
runtime/debug.Stack(0xc42003f8e0, 0x25bb60, 0xc420014840)
    /usr/local/go/src/runtime/debug/stack.go:24 +0x79
main.mwPanic.func1.1(0xc4200d01e0)
    /Users/sethammons/workspace/go/src/github.com/sethgrid/cloudlock/main.go:73 +0x76
panic(0x25bb60, 0xc420014840)
    /usr/local/go/src/runtime/panic.go:458 +0x243
main.mwLog.func1(0x3b8d40, 0xc420075930, 0xc4200d01e0)
    /Users/sethammons/workspace/go/src/github.com/sethgrid/cloudlock/main.go:126 +0x8fc
net/http.HandlerFunc.ServeHTTP(0xc4200cca60, 0x3b8d40, 0xc420075930, 0xc4200d01e0)
    /usr/local/go/src/net/http/server.go:1726 +0x44
main.mwPanic.func1(0x3b8d40, 0xc420075930, 0xc4200d01e0)
    /Users/sethammons/workspace/go/src/github.com/sethgrid/cloudlock/main.go:76 +0x8b
net/http.HandlerFunc.ServeHTTP(0xc4200cca80, 0x3b8d40, 0xc420075930, 0xc4200d01e0)
    /usr/local/go/src/net/http/server.go:1726 +0x44
net/http.(*ServeMux).ServeHTTP(0x3d3720, 0x3b8d40, 0xc420075930, 0xc4200d01e0)
    /usr/local/go/src/net/http/server.go:2022 +0x7f
net/http.serverHandler.ServeHTTP(0xc420092500, 0x3b8d40, 0xc420075930, 0xc4200d01e0)
    /usr/local/go/src/net/http/server.go:2202 +0x7d
net/http.(*conn).serve(0xc420092580, 0x3b92c0, 0xc420014740)
    /usr/local/go/src/net/http/server.go:1579 +0x4b7
created by net/http.(*Server).Serve
    /usr/local/go/src/net/http/server.go:2293 +0x44d
"
 }
`
	var b []byte
	buf := bytes.NewBuffer(b)
	// change input to a new scanner
	formatter(bufio.NewScanner(strings.NewReader(input)), buf)

	output, err := ioutil.ReadAll(buf)
	if err != nil {
		t.Errorf("error reading output - %v", err)
	}
	if string(output) != want {
		t.Errorf("got: %v\n\nwant: %v", string(output), want)
	}
}
