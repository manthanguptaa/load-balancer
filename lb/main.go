package main

import (
	"fmt"
	"net/http"
	"time"
)

type server struct {
	url       string
	is_active bool
}

var servers []*server = []*server{
	{"http://localhost:8080", true},
	{"http://localhost:8081", true},
	{"http://localhost:8082", true},
}
var round_robin int = -1

const healthcheck_after = 5 * time.Second

func (s *server) disable() {
	s.is_active = false
}

func (s *server) activate() {
	s.is_active = true
}

func getServer() (*server, error) {
	for numOfRetries := len(servers); numOfRetries >= 0; numOfRetries-- {
		round_robin = (round_robin + 1) % len(servers)
		if servers[round_robin].is_active {
			return servers[round_robin], nil
		}
	}
	return nil, fmt.Errorf("no server is active right now")
}

func forwardRequest(r *http.Request) {
	server, err := getServer()
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(r.Method, server.url, nil)
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
		if server.is_active {
			server.disable()
		}
	}

	defer resp.Body.Close()

	fmt.Printf("Response from server: %s %s\n\n", resp.Proto, resp.Status)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
	fmt.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Proto)
	fmt.Printf("Host: %s\n", r.Host)
	fmt.Printf("User-Agent: %s\n", r.UserAgent())
	fmt.Printf("Accept: %s\n", r.Header.Get("Accept"))

	forwardRequest(r)
}

func periodicHealthcheck() {
	for {
		for index := 0; index < len(servers); index++ {
			req, err := http.NewRequest("GET", servers[index].url+"/healthcheck", nil)
			if err != nil {
				fmt.Println("Error creating request:", err)
				continue
			}
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				if servers[index].is_active {
					servers[index].disable()
				}
			} else {
				if resp.StatusCode < 200 || resp.StatusCode >= 300 {
					if servers[index].is_active {
						servers[index].disable()
					}
				} else {
					if !servers[index].is_active {
						servers[index].activate()
					}
				}
				resp.Body.Close()
			}
		}
		time.Sleep(healthcheck_after)
	}

}

func main() {
	fmt.Println("booting load balancer...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go requestHandler(w, r)
	})

	go periodicHealthcheck()

	fmt.Println("load balancer is ready to serve...")

	http.ListenAndServe(":80", nil)
}
