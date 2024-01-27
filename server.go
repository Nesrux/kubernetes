package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)

}
func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	w.Write([]byte("<h1>Hello Kubernets!!!</h1>"))
	fmt.Fprintf(w, "Olá eu sou %s minha idade é %s", name, age)
}
