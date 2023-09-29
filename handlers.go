package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func uploadVideo(c *gin.Context) {
	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse request",
			"error":   err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	// TODO: check if the mimetype of the file is a valid video one

	// generate a unique name
	videoName := generateUniqueName(file.Filename)

	// path to store video
	videoPath := filepath.Join(temp_dir, videoName)

	if c.SaveUploadedFile(file, videoPath) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to upload file",
			"error":   err,
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"url":     fmt.Sprintf("%s/api/v1/videos/%s", c.Request.Host, videoName),
		"status":  http.StatusOK,
	})
}
