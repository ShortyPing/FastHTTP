package httpserver

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Error struct {
	message string
}

type HttpServer struct {
	Port int
}

func throw(error *Error) string {
	fmt.Println("[FASTHTTPCLI/HTTPSERVER] an unexpected error occured (" + error.message + ")")
	os.Exit(2)
	return error.message
}

func Run(server HttpServer, scope string) {
	fmt.Println("Starting HTTP-Server on", server.Port, "scope ", scope)

	if scope == "demo" {
		startDemo()
	}
	if scope == "prod" {
		startProduction()
	}

	str := strconv.Itoa(server.Port)
	err := http.ListenAndServe(":"+str, nil)

	if err != nil {
		throw(&Error{message: err.Error()})
	}
}

func startProduction() {
	dir, _ := os.Getwd()
	fmt.Println("Handling HTTP-Server for static files in directory", dir)
	http.Handle("/", http.FileServer(http.Dir(dir)))
}

func startDemo() {
	http.HandleFunc("/", handleDemo)
}

func handleDemo(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1></h1> HTTP Demo Server <br> Please use <code>fasthttp prod</code> to open a production server")
}
