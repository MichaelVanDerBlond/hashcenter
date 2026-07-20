package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/MichaelVanDerBlond/hashcenter/internal/attack"
)

func Attack() error {

	fs := flag.NewFlagSet("attack", flag.ContinueOnError)

	hashes := fs.String("hashes", "", "Path to .22000 file")
	dict := fs.String("dict", "", "Dictionary file")
	mode := fs.Int("mode", 22000, "Hash mode")
	attackMode := fs.Int("attack", 0, "Hashcat attack mode")

	if err := fs.Parse(os.Args[2:]); err != nil {
		return err
	}

	if *hashes == "" {
		return errors.New("missing --hashes")
	}

	if *dict == "" {
		return errors.New("missing --dict")
	}

	job := attack.Job{
		HashFile:   *hashes,
		Dictionary: *dict,
		HashMode:   *mode,
		AttackMode: *attackMode,
		Extra:      fs.Args(),
	}

	fmt.Println("Hash file :", job.HashFile)
	fmt.Println("Dictionary:", job.Dictionary)

	return attack.Execute(context.Background(), job)
}
