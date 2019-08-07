package handler

import (
	"fmt"
	"encoding/json"
	"net/http"
	"image"
	_ "image/png"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

type Response struct {
	Message string
	Content string
}

func ReadQR(w http.ResponseWriter, r *http.Request) {
	// Set response header
	w.Header().Set("Content-Type", "application/json")
	var respBody []byte

	// Check if request is POST
	if r.Method != "POST" {
		respBody, _ = json.Marshal(&Response{Message: "Only POST is Supported"})
		w.Write(respBody)
		return
	}

	// Parse form from request
	r.ParseMultipartForm(0)

	// Get File
	fileUploaded, _, err := r.FormFile("file")
	if err != nil {
		respBody, _ = json.Marshal(&Response{Message: "Fail to upload file"})
		w.Write(respBody)
		return
	}

	// Decode to Image
	img, _, err := image.Decode(fileUploaded)
	if err != nil {
		respBody, _ = json.Marshal(&Response{Message: "Error to decode file to image"})
		fmt.Println(err)
		w.Write(respBody)
		return
	}

	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		respBody, _ = json.Marshal(&Response{Message: "Error to create binary bitmap"})
		w.Write(respBody)
		return
	}

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		respBody, _ = json.Marshal(&Response{Message: "Fail to decode QR"})
		w.Write(respBody)
		return
	}

	// Create Return body
	respBody, err = json.Marshal(&Response{Content: result.GetText()})
	if err != nil {
		respBody, _ = json.Marshal(&Response{Message: "Fail to build response body"})
		w.Write(respBody)
		return
	}
	w.Write(respBody)
	return
}