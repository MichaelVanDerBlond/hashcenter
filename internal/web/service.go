package web

import (
	"github.com/MichaelVanDerBlond/hashcenter/internal/database"
	"github.com/MichaelVanDerBlond/hashcenter/internal/jobs"
	"github.com/MichaelVanDerBlond/hashcenter/internal/repository"
	"github.com/MichaelVanDerBlond/hashcenter/internal/task"
)

func LoadIndexData() (IndexData, error) {

	db, err := database.Open()
	if err != nil {
		return IndexData{}, err
	}
	defer db.Close()

	repo := repository.NewDictionaryRepository(db)

	uploads, err := LoadUploads()
	if err != nil {
		return IndexData{}, err
	}

	favorites, err := repo.Favorites("")
	if err != nil {
		return IndexData{}, err
	}

	others, err := repo.Others("")
	if err != nil {
		return IndexData{}, err
	}

	favoriteCount, err := repo.FavoriteCount()
	if err != nil {
		return IndexData{}, err
	}

	otherCount, err := repo.OtherCount()
	if err != nil {
		return IndexData{}, err
	}

	return IndexData{
		Title: "HashCenter",

		Jobs: jobs.Active(),

		Queue: task.DefaultManager().All(),

		Uploads: uploads,

		Favorites: favorites,
		Others:    others,

		FavoriteCount: favoriteCount,
		OtherCount:    otherCount,
	}, nil
}
