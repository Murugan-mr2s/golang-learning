package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"pixels.com/firstapi/handlers"
)

/*
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func goodByeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("goodbye request received")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Hello ", string(data))
}

func main() {

	fmt.Println("restful api basics")
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/goodbye", goodByeHandler)
	fmt.Println("Server is listening on 8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
*/

func main() {

	l := log.New(os.Stdout, "api-logs", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	ph := handlers.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	sm.Handle("/products", ph)
	//http.ListenAndServe(":8080", sm)

	s := &http.Server{Addr: ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	//s.ListenAndServe()

	//graceful server shutdown

	go func() {

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println(" server graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
