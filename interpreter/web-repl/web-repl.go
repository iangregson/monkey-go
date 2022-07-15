package main

import (
	"fmt"
	"monkey-go/interpreter/repl"
	"syscall/js"
)

var console js.Value
var hasData chan bool

func main() {
	hasData = make(chan bool, 1)
	js.Global().Set("initMonkeyRepl", terminalWrapper())
	<-make(chan bool)
}

type Terminal struct {
	terminal js.Value
	reader   js.Value
	writer   js.Value
}

func (t Terminal) Write(v []byte) (int, error) {
	t.terminal.Call("write", string(v))
	return len(v), nil
}

func (t *Terminal) Read(p []byte) (n int, err error) {
	<-hasData
	bytes := t.reader.Call("readString").String()
	copy(p, []byte(bytes))
	n = len(bytes)
	return
}

func terminalWrapper() js.Func {
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 3 {
			return "Invalid arguments."
		}

		terminal := &Terminal{
			terminal: args[0],
			reader:   args[1],
			writer:   args[2],
		}

		fmt.Fprintf(terminal, "Hello! This is the Monkey programming language!\r\n")
		fmt.Fprintf(terminal, "Feel free to type in some commands (you can learn about it at https://monkeylang.org/ \r\n")

		onReadable := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			hasData <- true
			return nil
		})

		args[1].Call("onReadable", onReadable)

		go repl.Start(terminal, terminal)

		return nil
	})

	return jsFunc
}
