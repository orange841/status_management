package main

import (
  "fmt"
  "log"
  "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "ホームページ来たよ")
  fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
  http.HandleFunc("/", homePage)
  log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
    handleRequests()
}
