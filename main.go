package main

import (
	"github.com/gopherjs/gopherjs/js"
)
type Widget struct {
	Doc *js.Object
	Elem string
	Id string
}

func MakeWidget(Doc *js.Object,Elem , Id string)(w *Widget){
	w.Doc = Doc
	w.Elem = Elem
	w.Id = Id
	return w
}

func main() {
}
