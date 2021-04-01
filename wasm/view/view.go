package view

import (
	"Gestor-de-prazos/tool/web/dom"
	"syscall/js"
)

var el map[string]js.Value

func init() {
	el = make(map[string]js.Value, 0)
}

func Login() {

	el["auth"] = dom.CreateElement(dom.GetBody(),
		"div", "",
		"class", "auth")

	container := dom.CreateElement(el["auth"],
		"div", "",
		"class", "container pt-5")

	row := dom.CreateElement(container,
		"div", "",
		"class", "row")

	col := dom.CreateElement(row,
		"div", "",
		"class", "col-md-8 col-sm-12 mx-auto")

	card := dom.CreateElement(col,
		"div", "",
		"class", "card pt-4")

	cardBody := dom.CreateElement(card,
		"div", "",
		"class", "card-body")

	divImg := dom.Create(cardBody,
		"div", "",
		"class", "text-center mb-5")

	dom.Create(divImg,
		"img", "",
		"src", "static/assets/images/favicon.svg",
		"height", "48",
		"class", "mb-4")

	el["titlePrimary"] = dom.Create(divImg,
		"h3", "Entrar")

	pDivImg := dom.Create(divImg,
		"p", "")

	fontImg := dom.Create(pDivImg,
		"font", "",
		"style", "vertical-align: inherit;")

	el["titleSecondary"] = dom.Create(fontImg,
		"font", "Faça login para continuar no Gestor.",
		"style", "vertical-align: inherit;")

	el["form"] = dom.Create(cardBody,
		"div", "")

	formUser := dom.Create(el["form"],
		"div", "",
		"class", "form-group position-relative has-icon-left")

	dom.Create(formUser,
		"label", "Usuario",
		"for", "username")

	divInpUser := dom.Create(formUser,
		"div", "",
		"class", "position-relative")

	dom.Create(divInpUser,
		"input", "",
		"type", "text",
		"class", "form-control",
		"id", "username")

	divIconUser := dom.Create(divInpUser,
		"div", "",
		"class", "form-control-icon")

	dom.Create(divIconUser,
		"i", "",
		"class", "fi-rr-user")

	formPass := dom.Create(el["form"],
		"div", "",
		"class", "form-group position-relative has-icon-left")

	divForgotPass := dom.Create(formPass,
		"div", "",
		"class", "clearfix")

	dom.Create(divForgotPass,
		"label", "Senha",
		"for", "password")

	el["aForgotPass"] = dom.Create(divForgotPass,
		"a", "",
		"class", "float-end")
	el["aForgotPass"].Set("onclick", js.FuncOf(forgotPassword))

	dom.Create(el["aForgotPass"],
		"small", "Esqueceu sua senha?",
		"style", "cursor: pointer;")

	divInpPass := dom.Create(formPass,
		"div", "",
		"class", "position-relative")

	dom.CreateElement(divInpPass,
		"input", "",
		"type", "password",
		"class", "form-control",
		"id", "password")

	divIconPass := dom.Create(divInpPass,
		"div", "",
		"class", "form-control-icon")

	dom.Create(divIconPass,
		"i", "",
		"class", "fi-rr-key")

	divFormCheck := dom.Create(el["form"],
		"div", "",
		"class", "form-check clearfix my-4")

	divCheck := dom.Create(divFormCheck,
		"div", "",
		"class", "checkbox float-start")

	dom.Create(divCheck,
		"input", "",
		"type", "checkbox",
		"id", "checkbox1",
		"class", "form-check-input")

	dom.Create(divCheck,
		"label", "Lembre de mim",
		"for", "checkbox1")

	divRegister := dom.Create(divFormCheck,
		"div", "",
		"class", "float-end")

	el["aRegister"] = dom.Create(divRegister,
		"a", "Não tem conta?",
		"style", "cursor: pointer; color: #2472fd;")
	el["aRegister"].Set("onclick", js.FuncOf(register))

	divButton := dom.Create(el["form"],
		"div", "",
		"class", "clearfix")

	buttonLogin := Button(divButton, "Enviar", "btn btn-primary float-end")
	buttonLogin.Set("onclick", js.FuncOf(InitAplication))
}

func register(this js.Value, args []js.Value) interface{} {
	if dom.IsElement(this) {

		dom.SetText(el["titlePrimary"], "Inscrever-se")
		dom.SetText(el["titleSecondary"], "Por favor, preencha o formulário para se registrar.")

		dom.Clear(el["form"])

		form := dom.Create(el["form"],
			"div", "")

		Input(form, "PRIMEIRO NOME", "text")

		Input(form, "ÚLTIMO NOME", "text")

		Input(form, "NOME DO USUÁRIO", "text")

		Input(form, "O EMAIL", "text")

		divButton := dom.Create(el["form"],
			"div", "",
			"class", "clearfix")

		Button(divButton, "Enviar", "btn btn-primary float-end")
	}
	return nil
}

func forgotPassword(this js.Value, args []js.Value) interface{} {
	if dom.IsElement(this) {
		dom.SetText(el["titlePrimary"], "Esqueceu sua senha")
		dom.SetText(el["titleSecondary"], "Por favor, digite seu e-mail para receber o link de redefinição de senha.")

		dom.Clear(el["form"])

		formEmail := dom.Create(el["form"],
			"div", "",
			"class", "form-group")

		dom.Create(formEmail,
			"label", "Email",
			"for", "email")

		dom.Create(formEmail,
			"input", "",
			"type", "email",
			"class", "form-control",
			"id", "email")

		divButton := dom.Create(el["form"],
			"div", "",
			"class", "clearfix")

		Button(divButton, "Enviar", "btn btn-primary float-end")

	}
	return nil
}

func Button(element js.Value, nome, class string) (button js.Value) {
	button = dom.Create(element,
		"button", nome,
		"class", class)
	return
}

func Input(element js.Value, label, tipo string) (input js.Value) {

	divInput := dom.Create(element,
		"div", "",
		"class", "form-group")

	dom.Create(divInput,
		"label", label)

	input = dom.Create(divInput,
		"input", "",
		"type", "email",
		"class", "form-control")

	return
}

func InitAplication(this js.Value, args []js.Value) interface{} {
	if dom.IsElement(this) {
		dom.Remove(el["auth"])
		Aplication()
	}

	return nil
}

func Aplication() {

	app := dom.Create(dom.GetBody(),
		"div", "",
		"id", "app")

	sideBar(app)

	el["divMain"] = dom.Create(app,
		"div", "",
		"id", "main",
		"style", "margin-left: 0px; background: rgb(0 0 0 / 10%);")

	navBar()
	content()
	dataTable()
	footer()
}

func content() {

	mainContent := dom.Create(el["divMain"],
		"div", "",
		"class", "main-content container-fluid")

	pageTitle := dom.Create(mainContent,
		"div", "",
		"class", "page-title")

	row := dom.Create(pageTitle,
		"div", "",
		"class", "row")

	colTitle := dom.Create(row,
		"div", "",
		"class", "col-12 col-md-6 order-md-1 order-last")

	dom.Create(colTitle,
		"h3", "Teste")

	dom.Create(colTitle,
		"p", "There's a lot of form layout that you can use",
		"class", "text-subtitle text-muted")

	colBreadCrumb := dom.Create(row,
		"div", "",
		"class", "col-12 col-md-6 order-md-2 order-first")

	navBreadCrumb := dom.Create(colBreadCrumb,
		"nav", "",
		"aria-label", "breadcrumb",
		"class", "breadcrumb-header")

	olBreadCrumb := dom.Create(navBreadCrumb,
		"ol", "",
		"class", "breadcrumb")

	liPrimary := dom.Create(olBreadCrumb,
		"li", "",
		"class", "breadcrumb-item")

	dom.Create(liPrimary,
		"a", "primary")

	liSecondary := dom.Create(olBreadCrumb,
		"li", "",
		"class", "breadcrumb-item active")

	dom.Create(liSecondary,
		"a", "secondary")

	section := dom.Create(mainContent,
		"section", "")

	divRow := dom.Create(section,
		"div", "",
		"class", "row match-height")

	col := dom.Create(divRow,
		"div", "",
		"class", "col-12")

	card := dom.Create(col,
		"div", "",
		"class", "card")

	cardHeader := dom.Create(card,
		"div", "",
		"class", "card-header")

	dom.Create(cardHeader,
		"h4", "Cadastro",
		"class", "card-title")

	cardContent := dom.Create(card,
		"div", "",
		"class", "card-content")

	cardBody := dom.Create(cardContent,
		"div", "",
		"class", "card-body")

	form := dom.Create(cardBody,
		"div", "",
		"class", "form")

	rowForm := dom.Create(form,
		"div", "",
		"class", "row")

	colPessoa := dom.Create(rowForm,
		"div", "",
		"class", "col-md-6 col-12")
	Input(colPessoa, "Pessoa", "text")

	colVenc := dom.Create(rowForm,
		"div", "",
		"class", "col-md-6 col-12")
	Input(colVenc, "Vencimento", "text")

	colNotif := dom.Create(rowForm,
		"div", "",
		"class", "col-md-6 col-12")
	Input(colNotif, "Notificacao", "text")

	colTipoDoc := dom.Create(rowForm,
		"div", "",
		"class", "col-md-6 col-12")
	Input(colTipoDoc, "Tipo do documento", "text")

	colDescricao := dom.Create(rowForm,
		"div", "",
		"class", "col-md-6 col-12")
	Input(colDescricao, "Descricao", "text")

	colButton := dom.Create(rowForm,
		"div", "",
		"class", "col-12 d-flex justify-content-end")

	Button(colButton, "Adicionar", "btn btn-primary")
}

func navBar() {

	navBar := dom.Create(el["divMain"],
		"nav", "",
		"class", "navbar navbar-header navbar-expand navbar-light")

	aSideBar := dom.Create(navBar,
		"a", "",
		"class", "sidebar-toggler")
	aSideBar.Set("onclick", js.FuncOf(openSidebar))

	dom.Create(aSideBar,
		"i", "",
		"class", "fi-rr-menu-burger")

	navBarCollapse := dom.Create(navBar,
		"div", "",
		"class", "collapse navbar-collapse",
		"id", "navbarSupportedContent")

	ulNavBar := dom.Create(navBarCollapse,
		"ul", "",
		"class", "navbar-nav d-flex align-items-center navbar-light ms-auto")

	liDropDown := dom.Create(ulNavBar,
		"li", "",
		"class", "dropdown")

	aDropDown := dom.Create(liDropDown,
		"a", "",
		"data-bs-toogle", "dropdown",
		"class", "nav-link dropdown-toggle nav-link-lg nav-link-user",
		"aria-expanded", "false")

	divAvatar := dom.Create(aDropDown,
		"div", "",
		"class", "avatar me-1")

	dom.Create(divAvatar,
		"img", "",
		"src", "static/assets/images/avatar/avatar-s-1.png")

	divMenuDrop := dom.Create(liDropDown,
		"div", "",
		"class", "dropdown-menu dropdown-menu-end")

	aDropDown.Set("divDropdown", js.ValueOf(divMenuDrop))
	aDropDown.Set("onclick", js.FuncOf(ShowDropDown))

	itemDropConta := dom.Create(divMenuDrop,
		"a", "",
		"class", "dropdown-item")

	dom.Create(itemDropConta,
		"i", "",
		"class", "fi-rr-user")

	dom.Create(itemDropConta,
		"span", "  Conta",
		"style", "vertical-align: inherit;")

	itemDropMsg := dom.Create(divMenuDrop,
		"a", "",
		"class", "dropdown-item")

	dom.Create(itemDropMsg,
		"i", "",
		"class", "fi-rr-envelope")

	dom.Create(itemDropMsg,
		"span", "  Mensagens",
		"style", "vertical-align: inherit;")

	itemDropConfig := dom.Create(divMenuDrop,
		"a", "",
		"class", "dropdown-item")

	dom.Create(itemDropConfig,
		"i", "",
		"class", "fi-rr-settings-sliders")

	dom.Create(itemDropConfig,
		"span", "  Configurações",
		"style", "vertical-align: inherit;")

	itemDropLogout := dom.Create(divMenuDrop,
		"a", "",
		"class", "dropdown-item",
		"onclick", "window.location.reload();")

	dom.Create(itemDropLogout,
		"i", "",
		"class", "fi-rr-sign-out")

	dom.Create(itemDropLogout,
		"span", "  Sair",
		"style", "vertical-align: inherit;")
}

var dropdown bool = true

func ShowDropDown(this js.Value, args []js.Value) interface{} {
	if dom.IsElement(this) {
		if dropdown {
			dom.AddClass(this.Get("divDropdown"), "show")
			dropdown = false
			return nil
		}
		dom.RemoveClass(this.Get("divDropdown"), "show")
		dropdown = true
	}
	return nil
}

func footer() {
	footer := dom.Create(el["divMain"],
		"footer", "")

	divFooter := dom.Create(footer,
		"div", "",
		"class", "footer clearfix mb-0 text-muted")

	divFooterStart := dom.Create(divFooter,
		"div", "",
		"class", "float-start")

	pFooterStart := dom.Create(divFooterStart,
		"p", "",
		"style", "vertical-align: inherit;")

	dom.Create(pFooterStart,
		"font", "2021 © Gestor de Prazos",
		"style", "vertical-align: inherit;")

	divFooterEnd := dom.Create(divFooter,
		"div", "",
		"class", "float-end")

	pFooterEnd := dom.Create(divFooterEnd,
		"p", "",
		"style", "vertical-align: inherit;")

	fontEnd := dom.Create(pFooterEnd,
		"font", "Criado com ",
		"style", "vertical-align: inherit;")

	dom.Create(fontEnd,
		"i", "",
		"class", "fi-rr-heart",
		"style", "color: red;")

	dom.CreateTextNode(fontEnd, " por ")

	dom.Create(fontEnd,
		"a", "Maxwel",
		"style", "color: blue;")
}

func sideBar(element js.Value) {
	el["sidebar"] = dom.Create(element,
		"div", "",
		"id", "sidebar",
		"class", "")

	sidebarWrapper := dom.Create(el["sidebar"],
		"div", "",
		"class", "sidebar-wrapper active ps")

	sidebarHeader := dom.Create(sidebarWrapper,
		"div", "",
		"class", "sidebar-header")

	dom.Create(sidebarHeader,
		"img", "",
		"src", "static/assets/images/logo.svg")

	sidebarMenu := dom.Create(sidebarWrapper,
		"div", "",
		"class", "sidebar-menu")

	ulMenu := dom.Create(sidebarMenu,
		"ul", "",
		"class", "menu")

	liTitleMain := dom.Create(ulMenu,
		"li", "",
		"class", "sidebar-title")

	dom.Create(liTitleMain,
		"font", "Menu principal",
		"style", "vertical-align: inherit;")

	liItemPainel := dom.Create(ulMenu,
		"li", "",
		"class", "sidebar-item")

	aPainel := dom.Create(liItemPainel,
		"a", "",
		"class", "sidebar-link")

	fontPainel := dom.Create(aPainel,
		"font", "",
		"class", "text-secondary",
		"style", "vertical-align: inherit;")

	dom.Create(fontPainel,
		"i", "",
		"style", "color: #0d6efd; margin-right: 20px;",
		"class", "fi-rr-home")
	dom.CreateTextNode(fontPainel, "Painel")

	liItemTest := dom.Create(ulMenu,
		"li", "",
		"class", "sidebar-item has-sub")

	aTest := dom.Create(liItemTest,
		"a", "",
		"class", "sidebar-link")

	fontTest := dom.Create(aTest,
		"font", "",
		"class", "text-secondary",
		"style", "vertical-align: inherit;")

	dom.Create(fontTest,
		"i", "",
		"style", "color: #0d6efd; margin-right: 20px;",
		"class", "fi-rr-incognito")
	dom.CreateTextNode(fontTest, "Test")

	ulSubmenu := dom.Create(ulMenu,
		"ul", "",
		"class", "submenu")

	aTest.Set("submenu", js.ValueOf(ulSubmenu))
	aTest.Set("onclick", js.FuncOf(openSubMenu))

	litemTest01 := dom.Create(ulSubmenu,
		"li", "")

	aItemTest01 := dom.Create(litemTest01,
		"a", "")

	dom.Create(aItemTest01,
		"font", "test 01",
		"style", "vertical-align: inherit;")

	litemTest02 := dom.Create(ulSubmenu,
		"li", "")

	aItemTest02 := dom.Create(litemTest02,
		"a", "")

	dom.Create(aItemTest02,
		"font", "test 02",
		"style", "vertical-align: inherit;")

	buttonSidebarOut := dom.Create(sidebarWrapper,
		"button", "",
		"class", "sidebar-toggler btn x")
	buttonSidebarOut.Set("onclick", js.FuncOf(closeSidebar))

	dom.Create(buttonSidebarOut,
		"i", "",
		"class", "fi-rr-cross")

	railX := dom.Create(sidebarWrapper,
		"div", "",
		"class", "ps_rail-x",
		"style", "left: 0px; bottom: 0px;")

	dom.Create(railX,
		"div", "",
		"class", "ps_thumb-x",
		"tabindex", "0",
		"style", "left: 0px; width: 0px;")

	railY := dom.Create(sidebarWrapper,
		"div", "",
		"class", "ps_rail-Y",
		"style", "top: 0px; height: 698px; right: 0px;")

	dom.Create(railY,
		"div", "",
		"class", "ps_thumb-Y",
		"tabindex", "0",
		"style", "top: 0px; height: 0px;")
}

var sidebarStatus bool = true

func closeSidebar(this js.Value, args []js.Value) interface{} {
	if dom.IsElement(this) {
		dom.RemoveClass(el["sidebar"], "active")
		dom.SetAttr(el["divMain"], "style", "margin-left: 0px; background: rgb(0 0 0 / 10%);")
		sidebarStatus = true
	}
	return nil
}

func openSidebar(this js.Value, args []js.Value) interface{} {
	if dom.IsElement(this) {
		if sidebarStatus {
			dom.AddClass(el["sidebar"], "active")
			if width := dom.GetWindow().Get("innerWidth").Int(); width > 759 {
				dom.SetAttr(el["divMain"], "style", "margin-left: 260px; background: rgb(0 0 0 / 10%);")
			}
			sidebarStatus = false
			return nil
		}
		dom.RemoveClass(el["sidebar"], "active")
		dom.SetAttr(el["divMain"], "style", "margin-left: 0px; background: rgb(0 0 0 / 10%);")
		sidebarStatus = true
	}
	return nil
}

var subMenuStatus bool = true

func openSubMenu(this js.Value, args []js.Value) interface{} {
	if dom.IsElement(this) {
		if subMenuStatus {
			dom.AddClass(this.Get("submenu"), "active")
			subMenuStatus = false
			return nil
		}
		dom.RemoveClass(this.Get("submenu"), "active")
		subMenuStatus = true
	}
	return nil
}

func dataTable() {

	mainContent := dom.Create(el["divMain"],
		"div", "",
		"class", "main-content container-fluid")

	section := dom.Create(mainContent,
		"section", "",
		"class", "section")

	card := dom.Create(section,
		"div", "",
		"class", "card")

	dom.Create(card,
		"div", "Tarefas",
		"class", "card-header")

	cardBody := dom.Create(card,
		"div", "",
		"class", "card-body")

	dataTable := dom.Create(cardBody,
		"div", "",
		"class", "dataTable-wrapper dataTable-loading no-footer sortable searchable fixed-columns")

	dataTableTop := dom.Create(dataTable,
		"div", "",
		"class", "dataTable-top")

	dataTableSearch := dom.Create(dataTableTop,
		"div", "",
		"class", "dataTable-search")

	dom.Create(dataTableSearch,
		"input", "",
		"class", "dataTable-input",
		"placeholder", "Pesquisa...",
		"type", "text")

	dataTableContainer := dom.Create(dataTable,
		"div", "",
		"class", "dataTable-container")

	table := dom.Create(dataTableContainer,
		"table", "",
		"class", "table table-striped dataTable-table")

	thead := dom.Create(table,
		"thead", "")

	trHead := dom.Create(thead,
		"tr", "")

	dom.Create(trHead, "th", "Nome", "style", "width: 28.7517%;")
	dom.Create(trHead, "th", "Data vencimento", "style", "width: 28.7517%;")
	dom.Create(trHead, "th", "Data notificacao", "style", "width: 28.7517%;")
	dom.Create(trHead, "th", "Tipo do documento", "style", "width: 28.7517%;")
	dom.Create(trHead, "th", "Status", "style", "width: 28.7517%;")

	tBody := dom.Create(table,
		"thead", "")

	trBody := dom.Create(tBody,
		"tr", "")

	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	tdStatus := dom.Create(trBody, "td", "")
	dom.Create(tdStatus, "span", "Active", "class", "badge bg-success")

	trBody = dom.Create(tBody,
		"tr", "")

	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	tdStatus = dom.Create(trBody, "td", "")
	dom.Create(tdStatus, "span", "Active", "class", "badge bg-success")

	trBody = dom.Create(tBody,
		"tr", "")

	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	tdStatus = dom.Create(trBody, "td", "")
	dom.Create(tdStatus, "span", "Active", "class", "badge bg-success")

	trBody = dom.Create(tBody,
		"tr", "")

	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	tdStatus = dom.Create(trBody, "td", "")
	dom.Create(tdStatus, "span", "Active", "class", "badge bg-success")

	trBody = dom.Create(tBody,
		"tr", "")

	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	dom.Create(trBody, "td", "teste")
	tdStatus = dom.Create(trBody, "td", "")
	dom.Create(tdStatus, "span", "Active", "class", "badge bg-success")

	dataTableBottom := dom.Create(dataTable,
		"div", "",
		"class", "dataTable-bottom")

	dom.Create(dataTableBottom,
		"div", "Quantidade de itens 1",
		"class", "dataTable-info")
}
