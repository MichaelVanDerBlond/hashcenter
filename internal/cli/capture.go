package cli

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/MichaelVanDerBlond/hashcenter/internal/capture"
	"github.com/MichaelVanDerBlond/hashcenter/internal/session"
)

func Capture() error {

	fs := flag.NewFlagSet("capture", flag.ExitOnError)

	channel := fs.Int("channel", 1, "Wi-Fi channel")
	duration := fs.Int("time", 30, "Capture duration (seconds)")
	output := fs.String("output", "capture", "Output filename")

	if err := fs.Parse(os.Args[2:]); err != nil {
		return err
	}

	s := session.New()

	result, err := capture.Execute(
		s,
		*channel,
		*output,
		time.Duration(*duration)*time.Second,
	)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("========== SESSION ==========")
	fmt.Println("ID        :", s.ID)
	fmt.Println("State     :", s.State)
	fmt.Println("Interface :", s.Interface)
	fmt.Println("Channel   :", s.Channel)
	fmt.Println("Backend   :", s.Backend)
	fmt.Println("Capture   :", s.CaptureFile)
	fmt.Println("=============================")
	fmt.Println()

	fmt.Println("Capture completed")
	fmt.Printf("File     : %s\n", result.Output)
	fmt.Printf("Duration : %s\n", result.Duration.Round(time.Second))
	fmt.Printf("Success  : %v\n", result.Success)

	return nil
}
