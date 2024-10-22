package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/test", handler)
	log.Fatal(http.ListenAndServe(":9876", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	doSomething(10000)
	buff := genSomeBytes()
	b, err := ioutil.ReadAll(buff)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}

func doSomething(times int) {
	for i := 0; i < times; i++ {
		for j := 0; j < times; j++ {
		}
	}
}

func genSomeBytes() *bytes.Buffer {
	var buff bytes.Buffer
	for i := 1; i < 20000; i ++ {
		buff.Write([]byte{'0' + byte(rand.Intn(10))})
	}
	return &buff
}
