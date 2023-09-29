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

	// get the mime type
	mime, err := getMimeType(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse request",
			"error":   err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	// check if mime type is valid
	// returns false if it isnt valid
	if !isVideoMimeType(mime) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to parse request",
			"error":   "tut-tut, thats not a video file. upload a video padawan",
			"status":  http.StatusBadRequest,
		})
		return
	}

	// generate a unique name
	videoName := generateUniqueName(file.Filename)

	// path to store video
	videoPath := filepath.Join(temp_dir, videoName)

	if err := c.SaveUploadedFile(file, videoPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to upload file",
			"error":   err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    gin.H{"url": fmt.Sprintf("%s/api/v1/videos/%s", c.Request.Host, videoName)},
		"status":  http.StatusOK,
	})
}
