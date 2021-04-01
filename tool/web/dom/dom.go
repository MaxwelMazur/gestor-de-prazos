package dom

import (
	"Gestor-de-prazos/tool/dice"
	"strings"
	"syscall/js"
)

// [AppendChild pega o elemento pai e adiciona um filho a este mesmo]
func AppendChild(element, elementChild js.Value) {
	if IsElement(element) && IsElement(elementChild) {
		element.Call("appendChild", elementChild)
	}
}

// [QuerySelector selectiona um element uso qualquer selector]
func QuerySelector(selectors string) js.Value {
	return js.Global().Get("document").Call("querySelector", selectors)
}

// [QuerySelector selectiona um slice de elemento uso qualquer selector]
func QuerySelectorAll(selectors string) js.Value {
	return js.Global().Get("document").Call("querySelectorAll", selectors)
}

// [AddEventListener adiciona evento a um elemento]
func AddEventListener(element js.Value, event string, funcao js.Func) {
	if IsElement(element) {
		element.Call("addEventListener", event, funcao)
	}
}

// [CreateElement cria um elemento de qualquer tag com a quantidade de atributos necessarios]
func CreateElement(element js.Value, elementName, elementCharacter string, elementAtributos ...string) js.Value {
	name := strings.ToLower(elementName)
	tagHtml := js.Global().Get("document").Call("createElement", name)
	numberAtributos := len(elementAtributos)
	if numberAtributos > 0 {
		for i := 0; i < numberAtributos; i += 2 {
			if i+1 >= numberAtributos {
				break
			}
			tagHtml.Call("setAttribute", elementAtributos[i], elementAtributos[i+1])
		}
	}

	if len(elementCharacter) > 0 {
		if name == "input" || name == "textarea" {
			tagHtml.Set("value", elementCharacter)
		} else {
			tagHtml.Set("textContent", elementCharacter)
		}
	}
	AppendChild(element, tagHtml)
	return tagHtml
}

// [IsElement verifica se e um elemento javascript]
func IsElement(element js.Value) bool {
	if element.String() == "<object>" {
		return true
	}
	return false
}

// [SetContent adiciona um novo valor ao elemento]
func SetContent(element js.Value, content string) {
	elementName := ""
	if IsElement(element) {
		elementName = element.Get("nodeName").String()
		if elementName == "INPUT" || elementName == "TEXTAREA" {
			element.Set("value", content)
		} else {
			element.Set("textContent", content)
		}
	}
}

// [GetContent pega o valor do elemento]
func GetContent(element js.Value) (content string) {
	elementName := ""
	content = ""
	if IsElement(element) {
		elementName = element.Get("nodeName").String()
		if elementName == "INPUT" || elementName == "TEXTAREA" {
			content = element.Get("value").String()
		} else {
			content = element.Get("textContent").String()
		}
	}
	return content
}

// [SetTimeOut chama uma funcao apos um tempo determinado]
func SetTimeOut(fn js.Func, milisec int) {
	fnTime := js.Global().Get("setTimeout")
	fnTime.Invoke(fn, milisec)
}

// [Hide esconde o elemento]
func Hide(element js.Value) {
	if IsElement(element) {
		element.Get("style").Set("display", "none")
	}
}

// [Show mostra o elemento]
func Show(element js.Value) {
	if IsElement(element) {
		element.Get("style").Set("display", "")
	}
}

// [AddClass adiciona class]
func AddClass(element js.Value, class string) {
	if IsElement(element) {
		element.Get("classList").Call("add", class)
	}
}

// [RemoveClass remove class]
func RemoveClass(element js.Value, class string) {
	if IsElement(element) {
		element.Get("classList").Call("remove", class)
	}
}

// [CreateTextNode pego um valor e jogo dentro do elemento]
func CreateTextNode(element js.Value, content string) js.Value {
	tag := js.Global().Get("document").Call("createTextNode", content)
	AppendChild(element, tag)
	return tag
}

// StyleCSS seta style de css
func StyleCSS(css string) {
	style := CreateElement(js.Global().Get("document").Get("head"), "style", "")
	CreateTextNode(style, css)
}

//0000000000000000000000000000000000000000000000000000000000000000000000000
// nao revisadas

func IsHide(el js.Value) bool {
	if IsElement(el) {
		isret := el.Get("style").Get("display").String()
		if isret == "none" {
			return true
		}
	}
	return false
}

func Focus(el js.Value) {
	if IsElement(el) {
		el.Call("focus")
	}
}

func Remove(el js.Value) {
	if IsElement(el) {
		el.Call("remove")
	}
}

func Clear(el js.Value) {
	if IsElement(el) {
		el.Set("textContent", " ")
	}
}

func ExecJs(script_js string) js.Value {
	return js.Global().Call("eval", "(function() {"+script_js+"; return 0;})()")
}

//~ package adom

//~ import "syscall/js"
//~ import "strings"
//~ import "../../../arus/web/awebw"

var IsMobile bool = false

func GetDoc() js.Value {
	return js.Global().Get("document")
}

func GetWindow() js.Value {
	return js.Global().Get("window")
}

func GetBody() js.Value {
	return js.Global().Get("document").Get("body")
}

func GetClientWidth() int {
	return js.Global().Get("document").Get("documentElement").Get("clientWidth").Int()
}

func GetClientHeight() int {
	return js.Global().Get("document").Get("documentElement").Get("clientHeight").Int()
}

func GetUserAgente() string {
	return js.Global().Get("navigator").Get("userAgent").String()
}

func GetPlataformaOs() string {
	return js.Global().Get("navigator").Get("platform").String()
}

func GetFromJs(script_js string) js.Value {
	return js.Global().Call("eval", "(function() { var retstr=''; retstr="+script_js+"; return String(retstr);})()")
}

func GetElementById(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

func DispatchEvent(el js.Value, eventName string) {
	if IsElement(el) {
		ev := js.Global().Get("Event").New(eventName)
		el.Call("dispatchEvent", ev)
	}
}

func AddEvent(el js.Value, event string, fn js.Func) { //("addEventListener", "click", procbtn_004)
	if IsElement(el) {
		//el.Call("removeEventListener", event, fn)
		el.Call("addEventListener", event, fn)
	}
}

func RemoveEvent(el js.Value, event string, fn js.Func) {
	if IsElement(el) {
		el.Call("removeEventListener", event, fn)
	}
}

func RemoveAtrr(el js.Value, atrnome string) {
	if IsElement(el) {
		el.Call("removeAttribute", atrnome)
	}
}

func NewEventTimer() js.Value {
	return js.Global().Get("setInterval")
}

func AddEventTimer(objtimer js.Value, fn js.Func, milisec int) {
	if objtimer.String() == "<function>" {
		objtimer.Invoke(fn, milisec)
	}
}

//~ func RemoveEventTimer(objtimer js.Value) {
//~ if objtimer.String() == "<function>" {
//~ js.Global().Get("clearInterval").Invoke(objtimer)
//~ }
//~ }

func GetEncodeURI(lUrl string) (retStr string) {
	retStr = js.Global().Call("encodeURIComponent", lUrl).String()
	return retStr
}

// retorna se Ã© 'INPUT' 'BUTTON'ETC..
func GetTipoName(el js.Value) string {
	if IsElement(el) {
		return el.Get("nodeName").String()
	}
	return ""
}

//~ func AppendChild(el, child_el js.Value) {
//~ if IsElement(el) && IsElement(child_el) {
//~ el.Call("appendChild", child_el)
//~ }
//~ }

//~ func CreateElement(el_parent js.Value, tag_name, id, class, conteudo string) js.Value {
//~ tag := js.Global().Get("document").Call("createElement", tag_name)
//~ if len(id) > 0 {
//~ tag.Call("setAttribute", "id", id)
//~ }
//~ if len(class) > 0 {
//~ tag.Call("setAttribute", "class", class)
//~ }
//~ if len(conteudo) > 0 {
//~ tag.Set("textContent", conteudo)
//~ }
//~ AppendChild(el_parent, tag)
//~ return tag
//~ }

//~ func CreateTextNode(el_parent js.Value, conteudo string) js.Value {
//~ tag := js.Global().Get("document").Call("createTextNode", conteudo)
//~ AppendChild(el_parent, tag)
//~ return tag
//~ }

func Div(el_parent js.Value, id, class, conteudo string) js.Value {
	return CreateElement(el_parent, "div", id, class, conteudo)
}

func Span(el_parent js.Value, id, class, conteudo string) js.Value {
	return CreateElement(el_parent, "span", id, class, conteudo)
}

func Input(el_parent js.Value, id, class, tipo string) js.Value { //
	input := CreateElement(el_parent, "input", id, class, "")
	if len(tipo) > 1 {
		input.Call("setAttribute", "type", tipo)
	} else {
		input.Call("setAttribute", "type", "text")
	}
	return input
}

func Br(el_parent js.Value) js.Value { //
	return CreateElement(el_parent, "BR", "", "", "")
}

func Create(el_parent js.Value, el_name, el_text string, el_atributos ...string) js.Value {
	u_el_name := strings.ToLower(el_name)
	html_tag := js.Global().Get("document").Call("createElement", u_el_name)
	nrat := len(el_atributos)
	if nrat > 0 {
		for p := 0; p < nrat; p += 2 {
			if p+1 >= nrat {
				break
			}
			html_tag.Call("setAttribute", el_atributos[p], el_atributos[p+1])
		}
	}

	if len(el_text) > 0 {
		if u_el_name == "input" || u_el_name == "textarea" {
			html_tag.Set("value", el_text)
		} else {
			html_tag.Set("textContent", el_text)
		}
	}
	AppendChild(el_parent, html_tag)
	return html_tag
}

/*************************** Table ***************/
func Table(el_parent js.Value, id, class string) js.Value {
	return CreateElement(el_parent, "TABLE", id, class, "")
}

// cria colunas na table
func TableAddThead(table js.Value) js.Value {
	thead := CreateElement(table, "THEAD", "", "", "")
	return CreateElement(thead, "TR", "", "", "")
}

// cria footer na table
func TableAddTfoot(table js.Value) js.Value {
	thead := CreateElement(table, "TFOOT", "", "", "")
	return CreateElement(thead, "TR", "", "", "")
}

//adiciona e cria TH texto nas colunas
func TableSetTextThead(thead js.Value, text, class string) js.Value {
	th := CreateElement(thead, "TH", "", class, "")
	CreateTextNode(th, text)
	return th
}

// cria o body da table
func TableAddTbody(table js.Value, id, class string) js.Value {
	return CreateElement(table, "TBODY", id, class, "")
}

// cria row na table
func TableAddRowTbody(tbody js.Value, id, class string) js.Value {
	return CreateElement(tbody, "TR", id, class, "")
}

//adiciona texto no Row e Tfoot
func TableSetTextRow(row js.Value, text, id, class string) js.Value {
	td := CreateElement(row, "TD", id, class, "")
	CreateTextNode(td, text)
	return td
}

//Get texto do Row TR
func TableGetTextRow(row js.Value, index int) (rwstr string) {
	rwstr = ""
	ntr := row.Get("cells").Length()
	if index >= 0 && ntr > 0 && index < ntr {
		rwstr = row.Get("cells").Call("item", index).Get("textContent").String()
	}
	return rwstr
}

//Get Element Ptr do Row TR col
func TableGetRowElement(row js.Value, index int) js.Value {
	var el js.Value
	ntr := row.Get("cells").Length()
	if index >= 0 && ntr > 0 && index < ntr {
		el = row.Get("cells").Call("item", index)
	}
	return el
}

//Get Update text do Row TR
func TableUpdateTextRow(row js.Value, index int, text string) {
	ntr := row.Get("cells").Length()
	if index >= 0 && ntr > 0 && index < ntr {
		row.Get("cells").Call("item", index).Set("textContent", text)
	}
}

//adiciona Ativado Desativado icon no Row
func TableSetTextRowStatus(row js.Value, text, id, class string) js.Value {
	td := CreateElement(row, "TD", id, class, "")
	if dice.StringToInt(text) < 1 {
		CreateElement(td, "i", "", "fas fa-check text-success", " ")
		text = " Ativado "
		AddClass(td, "text-success")
	} else {
		CreateElement(td, "i", "", "fas fa-times text-danger", " ")
		text = " Desativado "
		AddClass(td, "text-danger")
	}
	CreateTextNode(td, text)
	return td
}

// nr rows in tbory
func GetTbodyNrows(tbody js.Value) int {
	return tbody.Get("rows").Length()
}

// get index element row  do tbody
func TbodyGetRow(tbody js.Value, index int) js.Value {
	if index < 0 {
		index = 0
	}
	return tbody.Get("rows").Index(index)
}

func SetText(el js.Value, text string) {
	elnome := ""
	if IsElement(el) {
		elnome = GetTipoName(el)
		if elnome == "INPUT" || elnome == "TEXTAREA" {
			el.Set("value", text)
		} else {
			el.Set("textContent", text)
		}
	}
}

func GetText(el js.Value) (text string) {
	elnome := ""
	text = ""
	if IsElement(el) {
		elnome = GetTipoName(el)
		if elnome == "INPUT" || elnome == "TEXTAREA" {
			text = el.Get("value").String()
		} else {
			text = el.Get("textContent").String()
		}
	}
	return text
}

func GetChecked(el js.Value) bool {
	if IsElement(el) {
		return el.Get("checked").Bool()
	}
	return false
}

func SetChecked(el js.Value, stat bool) {
	if IsElement(el) {
		el.Set("checked", stat)
		//DispatchEvent(el, "change")
	}
}

func SetDisabled(el js.Value, stat bool) {
	if IsElement(el) {
		el.Set("disabled", stat)
	}
}

func SetAttr(el js.Value, atrnome, atrvalue string) {
	if IsElement(el) {
		el.Call("setAttribute", atrnome, atrvalue)
	}
}

func GetAttr(el js.Value, atrnome string) string {
	if IsElement(el) {
		att := el.Call("getAttribute", atrnome).String()
		if att != "<null>" {
			return att
		}
	}
	return ""
}

func GetId(el js.Value) (strid string) {
	strid = ""
	if IsElement(el) {
		strid = el.Get("id").String()
	}
	return strid
}

func SetStyle(el js.Value, stynome, styvalue string) {
	if IsElement(el) {
		el.Get("style").Set(stynome, styvalue)
	}
}

func CloseWindow() {
	ExecJs("window.close();")
}

/*********media ***********************/
//igual @media screen and (max-width: 700px)  no css
//Ex matchMedia;   "max-width: 700px", "orientation: portrait"
func IsWindowMedia(matchMedia string) bool {
	return js.Global().Get("window").Call("matchMedia", "("+matchMedia+"))").Get("matches").Bool()
}

/********************disable keys*****************/
func DisableKeysSystem(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 && !args[0].IsUndefined() {
		key_code := args[0].Get("keyCode").Int()
		ctrlKey := args[0].Get("ctrlKey").Bool()
		//altKey   := args[0].Get("altKey").Bool()

		//println(key_code)

		if ctrlKey && key_code == 74 { // ctrl J
			args[0].Call("preventDefault")
		}

		switch key_code {
		case 112, 116, 121, 123:
			args[0].Call("preventDefault")
		}
	}
	return nil
}
