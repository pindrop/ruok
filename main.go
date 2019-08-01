package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Print("Starting server")
	srv := &http.Server{
		Addr:         ":8081",
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  120 * time.Second,
		Handler:      http.HandlerFunc(EchoHandler),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("Stopped serving:", err.Error())
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Shutdown error occurred:", err.Error())
	}
	log.Print("Shutting down")
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v%v",
		r.RemoteAddr,
		r.Method,
		r.URL.Host,
		r.URL.Path)
	w.Header().Set("Server", "ruok")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if _, err := w.Write([]byte("{\"OK\": true}")); err != nil {
		log.Print("Failed to write response:", err.Error())
	}
}
