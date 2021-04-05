package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
}

func (h *Hello) ServerHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello world")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	}
	log.Printf("Data %s\n", d)
	fmt.Fprintf(rw, "Hello %s", d)
}