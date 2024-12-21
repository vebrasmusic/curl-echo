package util

import (
	"fmt"
	"time"
)

func ShowLoading(loading *bool) {
	spinner := []string{"|", "/", "-", "\\"}
	i := 0
	for *loading {
		fmt.Printf("\r%s", spinner[i%len(spinner)])
		i++
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Print("\r") // Clear spinner when done
}
