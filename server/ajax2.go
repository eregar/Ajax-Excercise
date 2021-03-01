package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type cancion struct {
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Link   string `json:"link"`
}

//Arreglo lista de elementos
type Arreglo struct {
	Lista []string `json:"lista"`
}

var lista Arreglo

func mostrarHTML(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "ajax02.html")
}

func actualizarArreglo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		arrlenstr := r.Form.Get("len")
		arrlen, err := strconv.Atoi(arrlenstr)
		if err == nil {
			fmt.Println(len(lista.Lista) - arrlen)
			if arrlen < len(lista.Lista) {
				newData := Arreglo{lista.Lista[arrlen:]}
				json.NewEncoder(w).Encode(newData)
			} else {
				fmt.Fprintf(w, "")
			}
		} else {
			http.Error(w, "algo salio mal", http.StatusNotAcceptable)
		}
	}
}

func datosArreglo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(lista)
	} else {
		r.ParseForm()
		yrev := r.Form.Get("y")
		nrev := r.Form.Get("n")
		y, err := strconv.Atoi(yrev)
		n, err2 := strconv.Atoi(nrev)
		if err == nil && err2 == nil {
			temp := lista.Lista[y]
			lista.Lista[y] = lista.Lista[n]
			lista.Lista[n] = temp
			fmt.Fprintf(w, "OK")
		} else {
			fmt.Fprintf(w, "NOT OK")
		}
	}
}

func darMensaje(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	x := r.Form.Get("y")

	fmt.Printf("%s", x)

	chanson := cancion{
		Titulo: "Ocean man",
		Autor:  "Ween",
		Link:   "https://www.youtube.com/watch?v=6E5m_XtCX3c",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(chanson)

}

func main() {
	lista = Arreglo{[]string{"primero", "segundo", "tercero", "cuarto", "quinto", "nuevo"}}

	http.HandleFunc("/dato", darMensaje)
	http.HandleFunc("/array", datosArreglo)
	http.HandleFunc("/refresh", actualizarArreglo)
	http.HandleFunc("/", mostrarHTML)

	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}
