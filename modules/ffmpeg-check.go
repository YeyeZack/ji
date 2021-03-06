package modules

import (
	"fmt"
	"os"
	"os/exec"
)

// CheckForFFmpeg looks for ffmpeg in the path
func CheckForFFmpeg() (string) {
	path, err := exec.LookPath("ffmpeg")

	if err != nil {
		fmt.Println("Could not find FFmpeg, maybe you don't have it installed or on your PATH")
		os.Exit(1)
	}
	return path
}
