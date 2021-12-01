package main

import (
	"fmt"
	"log"
	"net/http"

	handlers "github.com/arahkya/logbook/net/handlers"
)

func main() {
	fmt.Println("Server start and listen at port 443")

	http.HandleFunc("/", handlers.IndexHandler)

	log.Fatalln(http.ListenAndServeTLS(":443", "ssl/logbook.pem", "ssl/logbook.key", nil))
}
