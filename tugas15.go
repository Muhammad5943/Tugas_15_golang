package main

import "fmt"
import "net/http"

const PORT = ":8080"

func index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintln(w, "Hasil Perulangan 1-50 :")
  for i := 1 ; i <= 50; i++ {
    fmt.Fprint(w, i, " ")

  }
}

func main(){
  http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Tugas Ke 15 Golang \n\n","Silahkan ke tulis /index pada URL untuk melihat hasil perulangannya")

  })
  http.HandleFunc("/index", index)
  fmt.Println("Starting Web Server pada localhost:8080")
  http.ListenAndServe(PORT,nil)
}
