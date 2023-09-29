package main

import (
	"path/filepath"

	"github.com/google/uuid"
)

// generate unique name
func generateUniqueName(filename string) string {
	extension := filepath.Ext(filename)
	uniqueID := uuid.New().String()
	return uniqueID + extension
}
