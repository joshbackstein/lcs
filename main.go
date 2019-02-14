package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const maxUploadSize = 2 * 1024 * 1024 // 2 mb
const uploadPath = "./uploads"
const outputPathRoot = "./pdf"
const inputExtension = ".tex"
const outputExtension = ".pdf"

func main() {
	http.HandleFunc("/upload", uploadFileHandler())

	fs := http.FileServer(http.Dir(outputPathRoot))
	http.Handle("/pdf/", http.StripPrefix("/pdf", fs))

	log.Print("Server started on localhost:8080, use /upload for uploading files and /pdf/{fileName} for downloading")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate file size
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}

		// parse and validate file and post parameters
		file, _, err := r.FormFile("uploadedFile")
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}
		defer file.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID_FILE", http.StatusBadRequest)
			return
		}

		// generate random name for file
		fileName := randToken(12)
		newPath := filepath.Join(uploadPath, fileName+inputExtension)

		// write file
		newFile, err := os.Create(newPath)
		if err != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}
		defer newFile.Close() // idempotent, okay to call twice
		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
			return
		}

		// convert file
		outputFilePath := filepath.Join(outputPathRoot, fileName+outputExtension)
		fmt.Printf("Uploaded file:  %s\n", newPath)
		fmt.Printf("Converted file: %s\n", outputFilePath)
		cmd := exec.Command("pandoc", newPath, "-o", outputFilePath)
		err = cmd.Run()
		if err != nil {
			renderError(w, "CANT_CONVERT_FILE", http.StatusInternalServerError)
			return
		}

		// return response
		w.Write([]byte("/" + outputFilePath))
	})
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("ERROR: " + message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
