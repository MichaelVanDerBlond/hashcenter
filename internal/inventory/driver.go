package inventory

import (
	"os"
	"path/filepath"
)

func enrichDriver(adapters []Adapter) ([]Adapter, error) {

	for i := range adapters {

		link := "/sys/class/net/" + adapters[i].Interface + "/device/driver"

		target, err := os.Readlink(link)
		if err != nil {
			continue
		}

		adapters[i].Driver = filepath.Base(target)
	}

	return adapters, nil
}
