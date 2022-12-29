package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
    r := mux.NewRouter()

    r.HandleFunc("/books/{title}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]

        fmt.Fprintf(w, "you have requested  the book %s\n", title)
    })
    http.ListenAndServe(":80",r)
}
