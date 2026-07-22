package handlers

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const UploadDir = "/hashcat/hashes/incoming"

func UploadHash(c *gin.Context) {

	if err := os.MkdirAll(UploadDir, 0755); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.String(http.StatusBadRequest, "no files uploaded")
		return
	}

	for _, file := range files {

		dst := filepath.Join(UploadDir, filepath.Base(file.Filename))

		if _, err := os.Stat(dst); err == nil {
			c.String(http.StatusConflict, "file already exists: %s", file.Filename)
			return
		} else if !errors.Is(err, os.ErrNotExist) {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}

	c.Redirect(http.StatusSeeOther, "/")
}
