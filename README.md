# Discord-Exploits
A program for creating exploited media files for discord written in Go.

## Usage
###### discord-exploits is a command line utility, meaning you have to run it through a terminal!

### Getting Help

`discord-exploits -h` will show you what commands can be used and what you have to do

### Creating an Expanding Video file

`discord-exploits -m expandingvideo -i <input file>.webm` 

The Program only supports .webm files

##### The video will be saved with a random file name in the directory in which you ran the command
##### You can use [ffmpeg](https://ffmpeg.org) to convert a video to .webm

### Creating a "Virus" Image

`discord-exploits -m virusimage -i <input file>.png` 

The Program currently only supports PNG, other files will be implemented in the future

##### The image will be saved with a random file name in the directory in which you ran the command
##### The image may get flagged by windows defender and will get removed. to restore the file, go to Settings > Update & Security > Windows Security > Virus & Threat protection and restore the file


## Installation
### Via releases

go to [the releases page](https://github.com/Schmenn/discord-exploits/releases) and download `discord-exploits.exe` 
###### only windows amd64, pre-compiled executables for other systems are getting added in future versions


### Compiling it yourself
###### make sure you have [Go](https://golang.org) installed and in your path
1. Clone this repository
    
    `git clone https://github.com/Schmenn/discord-exploits`

2. Go into the folder

    `cd discord-exploits`

3. Compile it

    `go build`

###### the executable will have the name `discord-exploits`

## Features
### Current Features
* Feature for creating a video that, when played on discord, will seem like infinite, because the duration keeps getting longer
* Feature for creating an image then triggers other users' windows defender after being cached

### Upcoming Features
* Feature for creating a video that, when played on discord, will look like it has got a huge negative duration

* Feature for creating a video that, when played on discord, will look like it has got a constant duration of 0

##### If you have any questions, feel free to ask me on Discord (Schmenn#1088)
