package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/yiaw/AdmissionController/cmd/app"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello handler call\n"))
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/validate", app.ValidatingWebHook)
	http.HandleFunc("/mutate", app.MutatingWebHook)

	log.Println("Start HTTPS Server")

	go log.Fatal(http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil))

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("shutdown signal, shutting down webhook server")
}
