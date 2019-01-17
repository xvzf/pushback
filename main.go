package main

import (
	"log"
	"net/http"
	"time"

	"github.com/xvzf/pushback/pushback"
)

// Pushback Entrypoint
func main() {
	c := &pushback.HandlerConfig{
		Path: "/tmp/",
	}

	// @TODO
	http.HandleFunc("/pushback/test", pushback.NewHandler(c))

	// Test if the server is operating properly
	go func() {
		time.Sleep(time.Second * 2)
		r, err := pushback.PushFile("http://localhost:1337/pushback/test", "./main.go")
		if err != nil {
			panic(err)
		}
		log.Print(r)
	}()

	err := http.ListenAndServe("[::]:1337", nil)
	if err != nil {
		log.Fatal(err)
	}
}
