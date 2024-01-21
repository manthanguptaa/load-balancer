package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
	fmt.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Proto)
	fmt.Printf("Host: %s\n", r.Host)
	fmt.Printf("User-Agent: %s\n", r.UserAgent())
	fmt.Printf("Accept: %s\n\n", r.Header.Get("Accept"))
	fmt.Println("Replied with a hello message")
}

func main() {
	port := ":8080"
	http.HandleFunc("/", helloHandler)

	for {
		err := http.ListenAndServe(port, nil)

		if err != nil {
			port = fmt.Sprintf(":%d", nextAvailablePort(port))
			fmt.Println("New port: ", port)
		} else {
			break
		}
	}
}

func nextAvailablePort(port string) int {
	portNumber, err := strconv.Atoi(strings.TrimPrefix(port, ":"))
	if err != nil {
		panic("invalid port number!")
	}

	portNumber++

	return portNumber
}
