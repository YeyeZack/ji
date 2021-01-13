package modules

import (
	"os/exec"
)

// CheckForFFmpeg looks for ffmpeg in the path
func CheckForFFmpeg() string {
	path, err := exec.LookPath("ffmpeg")

	if err != nil {
		panic(err)
	}

	return path
}
