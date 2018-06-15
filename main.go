package myserver

import (
	"github.com/julienschmidt/httprouter"
	"strconv"
	"fmt"
	"net/http"
	"net/http/httputil"
	"log"
)

func print(s interface{}) {
	fmt.Println(s)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("The landing paj"))
}

func mirror(w http.ResponseWriter, r *http.Request,  _ httprouter.Params) {
	msg, err := httputil.DumpRequest(r, true)
	if err != nil {
		print(err)
		w.Write([]byte(err.Error()))
	}

	print(string(msg))
	w.Write([]byte(string(msg)))
}

func squareHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	x, err := strconv.Atoi(ps.ByName("x"))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	x = Square(x)
	w.Write([]byte(strconv.Itoa(x)))
}

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("pong"))
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/ping", ping)
	router.GET("/mirror", mirror)
	router.GET("/squared/:x", squareHandler)
	router.ServeFiles("/explore/*filepath", http.Dir("./cool_files"))
	
	log.Fatal(http.ListenAndServe(":80", router)) 
}

