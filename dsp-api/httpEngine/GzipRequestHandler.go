package httpEngine

import (
	"compress/gzip"
	"io"
	"net/http"
)

func readRequestBody(r *http.Request) ([]byte, error) {
	var reader io.Reader = r.Body

	if r.Header.Get("Content-Encoding") == "gzip" {
		gz, err := gzip.NewReader(r.Body)
		if err != nil {
			return nil, err
		}
		defer gz.Close()
		reader = gz
	}

	return io.ReadAll(reader)
}
