package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./site/")))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
