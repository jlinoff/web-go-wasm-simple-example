// https://golang.org/pkg/crypto/cipher/#Block
package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func main() {
	// https://github.com/golang/go/blob/master/src/syscall/js/js.go
	// https://blog.owulveryck.info/2018/06/08/some-notes-about-the-upcoming-webassembly-support-in-go.html
	// https://medium.zenika.com/go-1-11-webassembly-for-the-gophers-ae4bb8b1ee03
	clog("Hello world!")
	alert("wow, wasm loaded!")

	// populate the HTML in the wasm-area element.
	item := getElementById("wasm-area")
	clog("Found element 'wasm-area': %v.", item)
	item.Set("innerHTML", "<p>hello, <b>world</b>!</p><p>"+ time.Now().String() + "</p>")

	// change the run button size to show how to set attributes.
	item1 := getElementById("runButton")
	item1.Call("setAttribute", "style", "font-size: 24pt; color: green")
	clog("All done.")
}


// alert pops up an alert dialogue with a message.
func alert(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	alert := js.Global().Get("alert")
	alert.Invoke(msg)
}

// clog writes a message to the console log.
func clog(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	js.Global().Get("console").Call("log", msg)
}

// getElementById gets an element by id.
// The hash mark is not needed.
func getElementById(id string) js.Value {
	item := js.Global().Get("document").Call("getElementById", id)
	if item == js.Null() {
		clog("ERROR: unable to find element id %q", id)
	}
	return item
}
