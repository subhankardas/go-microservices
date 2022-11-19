package middleware

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GZipHandler struct{}

func NewGZipHandler() *GZipHandler {
	return &GZipHandler{}
}

func (gziph *GZipHandler) GZipperMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// If accept encoding header has gzip, use gzip response writer
		if strings.Contains(request.Header.Get("Accept-Encoding"), "gzip") {
			gzipResponse := NewGZipWriter(response)
			gzipResponse.Header().Set("Content-Encoding", "gzip")

			// Serve request with gzip response i.e. an implementation of ResponseWriter interface
			next.ServeHTTP(gzipResponse, request)
			defer gzipResponse.Flush()

			return
		}

		// Else serve request as usual
		next.ServeHTTP(response, request)
	})
}

type GZipWriter struct {
	response   http.ResponseWriter // Actual response writer
	gzipWriter *gzip.Writer        // Our custom gzip writer
}

func NewGZipWriter(response http.ResponseWriter) *GZipWriter {
	gzipResponse := gzip.NewWriter(response)
	return &GZipWriter{response: response, gzipWriter: gzipResponse}
}

// Implement methods defined by ResponseWriter interface. //

// Returns map of headers from our actual response writer.
func (gzipW *GZipWriter) Header() http.Header {
	return gzipW.response.Header()
}

// Write to our response using our gzip writer.
func (gzipW *GZipWriter) Write(data []byte) (int, error) {
	return gzipW.gzipWriter.Write(data)
}

// Write header to our actual response writer.
func (gzipW *GZipWriter) WriteHeader(statusCode int) {
	gzipW.response.WriteHeader(statusCode)
}

// Method to close our gzip writer stream.
func (gzipW *GZipWriter) Flush() {
	gzipW.gzipWriter.Flush()
	gzipW.gzipWriter.Close()
}
