package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func LogMessage(msg string) {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset() 
	defer bufferPool.Put(buf)

	buf.WriteString("[LOG] ")
	buf.WriteString(msg)
	fmt.Println(buf.String())
}

func main() {
	for i := 0; i < 3; i++ {
		LogMessage(fmt.Sprintf("Event entry #%d", i))
	}
}