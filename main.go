package main

import (
	"io"
	"net/http"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("127.0.0.1:12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}