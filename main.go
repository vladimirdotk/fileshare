package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func validateFlags(ip *string, port *int, file *string) {
	if net.ParseIP(*ip) == nil {
		log.Fatal("Failed to get valid IP!")
	}

	if *port < 1 || *port > 65536 {
		log.Fatal("Failed to validate port!")
	}

	if _, err := os.Stat(*file); os.IsNotExist(err) {
		log.Fatal("File doesn't exists")
	}
}

func main() {

	ip := flag.String("ip", "127.0.0.1", "Server ip")
	port := flag.Int("port", 8088, "Server port")
	file := flag.String("f", "notValid", "File to share")

	flag.Parse()

	validateFlags(ip, port, file)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *file)
	})

	ipPort := fmt.Sprintf("%s:%s", *ip, strconv.Itoa(*port))

	log.Println(fmt.Sprintf("Serving file %s on http://%s", *file, ipPort))
	log.Fatal(http.ListenAndServe(ipPort, nil))
}
