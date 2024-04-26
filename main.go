package main

import (
	"fmt"
	"go_final/database"
	"log"
	"net/http"
)

func main() {
	database.EnsureDBExists()

	webDir := "../web/"
	http.Handle("/", http.FileServer(http.Dir(webDir)))
	fmt.Println("Runs")
	log.Fatal(http.ListenAndServe(":7540", nil))
}
