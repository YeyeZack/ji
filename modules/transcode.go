package modules

import (
	"fmt"
	"os/exec"
	"strings"
)

//Transcode transcodes videos and images
func Transcode(input string, to string) string {
	path := CheckForFFmpeg()
	output := CreateName(to)

	if strings.HasSuffix(input, ".ogg") || strings.HasSuffix(input, ".mp3") || strings.HasSuffix(input, ".m4a") {
		out, err := exec.Command(path, "-i", input, "-ar", "44100", "-y", "-c:a", "libvorbis", output).CombinedOutput()
		Check(err)

		fmt.Println("temporarily saving transcoded file to " + output)

		fmt.Println(string(out))

		return output
	}

	out, err := exec.Command(path, "-i", input, output).CombinedOutput()

	Check(err)

	fmt.Println("temporarily saving transcoded file to " + output)

	fmt.Println(string(out))

	return output
}
