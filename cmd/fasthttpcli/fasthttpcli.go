package main

import (
	"fasthttp/internal/httpserver"
	"os"
	"strconv"
)
import "fmt"

func main() {
	startVariables := os.Args[1:]
	if len(startVariables) == 0 {
		SendHelpMessage()
		return
	}
	if len(startVariables) == 2 {
		option := startVariables[0]
		port, err := strconv.Atoi(startVariables[1])

		if err != nil || port <= 0 || port > 65535 {
			fmt.Println("Please input a valid port 1-65535")
			return
		}
		server := httpserver.HttpServer{
			Port: port,
		}
		if option == "http" {
			httpserver.Run(server, "demo")
		}
		if option == "prod" {
			httpserver.Run(server, "prod")
		}
	}
}

func SendHelpMessage() {
	fmt.Println(
		"~~~~~~~~~~~~~~ FastHTTP ~~~~~~~~~~~~~~" +
			"\nUSAGE: fasthttp [OPTION] port" +
			"\n" +
			"\nOPTIONS:" +
			"\n		http (port) - Opens a demo http server on given port" +
			"\n		prod (port) - Opens a http server for current folder")

}
