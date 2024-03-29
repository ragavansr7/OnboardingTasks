package main

import (
	"fmt"
	"net/http"
	"time"
)

type Store interface {
	Fetch() string
	Cancel()
}

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

//	func Server(store Store) http.HandlerFunc {
//		return func(w http.ResponseWriter, r *http.Request) {
//			fmt.Fprint(w, store.Fetch())
//		}
//	}
//
//	func Server(store Store) http.HandlerFunc {
//		return func(w http.ResponseWriter, r *http.Request) {
//			store.Cancel()
//			fmt.Fprint(w, store.Fetch())
//		}
//	}
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

func main() {
	fmt.Println("")
	var asds Store
	asd := Server(asds)
	fmt.Println("asd---", asd)
}
