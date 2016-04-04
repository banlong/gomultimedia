package main

import (
"github.com/giorgisio/goav/avformat"
"log"

)

func main() {

	filename := "videos/sample.mp4"
	var ctxtFormat    *avformat.Context

	 //Register all formats and codecs
	avformat.AvRegisterAll()

	// Open video file
	if avformat.AvformatOpenInput(&ctxtFormat, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctxtFormat.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")
		return
	}
}
