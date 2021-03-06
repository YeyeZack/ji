package modules

import (
	"fmt"
)

//Help program usage
func Help(progName string) {
	fmt.Println("Discord-Exploits Help")
	fmt.Println("  Usage: " + progName + " -i <input file> -m <mode> [-q]")
	fmt.Println("")
	fmt.Println("--quiet -q   don't show welcome screen")
	fmt.Println("-i <file>    provide input file")
	fmt.Println("-m <mode>    specify mode")
	fmt.Println("")
	fmt.Println("modes:")
	fmt.Println("  video:")
	fmt.Println("    e     takes input video (.webm or .mp4) and edits it so discord will keep making it longer")
	fmt.Println("    n     takes input video (.webm or .mp4) and edits it so discord will think it has got a huge negative duration")
	fmt.Println("    z, 0  takes input video (.webm or .mp4) and edits it so discord will think it has got a 0s duration")
	fmt.Println("    c     takes input video (.webm or .mp4) and edits it so discord will crash when you play it to the end (only works on some PCs)")
	fmt.Println("  image:")
	fmt.Println("    v     takes an image (.png, .jpg or .jpeg) and makes other users' windows defender think it's a virus")
	fmt.Println("  audio:")
	fmt.Println("    t     takes an audio file (.ogg, .m4a or .mp3) and edits it so discord will play a second audio track when you replay it.")
	fmt.Println("")
	fmt.Println("the following file formats require ffmpeg: .mp4, .jpg, .jpeg, .ogg, .m4a and .mp3")
	fmt.Println("")

}
