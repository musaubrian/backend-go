package main

import (
	"fmt"
	"log"
	"net/http"

//	"github.com/gorilla/mux"
)
 
func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Yay, you found the hompage")
}
func handleRequests() {
    http.HandleFunc("/", homePage)
    http.ListenAndServe(":8000", nil)
    log.Fatal(http.ListenAndServe(":8000", nil))

}
 

func main() {
   handleRequests()
}
