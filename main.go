package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server start and listen at port 443")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello world")
	})

	log.Fatalln(http.ListenAndServeTLS(":443", "ssl/logbook.pem", "ssl/logbook.key", nil))
}
