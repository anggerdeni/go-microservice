package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Echo struct {
	log *log.Logger
}

func NewEcho(l *log.Logger) *Echo {
	return &Echo{l}
}

func (e *Echo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	e.log.Println("Echo is called")

	requestBody, errReadingRequestBody := ioutil.ReadAll(r.Body)

	if errReadingRequestBody != nil {
		http.Error(rw, "Error reading request body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "%s", requestBody)
}
