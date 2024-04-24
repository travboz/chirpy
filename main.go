package main

import (
	"log"
	"net"
	"net/http"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		writer.WriteHeader(404)
	})
	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	log.Printf("Serving on port: %s\n", port)
	srv.Serve(ln)

}

// Required to use CORS middleware for testing and compatibility
// function provided by boot.dev
func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
