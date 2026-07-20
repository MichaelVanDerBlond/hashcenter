package web

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/MichaelVanDerBlond/hashcenter/internal/models"
)

const uploadDirectory = "/hashcat/hashes/incoming"

func LoadUploads() ([]models.Upload, error) {

	if err := os.MkdirAll(uploadDirectory, 0755); err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(uploadDirectory)
	if err != nil {
		return nil, err
	}

	var uploads []models.Upload

	for _, entry := range entries {

		if entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			continue
		}

		uploads = append(uploads, models.Upload{
			Name:   info.Name(),
			Size:   info.Size(),
			Type:   strings.ToLower(filepath.Ext(info.Name())),
			Status: "Не запускался",
		})
	}

	return uploads, nil
}
