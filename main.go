package main

import (
	"image/png"
	"log"
	"net/http"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Page struct {
	Title string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "QR code generator3"}

	t, _ := template.ParseFiles("generate.html")
	t.Execute(w, p)
}

func CodePage(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("dataString")
	qrCode, _ := qr.Encode(dataString, qr.L, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, 512, 512)
	png.Encode(w, qrCode)
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/generate/", CodePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
