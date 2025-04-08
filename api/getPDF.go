package main

import (
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Use the exact file path and name
	file, err := os.Open("files/sanjarbek_backend_cv.pdf")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=Sanjarbek_Abdurahmonov.pdf")

	io.Copy(w, file)
}
