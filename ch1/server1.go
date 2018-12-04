// Server1 是一个迷你回声服务器
package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainserver1() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
