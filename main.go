package main

import (
	//	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

func print(s interface{}) {
	fmt.Println(s)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The landing paj"))
}

func mirror(w http.ResponseWriter, r *http.Request) {
	msg, err := httputil.DumpRequest(r, true)
	if err != nil {
		print(err)
		w.Write([]byte(err.Error()))
	}

	print(string(msg))
	w.Write([]byte(string(msg)))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/mirror", mirror)
	http.Handle("/explore/",
		http.StripPrefix("/explore/", http.FileServer(http.Dir("./cool_files"))))

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

/*
func main() {
	http.Handle("/explore", http.FileServer(http.Dir("./src")))
	//http.HandleFunc("/", index)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

*/
