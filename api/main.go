package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	ip, err := getIP()
	if err != nil {
		fmt.Println("Error getting IP address:", err)
		return
	}

	fmt.Println("Server running at IP:", ip)

	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/ping", PingPongServer) // Ping-pong endpoint

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	err = http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", os.Getenv("NAME"))
}

func PingPongServer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "pong")
	}
}

func getIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			return ipNet.IP.String(), nil
		}
	}

	return "", fmt.Errorf("no valid IP address found")
}
