package main

import (
	"fmt"
	"log"
	"net/http"
)

func Temperature(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hola marica")
}
func main() {
	http.HandleFunc("/", Temperature)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
