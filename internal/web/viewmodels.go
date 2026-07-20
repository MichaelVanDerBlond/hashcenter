package web

import (
	"github.com/MichaelVanDerBlond/hashcenter/internal/jobs"
	"github.com/MichaelVanDerBlond/hashcenter/internal/models"
)

type IndexData struct {
	Title string

	Jobs []*jobs.Runtime

	Uploads []models.Upload

	Favorites []models.Dictionary
	Others    []models.Dictionary

	FavoriteCount int
	OtherCount    int
}
