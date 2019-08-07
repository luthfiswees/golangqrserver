package main

import (
	"log"
	"net/http"

	"github.com/subosito/gotenv"

	"github.com/luthfiswees/golangqrserver/handler"
)

func main() {
	gotenv.Load()
	
	http.HandleFunc("/readqr", handler.ReadQR)

	log.Println("QRServer Now Serving in 15151")
	http.ListenAndServe(":15151", nil)
}