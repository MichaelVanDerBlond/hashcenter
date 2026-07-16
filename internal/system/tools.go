package system

import (
	"os/exec"
	"strings"
)

type ToolInfo struct {
	Name    string
	Version string
	Present bool
}

func collectTools(info *Info) {

	tools := []struct {
		Name string
		Cmd  string
		Args []string
	}{
		{"Hashcat", "hashcat", []string{"--version"}},
		{"Aircrack-ng", "aircrack-ng", []string{"--help"}},
		{"hcxpcaptool", "hcxpcaptool", []string{"--help"}},
		{"hcxhashtool", "hcxhashtool", []string{"--help"}},
		{"TShark", "tshark", []string{"--version"}},
	}

	for _, t := range tools {

		out, err := exec.Command(t.Cmd, t.Args...).CombinedOutput()

		item := ToolInfo{
			Name: t.Name,
		}

		if err == nil {
			item.Present = true

			lines := strings.Split(strings.TrimSpace(string(out)), "\n")

			if len(lines) > 0 {
				item.Version = lines[0]
			}
		}

		info.Tools = append(info.Tools, item)
	}
}
