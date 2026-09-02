package main

import (
	"net/http"
)

func handlerReady(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	filepathRoot := "."
	mux := http.NewServeMux()
	svr := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.HandleFunc("/healthz", handlerReady)
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	assetsHandler := http.StripPrefix(
		"/assets/",
		http.FileServer(http.Dir("assets")),
	)
	mux.Handle("/assets", assetsHandler)
	svr.ListenAndServe()
}
