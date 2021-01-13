package modules

import(
	"os/exec"
)

//Transcode transcodes video to webm
func Transcode (input string, to string) string {
	path := CheckForFFmpeg()
	output := CreateName(to)

	cmd := exec.Command(path, "-i", input, output)
	err := cmd.Run()
	Check(err)
	return output
}
