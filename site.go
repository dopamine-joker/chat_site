package main

import (
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

func main() {
	addr := strconv.Itoa(8080)
	logrus.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir("./site/"))))
}
