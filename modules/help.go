package modules

import (
	"fmt"
)

//Help program usage
func Help(progName string) {
	fmt.Println("Discord-Exploits Help")
	fmt.Println("  Usage: " + progName + " -i <input file> -m <mode> [-q]")
	fmt.Println("")
	fmt.Println("--quiet -q   doesn't show welcome screen")
	fmt.Println("-i <file>    provide input file")
	fmt.Println("-m <mode>    specify mode")
	fmt.Println("")
	fmt.Println("modes:")
	fmt.Println("  video:")
	fmt.Println("    expandingvideo   takes input video (.webm) and edits it so discord will keep making it longer")
	fmt.Println("    negativevideo    takes input video (.webm) and edits it so discord will think it has got a huge negative duration")
	fmt.Println("    zerovideo        takes input video (.webm) and edits it so discord will think it has got a 0s duration")
	fmt.Println("  image:")
	fmt.Println("    virusimage       takes an image (.png) and makes other users' windows defender think it's a virus\n ")
}