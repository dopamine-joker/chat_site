package site

import (
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Site struct {
}

func New() *Site {
	return &Site{}
}

func (s *Site) Run() {
	addr := strconv.Itoa(8080)
	logrus.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir("./site/"))))
}
