package main

import (
	"Gestor-de-prazos/wasm/view"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Web Assembly aqui", js.Global().Get("navigator").Get("platform").String())
	view.Login()
	<-make(chan bool)
}
