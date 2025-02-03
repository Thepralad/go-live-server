package main

import (
	"fmt"
	"net/http"
	"flag"
)

func main(){
	port := flag.String("port", "8080", "to add a custom port")
	path := flag.String("path", ".", "specifing path")
	flag.Parse()
	file := http.FileServer(http.Dir(*path))
	http.Handle("/", file)
	fmt.Println("Starting GO Live Server on http://localhost:" + *port);
	http.ListenAndServe(":"+*port, nil)
}
