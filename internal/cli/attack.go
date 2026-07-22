package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/MichaelVanDerBlond/hashcenter/internal/attack"
)

func Attack() error {

	fs := flag.NewFlagSet("attack", flag.ContinueOnError)

	hashes := fs.String("hashes", "", "Path to .22000 file")
	dict := fs.String("dict", "", "Dictionary file or comma-separated list")

	mode := fs.Int("mode", 22000, "Hash mode")
	attackMode := fs.Int("attack", 0, "Hashcat attack mode")

	rule := fs.String("rule", "", "Rule file")
	mask := fs.String("mask", "", "Mask")

	session := fs.String("session", "", "Hashcat session name")
	device := fs.String("device", "", "GPU device list")
	workload := fs.Int("workload", 0, "Hashcat workload profile")

	if err := fs.Parse(os.Args[2:]); err != nil {
		return err
	}

	if *hashes == "" {
		return errors.New("missing --hashes")
	}

	if *dict == "" && *attackMode != 3 {
		return errors.New("missing --dict")
	}

	var dictionaries []string

	if *dict != "" {
		for _, d := range strings.Split(*dict, ",") {
			d = strings.TrimSpace(d)
			if d != "" {
				dictionaries = append(dictionaries, d)
			}
		}

		if len(dictionaries) == 0 {
			return errors.New("no valid dictionaries specified")
		}
	}

	if len(dictionaries) <= 1 {

		job := attack.Job{
			HashFile:    *hashes,
			Dictionary:  *dict,
			HashMode:    *mode,
			AttackMode:  *attackMode,
			Rule:        *rule,
			Mask:        *mask,
			SessionName: *session,
			Device:      *device,
			Workload:    *workload,
			Extra:       fs.Args(),
		}

		return attack.Execute(context.Background(), job)
	}

	jobs := make([]attack.Job, 0, len(dictionaries))

	for _, dictionary := range dictionaries {

		jobs = append(jobs, attack.Job{
			HashFile:    *hashes,
			Dictionary:  dictionary,
			HashMode:    *mode,
			AttackMode:  *attackMode,
			Rule:        *rule,
			Mask:        *mask,
			SessionName: *session,
			Device:      *device,
			Workload:    *workload,
			Extra:       fs.Args(),
		})
	}

	result, err := attack.ExecuteMany(context.Background(), jobs)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("Completed :", result.Succeed)
	fmt.Println("Failed    :", result.Failed)

	return nil
}
