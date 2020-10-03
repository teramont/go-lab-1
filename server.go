package main

import (
	"net/http"
)

type Time struct {
	Now string
}

func main() {
	http.HandleFunc("/time", mainHandler)
	http.ListenAndServe(":8795", nil)
}

//end1
