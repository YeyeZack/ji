package modules

import (
	"fmt"
	"os/exec"
)

//Transcode transcodes videos and images
func Transcode(input string, to string) string {
	path := CheckForFFmpeg()
	output := CreateName(to)

	out, err := exec.Command(path, "-i", input, output).CombinedOutput()

	Check(err)

	fmt.Println("temporarily saving transcoded file to " + output)

	fmt.Println(string(out))

	return output
}
