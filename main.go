package main

import (
	"syscall/js"
)

func main() {
	c := make(chan bool, 0)

	add := func (this js.Value, inputs []js.Value) interface{} {
		return js.ValueOf(inputs[0].Int() + inputs[1].Int())
	}

	closeGo := func(this js.Value, inputs []js.Value) interface {} {
		c <- true
		return nil
	}

	setGlobal42 := func (this js.Value, inputs []js.Value) interface {} {
		js.Global().Set("$42", js.ValueOf(42))
		return nil
	}

	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("closeGo", js.FuncOf(closeGo))
	js.Global().Set("setGlobal42", js.FuncOf(setGlobal42))

	<-c
}
