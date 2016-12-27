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

	//TEST EXTRACT & MUX FOR MOBILE
	//outputDir := "mobi/out/"
	//video := outputDir + "video.mp4"
	//audio := outputDir+ "audio.aac"
	//tools.CreateDir(outputDir)
	//ffmpeg.ExtractAV("mobi/1.mp4", video, audio )
	//ffmpeg.MP4BoxAudioMux(video, audio)


	ProduceDashAdaptiveFromDir()
}

func ProduceDashAdaptiveFromDir()  error{
	processStart := time.Now()
	tempDir := "adaptive/"
	videoDir := tempDir +"mp4/"
	hqDir := videoDir + "hq/"
	mqDir := videoDir + "mq/"
	lqDir := videoDir + "lq/"
	audioDir := tempDir + "segments/"
	dashDir := tempDir + "dash/"
	mpdDir := dashDir + "mpd/"
	inputDir1 :=  dashDir + "input1/"
	inputDir2 :=  dashDir + "input2/"
	inputDir3 :=  dashDir + "input3/"

	tools.CreateDir(dashDir)
	tools.CreateDir(mpdDir)
	tools.CreateDir(inputDir1)
	tools.CreateDir(inputDir2)
	tools.CreateDir(inputDir3)
	names := tools.GetFileNames(mqDir)

	videoInput1 := inputDir1 + "video1.mp4"
	videoInput2 := inputDir2 + "video2.mp4"
	videoInput3 := inputDir3 + "video3.mp4"
	audioInput := inputDir1 + "audio.aac"

	nameElement := names.Front()
	for (nameElement != nil){
		videoTrack1 := mqDir + nameElement.Value.(string)
		videoTrack2 := hqDir + nameElement.Value.(string)
		videoTrack3 := lqDir + nameElement.Value.(string)

		baseName := path.Base(videoTrack1)
		parts := strings.Split(baseName, ".")
		name := parts[0]
		audioTrack := audioDir + name + ".aac"

		//copy video file to an input folder
		video1, _ := tools.GetBytes(videoTrack1)
		tools.SaveBinFile2Disk(video1, inputDir1, "video1.mp4")
		video2, _ := tools.GetBytes(videoTrack2)
		tools.SaveBinFile2Disk(video2, inputDir2, "video2.mp4")
		video3, _ := tools.GetBytes(videoTrack3)
		tools.SaveBinFile2Disk(video3, inputDir3, "video3.mp4")
		audio, _ := tools.GetBytes(audioTrack)
		tools.SaveBinFile2Disk(audio, inputDir1, "audio.aac")

		//Mux video
		ffmpeg.MP4BoxAudioMux(videoInput1, audioTrack)

		input:= ffmpeg.MP4BoxParameter{
			Video_Track1: 		videoInput1,
			Video_Track2: 		videoInput2,
			Video_Track3: 		videoInput3,
			Audio_Track: 		audioInput,
			DashDuration:		"3000",
			MpdDirectory:   	mpdDir,
			MpdName: 			"bug.mpd",
			UseSegmentTimeline:	true,
			DashCTX: 			true,
			Profile:			"live",
			FragDuration:		"3000",
			RandomAccess: 		true,
			BitstreamSwitch: 	"merge",
		}
		ffmpeg.CreateDashifFromMuxSeg(input)


		nameElement = nameElement.Next()
	}

	processTime := time.Since(processStart)
	log.Println("PROCESSING TIME - ", processTime)



	return nil
}


//Test Bug: dash cannot play on chrome, error decoding
//This module test to find the cause of the above issue
func DashProduceFromDir()  error{
	processStart := time.Now()
	tempDir := "bug/"
	videoDir := tempDir +"video/"
	audioDir := tempDir + "audio/"
	mpdDir := tempDir + "mpd/"
	inputDir :=  tempDir + "input/"

	tools.CreateDir(mpdDir)
	tools.CreateDir(inputDir)
	names := tools.GetFileNames(videoDir)

	videoInput := inputDir + "seg.mp4"
	nameElement := names.Front()
	for (nameElement != nil){
		videoTrack := videoDir + nameElement.Value.(string)
		ffmpeg.Transcode(videoTrack, inputDir,  "seg.mp4", "ultrafast", "mq")

		baseName := path.Base(videoTrack)
		parts := strings.Split(baseName, ".")
		name := parts[0]
		audioTrack := audioDir + name + ".aac"

		//copy video file to an input folder
		//video, _ := tools.GetBytes(videoTrack)
		//tools.SaveBinFile2Disk(video, inputDir, "seg.mp4")
		//videoInput := inputDir + "seg.mp4"

		//Mux video
		ffmpeg.MP4BoxAudioMux(videoInput, audioTrack)

		input:= ffmpeg.MP4BoxParameter{
			Video_Track1: 		videoInput,
			DashDuration:		"3000",
			MpdDirectory:   	mpdDir,
			MpdName: 			"bug.mpd",
			UseSegmentTimeline:	true,
			DashCTX: 			true,
			Profile:			"live",
			FragDuration:		"3000",
			RandomAccess: 		true,
		}

		ffmpeg.CreateDashifFromMuxSeg(input)

		nameElement = nameElement.Next()
	}

	processTime := time.Since(processStart)
	log.Println("PROCESSING TIME - ", processTime)



	return nil
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
	duration := "3"


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	sPam := ffmpeg.FfmpegParam{
		InputVideo: 		srcFile,
		SegmentDuration: 	duration,
		SegmentExt: 		".mp4",
		SegmentName: 		"",
		OutputLocation: 	tempDir,
		FrameRate: 			"30",
		Debug: 				false,
		SegmentList: 		tempDir + "list.txt",
	}
	fObj := ffmpeg.NewFfmpeg()
	names, _ := fObj.Split(sPam)
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

		input:= ffmpeg.MP4BoxParameter{
			Video_Track1: 		videoInput,
			DashDuration:		"3000",
			MpdDirectory:   	mpdDir,
			MpdName: 			"bug.mpd",
			UseSegmentTimeline:	true,
			DashCTX: 			true,
			Profile:			"live",
			FragDuration:		"3000",
			RandomAccess: 		true,
		}

		ffmpeg.CreateDashifFromMuxSeg(input)
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
	duration := "10"


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	sPam := ffmpeg.FfmpegParam{
		InputVideo: 		srcFile,
		SegmentDuration: 	duration,
		SegmentExt: 		".mp4",
		SegmentName: 		"sample",
		OutputLocation: 	segDir,
		FrameRate: 			"30",
		Debug: 				false,
		SegmentList: 		segDir + "list.txt",
	}
	fObj := ffmpeg.NewFfmpeg()
	names, err := fObj.Split(sPam)
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
	duration := "3"


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	sPam := ffmpeg.FfmpegParam{
		InputVideo: 		srcFile,
		SegmentDuration: 	duration,
		SegmentExt: 		".mp4",
		SegmentName: 		"",
		OutputLocation: 	segDir,
		FrameRate: 			"30",
		Debug: 				false,
		SegmentList: 		segDir + "list.txt",
	}
	fObj := ffmpeg.NewFfmpeg()
	names, err := fObj.Split(sPam)
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

	srcFile := "videos/hy.mp4"
	tempDir := "factory/"
	segDir := tempDir + "segments/"
	videoDir := tempDir + "video/"
	audioDir := tempDir + "audio/"
	encodeDir := tempDir + "encode/"
	mpdDir := tempDir + "mpd/"
	inputDir :=  tempDir + "input/"
	audioExt := ".aac"
	splitDuration := "3"

	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	sPam := ffmpeg.FfmpegParam{
		InputVideo: 		srcFile,
		SegmentDuration: 	splitDuration,
		SegmentExt: 		".mp4",
		SegmentName: 		"",
		OutputLocation: 	segDir,
		FrameRate: 			"30",
		Debug: 				false,
		SegmentList: 		segDir + "list.txt",
	}
	fObj := ffmpeg.NewFfmpeg()
	names, _ := fObj.Split(sPam)
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
		//ffmpeg.ProduceDashif(dashDuration, mpdDir, "kd2.mpd",  videoInput, audioInput)

		//Mux
		ffmpeg.MP4BoxAudioMux(videoInput, audioInput)

		//Dash
		input:= ffmpeg.MP4BoxParameter{
			Video_Track1: 		videoInput,
			DashDuration:		"3000",
			MpdDirectory:   	mpdDir,
			MpdName: 			"bug.mpd",
			UseSegmentTimeline:	true,
			DashCTX: 			true,
			Profile:			"live",
			FragDuration:		"3000",
			RandomAccess: 		true,
		}

		ffmpeg.CreateDashifFromMuxSeg(input)


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
	duration := "3"


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	sPam := ffmpeg.FfmpegParam{
		InputVideo: 		srcFile,
		SegmentDuration: 	duration,
		SegmentExt: 		".mp4",
		SegmentName: 		"",
		OutputLocation: 	segDir,
		FrameRate: 			"30",
		Debug: 				false,
		SegmentList: 		segDir + "list.txt",
	}
	fObj := ffmpeg.NewFfmpeg()
	names, err := fObj.Split(sPam)

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
	duration := "3"


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	sPam := ffmpeg.FfmpegParam{
		InputVideo: 		srcFile,
		SegmentDuration: 	duration,
		SegmentExt: 		".mp4",
		SegmentName: 		"",
		OutputLocation: 	segDir,
		FrameRate: 			"30",
		Debug: 				false,
		SegmentList: 		segDir + "list.txt",
	}
	fObj := ffmpeg.NewFfmpeg()
	names, err := fObj.Split(sPam)
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

//SPLIT WITH NEW FEATURE
func TestNewSplit()  {
		sPam := ffmpeg.FfmpegParam{
			InputVideo: "videos/trutinh.mp4",
			SegmentDuration: "3",
			SegmentExt: ".mp4",
			SegmentName: "",
			OutputLocation: "factory/trutinh/",
			FrameRate: "30",
			Debug: false,
			SegmentList:"factory/trutinh/list.txt",
		}
		fObj := ffmpeg.NewFfmpeg()

		fObj.Split(sPam)
}

func TestBentoDash()  {
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