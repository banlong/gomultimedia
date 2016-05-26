//CONVERT A VIDEO
//Have to install the ffmpeg build version, set path to its \bin so that we can call ffmpeg

package main
import (
	"log"
	"gomultimedia/tools"
	"gomultimedia/transcode"
	"path"
	"time"
	"strings"

)

func main(){
	//VOD
	//ProduceDashIfFromEncodedSeg()
	//ProduceDash()
	//ProduceVODDash()


	//LIVE
	//ProduceDashIfLive()

	//ffmpeg.SplitVideo("videos/duck.mp4", "3", "factory/segments/")

	//Dash Using Bento
	pam := ffmpeg.BentoParam{
		InputVideo: 		"videos/asia-fragmented.mp4",
		OutputDir:			"factory/mpd/",
		Profile:			"live",
		UseSegmentTimeLine:	true,
		Debug:				false,
		NoSplit: 			false,
		UseExistingDir:     true,
	}
	ffmpeg.BentoDashIf(pam)
}


//THE DASH-IF MODULE CANNOT BE PLAYED TO THE END (LOST ABOUT 10%) VIDEO
func ProduceDashIfFromSplitSeg() error{
	//Play on DASH IF
	//This module produce DASH resource that be played on DASH IF player
	//INPUT: Using the split segments from long video

	processStart := time.Now()

	srcFile := "videos/kyduyen.mp4"
	tempDir := "factory/"
	segDir := tempDir + "segments/"
	mpdDir := tempDir + "mpd/"
	inputDir :=  tempDir + "input/"
	duration := 3


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	names, _ := ffmpeg.Split(srcFile, duration , segDir, "", ".mp4")
	splitTime := time.Since(splitStart)

	//Extract Video
	nameElement := names.Front()
	tools.CreateDir(mpdDir)
	tools.CreateDir(inputDir)

	for (nameElement != nil){
		filename := segDir + nameElement.Value.(string)
		baseName := path.Base(filename)
		parts := strings.Split(baseName, ".")
		name := parts[0]
		videoTrack := segDir + name + ".mp4"

		//copy video file and audio to an input folder
		video, _ := tools.GetBytes(videoTrack)

		tools.SaveBinFile2Disk(video, inputDir, "seg.mp4")
		videoInput := inputDir + "seg.mp4"
		ffmpeg.ProduceDashifFromMuxSeg(3000, mpdDir, "kd2.mpd",  videoInput)
		nameElement = nameElement.Next()
	}

	processTime := time.Since(processStart)

	log.Println("PROCESSING TIME - ", processTime)
	log.Println("--------------------------------")
	log.Println("Split Time: ", splitTime)


	return nil
}
func ProduceDashIfFromEncodedSeg3() error{
	//Play on DASH IF
	//This module produce DASH resource that be played on DASH IF player
	//INPUT: Using the encoded segment & aac video to muxed into Video
	//Flow: Split - Extract - Transcode - Mux

	processStart := time.Now()

	srcFile := "videos/kyduyen.mp4"
	tempDir := "factory/"
	segDir := tempDir + "segments/"
	videoDir := tempDir + "video/"
	audioDir := tempDir + "audio/"
	encodeDir := tempDir + "encode/"
	mpdDir := tempDir + "mpd/"
	inputDir :=  tempDir + "input/"
	audioExt := ".aac"
	duration := 10


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	names, err := ffmpeg.Split(srcFile, duration , segDir, "", ".mp4")
	splitTime := time.Since(splitStart)

	//Extract Video
	nameElement := names.Front()
	tools.CreateDir(videoDir)
	tools.CreateDir(audioDir)
	tools.CreateDir(encodeDir)
	tools.CreateDir(mpdDir)
	tools.CreateDir(inputDir)

	for (nameElement != nil){
		filename := segDir + nameElement.Value.(string)
		baseName := path.Base(filename)
		parts := strings.Split(baseName, ".")
		name := parts[0]
		videoTrack := videoDir + name + ".mp4"
		audioTrack := audioDir + name + audioExt
		err = ffmpeg.ExtractAV(filename, videoTrack, audioTrack)
		if err != nil{
			return err
		}

		//encode video
		err = ffmpeg.Transcode(videoTrack, encodeDir , name + ".mp4", "ultrafast", "mq")
		if err != nil{
			return err
		}

		encodeTrack := encodeDir + name + ".mp4"
		//copy video file and audio to an input folder
		//video, _ := tools.GetBytes(encodeTrack)
		//audio, _ := tools.GetBytes(audioTrack)
		//tools.SaveBinFile2Disk(video, inputDir, "seg.mp4")
		//tools.SaveBinFile2Disk(audio, inputDir, "seg.aac")

		videoInput := inputDir + "seg.mp4"
		ffmpeg.Mux(encodeTrack, audioTrack,videoInput)
		ffmpeg.ProduceDashifFromMuxSeg(3000, mpdDir, "kd2.mpd",  videoInput)
		nameElement = nameElement.Next()
	}

	processTime := time.Since(processStart)

	log.Println("PROCESSING TIME - ", processTime)
	log.Println("--------------------------------")
	log.Println("Split Time: ", splitTime)


	return nil
}
func ProduceDashIfFromEncodedSeg2() error{
	//Play on DASH IF
	//This module produce DASH resource that be played on DASH IF player
	//INPUT: Using the encoded segment & aac video to muxed into Video
	//Flow: Split - Transcode - Extract- Mux

	processStart := time.Now()

	srcFile := "videos/kyduyen.mp4"
	tempDir := "factory/"
	segDir := tempDir + "segments/"
	videoDir := tempDir + "video/"
	audioDir := tempDir + "audio/"
	encodeDir := tempDir + "encode/"
	mpdDir := tempDir + "mpd/"
	inputDir :=  tempDir + "input/"
	audioExt := ".aac"
	duration := 3


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	names, err := ffmpeg.Split(srcFile, duration , segDir, "", ".mp4")
	splitTime := time.Since(splitStart)


	nameElement := names.Front()
	tools.CreateDir(videoDir)
	tools.CreateDir(audioDir)
	tools.CreateDir(encodeDir)
	tools.CreateDir(mpdDir)
	tools.CreateDir(inputDir)

	for (nameElement != nil){
		filename := segDir + nameElement.Value.(string)
		baseName := path.Base(filename)
		parts := strings.Split(baseName, ".")
		name := parts[0]

		//encode video
		err = ffmpeg.Transcode(filename, encodeDir , name + ".mp4", "ultrafast", "mq")
		if err != nil{
			return err
		}

		//Extract Video
		encodeSeg := segDir + name + ".mp4"
		videoTrack := videoDir + name + ".mp4"
		audioTrack := audioDir + name + audioExt
		err = ffmpeg.ExtractAV(encodeSeg, videoTrack, audioTrack)
		if err != nil{
			return err
		}

		//Dash File
		//copy video file and audio to an input folder
		video, _ := tools.GetBytes(videoTrack)
		audio, _ := tools.GetBytes(audioTrack)

		tools.SaveBinFile2Disk(video, inputDir, "seg.mp4")
		tools.SaveBinFile2Disk(audio, inputDir, "seg.aac")

		videoInput := inputDir + "seg.mp4"
		audioInput := inputDir + "seg.aac"
		ffmpeg.ProduceDashif("3000", mpdDir, "kd2.mpd",  videoInput, audioInput)
		nameElement = nameElement.Next()
	}

	processTime := time.Since(processStart)

	log.Println("PROCESSING TIME - ", processTime)
	log.Println("--------------------------------")
	log.Println("Split Time: ", splitTime)


	return nil
}
func ProduceDashIfFromEncodedSeg() error{
	//Play on DASH IF
	//This module produce DASH resource that be played on DASH IF player
	//INPUT: Using the encoded segment & aac video to muxed into Video
	//Flow: Split - Extract - Transcode - Mux

	processStart := time.Now()

	srcFile := "videos/talk.mp4"
	tempDir := "factory/"
	segDir := tempDir + "segments/"
	videoDir := tempDir + "video/"
	audioDir := tempDir + "audio/"
	encodeDir := tempDir + "encode/"
	mpdDir := tempDir + "mpd/"
	inputDir :=  tempDir + "input/"
	audioExt := ".aac"
	splitDuration := 3
	dashDuration :="3000"

	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	names, _ := ffmpeg.Split(srcFile, splitDuration , segDir, "", ".mp4")
	splitTime := time.Since(splitStart)

	//Extract Video
	nameElement := names.Front()
	tools.CreateDir(videoDir)
	tools.CreateDir(audioDir)
	tools.CreateDir(encodeDir)
	tools.CreateDir(mpdDir)
	tools.CreateDir(inputDir)

	for (nameElement != nil){
		filename := segDir + nameElement.Value.(string)
		baseName := path.Base(filename)
		parts := strings.Split(baseName, ".")
		name := parts[0]
		videoTrack := videoDir + name + ".mp4"
		audioTrack := audioDir + name + audioExt
		err := ffmpeg.ExtractAV(filename, videoTrack, audioTrack)
		if err != nil{
			return err
		}

		//encode video
		err = ffmpeg.Transcode(videoTrack, encodeDir , name + ".mp4", "ultrafast", "mq")
		if err != nil{
			return err
		}

		encodeTrack := encodeDir + name + ".mp4"
		//copy video file and audio to an input folder
		video, _ := tools.GetBytes(encodeTrack)
		audio, _ := tools.GetBytes(audioTrack)

		tools.SaveBinFile2Disk(video, inputDir, "seg.mp4")
		tools.SaveBinFile2Disk(audio, inputDir, "seg.aac")

		videoInput := inputDir + "seg.mp4"
		audioInput := inputDir + "seg.aac"
		//m4aAudioInput := inputDir + "seg.m4a"
		//ffmpeg.ConvertAAC2M4A(audioInput, m4aAudioInput)
		//duration := ffmpeg.GetDurationInMillisecond(videoInput)
		ffmpeg.ProduceDashif(dashDuration, mpdDir, "kd2.mpd",  videoInput, audioInput)


		nameElement = nameElement.Next()
	}

	processTime := time.Since(processStart)

	log.Println("PROCESSING TIME - ", processTime)
	log.Println("--------------------------------")
	log.Println("Split Time: ", splitTime)


	return nil
}

//LIVE DASH
func ProduceDashIfLive() error{
	//This module produce DASH resource that be played on DASH IF player
	//INPUT: Using the encoded segment & aac video to muxed into Video
	//Flow: Split - Extract - Transcode - Mux

	processStart := time.Now()

	srcFile := "videos/kyduyen.mp4"
	tempDir := "factory/"
	segDir := tempDir + "segments/"
	videoDir := tempDir + "video/"
	audioDir := tempDir + "audio/"
	encodeDir := tempDir + "encode/"
	mpdDir := tempDir + "mpdlive/"
	inputDir :=  tempDir + "input/"
	audioExt := ".aac"
	duration := 3


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	names, err := ffmpeg.Split(srcFile, duration , segDir, "", ".mp4")
	splitTime := time.Since(splitStart)

	//Extract Video
	nameElement := names.Front()
	tools.CreateDir(videoDir)
	tools.CreateDir(audioDir)
	tools.CreateDir(encodeDir)
	tools.CreateDir(mpdDir)
	tools.CreateDir(inputDir)

	for (nameElement != nil){
		filename := segDir + nameElement.Value.(string)
		baseName := path.Base(filename)
		parts := strings.Split(baseName, ".")
		name := parts[0]
		videoTrack := videoDir + name + ".mp4"
		audioTrack := audioDir + name + audioExt
		err = ffmpeg.ExtractAV(filename, videoTrack, audioTrack)
		if err != nil{
			return err
		}

		//encode video
		err = ffmpeg.Transcode(videoTrack, encodeDir , name + ".mp4", "ultrafast", "mq")
		if err != nil{
			return err
		}

		encodeTrack := encodeDir + name + ".mp4"
		//copy video file and audio to an input folder
		video, _ := tools.GetBytes(encodeTrack)
		audio, _ := tools.GetBytes(audioTrack)

		tools.SaveBinFile2Disk(video, inputDir, "seg.mp4")
		tools.SaveBinFile2Disk(audio, inputDir, "seg.aac")

		videoInput := inputDir + "seg.mp4"
		audioInput := inputDir + "seg.aac"
		//NON-MULTIPLEX
		ffmpeg.ProduceDashifLive(3000, mpdDir, "kd2.mpd",  videoInput, audioInput)

		//MULTIPLEX
		//ffmpeg.ProduceDashLive(3000, mpdDir, "kd2.mpd",  videoInput, audioInput)
		nameElement = nameElement.Next()
	}

	processTime := time.Since(processStart)

	log.Println("PROCESSING TIME - ", processTime)
	log.Println("--------------------------------")
	log.Println("Split Time: ", splitTime)


	return nil
}

//MULTIPLEX VOD DASH
func ProduceDash() error{
	//Multiplex Dash, not play on DashIF

	processStart := time.Now()

	srcFile := "videos/kyduyen.mp4"
	mpdName := "kd.mpd"
	segmentName:= "kd_"

	tempDir := "factory/"
	segDir := tempDir + "segments/"
	videoDir := tempDir + "video/"
	audioDir := tempDir + "audio/"
	encodeDir := tempDir + "encode/"
	mpdDir := tempDir + "mpd/"
	audioExt := ".aac"
	duration := 3


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	names, err := ffmpeg.Split(srcFile, duration , segDir, "", ".mp4")
	splitTime := time.Since(splitStart)

	//Extract Video
	extractStart := time.Now()
	nameElement := names.Front()
	tools.CreateDir(videoDir)
	tools.CreateDir(audioDir)
	for (nameElement != nil){
		filename := segDir + nameElement.Value.(string)  // videos/temp/segments/sample-0000.mp4
		baseName := path.Base(filename)			 // sample-0000.mp4
		parts := strings.Split(baseName, ".")		 // [sample-0000, mp4]
		name := parts[0]				 // sample-0000

		//log.Println("File: ", filename)
		//log.Println("Name: ", name)

		err = ffmpeg.ExtractAV(filename, videoDir + name + ".mp4" , audioDir + name + audioExt)
		if err != nil{
			return err
		}
		nameElement = nameElement.Next()
	}
	extractTime :=  time.Since(extractStart)

	//Encode videos
	encodeStart := time.Now()
	tools.CreateDir(encodeDir)
	nameElement = names.Front()
	for (nameElement != nil){
		filename := videoDir + nameElement.Value.(string)  // videos/temp/segments/sample-0000.mp4
		baseName := path.Base(filename)			 // sample-0000.mp4
		parts := strings.Split(baseName, ".")		 // [sample-0000, mp4]
		name := parts[0]				 // sample-0000

		//log.Println("File: ", filename)
		//log.Println("Name: ", name)

		err = ffmpeg.Transcode(filename, encodeDir , name + ".mp4", "ultrafast", "mq")
		if err != nil{
			return err
		}
		nameElement = nameElement.Next()
	}
	encodeTime :=  time.Since(encodeStart)

	//Dash files
	tools.CreateDir(mpdDir)
	nameElement = names.Front()
	for (nameElement != nil) {
		namecomp := strings.Split(nameElement.Value.(string), ".")
		audioFile := audioDir + namecomp[0] + audioExt
		videoFile := encodeDir + nameElement.Value.(string)

		ffmpeg.ProduceDash(3000, mpdDir, mpdName,  videoFile, audioFile, segmentName)
		nameElement = nameElement.Next()
	}


	processTime := time.Since(processStart)

	log.Println("PROCESSING TIME - ", processTime)
	log.Println("--------------------------------")
	log.Println("Split Time: ", splitTime)
	log.Println("Extract Time: ", extractTime)
	log.Println("Encode Time: ", encodeTime)

	return nil
}
