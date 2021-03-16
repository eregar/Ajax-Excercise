package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type cancion struct {
	Titulo   string `json:"titulo"`
	Autor    string `json:"autor"`
	Link     string `json:"editorial"`
	Cantidad int    `json:"cantidad"`
	Oferta   bool   `json:"oferta"`
}

//Arreglo lista de elementos
type Arreglo struct {
	Lista []cancion `json:"data"`
}

func mostrarHTML(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "vue01.html")
}

func darMensaje(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	x := r.Form.Get("y")

	fmt.Printf("%s", x)

	chanson := Arreglo{
		[]cancion{
			{
				Titulo:   "Ocean man",
				Autor:    "Ween",
				Link:     "https://www.youtube.com/watch?v=6E5m_XtCX3c",
				Cantidad: 12,
				Oferta:   false,
			},
			{
				Titulo:   "City of Angels",
				Autor:    "24kGoldn",
				Link:     "https://www.youtube.com/watch?v=OOmu1Yki-7g",
				Cantidad: 333,
				Oferta:   true,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(chanson)

}

func main() {
	//lista = Arreglo{[]string{"primero", "segundo", "tercero", "cuarto", "quinto", "nuevo"}}

	http.HandleFunc("/dato", darMensaje)
	http.HandleFunc("/", mostrarHTML)

	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}
