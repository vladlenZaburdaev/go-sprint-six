package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	htmlData, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "file read error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	w.Write(htmlData)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error receiving file:", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error reading file", http.StatusInternalServerError)
		return
	}

	receivedText := string(fileData)
	if receivedText == "" {
		http.Error(w, "empty file", http.StatusBadRequest)
		return
	}

	convertedText, err := service.DetectAndConvert(receivedText)
	if err != nil {
		http.Error(w, "conversion error", http.StatusInternalServerError)
		return
	}

	fileName := time.Now().UTC().String() + filepath.Ext(handler.Filename)
	err = writeToFile(fileName, convertedText)
	if err != nil {
		http.Error(w, "error saving file", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Конвертированный текст: %s", convertedText)
}

func writeToFile(fileName string, fileData string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fileData)
	if err != nil {
		return err
	}

	return nil
}
