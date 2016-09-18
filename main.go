package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	formatter(scanner, os.Stdout)
}

func formatter(s *bufio.Scanner, w io.Writer) {
	for s.Scan() {
		line := s.Text()
		jsonStart := 0
		for i, thisRune := range line {
			if string(thisRune) == "{" {
				jsonStart = i
				break
			}
		}
		date := line[:jsonStart]
		rawJSON := line[jsonStart:]
		// todo: handle arbitrary nesting
		var err error
		var buf []byte
		data := make(map[string]interface{})
		// todo: don't use strings, just bytes
		if err = json.Unmarshal([]byte(rawJSON), &data); err != nil {
			w.Write([]byte(line))
			w.Write([]byte("\n"))
			continue
		}
		if buf, err = json.MarshalIndent(data, " ", "    "); err != nil {
			w.Write([]byte(fmt.Sprintf("\n >> [ERROR] %v \n", err)))
		}
		w.Write([]byte(date))
		var escaping bool
		for _, thisRune := range buf {
			if string(thisRune) == `\` && !escaping {
				escaping = true
				continue
			} else if string(thisRune) == "n" && escaping {
				w.Write([]byte("\n"))
				escaping = false
				continue
			} else if string(thisRune) == "t" && escaping {
				w.Write([]byte("    "))
				escaping = false
				continue
			} else {
				escaping = false
			}
			w.Write([]byte{thisRune})
		}
		w.Write([]byte("\n"))
	}
}
