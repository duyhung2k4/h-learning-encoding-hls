package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type fileController struct{}

type FileController interface {
	GetAllFile(w http.ResponseWriter, r *http.Request)
	GetFile(w http.ResponseWriter, r *http.Request)
}

func (c *fileController) GetAllFile(w http.ResponseWriter, r *http.Request) {
	dir := chi.URLParam(r, "dir")
	videoDir := fmt.Sprintf("encoding/%s", dir)

	files, err := os.ReadDir(videoDir)
	if err != nil {
		http.Error(w, "Unable to read directory", http.StatusInternalServerError)
		return
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fileNames); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

func (c *fileController) GetFile(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")
	dir := chi.URLParam(r, "dir")
	imagePath := fmt.Sprintf("encoding/%s/%s", dir, filename)

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, imagePath)
}

func NewFileController() FileController {
	return &fileController{}
}
