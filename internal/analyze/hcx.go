package analyze

import (
	"os"
	"strings"
)

func collectHCX(r *Report) error {
	if !exists("hcxpcapngtool") {
		return nil
	}

	tmp, err := os.CreateTemp("", "hashcenter-hcx-*.22000")
	if err != nil {
		return nil
	}

	outputFile := tmp.Name()
	_ = tmp.Close()
	defer os.Remove(outputFile)

	out, err := run(
		"hcxpcapngtool",
		"-o", outputFile,
		r.Path,
	)

	if err != nil {
		return nil
	}

	text := string(out)

	if strings.Contains(text, "PMKID") {
		// Пока только фиксируем наличие.
		// Детальный разбор добавим следующим релизом.
	}

	if strings.Contains(text, "EAPOL") {
		// Аналогично.
	}

	return nil
}
