package analyze

import (
	"strings"
)

func collectHCX(r *Report) error {

	if !exists("hcxpcapngtool") {
		return nil
	}

	out, err := run(
		"hcxpcapngtool",
		"-o", "/dev/null",
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
