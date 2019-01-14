package main

import (
	"github.com/gopherjs/gopherjs/js"
)
type Widget struct {
	Doc *js.Object
	Elem string
	Id string
}

func MakeWidget(Doc, Elem, Id string)(w *Widget){
	w.Doc = js.Global.Get(Doc)
	w.Elem = Elem
	w.Id = Id
	return w
}

func (w *Widget)Twttr() {
	var t *js.Object
	window := js.Global.Get("window")
	fjs := w.Doc.Call("getElementByTagName", w.Elem).Index(0)
	if window.Get("twttr") != nil{
		t =  window.Get("twttr")
	}
	js := w.Doc.Call("getElementById",w.Id)
	js.Set("src","https://platform.twitter.com/widgets.js")
	
}

func main() {
}
