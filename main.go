package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

var VERSION = "unknown"

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	addrs, _ := net.InterfaceAddrs()
	var ipAddr string
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipAddr = ipnet.IP.String()
			}
		}
	}

	fmt.Fprintf(w, "<h1>Aplikacja Webowa</h1>")
	fmt.Fprintf(w, "<p><b>Wersja:</b> %s</p>", VERSION)
	fmt.Fprintf(w, "<p><b>Hostname:</b> %s</p>", hostname)
	fmt.Fprintf(w, "<p><b>Adres IP:</b> %s</p>", ipAddr)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Serwer uruchomiony na porcie 8080...")
	http.ListenAndServe(":8080", nil)
}