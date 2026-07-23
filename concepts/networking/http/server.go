package http

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func HttpServer() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp,"Hello Server")
	})
	const serverAddr string = "127.0.0.1:8000"
	cert := "cert.pen"
	key :=  "key.pen"
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}
	server := &http.Server{
		Addr:      serverAddr,
		TLSConfig: tlsConfig,
		Handler: nil,
		
	}
	log.Printf("Starting server at %s\n", serverAddr)
	http2.ConfigureServer(server, &http2.Server{
		MaxConcurrentStreams: 250,
	})
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server",err)
	}
}