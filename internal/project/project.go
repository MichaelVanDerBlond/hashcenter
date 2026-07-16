package project

import "github.com/MichaelVanDerBlond/hashcenter/internal/hashes"

func Load() ([]Network, error) {

	list, err := hashes.List("output/hashes")
	if err != nil {
		return nil, err
	}

	var out []Network

	for i, h := range list {

		out = append(out, Network{
			ID: i + 1,

			ESSID: h.ESSID,
			BSSID: h.BSSID,

			Handshake: h.Handshakes,
			PMKID:     h.PMKIDs,

			HashFile: h.Path,
		})
	}

	return out, nil
}
