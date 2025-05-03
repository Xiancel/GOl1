package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fileName := "doc.pdf"
	extension := filepath.Ext(fileName)
	fmt.Println(extension)
	if extension == ".jpg" || extension == ".png" || extension == ".gif" {
		fmt.Println("it's Picture WoW")
	} else if extension == ".mp4" || extension == ".mov" || extension == ".avi" {
		fmt.Println("it's video Wow")
	} else if extension == ".mp3" || extension == ".wav" || extension == ".flac" {
		fmt.Println("it's music woW")
	} else if extension == ".doc" || extension == ".pdf" || extension == ".txt" {
		fmt.Println("it's document WooooooW")
	} else {
		fmt.Println("Unknow")
	}
}
