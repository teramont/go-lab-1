package main

import (
  "encoding/json"
  "net/http"
  "time"
)

type Time struct {
  Now string
}

func main() {
  http.HandleFunc("/time", mainHandler)
  http.ListenAndServe(":8795", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {

  now := time.Now()
  timeFormated := now.Format(time.RFC3339)

  timeObj := &Time{timeFormated}

  timeJson, _ := json.Marshal(timeObj)

  w.Header().Set("Content-Type", "application/json")
  w.Write(timeJson)
}
//test