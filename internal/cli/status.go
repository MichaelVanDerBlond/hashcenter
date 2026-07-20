package cli

import (
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/database"
	"github.com/MichaelVanDerBlond/hashcenter/internal/repository"
	"github.com/MichaelVanDerBlond/hashcenter/internal/status"
)

func Status() error {

	db, err := database.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	dictionaries := repository.NewDictionaryRepository(db)

	service := status.New(dictionaries)

	snapshot, err := service.Snapshot()
	if err != nil {
		return err
	}

	fmt.Println("HashCenter Status")
	fmt.Println("-----------------")
	fmt.Printf("Active sessions      : %d\n", snapshot.ActiveSessions)
	fmt.Printf("Favorite dictionaries: %d\n", snapshot.FavoriteDictionaries)
	fmt.Printf("Other dictionaries   : %d\n", snapshot.OtherDictionaries)

	return nil
}
