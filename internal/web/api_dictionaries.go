package web

import (
	"net/http"

	"github.com/MichaelVanDerBlond/hashcenter/internal/database"
	"github.com/MichaelVanDerBlond/hashcenter/internal/repository"
	"github.com/gin-gonic/gin"
)

type dictionariesResponse struct {
	Favorites any `json:"favorites"`
	Others    any `json:"others"`
}

func apiDictionaries(c *gin.Context) {

	db, err := database.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer db.Close()

	repo := repository.NewDictionaryRepository(db)

	favorites, err := repo.Favorites("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	others, err := repo.Others("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dictionariesResponse{
		Favorites: favorites,
		Others:    others,
	})
}
