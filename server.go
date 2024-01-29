package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/configMap", configMap)

	http.ListenAndServe(":8000", nil)

}
func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	w.Write([]byte("<h1>Hello Kubernets!!!</h1>"))
	fmt.Fprintf(w, "Olá eu sou %s minha idade é %s", name, age)
}
func configMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./myfamily/family.txt")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: ", err)
	}
	fmt.Fprintf(w, "minha familia %s.", string(data))
}
