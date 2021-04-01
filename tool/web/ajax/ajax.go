package ajax

import (
	"syscall/js"
)

func Stringify(obj interface{}) js.Value {
	return js.Global().Get("JSON").Call("stringify", js.ValueOf(obj))
}

func ParseResp(obj interface{}) (js.Value, int) {
	objJson := Stringify(obj)
	resp := js.Global().Get("JSON").Call("parse", objJson)
	return resp, resp.Length()
}

func RequestAjax(phost string, req interface{}, respFunc js.Func) {
	objrqs := Stringify(req)
	xhttp := js.Global().Get("XMLHttpRequest").New()
	xhttp.Call("open", "POST", phost, true)
	xhttp.Call("setRequestHeader", "Content-Type", "application/json;charset=UTF-8")
	//~ xhttp.Call("setRequestHeader", "Access-Control-Allow-Origin", "*")
	xhttp.Call("send", objrqs)

	xhttp.Set("onreadystatechange", respFunc)

	println("Envio - Json")
	println(objrqs.String())
	println("-------------------------------------")
}

func ResponseAjax(this js.Value, args []js.Value) (js.Value, bool) {
	xReadyState := this.Get("readyState").Int()
	xStatus := this.Get("status").Int()
	if xReadyState == 4 && xStatus == 200 {
		if !this.IsUndefined() && !this.IsNull() {
			return js.Global().Get("JSON").Call("parse", this.Get("responseText")), true
		} else {
			println("Sem Retorno Do Server...")
		}
	}
	return js.Null(), false
}

func HttpGetResponse(this js.Value, args []js.Value) (js.Value, bool) {
	xReadyState := this.Get("readyState").Int()
	xStatus := this.Get("status").Int()
	if xReadyState == 4 && xStatus == 200 {
		return this.Get("responseText"), true
	}
	return js.Null(), false
}
