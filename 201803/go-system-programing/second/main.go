package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encodeing", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}

	gzipWriter := gzip.NewWriter(w)
	gzipWriter.Header.Name = "test.json"
	multiWriter := io.MultiWriter(gzipWriter, os.Stdout)

	jsonEncoder := json.NewEncoder(multiWriter)
	jsonEncoder.SetIndent("", "\t")
	jsonEncoder.Encode(source)
	gzipWriter.Flush()
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
