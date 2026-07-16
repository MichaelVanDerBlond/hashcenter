package scanner

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/MichaelVanDerBlond/hashcenter/internal/models"
)

const DictionaryRoot = "/hashcat/dict"

func ScanDictionaries() ([]models.Dictionary, error) {

	var list []models.Dictionary

	err := filepath.WalkDir(DictionaryRoot, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		switch strings.ToLower(filepath.Ext(path)) {

		case ".txt", ".lst", ".dic":

		default:
			return nil

		}

		info, err := os.Stat(path)

		if err != nil {
			return nil
		}

		list = append(list, models.Dictionary{
			Path: path,
			Name: filepath.Base(path),
			Size: info.Size(),
		})

		return nil

	})

	sort.Slice(list, func(i, j int) bool {
		return list[i].Size < list[j].Size
	})

	return list, err

}
