[![GitHub stars](https://img.shields.io/github/stars/Schmenn/discord-exploits?color=FFFFFF&label=stars&style=flat-square)](https://github.com/Schmenn/discord-exploits/stargazers)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Schmenn/discord-exploits?color=FFFFFF&style=flat-square) 
![GitHub issues](https://img.shields.io/github/issues/Schmenn/discord-exploits?color=FFFFFF&style=flat-square)
[![GitHub forks](https://img.shields.io/github/forks/Schmenn/discord-exploits?color=FFFFFF&style=flat-square)](https://github.com/Schmenn/discord-exploits/network)
![Made with](https://img.shields.io/badge/made%20with-Go-29BEB0?style=flat-square)

# Discord-Exploits
A program for creating exploited media files for discord written in Go.

##### If you have any questions, feel free to ask me on ![Discord](https://img.shields.io/badge/Discord-Schmenn%231088-7289DA?style=flat-square) or join the server: [![Discord](https://img.shields.io/discord/809503251455148063?label=discord&style=flat-square)](https://discord.gg/QQfE4QtzFJ)

<b>DO NOT DM ME ASKING WHY THE WINDOW CLOSES INSTANTLY WHEN YOU DOUBLE CLICK IT, 

DISCORD-EXPLOITS IS A COMMAND-LINE UTILITY MEANING YOU HAVE TO USE IT FROM A TERMINAL

IF THE PROGRAM SHOWS AN ERROR, READ THE ERROR BEFORE MESSAGING ME</b>



### Getting Help

`discord-exploits -h` will show you what commands can be used and what you have to do

### Creating an Expanding Video file

`discord-exploits -m <mode> -i <input file> [-q]` 

the mode `c` stands for "crash", this file will crash almost any desktop discord client when it is played to the end

the mode `v` stands for virus image

the mode `e` stands for expanding video duration

the mode `n` stands for negative video duration

the mode `0` or `z` stands for video with 0s duration

the mode `t` stands for "twice", this file will play a different audio when it is replayed (still experimental)

The Program only supports `webm` files for video, `png` files for images and `ogg` files for audio (t)

##### The file will be saved with a random file name in the directory in which you ran the command
##### You can use [FFmpeg](https://ffmpeg.org) to convert a video to .webm or to convert an image to png (better than online converters)
##### The "virus" image may get flagged by windows defender and will get removed. to restore the file, go to Settings > Update & Security > Windows Security > Virus & Threat protection and restore the file

## Installation
### Via releases

go to [the releases page](https://github.com/Schmenn/discord-exploits/releases) and download either the windows or linux version

###### availabe systems are:
`windows-64-bit
windows-32-bit`

`linux-64-bit
linux-32-bit`

##### The Program was not yet tested on linux


### Compiling it yourself
###### make sure you have [Go](https://golang.org) installed and in your path
1. Clone this repository
    
    `git clone https://github.com/Schmenn/discord-exploits`

###### If you don't have git installed, install it [here](https://git-scm.com) or download the code as a zip
2. Go into the folder

    `cd discord-exploits`

3. Build it

    `go build`

###### the executable will have the name `discord-exploits`

## Features
### Current Features
* Feature for creating a video that, when played on discord, will seem like infinite, because the duration keeps getting longer
* Feature for creating a video that, when played on discord, will look like it has got a huge negative duration
* Feature for creating a video that, when played on discord, will look like it has got a constant duration of 0
* Feature for creating an image then triggers other users' windows defender after being cached
* Feature for creating an audio file that plays a different track when you replay it ([always requires FFmpeg](https://ffmpeg.org))
* Feature for creating a video that, when played on a discord client, will crash it ([always requires FFmpeg](https://ffmpeg.org))
* Transcoding from `mp4` to `webm`, from `jpeg` and `jpg` to `png` and from `mp4` and `m4a` to `ogg` ([requires FFmpeg](https://ffmpeg.org))

### Disclaimer
The t/twice mode requires ffmpeg with all compatible file formats, even `ogg`, because the current version of the exploit has to convert the audio to a sample rate of 44100.

The c/crash mode also requires ffmpeg will all compatible file formats, because the input file needs to be concatenated with another file.

### Upcoming Features are:
not yet decided.


