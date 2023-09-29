package main

import (
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// generate unique name
func generateUniqueName(filename string) string {
	extension := filepath.Ext(filename)
	uniqueID := uuid.New().String()
	return uniqueID + extension
}

// get the mime type
func getMimeType(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	buffer := make([]byte, 512) // Read the first 512 bytes to detect mime type
	_, err = io.ReadFull(src, buffer)
	if err != nil {
		return "", err
	}

	mime := http.DetectContentType(buffer)
	return mime, nil
}

// checks if a mime type starts with "video"
func isVideoMimeType(mime string) bool {
	return strings.HasPrefix(mime, "video/")
}
