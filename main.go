package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["item"]

		if !ok || len(keys[0]) < 0 {
			rw.Write([]byte("Aruvadai Arambam"))
			return
		}
		rw.Write([]byte(keys[0]))
	})

	fmt.Println("server started...")
	http.ListenAndServe(":8080", nil)
	fmt.Println("server ending")
}
