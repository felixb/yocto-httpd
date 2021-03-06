package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const DEFAULT_PORT = 8080

func logRequest(req *http.Request, statusCode int) {
	log.Printf("%s %s %s %d", req.RemoteAddr, req.Method, req.URL.Path, statusCode)
}

func serveRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		logRequest(req, http.StatusMethodNotAllowed)
		w.Header().Add("Allow", "GET")
		http.Error(w, "Only GET supported.", http.StatusMethodNotAllowed)
		return
	}

	logRequest(req, http.StatusOK)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "OK")
}

func listen(port uint) error {
	log.Printf("starting httpd on port %d", port)
	addr := fmt.Sprintf(":%d", port)
	http.Handle("/", http.HandlerFunc(serveRequest))
	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}
	log.Printf("stopped httpd")
	return nil
}

func resolveDefaultPort() uint {
	if portEnvVar, ok := os.LookupEnv("PORT"); !ok {
		return DEFAULT_PORT
	} else {
		if port, err := strconv.ParseUint(portEnvVar, 10, 32); err != nil {
			log.Fatalf("error parsing port from environent variable: PORT=%s, %s", portEnvVar, err)
			return 0 // dead code!!1 omfg zombiess
		} else {
			return uint(port)
		}
	}
}

func main() {
	var port = flag.Uint("port", resolveDefaultPort(), "Port to listen for connections, also specified via enviroment vairlable PORT")
	flag.Parse()

	err := listen(*port)
	if err != nil {
		log.Fatalf("error starting httpd: %s", err)
	}
}
