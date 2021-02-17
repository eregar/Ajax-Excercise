package main

import (
	"encoding/json"
	"net/http"
)

//Lista elements
type Lista struct {
	Name     string //`json:"name"`
	Elements []int  //`json:"elements"`
}

func darMensaje(w http.ResponseWriter, r *http.Request) {
	arreglo := Lista{"results", []int{25, 44, 33}}
	//arreglojson, _ := json.Marshal(arreglo)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	//w.Write(arreglojson)

	json.NewEncoder(w).Encode(arreglo)
	//fmt.Printf("%s", r)
	//fmt.Printf(string(value))

}

func main() {

	http.HandleFunc("/", darMensaje)
	err := http.ListenAndServe("localhost"+":"+"8080", nil)
	if err != nil {
		return
	}
}

/*
import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/eljson", ServeNoJson)
	http.HandleFunc("/secondfile", ServeSecondFile)
	http.HandleFunc("/testjson", ServeTestJson)
	http.HandleFunc("/", MainHtml)
	log.Printf("Starting server on localhost:8080...\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MainHtml(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func ServeNoJson(w http.ResponseWriter, r *http.Request) {
	log.Printf("request nojson file: %s", r.Host)
	http.ServeFile(w, r, "nojson.txt")
}

func ServeSecondFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("request secondfile: %s", r.Host)
	http.ServeFile(w, r, "secondfile.txt")
}

func ServeTestJson(w http.ResponseWriter, r *http.Request) {
	log.Printf("request testjson: %s", r.Host)
	http.ServeFile(w, r, "testjson.json")
}

*/
