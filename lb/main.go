package main

import (
	"fmt"
	"net/http"
)

const url = "http://localhost:8080"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
	fmt.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Proto)
	fmt.Printf("Host: %s\n", r.Host)
	fmt.Printf("User-Agent: %s\n", r.UserAgent())
	fmt.Printf("Accept: %s\n", r.Header.Get("Accept"))

	req, err := http.NewRequest(r.Method, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header = r.Header
	req.Host = r.Host
	req.RemoteAddr = r.RemoteAddr

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Response from server: ", resp.Proto, resp.Status)
	fmt.Println("Hello from Backend Server")
	fmt.Fprintf(w, "Hello from Backend Server")
}

func main() {
	http.HandleFunc("/", helloHandler)

	http.ListenAndServe(":80", nil)
}
