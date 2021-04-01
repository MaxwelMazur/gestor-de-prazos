package dice

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
)

// SliceToString pega um lista e converte para string
func SliceToString(list []string) string {
	var completo string
	for i := 0; i < len(list); i++ {
		completo += list[i]
	}
	return completo
}

// IntToString converte um int64 para string
func IntToString(number int64) string {
	return strconv.FormatInt(number, 10)
}

// FloatToString converte um float64 para string
func FloatToString(number float64, floating int) string {
	return strconv.FormatFloat(number, 'f', floating, 64)
}

// StringToInt converte uma string para int64
func StringToInt(character string) int64 {
	number, err := strconv.ParseInt(character, 10, 64)
	if err != nil {
		return 0
	} else {
		return number
	}
}

// StringToFloat converte uma string para float64
func StringToFloat(character string) float64 {
	number, err := strconv.ParseFloat(character, 64)
	if err != nil {
		return 0.0
	} else {
		return number
	}
}

// InterfaceToString tipo interface para string
func InterfaceToString(value interface{}) string {
	character := fmt.Sprintf("%v", value)
	return character
}

// EncrypterString pega uma string legivel deixando nao legivel
func EncrypterString(text string) string {
	key, _ := hex.DecodeString("417275734b65792d323032302e4e6577")
	plaintext := []byte(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return ""
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	strStdoux := fmt.Sprintf("%x", ciphertext)
	return strStdoux
}

// DecrypterString pega uma string nao legivel deixando legivel
func DecrypterString(text string) string {
	key, _ := hex.DecodeString("417275734b65792d323032302e4e6577")
	ciphertext, _ := hex.DecodeString(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}
	if len(ciphertext) < aes.BlockSize {
		return ""
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	strStdoux := string(ciphertext)
	return strStdoux
}

//~ package awebw

//~ import (
//~ "fmt"
//~ "strconv"
//~ "strings"
//~ "syscall/js"
//~ "regexp"
//~ "math/rand"
//~ "time"
//~ "encoding/hex"
//~ "crypto/aes"
//~ "crypto/cipher"
//~ cryRand "crypto/rand"
//~ "io"
//~ )

//~ func Right(strvalue string, nlen int) string {
//~ stlen := len(strvalue)
//~ if stlen > 0 && nlen >= 0 {
//~ return strvalue[stlen-nlen:]
//~ } else {
//~ return ""
//~ }
//~ }

//~ func Left(strvalue string, nlen int) string {
//~ stlen := len(strvalue)
//~ if stlen > 0 && nlen >= 0 {
//~ return strvalue[:stlen-(stlen-nlen)]
//~ } else {
//~ return ""
//~ }
//~ }

//~ //MID
//~ func SubStr(strvalue string, idx, nlen int) string {
//~ stlen := len(strvalue)
//~ if stlen > 0 && nlen > 0 {
//~ if idx < 0 {
//~ idx = 0
//~ }
//~ nto := (idx + nlen)
//~ if nto > stlen {
//~ nto = stlen
//~ }
//~ return strvalue[idx:nto]
//~ }
//~ return ""
//~ }

//~ func Str(ivalue int) string {
//~ return strconv.Itoa(ivalue)
//~ }

//~ func StrF(fvalue float64, ndig int) (retstr string) {
//~ retstr = "0.0"
//~ if ndig < 1 {
//~ retstr = fmt.Sprintf("%g", fvalue)
//~ } else {
//~ retstr = fmt.Sprintf("%."+Str(ndig)+"f", fvalue)
//~ }
//~ return retstr
//~ }

//~ func InterfaceStr(value interface{}) (retstr string) {
//~ retstr = fmt.Sprintf("%v", value)
//~ return retstr
//~ }

//~ func Val(strvalue string) int {
//~ iret, err := strconv.Atoi(strvalue)
//~ if err != nil {
//~ return 0
//~ } else {
//~ return iret
//~ }
//~ }

//~ func ValF(strvalue string) float64 {
//~ fret, err := strconv.ParseFloat(strvalue, 64)
//~ if err != nil {
//~ return 0.0
//~ } else {
//~ return fret
//~ }
//~ }

//~ func pads(str string, n int) string {
//~ if n <= 0 {
//~ return ""
//~ }
//~ return strings.Repeat(str, n)
//~ }

//~ // Rset old
//~ func Lpad(str string, length int, pad string) string {
//~ return pads(pad, length-len(str)) + str
//~ }

//~ // Lset Old
//~ func Rpad(str string, length int, pad string) string {
//~ return str + pads(pad, length-len(str))
//~ }

//~ func FormatMoeda(strfloat string) string {
//~ if len(strfloat) < 1 {
//~ return "0.00"
//~ }
//~ if strings.Index(strfloat, ".") < 0 {
//~ strfloat += ".00"
//~ }

//~ strdec := strings.Split(strfloat, ".")
//~ decimais := strdec[0]
//~ if len(decimais) < 1 {
//~ decimais = "0"
//~ }
//~ centavos := strdec[1]

//~ var vlrfmt, vlr = "", ""
//~ n := 3

//~ for {
//~ if len(decimais) > 3 {
//~ vlrfmt = "." + decimais[len(decimais)-n:] + vlrfmt
//~ vlr = decimais[:len(decimais)-n]
//~ decimais = vlr
//~ } else {
//~ break
//~ }
//~ }

//~ vlrf := decimais + vlrfmt + "," + centavos
//~ return vlrf
//~ }

//~ func FormatMoedaF(fvalue float64, ndig int) string {
//~ strfloat := StrF(fvalue, ndig)
//~ return FormatMoeda(strfloat)
//~ }

//~ func RemoveMascara(text string) string {
//~ var text_clean string

//~ text = strings.Replace(text, ".", "", -1)
//~ text = strings.Replace(text, "-", "", -1)
//~ text = strings.Replace(text, "_", "", -1)
//~ text = strings.Replace(text, "/", "", -1)
//~ text = strings.Replace(text, ")", "", -1)
//~ text = strings.Replace(text, "(", "", -1)
//~ text = strings.Replace(text, ",", "", -1)
//~ text = strings.Replace(text, " ", "", -1)
//~ text = strings.Replace(text, ":", "", -1)
//~ text_clean = text
//~ return text_clean
//~ }

//~ func RemoveMascaraMoeda(text string) float64 {
//~ var text_clean string=""
//~ text = strings.Replace(text, ".", "", -1)
//~ text = strings.Replace(text, ",", ".", -1)
//~ text_clean = text
//~ return ValF(text_clean)
//~ }

//~ // enc
//~ func Encrypter(text string) string {
//~ key, _ := hex.DecodeString("417275734b65792d323032302e4e6577")
//~ plaintext := []byte(text)
//~ block, err := aes.NewCipher(key)
//~ if err != nil {
//~ return ""
//~ }
//~ ciphertext := make([]byte, aes.BlockSize+len(plaintext))
//~ iv := ciphertext[:aes.BlockSize]
//~ if _, err := io.ReadFull(cryRand.Reader, iv); err != nil {
//~ return ""
//~ }
//~ stream := cipher.NewCFBEncrypter(block, iv)
//~ stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
//~ strStdoux := fmt.Sprintf("%x", ciphertext)
//~ return strStdoux
//~ }

//~ //dec
//~ func Decrypter(text string) string {
//~ key, _ := hex.DecodeString("417275734b65792d323032302e4e6577")
//~ ciphertext, _ := hex.DecodeString(text)
//~ block, err := aes.NewCipher(key)
//~ if err != nil {
//~ return ""
//~ }
//~ if len(ciphertext) < aes.BlockSize {
//~ return ""
//~ }
//~ iv := ciphertext[:aes.BlockSize]
//~ ciphertext = ciphertext[aes.BlockSize:]
//~ stream := cipher.NewCFBDecrypter(block, iv)
//~ stream.XORKeyStream(ciphertext, ciphertext)
//~ strStdoux := string(ciphertext)
//~ return strStdoux
//~ }

//~ /*******************DATE TIME******************************/

//~ // Pega o tempo atual do computador
//~ func TimeSystem() time.Time {
//~ return time.Now()
//~ }

//~ //pega a data atual em formato banco de dados
//~ func Now2db() (vdata, vtime string) {
//~ t := TimeSystem()
//~ vdata = t.Format("2006-01-02")
//~ vtime = t.Format("15:04:05")
//~ return vdata, vtime
//~ }

//~ //pega a data atual em formato Brasileiro
//~ func Now2br() (vdate, vtime string) {

//~ t := TimeSystem()

//~ vdate = t.Format("02/01/2006")
//~ vtime = t.Format("15:04:05")

//~ return vdate, vtime
//~ }

//~ // Essa função converte formato a data BR para banco de dados
//~ func Br2db(databr string) (vdata, vtime string) {

//~ mskbr := "02/01/2006"

//~ if len(databr) > 10 {
//~ mskbr = "02/01/2006 15:04:05"
//~ }

//~ t, err := time.Parse(mskbr, databr)

//~ if err != nil {
//~ return "", ""
//~ }

//~ vdata = t.Format("2006-01-02")
//~ vtime = t.Format("15:04:05")

//~ return vdata, vtime
//~ }

//~ // Essa função converte formato de data DB para data Brasileira
//~ func Db2br(datadb string) (vdata, vtime string) {

//~ mskdb := "2006-01-02"

//~ if len(datadb) > 10 {
//~ mskdb = "2006-01-02 15:04:05"
//~ }

//~ t, err := time.Parse(mskdb, datadb)

//~ if err != nil {
//~ return "", ""
//~ }

//~ vdata = t.Format("02/01/2006")
//~ vtime = t.Format("15:04:05")

//~ return vdata, vtime
//~ }

//~ func ParseData(data_in string) (time.Time, bool) {
//~ mskdb := "2006-01-02"
//~ if len(data_in) > 10 {
//~ mskdb = "2006-01-02 15:04:05"
//~ }

//~ if strings.Index(data_in, "-") < 0 {
//~ mskdb = "02/01/2006"
//~ if len(data_in) > 10 {
//~ mskdb = "02/01/2006 15:04:05"
//~ }
//~ }

//~ pdtime, erri := time.Parse(mskdb, data_in)
//~ if erri != nil {
//~ return pdtime, false
//~ }
//~ return pdtime, true
//~ }

//~ func IsDataIgual(data_ini, data_fin string) bool {

//~ pdata_ini, isok := ParseData(data_ini)
//~ if isok == false {
//~ return false
//~ }

//~ pdata_fin, isok := ParseData(data_fin)
//~ if isok == false {
//~ return false
//~ }

//~ if pdata_ini.Equal(pdata_fin) { // igual
//~ return true
//~ }
//~ return false
//~ }

//~ func IsDataMaior(data_ini, data_fin string) bool {
//~ pdata_ini, isok := ParseData(data_ini)
//~ if isok == false {
//~ return false
//~ }

//~ pdata_fin, isok := ParseData(data_fin)
//~ if isok == false {
//~ return false
//~ }

//~ if pdata_ini.After(pdata_fin) { // Maior
//~ return true
//~ }
//~ return false
//~ }

//~ func IsDataMenor(data_ini, data_fin string) bool {
//~ pdata_ini, isok := ParseData(data_ini)
//~ if isok == false {
//~ return false
//~ }

//~ pdata_fin, isok := ParseData(data_fin)
//~ if isok == false {
//~ return false
//~ }

//~ if pdata_ini.Before(pdata_fin) { // Menor
//~ return true
//~ }
//~ return false
//~ }

//~ func DaysOfWeek(data_in string) int {
//~ mskdb := "2006-01-02"
//~ if len(data_in) > 10 {
//~ mskdb = "2006-01-02 15:04:05"
//~ }

//~ if strings.Index(data_in, "-") < 0 {
//~ mskdb = "02/01/2006"
//~ if len(data_in) > 10 {
//~ mskdb = "02/01/2006 15:04:05"
//~ }
//~ }

//~ pdtime, erri := time.Parse(mskdb, data_in)
//~ if erri != nil {
//~ return -1
//~ }
//~ Week := pdtime.Weekday()

//~ return int(Week)+1
//~ }

//~ func AddData(data_in string, ano,mes,dia int) (vdata, vtime string) {
//~ mskfull := "2006-01-02"
//~ mskdata := mskfull
//~ if len(data_in) > 10 {
//~ mskfull = "2006-01-02 15:04:05"
//~ }

//~ if strings.Index(data_in, "-") < 0 {
//~ mskfull = "02/01/2006"
//~ mskdata = mskfull
//~ if len(data_in) > 10 {
//~ mskfull = "02/01/2006 15:04:05"
//~ }
//~ }

//~ t, err := time.Parse(mskfull, data_in)

//~ if err != nil {
//~ return "", ""
//~ }

//~ nt := t.AddDate(ano,mes,dia)
//~ vdata = nt.Format(mskdata)
//~ vtime = nt.Format("15:04:05")
//~ return vdata, vtime
//~ }

//~ func AddMesBr(datainicialbr string, qtmes int) (vdata, vtime string) {

//~ mskdb := "02/01/2006"

//~ t, err := time.Parse(mskdb, datainicialbr)

//~ if err != nil {
//~ return "", ""
//~ }

//~ nt := t.AddDate(0, qtmes, 0)

//~ vdata = nt.Format("02/01/2006")
//~ vtime = nt.Format("15:04:05")

//~ return vdata, vtime
//~ }

//~ func RemoveDiasBr(databr string, qttdia int) (vdate, vtime string) {

//~ mascara := "02/01/2006"

//~ t, err := time.Parse(mascara, databr)

//~ if err != nil {
//~ return "", ""
//~ }

//~ nt := t.AddDate(0, 0, -qttdia)

//~ vdate = nt.Format("02/01/2006")
//~ vtime = nt.Format("15:04:05")

//~ return vdate, vtime
//~ }

//~ func SemanaBr(semana int) (rm string) {

//~ switch semana {
//~ case 0:
//~ rm = "Domingo"
//~ case 1:
//~ rm = "Segunda-Feira"
//~ case 2:
//~ rm = "Terça-Feira"
//~ case 3:
//~ rm = "Quarta-Feira"
//~ case 4:
//~ rm = "Quinta-Feira"
//~ case 5:
//~ rm = "Sexta-Feira"
//~ case 6:
//~ rm = "Sábado"
//~ default:
//~ rm = "Dia semana inexistente."
//~ }

//~ return rm
//~ }

//~ func SemanaExtenso(dia_semana int) string {
//~ switch dia_semana {
//~ case 1:
//~ return "Domingo"
//~ case 2:
//~ return "Segunda"
//~ case 3:
//~ return "Terça"
//~ case 4:
//~ return "Quarta"
//~ case 5:
//~ return "Quinta"
//~ case 6:
//~ return "Sexta"
//~ case 7:
//~ return "Sábado"
//~ }
//~ return ""
//~ }

//~ func MesBr(mes int) (rm string) {

//~ switch mes {
//~ case 1:
//~ rm = "Janeiro"
//~ case 2:
//~ rm = "Fevereiro"
//~ case 3:
//~ rm = "Março"
//~ case 4:
//~ rm = "Abril"
//~ case 5:
//~ rm = "Maio"
//~ case 6:
//~ rm = "Junho"
//~ case 7:
//~ rm = "Julho"
//~ case 8:
//~ rm = "Agosto"
//~ case 9:
//~ rm = "Setembro"
//~ case 10:
//~ rm = "Outubro"
//~ case 11:
//~ rm = "Novembro"
//~ case 12:
//~ rm = "Dezembro"
//~ }

//~ return rm
//~ }

//~ /*****************************************/

//~ func SetMasKMoedaInput(el_parent js.Value) {
//~ FuncInputMaskMoeda := js.FuncOf( func(this js.Value, args []js.Value) interface{} {
//~ if adom.IsElement(this) {
//~ valor := adom.GetText(el_parent)
//~ valorlimpo := RemoveMascara(valor)
//~ if len(valorlimpo) > 10 {
//~ valorlimpo = Left(valorlimpo,len(valorlimpo)-1)
//~ }
//~ valor_float := ValF(valorlimpo) / float64(100.0)
//~ adom.SetText(el_parent,FormatMoedaF(valor_float,2))
//~ }
//~ return nil
//~ })
//~ adom.AddEvent(el_parent,"input",FuncInputMaskMoeda)
//~ }

//~ func TimeSystem() js.Value {
//~ return js.Global().Get("Date").New() //Date(ano, mês - 1, dia, hora, minuto, segundo, milissegundo);
//semana := dt.Call("getDay").Int()// dia da semana
//~ }

//pega a data atual em formato banco de dados
//~ func Now2db() (vdata, vtime string) {
//~ dt := TimeSystem()
//~ ano := dt.Call("getFullYear").Int()
//~ mes := dt.Call("getMonth").Int() + 1
//~ dia := dt.Call("getDate").Int()
//~ hor := dt.Call("getHours").Int()
//~ min := dt.Call("getMinutes").Int()
//~ seg := dt.Call("getSeconds").Int()
//~ vdata = Str(ano) + "-" + Lpad(Str(mes), 2, "0") + "-" + Lpad(Str(dia), 2, "0")
//~ vtime = Lpad(Str(hor), 2, "0") + ":" + Lpad(Str(min), 2, "0") + ":" + Lpad(Str(seg), 2, "0")
//~ return vdata, vtime
//~ }

//pega a data atual em formato Brasileiro
//~ func Now2br() (vdata, vtime string) {
//~ dt := TimeSystem()
//~ ano := dt.Call("getFullYear").Int()
//~ mes := dt.Call("getMonth").Int() + 1
//~ dia := dt.Call("getDate").Int()
//~ hor := dt.Call("getHours").Int()
//~ min := dt.Call("getMinutes").Int()
//~ seg := dt.Call("getSeconds").Int()
//~ vdata = Lpad(Str(dia), 2, "0") + "/" + Lpad(Str(mes), 2, "0") + "/" + Str(ano)
//~ vtime = Lpad(Str(hor), 2, "0") + ":" + Lpad(Str(min), 2, "0") + ":" + Lpad(Str(seg), 2, "0")
//~ return vdata, vtime
//~ }

//~ func DateFormatDb(dt js.Value) (vdata, vtime string) {
//~ ano := dt.Call("getFullYear").Int()
//~ mes := dt.Call("getMonth").Int() + 1
//~ dia := dt.Call("getDate").Int()
//~ hor := dt.Call("getHours").Int()
//~ min := dt.Call("getMinutes").Int()
//~ seg := dt.Call("getSeconds").Int()
//~ vdata = Str(ano) + "-" + Lpad(Str(mes), 2, "0") + "-" + Lpad(Str(dia), 2, "0")
//~ vtime = Lpad(Str(hor), 2, "0") + ":" + Lpad(Str(min), 2, "0") + ":" + Lpad(Str(seg), 2, "0")
//~ return vdata, vtime
//~ }

// Essa função converte formato a data BR para banco de dados
//~ func Br2Db(databr string) (vdata, vtime string) {
//~ vtime ="00:00:00"
//~ vdata =""
//~ if len(databr) > 11 {
//~ horastmp := strings.Split(databr, " ")
//~ if len(horastmp)> 0{
//~ vtime = horastmp[1]
//~ }
//~ }
//~ databrs := strings.Split(databr, "/")

//~ if len(databrs) > 1 {
//~ ano := Val(Left(databrs[2],4))
//~ mes := Val(databrs[1])
//~ dia := Val(databrs[0])
//~ if (ano > 1000 && mes > 0 && mes < 13 && dia > 0 && dia < 32){
//~ vdata = Str(ano) + "-" + Lpad(Str(mes), 2, "0") + "-" + Lpad(Str(dia), 2, "0")
//~ }else{
//~ vtime=""
//~ }
//~ }else{
//~ vtime=""
//~ }
//~ return vdata, vtime
//~ }

//~ // Essa função converte formato de data DB para data Brasileira
//~ func Db2Br(datadb string) (vdata, vtime string) {
//~ vtime ="00:00:00"
//~ vdata =""
//~ if len(datadb) > 11 {
//~ horastmp := strings.Split(datadb, " ")
//~ if len(horastmp)> 0{
//~ vtime = horastmp[1]
//~ }
//~ }
//~ databrs := strings.Split(datadb, "-")

//~ if len(databrs) > 1 {
//~ ano := Val(Left(databrs[0],4))
//~ mes := Val(databrs[1])
//~ dia := Val(Left(databrs[2],2))
//~ if (ano > 1000 && mes > 0 && mes < 13 && dia > 0 && dia < 32){
//~ vdata = Lpad(Str(dia), 2, "0") + "/" + Lpad(Str(mes), 2, "0") + "/" + Str(ano)
//~ }else{
//~ vtime=""
//~ }
//~ }else{
//~ vtime=""
//~ }
//~ return vdata, vtime
//~ }

//~ // Parsing data para Objeto comparavel
//~ func ParseDb(datadb string) (js.Value, float64) {
//~ ano :=0
//~ mes :=0
//~ dia :=0
//~ hora:=0
//~ minu:=0
//~ segu:=0
//~ vtime :="00:00:00"
//~ if len(datadb) > 11 {
//~ horastmp := strings.Split(datadb, " ")
//~ if len(horastmp)> 0{
//~ vtime = horastmp[1]
//~ }
//~ }
//~ databrs := strings.Split(datadb, "-")
//~ timebrs := strings.Split(vtime, ":")

//~ if len(databrs) > 1 {
//~ ano = Val(Left(databrs[0],4))
//~ mes = Val(databrs[1])
//~ dia = Val(Left(databrs[2],2))
//~ if (ano > 1000 && mes > 0 && mes < 13 && dia > 0 && dia < 32){
//~ if len(timebrs)> 1 {
//~ hora = Val(Left(timebrs[0],2))
//~ minu = Val(Left(timebrs[1],2))
//~ segu = Val(Left(timebrs[2],2))
//~ }
//~ }else{
//~ return js.Null(), 0.0
//~ }
//~ }
//~ dd:=js.Global().Get("Date").New(ano,(mes-1),dia,hora,minu,segu,0)
//~ return dd, float64(dd.Call("getTime").Float())
//~ }

//~ /*************************** AJAX **********************************/
//~ // send post
//~ func HttpSendPost(phost string, req interface{}, retfn js.Func) {
//~ objrqs := js.Global().Get("JSON").Call("stringify", js.ValueOf(req))
//~ xhttp := js.Global().Get("XMLHttpRequest").New()
//~ //xhttp.Set("seenBytes", 0)
//~ //xhttp.Set("onreadystatechange", retfn)
//~ xhttp.Set("onload", retfn)
//~ //xhttp.Set("timeout", (60*60*1000))
//~ xhttp.Call("open", "POST", phost, true)
//~ xhttp.Call("setRequestHeader", "Content-Type", "application/json;charset=UTF-8")
//~ xhttp.Call("send", objrqs)

//~ println("Envio - Json")
//~ println(objrqs.String())
//~ println("-------------------------------------")
//~ }

//~ func HttpGetResponseJson(this js.Value, args []js.Value) (js.Value, bool) {
//~ xReadyState := this.Get("readyState").Int()
//~ xStatus := this.Get("status").Int()
//~ if xReadyState == 4 && xStatus == 200 {
//~ if !this.IsUndefined() && !this.IsNull() {
//~ return js.Global().Get("JSON").Call("parse", this.Get("responseText")), true
//~ }else{
//~ println("Sem Retorno Do Server...")
//~ }
//~ }
//~ return js.Null(), false
//~ }

//~ func HttpGetResponse(this js.Value, args []js.Value) (js.Value, bool) {
//~ xReadyState := this.Get("readyState").Int()
//~ xStatus := this.Get("status").Int()
//~ if xReadyState == 4 && xStatus == 200 {
//~ return this.Get("responseText"), true
//~ }
//~ return js.Null(), false
//~ }

//~ /*******************************************************************************/

//~ func HttpGetRetId(wreq js.Value) (retid int) {
//~ retid = 0
//~ if !wreq.IsUndefined() && !wreq.IsNull() {
//~ objs:=wreq.Get("ret_id")
//~ if !objs.IsUndefined() && !objs.IsNull() {
//~ retid = objs.Int()
//~ }
//~ }
//~ return retid
//~ }

//~ func GetTable2(wreq js.Value, nome_table string) (js.Value, int){
//~ var r,nrws int = 0,0

//~ if HttpGetRetId(wreq) < 1{
//~ return js.Null(), nrws
//~ }

//~ if len(nome_table) > 1 {
//~ nrtb := wreq.Get("tables").Length();
//~ if nrtb > 0 {
//~ for r = 0; r < nrtb; r++ {
//~ if wreq.Get("tables").Index(r).Get("nome").String() == nome_table {
//~ break
//~ }
//~ }
//~ }
//~ }

//~ if wreq.Get("tables").Index(r).Get("rows").String() != "<null>" {
//~ nrws = wreq.Get("tables").Index(r).Get("rows").Length()
//~ }
//~ return wreq.Get("tables").Index(r), nrws
//~ }

//~ func GetColTable(tbl js.Value, nome_col string) int {
//~ var r int = 0
//~ if len(nome_col) > 1 {
//~ nrtb := tbl.Get("cols").Length();
//~ if nrtb > 0 {
//~ for r = 0; r < nrtb; r++ {
//~ if tbl.Get("cols").Index(r).String() == nome_col {
//~ return r
//~ break
//~ }
//~ }
//~ }
//~ }

//~ println("Coluna [ "+nome_col+ "] Inexitente!");
//~ return -1
//~ }

//~ func GetRowValue(tbl js.Value,idxcol, rowidx int) (rstr string) {
//~ rstr = ""
//~ if idxcol >= 0 {
//~ rows := tbl.Get("rows").Index(rowidx)
//~ rstr = rows.Get("r").Index(idxcol).String()
//~ }
//~ return rstr
//~ }

//~ func TableGetParams(tbl js.Value, params string) (rstr string) {
//~ rstr=""
//~ if !tbl.IsUndefined() && !tbl.IsNull(){
//~ rstr = tbl.Get("params").Get(params).String()
//~ }
//~ return rstr
//~ }

//~ func RqGetParams(wreq js.Value, params string) (rstr string) {
//~ rstr=""
//~ if !wreq.IsUndefined() && !wreq.IsNull(){
//~ rstr = wreq.Get("params").Get(params).String()
//~ }
//~ return rstr
//~ }
//~ /*******************WEBSOCKETS********************************/

//~ func CreateSocket(wshost string, onopenfn, onmessagefn, onclosefn, onerrorfn js.Func) (js.Value, bool) {
//~ websocket := js.Global().Get("WebSocket").New(wshost)
//~ ////if js.Global().String() == "<object>" {
//~ if websocket.String() == "<object>" {
//~ websocket.Set("onopen", onopenfn)
//~ websocket.Set("onmessage", onmessagefn)
//~ websocket.Set("onclose", onclosefn)
//~ websocket.Set("onerror", onerrorfn)
//~ return websocket, true
//~ } else {
//~ return js.Null(), false
//~ }
//~ }

//~ func GetIsMobile() (int, string) {
//~ UserAgente:= js.Global().Get("navigator").Get("userAgent").String()
//~ agent := strings.ToLower(UserAgente)
//~ //println(agent)
//~ if strings.Index(agent, "iphone")>0 {
//~ return 2, "IPHONE"
//~ }
//~ if strings.Index(agent, "ipad")>0 {
//~ return 2, "IPAD"
//~ }
//~ if strings.Index(agent, "ipod")>0 {
//~ return 2, "IPOD"
//~ }

//~ if strings.Index(agent, "android")>0 {
//~ return 1, "ANDROID"
//~ }
//~ if strings.Index(agent, "webos")>0 {
//~ return 1 , "WEBOS"
//~ }
//~ if strings.Index(agent, "blackberry")>0 {
//~ return 1, "ANDROID"
//~ }
//~ if strings.Index(agent, "blazer")>0 {
//~ return 1,"ANDROID"
//~ }
//~ if strings.Index(agent, "kindle")>0 {
//~ return 1,"KINDLE"
//~ }
//~ if strings.Index(agent, "symbianos")>0 {
//~ return 1,"ANDROID"
//~ }
//~ if strings.Index(agent, "opera mini")>0 {
//~ return 1,"ANDROID"
//~ }

//~ return 0, "DESKTOP"
//~ }

//~ func ValidarEmail(email string) bool {
//~ Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
//~ return Re.MatchString(email)
//~ }

//~ func ValidaCPF(cpf string) bool {
//~ cpf = strings.Replace(cpf, ".", "", -1)
//~ cpf = strings.Replace(cpf, "-", "", -1)
//~ if len(cpf) != 11 {
//~ return false
//~ }
//~ var eq bool
//~ var dig string
//~ for _, val := range cpf {
//~ if len(dig) == 0 {
//~ dig = string(val)
//~ }
//~ if string(val) == dig {
//~ eq = true
//~ continue
//~ }
//~ eq = false
//~ break
//~ }
//~ if eq {
//~ return false
//~ }

//~ i := 10
//~ sum := 0
//~ for index := 0; index < len(cpf)-2; index++ {
//~ pos, _ := strconv.Atoi(string(cpf[index]))
//~ sum += pos * i
//~ i--
//~ }

//~ prod := sum * 10
//~ mod := prod % 11
//~ if mod == 10 {
//~ mod = 0
//~ }
//~ digit1, _ := strconv.Atoi(string(cpf[9]))
//~ if mod != digit1 {
//~ return false
//~ }
//~ i = 11
//~ sum = 0
//~ for index := 0; index < len(cpf)-1; index++ {
//~ pos, _ := strconv.Atoi(string(cpf[index]))
//~ sum += pos * i
//~ i--
//~ }
//~ prod = sum * 10
//~ mod = prod % 11
//~ if mod == 10 {
//~ mod = 0
//~ }
//~ digit2, _ := strconv.Atoi(string(cpf[10]))
//~ if mod != digit2 {
//~ return false
//~ }

//~ return true
//~ }

//~ func ValidaCNPJ(cnpj string) bool {
//~ cnpj = strings.Replace(cnpj, ".", "", -1)
//~ cnpj = strings.Replace(cnpj, "-", "", -1)
//~ cnpj = strings.Replace(cnpj, "/", "", -1)
//~ if len(cnpj) != 14 {
//~ return false
//~ }

//~ algs := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
//~ var algProdCpfDig1 = make([]int, 12, 12)
//~ for key, val := range algs {
//~ intParsed, _ := strconv.Atoi(string(cnpj[key]))
//~ sumTmp := val * intParsed
//~ algProdCpfDig1[key] = sumTmp
//~ }
//~ sum := 0
//~ for _, val := range algProdCpfDig1 {
//~ sum += val
//~ }
//~ digit1 := sum % 11
//~ if digit1 < 2 {
//~ digit1 = 0
//~ } else {
//~ digit1 = 11 - digit1
//~ }
//~ char12, _ := strconv.Atoi(string(cnpj[12]))
//~ if char12 != digit1 {
//~ return false
//~ }
//~ algs = append([]int{6}, algs...)

//~ var algProdCpfDig2 = make([]int, 13, 13)
//~ for key, val := range algs {
//~ intParsed, _ := strconv.Atoi(string(cnpj[key]))

//~ sumTmp := val * intParsed
//~ algProdCpfDig2[key] = sumTmp
//~ }
//~ sum = 0
//~ for _, val := range algProdCpfDig2 {
//~ sum += val
//~ }

//~ digit2 := sum % 11
//~ if digit2 < 2 {
//~ digit2 = 0
//~ } else {
//~ digit2 = 11 - digit2
//~ }
//~ char13, _ := strconv.Atoi(string(cnpj[13]))
//~ if char13 != digit2 {
//~ return false
//~ }

//~ return true
//~ }

//~ // não usar em laço
//~ func GetUniqID() string {
//~ const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//~ rr := rand.NewSource(time.Now().UnixNano())
//~ rf := rand.New(rr)
//~ a:= string(letterBytes[rf.Intn(len(letterBytes))])
//~ b:= string(letterBytes[rf.Intn(len(letterBytes))])
//~ nkstr := a+Str(rf.Intn(9000))+b+Str(rf.Intn(7000))
//~ return nkstr
//~ }

//~ func DownLoadFilePdf(wreq js.Value) {
//~ filename := "P"+GetUniqID()+".pdf"
//~ go func() { js.Global().Call("DownLoadFilex","data:application/pdf;base64,"+wreq.String(), filename, "application/pdf") } ()
//~ }

//~ func ValidaNumeroCelular(numero string) bool {
//~ nmr := regexp.MustCompile(`(\(?\d{2}\)?\s)?(\d{4,5}\-\d{4})`)
//~ return nmr.MatchString(numero)
//~ }

//~ //1.234,00
//~ func IsCurrency(numero string) bool {
//~ only_cur := regexp.MustCompile(`^\d{1,3}(?:\.\d{3})*.\d{2}$`)
//~ if only_cur.MatchString( numero){
//~ return true
//~ }
//~ return false
//~ }
//~ //1.234,00
//~ func IsFloat(numero string) bool {
//~ only_cur := regexp.MustCompile(`^[+-]?([0-9]+([.][0-9]*)?|[.][0-9]+)$`)
//~ if only_cur.MatchString( numero){
//~ return true
//~ }
//~ return false
//~ }

//~ //123456
//~ func IsNumber(numero string) bool {
//~ only_nr := regexp.MustCompile(`^\d*$`)
//~ if only_nr.MatchString(numero) {
//~ return true
//~ }
//~ return false
//~ }

//~ func IsDataBr(data string) bool {
//~ d := regexp.MustCompile(`[0-9]+/[0-9]+/[0-9]{2,4}$`)
//~ return d.MatchString(data)
//~ }

//~ func main() {

//~ tb1 := NewTable("autenticar") //object{}
//~ rw := TableAddRwCols(tb1, "modalidade_id", "descricao", "max_colocacao")

//~ for i := 0; i < 22; i++ {
//~ rw = TableAddRows(tb1, rw, Str(i), Str(i+33), Str(i+9))
//~ }
//~ rw = TableAddRows(tb1, rw, "10", "20", "30")
//~ rw = TableAddRows(tb1, rw, "11", "21", "31")
//TableAddRows

//~ tb2 := NewTable("autenticar2") //object{}
//~ rw2 := TableAddRwCols(tb2, "modalidade_id", "descricao", "max_colocacao")
//~ rw2 = TableAddRows(tb2, rw2, "12", "22", "32")
//~ rw2 = TableAddRows(tb2, rw2, "12", "22", "30")
//~ rw2 = TableAddRows(tb2, rw2, "12", "22", "32")

//~ rqx := CreateAddTables("suniapi.GetContas", tb1,tb2)

//~ rjson, _ := json.MarshalIndent(rqx, " ", " ")
//~ //rjson, _ := json.Marshal(rqx)
//~ fmt.Println(string(rjson))

//~ fmt.Println(StrF(1.7890,0))
//~ fmt.Println(FormatMoedaF(0.89))

//~ package autil

//~ import (
//~ "bufio"
//~ "crypto/aes"
//~ "crypto/cipher"
//~ "crypto/rand"
//~ "encoding/hex"
//~ "fmt"
//~ "io"
//~ "log"
//~ "os"
//~ "os/exec"
//~ "strconv"
//~ "strings"
//~ "time"
//~ "regexp"
//~ "net/smtp"
//~ )

//~ var (
//~ EMAIL_USER string = "mailservice@suniweb.com"
//~ EMAIL_PWD  string = Decrypter("4542764d06b92004cd1f1b8888dda9011678482efe32c43604500250")
//~ EMAIL_HOST string = "mail.suniweb.com"// use porta 587
//~ )

//~ func Right(strvalue string, nlen int) string {
//~ stlen := len(strvalue)
//~ if stlen > 0 && nlen >= 0  && stlen > nlen{
//~ return strvalue[stlen-nlen:]
//~ } else {
//~ return strvalue
//~ }
//~ }

//~ func Left(strvalue string, nlen int) string {
//~ stlen := len(strvalue)
//~ if stlen > 0 && nlen >= 0 && stlen > nlen {
//~ return strvalue[:stlen-(stlen-nlen)]
//~ } else {
//~ return strvalue
//~ }
//~ }

//~ //MID
//~ func SubStr(strvalue string, idx, nlen int) string {
//~ stlen := len(strvalue)
//~ if stlen > 0 && nlen > 0 {
//~ if idx < 0 {
//~ idx = 0
//~ }
//~ nto := (idx + nlen)
//~ if nto > stlen {
//~ nto = stlen
//~ }
//~ return strvalue[idx:nto]
//~ }
//~ return ""
//~ }

//~ func Str(ivalue int64) string {
//~ return strconv.FormatInt(ivalue, 10)
//~ }

//~ func StrF(fvalue float64, ndig int) string {
//~ return strconv.FormatFloat(fvalue, 'f', ndig, 64)
//~ }

//~ func Val(strvalue string) int64 {
//~ iret, err := strconv.ParseInt(strvalue, 10, 64)
//~ if err != nil {
//~ return 0
//~ } else {
//~ return iret
//~ }
//~ }

//~ func ValF(strvalue string) float64 {
//~ fret, err := strconv.ParseFloat(strvalue, 64)
//~ if err != nil {
//~ return 0.0
//~ } else {
//~ return fret
//~ }
//~ }

//~ // Pega o tempo atual do computador assim que executado
//~ func TimeSystem() time.Time {
//~ return time.Now()
//~ }

//~ //pega a data atual em formato banco de dados
//~ func Now2db() (vdata, vtime string) {
//~ t := TimeSystem()
//~ vdata = t.Format("2006-01-02")
//~ vtime = t.Format("15:04:05")
//~ return vdata, vtime
//~ }

//~ //pega a data atual em formato Brasileiro
//~ func Now2br() (vdate, vtime string) {

//~ t := TimeSystem()

//~ vdate = t.Format("02/01/2006")
//~ vtime = t.Format("15:04:05")

//~ return vdate, vtime
//~ }

//~ // Essa função converte formato a data BR para banco de dados
//~ func Br2db(databr string) (vdata, vtime string) {

//~ mskbr := "02/01/2006"

//~ if len(databr) > 10 {
//~ mskbr = "02/01/2006 15:04:05"
//~ }

//~ t, err := time.Parse(mskbr, databr)

//~ if err != nil {
//~ return "", ""
//~ }

//~ vdata = t.Format("2006-01-02")
//~ vtime = t.Format("15:04:05")

//~ return vdata, vtime
//~ }

//~ // Essa função converte formato de data DB para data Brasileira
//~ func Db2br(datadb string) (vdata, vtime string) {

//~ mskdb := "2006-01-02"

//~ if len(datadb) > 10 {
//~ mskdb = "2006-01-02 15:04:05"
//~ }

//~ t, err := time.Parse(mskdb, datadb)

//~ if err != nil {
//~ return "", ""
//~ }

//~ vdata = t.Format("02/01/2006")
//~ vtime = t.Format("15:04:05")

//~ return vdata, vtime
//~ }

//~ func ParseData(data_in string) (time.Time, bool) {
//~ mskdb := "2006-01-02"
//~ if len(data_in) > 10 {
//~ mskdb = "2006-01-02 15:04:05"
//~ }

//~ if strings.Index(data_in, "-") < 0 {
//~ mskdb = "02/01/2006"
//~ if len(data_in) > 10 {
//~ mskdb = "02/01/2006 15:04:05"
//~ }
//~ }

//~ pdtime, erri := time.Parse(mskdb, data_in)
//~ if erri != nil {
//~ return pdtime, false
//~ }
//~ return pdtime, true
//~ }

//~ func IsDataIgual(data_ini, data_fin string) bool {

//~ pdata_ini, isok := ParseData(data_ini)
//~ if isok == false {
//~ return false
//~ }

//~ pdata_fin, isok := ParseData(data_fin)
//~ if isok == false {
//~ return false
//~ }

//~ if pdata_ini.Equal(pdata_fin) { // igual
//~ return true
//~ }
//~ return false
//~ }

//~ func IsDataMaior(data_ini, data_fin string) bool {
//~ pdata_ini, isok := ParseData(data_ini)
//~ if isok == false {
//~ return false
//~ }

//~ pdata_fin, isok := ParseData(data_fin)
//~ if isok == false {
//~ return false
//~ }

//~ if pdata_ini.After(pdata_fin) { // Maior
//~ return true
//~ }
//~ return false
//~ }

//~ func IsDataMenor(data_ini, data_fin string) bool {
//~ pdata_ini, isok := ParseData(data_ini)
//~ if isok == false {
//~ return false
//~ }

//~ pdata_fin, isok := ParseData(data_fin)
//~ if isok == false {
//~ return false
//~ }

//~ if pdata_ini.Before(pdata_fin) { // Menor
//~ return true
//~ }
//~ return false
//~ }

//~ func DaysOfWeek(data_in string) (dayOfWeek int64) {
//~ mskdb := "2006-01-02"
//~ if len(data_in) > 10 {
//~ mskdb = "2006-01-02 15:04:05"
//~ }

//~ if strings.Index(data_in, "-") < 0 {
//~ mskdb = "02/01/2006"
//~ if len(data_in) > 10 {
//~ mskdb = "02/01/2006 15:04:05"
//~ }
//~ }

//~ pdtime, erri := time.Parse(mskdb, data_in)
//~ if erri != nil {
//~ return -1
//~ }
//~ Week := pdtime.Weekday()

//~ dayOfWeek = int64(Week) + 1
//~ return dayOfWeek
//~ }

//~ func AddData(data_in string, ano,mes,dia int) (vdata, vtime string) {
//~ mskfull := "2006-01-02"
//~ mskdata := mskfull
//~ if len(data_in) > 10 {
//~ mskfull = "2006-01-02 15:04:05"
//~ }

//~ if strings.Index(data_in, "-") < 0 {
//~ mskfull = "02/01/2006"
//~ mskdata = mskfull
//~ if len(data_in) > 10 {
//~ mskfull = "02/01/2006 15:04:05"
//~ }
//~ }

//~ t, err := time.Parse(mskfull, data_in)

//~ if err != nil {
//~ return "", ""
//~ }

//~ nt := t.AddDate(ano,mes,dia)
//~ vdata = nt.Format(mskdata)
//~ vtime = nt.Format("15:04:05")
//~ return vdata, vtime
//~ }

//~ func AddDiasBr(datainicialbr string, qtddia int) (vdata, vtime string) {

//~ mskdb := "02/01/2006"

//~ t, err := time.Parse(mskdb, datainicialbr)

//~ if err != nil {
//~ return "", ""
//~ }

//~ nt := t.AddDate(0, 0, qtddia)

//~ vdata = nt.Format("02/01/2006")
//~ vtime = nt.Format("15:04:05")

//~ return vdata, vtime
//~ }

//~ func AddMesBr(datainicialbr string, qtmes int) (vdata, vtime string) {

//~ mskdb := "02/01/2006"

//~ t, err := time.Parse(mskdb, datainicialbr)

//~ if err != nil {
//~ return "", ""
//~ }

//~ nt := t.AddDate(0, qtmes, 0)

//~ vdata = nt.Format("02/01/2006")
//~ vtime = nt.Format("15:04:05")

//~ return vdata, vtime
//~ }

//~ func RemoveDiasBr(databr string, qttdia int) (vdate, vtime string) {

//~ mascara := "02/01/2006"

//~ t, err := time.Parse(mascara, databr)

//~ if err != nil {
//~ return "", ""
//~ }

//~ nt := t.AddDate(0, 0, -qttdia)

//~ vdate = nt.Format("02/01/2006")
//~ vtime = nt.Format("15:04:05")

//~ return vdate, vtime
//~ }

//~ func SemanaBr(semana int) (rm string) {

//~ switch semana {
//~ case 0:
//~ rm = "Domingo"
//~ case 1:
//~ rm = "Segunda-Feira"
//~ case 2:
//~ rm = "Terça-Feira"
//~ case 3:
//~ rm = "Quarta-Feira"
//~ case 4:
//~ rm = "Quinta-Feira"
//~ case 5:
//~ rm = "Sexta-Feira"
//~ case 6:
//~ rm = "Sábado"
//~ default:
//~ rm = "Dia semana inexistente."
//~ }

//~ return rm
//~ }

//~ func SemanaExtenso(dia_semana int64) string {
//~ switch dia_semana {
//~ case 1:
//~ return "Domingo"
//~ case 2:
//~ return "Segunda"
//~ case 3:
//~ return "Terça"
//~ case 4:
//~ return "Quarta"
//~ case 5:
//~ return "Quinta"
//~ case 6:
//~ return "Sexta"
//~ case 7:
//~ return "Sábado"
//~ }
//~ return ""
//~ }

//~ func MesBr(mes int) (rm string) {

//~ switch mes {
//~ case 1:
//~ rm = "Janeiro"
//~ case 2:
//~ rm = "Fevereiro"
//~ case 3:
//~ rm = "Março"
//~ case 4:
//~ rm = "Abril"
//~ case 5:
//~ rm = "Maio"
//~ case 6:
//~ rm = "Junho"
//~ case 7:
//~ rm = "Julho"
//~ case 8:
//~ rm = "Agosto"
//~ case 9:
//~ rm = "Setembro"
//~ case 10:
//~ rm = "Outubro"
//~ case 11:
//~ rm = "Novembro"
//~ case 12:
//~ rm = "Dezembro"
//~ }

//~ return rm
//~ }

//~ /*****************************************/

//~ func RemoveMascaraMoeda(text string) float64 {
//~ var text_clean string

//~ text = strings.Replace(text, ".", "", -1)
//~ text = strings.Replace(text, ",", ".", -1)
//~ text_clean = text
//~ return ValF(text_clean)
//~ }

//~ func RemoveMascara(text string) string {
//~ var text_clean string

//~ text = strings.Replace(text, ".", "", -1)
//~ text = strings.Replace(text, "-", "", -1)
//~ text = strings.Replace(text, "_", "", -1)
//~ text = strings.Replace(text, "/", "", -1)
//~ text = strings.Replace(text, ")", "", -1)
//~ text = strings.Replace(text, "(", "", -1)
//~ text = strings.Replace(text, ",", "", -1)
//~ text = strings.Replace(text, " ", "", -1)
//~ text = strings.Replace(text, ":", "", -1)
//~ text_clean = text
//~ return text_clean
//~ }

//~ func pads(str string, n int) string {
//~ if n <= 0 {
//~ return ""
//~ }
//~ return strings.Repeat(str, n)
//~ }

//~ // Left left-pads the string with pad up to len runes
//~ func Lpad(str string, length int, pad string) string {
//~ return pads(pad, length-len(str)) + str
//~ }

//~ // Right right-pads the string with pad up to len runes
//~ func Rpad(str string, length int, pad string) string {
//~ return str + pads(pad, length-len(str))
//~ }

//~ func SalvaLog(msg string) {
//~ pexe, _ := os.Executable()
//~ lst := strings.Count(pexe, "/")
//~ anome := strings.Split(pexe, "/")
//~ nomefull := os.TempDir() + "/" + anome[lst] + ".log"
//~ f, err := os.OpenFile(nomefull, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
//~ defer f.Close()
//~ if err == nil {
//~ //f.Seek(0, 1)
//~ f.WriteString(msg + "\n")
//~ log.SetOutput(f)
//~ }
//~ }

//~ func FormatMoeda(strfloat string) string {
//~ if len(strfloat) < 1 {
//~ return "0.00"
//~ }
//~ if strings.Index(strfloat, ".") < 0 {
//~ strfloat += ".00"
//~ }

//~ strdec := strings.Split(strfloat, ".")
//~ decimais := strdec[0]
//~ if len(decimais) < 1 {
//~ decimais = "0"
//~ }
//~ centavos := strdec[1]

//~ var vlrfmt, vlr = "", ""
//~ n := 3

//~ for {
//~ if len(decimais) > 3 {
//~ vlrfmt = "." + decimais[len(decimais)-n:] + vlrfmt
//~ vlr = decimais[:len(decimais)-n]
//~ decimais = vlr
//~ } else {
//~ break
//~ }
//~ }

//~ vlrf := decimais + vlrfmt + "," + centavos
//~ return vlrf
//~ }

//~ func FormatMoedaF(fvalue float64, ndig int) string {
//~ strfloat := StrF(fvalue, ndig)
//~ return FormatMoeda(strfloat)
//~ }

//~ // enc
//~ func Encrypter(text string) string {
//~ key, _ := hex.DecodeString("417275734b65792d323032302e4e6577")
//~ plaintext := []byte(text)
//~ block, err := aes.NewCipher(key)
//~ if err != nil {
//~ return ""
//~ }
//~ ciphertext := make([]byte, aes.BlockSize+len(plaintext))
//~ iv := ciphertext[:aes.BlockSize]
//~ if _, err := io.ReadFull(rand.Reader, iv); err != nil {
//~ return ""
//~ }
//~ stream := cipher.NewCFBEncrypter(block, iv)
//~ stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
//~ strStdoux := fmt.Sprintf("%x", ciphertext)
//~ return strStdoux
//~ }

//~ //dec
//~ func Decrypter(text string) string {
//~ key, _ := hex.DecodeString("417275734b65792d323032302e4e6577")
//~ ciphertext, _ := hex.DecodeString(text)
//~ block, err := aes.NewCipher(key)
//~ if err != nil {
//~ return ""
//~ }
//~ if len(ciphertext) < aes.BlockSize {
//~ return ""
//~ }
//~ iv := ciphertext[:aes.BlockSize]
//~ ciphertext = ciphertext[aes.BlockSize:]
//~ stream := cipher.NewCFBDecrypter(block, iv)
//~ stream.XORKeyStream(ciphertext, ciphertext)
//~ strStdoux := string(ciphertext)
//~ return strStdoux
//~ }

//~ //cryptografa com data dia pra expirar
//~ func EncrypterExData(text string) string{
//~ d,_ := Now2db()
//~ ntext:=d+";"+text
//~ return Encrypter(ntext)
//~ }

//~ //Descryptografa se dia hoje == ao da key
//~ func DecrypterExData(text string) (rStr string){
//~ rStr=""
//~ //d,_ := Now2db()
//~ ktext:=Decrypter(text)
//~ ntext:=strings.Split(ktext,";")
//~ if len(ntext) > 1 {
//~ //if d==ntext[0] {// se dia for igual a hoje OK else retorna vasio rStr
//~ rStr=ntext[1]
//~ // }
//~ }
//~ return rStr
//~ }

//~ // pega serial numero hd no Linux
//~ func GetSerialHD() (serialhd string) {
//~ //cmdName := "ls -all"
//~ serialhd = ""
//~ cmdName := "udevadm info --query=all --name=/dev/sda"
//~ cmdArgs := strings.Fields(cmdName)
//~ cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
//~ stdout, err := cmd.StdoutPipe()
//~ if err != nil {
//~ return serialhd
//~ }
//~ cmd.Start()
//~ buf := bufio.NewReader(stdout)
//~ for {
//~ line, _, _ := buf.ReadLine()
//~ if len(line) < 1 {
//~ break
//~ }
//~ strlin := string(line)
//~ okidx := strings.Index(strlin, "ID_SERIAL")
//~ if okidx >= 0 {
//~ serialhd = strlin[okidx:len(strlin)]
//~ break
//~ }
//~ }
//~ //os.Exit(0)
//~ //cmd.Wait()
//~ return serialhd
//~ }

//~ func ValidarEmail(email string) bool {
//~ Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
//~ return Re.MatchString(email)
//~ }

//~ func ValidaCPF(cpf string) bool {
//~ cpf = strings.Replace(cpf, ".", "", -1)
//~ cpf = strings.Replace(cpf, "-", "", -1)
//~ if len(cpf) != 11 {
//~ return false
//~ }
//~ var eq bool
//~ var dig string
//~ for _, val := range cpf {
//~ if len(dig) == 0 {
//~ dig = string(val)
//~ }
//~ if string(val) == dig {
//~ eq = true
//~ continue
//~ }
//~ eq = false
//~ break
//~ }
//~ if eq {
//~ return false
//~ }

//~ i := 10
//~ sum := 0
//~ for index := 0; index < len(cpf)-2; index++ {
//~ pos, _ := strconv.Atoi(string(cpf[index]))
//~ sum += pos * i
//~ i--
//~ }

//~ prod := sum * 10
//~ mod := prod % 11
//~ if mod == 10 {
//~ mod = 0
//~ }
//~ digit1, _ := strconv.Atoi(string(cpf[9]))
//~ if mod != digit1 {
//~ return false
//~ }
//~ i = 11
//~ sum = 0
//~ for index := 0; index < len(cpf)-1; index++ {
//~ pos, _ := strconv.Atoi(string(cpf[index]))
//~ sum += pos * i
//~ i--
//~ }
//~ prod = sum * 10
//~ mod = prod % 11
//~ if mod == 10 {
//~ mod = 0
//~ }
//~ digit2, _ := strconv.Atoi(string(cpf[10]))
//~ if mod != digit2 {
//~ return false
//~ }

//~ return true
//~ }

//~ func ValidaCNPJ(cnpj string) bool {
//~ cnpj = strings.Replace(cnpj, ".", "", -1)
//~ cnpj = strings.Replace(cnpj, "-", "", -1)
//~ cnpj = strings.Replace(cnpj, "/", "", -1)
//~ if len(cnpj) != 14 {
//~ return false
//~ }

//~ algs := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
//~ var algProdCpfDig1 = make([]int, 12, 12)
//~ for key, val := range algs {
//~ intParsed, _ := strconv.Atoi(string(cnpj[key]))
//~ sumTmp := val * intParsed
//~ algProdCpfDig1[key] = sumTmp
//~ }
//~ sum := 0
//~ for _, val := range algProdCpfDig1 {
//~ sum += val
//~ }
//~ digit1 := sum % 11
//~ if digit1 < 2 {
//~ digit1 = 0
//~ } else {
//~ digit1 = 11 - digit1
//~ }
//~ char12, _ := strconv.Atoi(string(cnpj[12]))
//~ if char12 != digit1 {
//~ return false
//~ }
//~ algs = append([]int{6}, algs...)

//~ var algProdCpfDig2 = make([]int, 13, 13)
//~ for key, val := range algs {
//~ intParsed, _ := strconv.Atoi(string(cnpj[key]))

//~ sumTmp := val * intParsed
//~ algProdCpfDig2[key] = sumTmp
//~ }
//~ sum = 0
//~ for _, val := range algProdCpfDig2 {
//~ sum += val
//~ }

//~ digit2 := sum % 11
//~ if digit2 < 2 {
//~ digit2 = 0
//~ } else {
//~ digit2 = 11 - digit2
//~ }
//~ char13, _ := strconv.Atoi(string(cnpj[13]))
//~ if char13 != digit2 {
//~ return false
//~ }

//~ return true
//~ }

//~ func EnviaEmail(email_from string, email_to []string ,title ,msg string) bool{
//~ strmsg := "From: " + email_from + "\n" +
//~ //"To: " + email_to[0] + "\n" +
//~ "Subject: "+title+"\n\n" + msg
//~ err := smtp.SendMail(EMAIL_HOST+":587", smtp.PlainAuth("", EMAIL_USER, EMAIL_PWD, EMAIL_HOST),email_from, email_to, []byte(strmsg))

//~ if err != nil {
//~ return false
//~ }
//~ return true
//~ //example:
//~ //emailto := make([]string,0)
//~ //emailto = append(emailto,"email1@destino.com")// destino 1
//~ //emailto = append(emailto,"email2@destino.com") // destino 2
//~ //msg:= "Ola teste 123 \r\n Linha1 \r\n Linha2 "
//~ //ret := EnviaEmail("origem@email.com", emailto ,"Oi", msg)
//~ }

//~ func Unidade(unidade_num int) string {

//~ var unidade string
//~ switch unidade_num {
//~ case 0:
//~ unidade = ""
//~ case 1:
//~ unidade = "Um"
//~ case 2:
//~ unidade = "Dois"
//~ case 3:
//~ unidade = "Tres"
//~ case 4:
//~ unidade = "Quatro"
//~ case 5:
//~ unidade = "Cinco"
//~ case 6:
//~ unidade = "Seis"
//~ case 7:
//~ unidade = "Sete"
//~ case 8:
//~ unidade = "Oito"
//~ case 9:
//~ unidade = "Nove"
//~ }

//~ return unidade
//~ }

//~ func DezenaDec(dezenadec_num int) string {
//~ var dezenadec string
//~ switch dezenadec_num {
//~ case 11:
//~ dezenadec = "Onze"
//~ case 12:
//~ dezenadec = "Doze"
//~ case 13:
//~ dezenadec = "Treze"
//~ case 14:
//~ dezenadec = "Quatorze"
//~ case 15:
//~ dezenadec = "Quinze"
//~ case 16:
//~ dezenadec = "Dezesseis"
//~ case 17:
//~ dezenadec = "Dezessete"
//~ case 18:
//~ dezenadec = "Dezoito"
//~ case 19:
//~ dezenadec = "Dezenove"
//~ }

//~ return dezenadec
//~ }

//~ func Dezena(dezena_num int) string {
//~ var dezena string
//~ switch dezena_num {
//~ case 10:
//~ dezena = "Dez"
//~ case 20:
//~ dezena = "Vinte"
//~ case 30:
//~ dezena = "Trinta"
//~ case 40:
//~ dezena = "Quarenta"
//~ case 50:
//~ dezena = "Cinquenta"
//~ case 60:
//~ dezena = "Sessenta"
//~ case 70:
//~ dezena = "Setenta"
//~ case 80:
//~ dezena = "Oitenta"
//~ case 90:
//~ dezena = "Noventa"
//~ }

//~ return dezena
//~ }

//~ func Centena(centena_num int) string {
//~ var centena string
//~ switch centena_num {
//~ case 100:
//~ centena = "Cem"
//~ case 200:
//~ centena = "Duzentos"
//~ case 300:
//~ centena = "Trezentos"
//~ case 400:
//~ centena = "Quatrocentos"
//~ case 500:
//~ centena = "Quinhentos"
//~ case 600:
//~ centena = "Seiscentos"
//~ case 700:
//~ centena = "Setecentos"
//~ case 800:
//~ centena = "Oitocentos"
//~ case 900:
//~ centena = "Novecentos"
//~ }

//~ return centena
//~ }

//~ func Milhares(milhares_num int) string {
//~ var milhares string

//~ switch milhares_num {
//~ case 100:
//~ milhares = "Mil"
//~ case 200:
//~ milhares = "Milhão"
//~ case 300:
//~ milhares = "Bilhão"
//~ case 400:
//~ milhares = "Trilhão"
//~ }

//~ return milhares
//~ }

//~ func ValorDezena(valor int) string {

//~ var dextenso string = ""
//~ vl := Str(int64(valor))

//~ if valor > 10 && valor < 20 {
//~ dextenso += DezenaDec(valor)
//~ } else if valor == 10 || (valor > 19 && valor < 100) {

//~ vl_right_dez := Right(vl, 1)
//~ int_vl_right_dez := Val(vl_right_dez)

//~ if int(int_vl_right_dez) == 0 {
//~ dextenso += Dezena(valor)
//~ } else {
//~ dezena_left := Left(vl, 1)
//~ unidade_right := Right(vl, 1)
//~ vl_dez_int := int(Val(dezena_left + "0"))
//~ vl_uni_int := int(Val(unidade_right))
//~ dextenso += Dezena(vl_dez_int) + " e " + Unidade(vl_uni_int)
//~ }
//~ }

//~ return dextenso
//~ }

//~ func ValorCentena(valor int) string {

//~ var cextenso string
//~ vl := Str(int64(valor))
//~ vl_dez_int := Right(vl, 2)
//~ vl_dez := Val(vl_dez_int)

//~ if valor > 100 && valor < 200 {

//~ cextenso += "Cento e " + ValorDezena(int(vl_dez))
//~ } else if valor == 100 || (valor > 199 && valor < 1000) {

//~ if vl_dez == 0 {
//~ cextenso += Centena(valor)
//~ } else {
//~ vl_cen_left := Left(vl, 1)
//~ cextenso += Centena(int(Val(vl_cen_left+"00"))) + " e " + ValorDezena(int(vl_dez))
//~ }
//~ }
//~ return cextenso
//~ }

//~ func ExtensoReal(valor string) string {
//~ //var maximo float64 = 99999999.99

//~ v := strings.Index(valor, ".")
//~ if v < 1  {
//~ valor += ".0"
//~ }

//~ var extenso string
//~ vl := strings.Split(valor, ".")
//~ inteiros := vl[0]
//~ centavos := vl[1]
//~ var a, b string

//~ switch len(inteiros) {
//~ case 1:
//~ if Val(inteiros) == int64(1) {
//~ vl_inteiros := Val(inteiros)
//~ vl_unidade := Unidade(int(vl_inteiros))
//~ extenso += vl_unidade + " Real"
//~ } else if Val(inteiros) > int64(1) {
//~ extenso += Unidade(int(Val(inteiros))) + " Reais"
//~ } else {
//~ extenso += ""
//~ }

//~ case 2:
//~ extenso = ValorDezena(int(Val(inteiros))) + " Reais"
//~ case 3:
//~ if Val(Right(inteiros, 2)) == int64(0) {
//~ extenso += ValorCentena(int(Val(inteiros))) + ValorDezena(int(Val(Right(inteiros, 2)))) + " Reais"
//~ } else {
//~ extenso += ValorCentena(int(Val(inteiros))) + " Reais"
//~ }

//~ case 4:
//~ if Val(Right(inteiros, 2)) > 0 {
//~ a = " " + ValorCentena(int(Val(Right(inteiros, 3))))
//~ } else {
//~ a = ""
//~ }

//~ if Val(Left(inteiros, 1)) > int64(0) {
//~ b = Unidade(int(Val(Left(inteiros, 1)))) + " Mil"
//~ } else {
//~ b = ""
//~ }

//~ extenso = b + a + " Reais"

//~ case 5:
//~ if Val(Right(inteiros, 2)) > 0 {
//~ a = " " + ValorCentena(int(Val(Right(inteiros, 3))))
//~ } else {
//~ a = ""
//~ }

//~ if Val(Left(inteiros, 1)) > int64(0) {
//~ b = DezenaDec(int(Val(Left(inteiros, 2)))) + " Mil"
//~ } else {
//~ b = ""
//~ }

//~ extenso = b + a + " Reais"
//~ //Mil
//~ case 6:

//~ if Val(Right(inteiros, 2)) > 0 {
//~ a = " " + ValorCentena(int(Val(Right(inteiros, 3))))
//~ } else {
//~ a = ""
//~ }

//~ if Val(Left(inteiros, 1)) > int64(0) {
//~ b = ValorCentena(int(Val(Left(inteiros, 3)))) + " Mil"
//~ } else {
//~ b = ""
//~ }

//~ extenso = b + a + " Reais"

//~ //Milhão
//~ case 7:

//~ vl := strings.Split(inteiros, "")
//~ vl_mil := vl[1]
//~ vl_mil += vl[2]
//~ vl_mil += vl[3]
//~ vl_mil += vl[4]
//~ vl_mil += vl[5]
//~ vl_mil += vl[6]

//~ if Val(Right(inteiros, 2)) > 0 {
//~ a = " " + ValorCentena(int(Val(Right(vl_mil, 3))))
//~ } else {
//~ a = ""
//~ }

//~ if Val(Left(inteiros, 1)) > int64(0) {
//~ b = ValorCentena(int(Val(Left(vl_mil, 3)))) + " Mil"
//~ } else {
//~ b = ""
//~ }

//~ extenso = Unidade(int(Val(Left(inteiros, 1)))) + " Milhão " + b + a + " Reais"

//~ case 8:

//~ vl := strings.Split(inteiros, "")
//~ vl_mil := vl[2]
//~ vl_mil += vl[3]
//~ vl_mil += vl[4]
//~ vl_mil += vl[5]
//~ vl_mil += vl[6]
//~ vl_mil += vl[7]

//~ if Val(Right(inteiros, 2)) > 0 {
//~ a = " " + ValorCentena(int(Val(Right(vl_mil, 3))))
//~ } else {
//~ a = ""
//~ }

//~ if Val(Left(inteiros, 1)) > int64(0) {
//~ b = ValorCentena(int(Val(Left(vl_mil, 3)))) + " Mil"
//~ } else {
//~ b = ""
//~ }

//~ extenso = ValorDezena(int(Val(Left(inteiros, 2)))) + " Milhões " + b + a + " Reais"

//~ default:
//~ extenso = ""
//~ centavos = ""
//~ }

//~ if Val(inteiros) > int64(0) && Val(centavos) > int64(0) {
//~ extenso += " e "
//~ }

//~ switch int(Val(centavos)) {
//~ case 0:
//~ extenso += ""
//~ case 1:
//~ extenso += Unidade(int(Val(centavos))) + " Centavo"
//~ default:
//~ if Val(centavos) < int64(10) {
//~ extenso += Unidade(int(Val(centavos))) + " Centavos"
//~ } else if Val(centavos) > int64(10) && Val(centavos) < int64(20) {
//~ extenso += DezenaDec(int(Val(centavos))) + " Centavos"
//~ } else if (Val(centavos) == int64(10)) || (Val(centavos) > int64(19) && Val(centavos) < int64(100)) {
//~ if Val(Right(centavos, 1)) == int64(0) {
//~ extenso += Dezena(int(Val(Left(centavos, 1)+"0"))) + " Centavos"
//~ } else {
//~ extenso += Dezena(int(Val(Left(centavos, 1)+"0"))) + " e " + Unidade(int(Val(Right(centavos, 1)))) + " Centavos"
//~ }
//~ }
//~ }

//~ return extenso
//~ }
