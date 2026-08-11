package web

import (
	"github.com/MichaelVanDerBlond/hashcenter/internal/jobs"
	"github.com/MichaelVanDerBlond/hashcenter/internal/models"
	"github.com/MichaelVanDerBlond/hashcenter/internal/task"
)

type IndexData struct {
	Title string

	Jobs []*jobs.Runtime

	Queue []*task.Task

	Uploads []models.Upload

	Favorites []models.Dictionary
	Others    []models.Dictionary

	FavoriteCount int
	OtherCount    int
}
