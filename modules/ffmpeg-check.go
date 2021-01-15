package modules

import (
	"os/exec"
)

// CheckForFFmpeg looks for ffmpeg in the path
func CheckForFFmpeg() string {
	path, err := exec.LookPath("ffmpeg")

	Check(err)

	return path
}
