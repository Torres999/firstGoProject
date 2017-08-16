package controller

import (
	"log"
	"net/http"
	"time"
)

type Controller struct {
}

func (t *Controller) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("【controller】process a request,snapshort:", time.Now())
	w.Write([]byte("【controller】process a request,snapshort:" + time.Now()))
}
