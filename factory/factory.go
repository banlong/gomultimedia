package factory

import (
	"log"
	"os"
	"sync"
	"gomultimedia/worker"
	"gomultimedia/transcode"
	"gomultimedia/tools"
	"strings"
	"path"
	"time"
	"flag"
	"gomultimedia/mp4"
	"io/ioutil"
	"strconv"
	"fmt"
	"gomultimedia/config"
	"os/exec"
)


func DashFromEqualDurationSegments(){
	//Flow: extract video -> insert key frames --> split at 3s
	//Expectation:
	// 1 - no black frame at the end & beginning of video, video play smooth
	// 2 - Split exactly at 3 seconds segments (usually split at key frames, -->unequally segments)
	// If this success, video dashing does not need segment timeline.

	// Extract Video
	//videoInputPath := "videos/TTGS.mp4"
	//videoOutputPath := "producevideo/TTGS-v.mp4"
	//audioOutput := "producevideo/TTGS-a.mp4"
	//ffmpeg.ExtractAV(videoInputPath, videoOutputPath, audioOutput)

	//Insert Keyframe


	//Split Video
	sPam := ffmpeg.FFMPEGParam{
		InputVideo: "videos/DP.mp4",
		SegmentDuration: "3",
		SegmentExt: ".mp4",
		SegmentName: "",
		SegmentList: "",
		OutputLocation: "producevideo/v-segments/",
		FrameRate: "30",
		Debug: false,
		DisplayCMD: true,
	}
	names, _ :=ffmpeg.Split2(sPam)

	//Encoding


	//Dashing
	tempDir := "producevideo/"
	segDir := tempDir + "v-segments/"
	mpdDir := tempDir + "mpd/"
	inputDir :=  tempDir + "temp/"

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
			DashDuration:		"5000",
			MpdDirectory:   	mpdDir,
			MpdName: 		"bug.mpd",
			UseSegmentTimeline:	false,
			DashCTX: 		true,
			Profile:		"live",
			FragDuration:		"5000",
			RandomAccess: 		true,
		}

		ffmpeg.CreateDashifFromMuxSeg(input)
		nameElement = nameElement.Next()
	}

	log.Println("Dash Completed")
}

func ProduceDashIf() error{
	//DashIF Producing, after produce had to modify the mpd
	//Only work with mp4 audio, so extract as mp4 audio

	processStart := time.Now()

	srcFile := "videos/kyduyen.mp4"
	tempDir := "factory/"
	segDir := tempDir + "segments/"
	videoDir := tempDir + "video/"
	audioDir := tempDir + "audio/"
	encodeDir := tempDir + "encode/"
	mpdDir := tempDir + "mpd/"
	audioExt := ".mp4"
	duration := "10"


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")

	sPam := ffmpeg.FFMPEGParam{
		InputVideo: 		srcFile,
		SegmentDuration: 	duration,
		SegmentExt: 		".mp4",
		SegmentName: 		"",
		OutputLocation: 	segDir,
		FrameRate: 		"30",
		Debug: 			false,
		SegmentList: 		segDir + "list.txt",
	}


	names, err := ffmpeg.Split(sPam)
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

		ffmpeg.ProduceDashif("3000", mpdDir, "kd2.mpd",  videoFile, audioFile)
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

func ProduceVODDash() error{
	//My experiment to produce dash-if

	processStart := time.Now()

	srcFile := "videos/quangha.mp4"
	tempDir := "factory/"
	segDir := tempDir + "segments/"
	videoDir := tempDir + "video/"
	audioDir := tempDir + "audio/"
	encodeDir := tempDir + "encode/"
	mpdDir := tempDir + "mpd/"
	duration := "10"


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	sPam := ffmpeg.FFMPEGParam{
		InputVideo: 		srcFile,
		SegmentDuration: 	duration,
		SegmentExt: 		".mp4",
		SegmentName: 		"sample",
		OutputLocation: 	segDir,
		FrameRate: 			"30",
		Debug: 				false,
		SegmentList: 		segDir + "list.txt",
	}


	names, err := ffmpeg.Split(sPam)
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

		err = ffmpeg.ExtractAV(filename, videoDir + name + ".mp4" , audioDir + name + ".mp4")
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
		videoFile := encodeDir + nameElement.Value.(string)
		audioFile := audioDir + nameElement.Value.(string)
		ffmpeg.ProduceDashif("3000", mpdDir, "seek.mpd",  videoFile, audioFile)
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

func ConvertAAC2MP4(dir string)  {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		aac := dir + file.Name()
		comps := strings.Split(aac, ".")
		mp4 := comps[0] + ".mp4"
		ffmpeg.Aac2mp4(aac, mp4)
	}

}

func TranscodeAndSplit()  {
	tempDir := "avisplit/"
	srcFile := "videos/avi/s02.avi"
	mp4filename := tempDir + "s02.mp4"
	dur, _ := ffmpeg.GetAviDuration(srcFile)
	start := time.Now()
	trxn := time.Now()

	tcmd := config.FFMPEG + " -i " + srcFile + " -c:v libx264 -c:a aac -y " + mp4filename
	log.Println("cmd:" ,tcmd)
	err := exec.Command("bash", "-c", tcmd).Run()
	if err != nil {
		log.Println("--Tran failure!, Id: ",  err.Error())
	}
	log.Println("--Tran completed ")
	trxnDur :=  time.Since(trxn)

	splitStart := time.Now()
	scmd := config.FFMPEG + " -i " + mp4filename + " -c copy -map 0 -segment_time 3 -f segment " + tempDir + "%04d.mp4"
	err = exec.Command("bash", "-c", scmd).Run()
	if err != nil {
		log.Println("--Split failure!, Id: ",  err.Error())
	}
	log.Println("--Split completed ")
	splitDur := time.Since(splitStart)

	end := time.Since(start)

	log.Println("File Size : ", tools.GetFileSize(srcFile))
	log.Println("Trxn time: ", trxnDur)
	log.Println("Split time: ", splitDur)
	log.Println("Total time: ", end)
	log.Println("Video Duration  : ", dur)

	//Result
	//File Size :  59206048
	//Trxn time:  56.3484465s
	//Split time:  544.1756ms
	//Total time:  56.8913655s
	//Duration  :  0:06:29.522856
	//Segment: 125
	//Quality: better than segment & transcode at the same time
	//Output size: 41.6M


}

func SplitBenchmark()  {
	//logifle, _ :=os.Create("log.txt")
	//log.SetOutput(logifle)

	tempDir := "avisplit/"
	filename := "videos/avi/s02.avi"
	//filename := "C:/Users/nghiepnds/Desktop/videos"
	dur, _ := ffmpeg.GetAviDuration(filename)
	start := time.Now()
	names, _ := ffmpeg.SplitAvi(filename, 3 , tempDir, "s02-", ".mp4", false)
	end := time.Since(start)

	log.Println("File Size : ", tools.GetFileSize(filename))
	log.Println("Split time: ", end)
	log.Println("Segments  : ", names.Len())
	log.Println("Duration  : ", dur)

	//File Size :  59206048
	//Total time:  1m36.3981337s
	//Segments  :  130
	//Duration  :  0:06:29.522856
	//Audio play early in the first segment
	//Slower than transcode & split
	//Output size: 42.6M
}

func produceDashLive(){
}

func ProduceDashSegment(input string, outTran string, outDir string, mpdFile string) error {
	//Transcode with fix frame rate
	//	cmd := fmt.Sprintln(config.FFMPEG + " -report -i " + input +
	//	" -y -c:v libx264 -preset ultrafast -crf 32 -threads 0 -c:a aac -strict -2 -b:v 64k -bufsize 64k -r 30 -ar 44100 ") + outDir + outTran

	cmd := fmt.Sprintln(config.FFMPEG + " -i " + input +
	" -y -c:v libx264 -preset ultrafast -crf 32 -threads 0 -c:a aac -strict -2 -b:v 64k -bufsize 64k -r 30 -ar 44100 " + outDir + outTran)

	log.Println(cmd)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("-- trans failed, ", err)
		return err
	}

	//Dash the file
	cmdB := fmt.Sprintln(config.MP4BOX + " -dash 3000 -frag 3000 -rap -segment-name s_ " + outDir +  outTran)
	log.Println(cmdB)
	err = exec.Command("bash", "-c", cmdB).Run()
	if err != nil {
		log.Println("-- dashed failed, ", err)
		return err
	}
	return nil
}

func CatAndFrag()  {
	//Concat segment File
	ffmpeg.MP4BoxConcat("trim/", "cat/", "cat.mp4" )

	//Frag file
	ffmpeg.DashPackage("cat/cat.mp4", "cat/frag.mp4")
}

func RemoveMoovs(dir string, fragDir string, nomoovDir string){
	//Read the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	//Create output folder
	err = tools.CreateDir(fragDir)
	err = tools.CreateDir(nomoovDir)
	if err != nil {
		log.Fatal(err)
	}
	//Write the file names to the file
	segId := 0
	for _, file := range files {
		names := strings.Split(file.Name(),".")
		name := names[0]
		segIdStr := strconv.Itoa(segId)
		segId++

		//Dash Fragment
		input:= dir +  file.Name()
		output:= fragDir + file.Name()
		ffmpeg.DashPackageWithSequence(input, output, "onDemand", segIdStr)

		//Remove atoms
		input = fragDir + name + "_dashinit.mp4"
		output = nomoovDir + file.Name()
		ffmpeg.RemoveMp4Moov(input, output)
	}

}

func ExtractElementaryStream()  {
	input := "videos/sample.mp4"
	output := "video.h264"
	err := ffmpeg.ExtractElementaryStream(input, output)
	if err != nil{
		log.Println(err.Error())
	}else{
		log.Println("finished")
	}
}

func ProduceVideos(input string)  {
	//Test Result: bigger file size-> processing time increases exponentially
	// 4.3m - 1.6GB
	//730ms - 1MB
	//7.8s  - 5MB

	tools.CreateDir("videos/produced/")

	//Create video at different resolutions
	proStart := time.Now()

	fastStart := time.Now()
	ffmpeg.Transcode(input, "videos/produced/", "high.mp4", "ultrafast", "hq")
	fastTime := time.Since(fastStart)

	mediumStart := time.Now()
	ffmpeg.Transcode(input, "videos/produced/", "medium.mp4", "ultrafast", "mq")
	mediumTime := time.Since(mediumStart)

	slowStart := time.Now()
	ffmpeg.Transcode(input, "videos/produced/", "low.mp4", "ultrafast", "lq")
	slowTime := time.Since(slowStart)

	proTime := time.Since(proStart)

	log.Println("PROCESSING TIME - ", proTime)
	log.Println("--------------------------------")
	log.Println("Fast Time: ", fastTime)
	log.Println("Medium Time: ", mediumTime)
	log.Println("Slow Time: ", slowTime)
}

func SplitExtractEncode() error{
	//This will split upload file ->  extract seg to a/v -> encode non-audio segs
	//To compare processing time between this method and Split-Encode-Extract

	processStart := time.Now()

	srcFile := "videos/tth.mp4"
	tempDir := "videos/temp/"
	segDir := tempDir + "segments/"
	videoDir := tempDir + "video/"
	audioDir := tempDir + "audio/"
	encodeDir := tempDir + "encode/"
	duration := "10"


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	sPam := ffmpeg.FFMPEGParam{
		InputVideo: 		srcFile,
		SegmentDuration: 	duration,
		SegmentExt: 		".mp4",
		SegmentName: 		"sample",
		OutputLocation: 	segDir,
		FrameRate: 			"30",
		Debug: 				false,
		SegmentList: 		segDir + "list.txt",
	}

	names, err := ffmpeg.Split(sPam)

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

		err = ffmpeg.ExtractAV(filename, videoDir + name + ".mp4" , audioDir + name + ".mp3")
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
	processTime := time.Since(processStart)

	log.Println("PROCESSING TIME - ", processTime)
	log.Println("--------------------------------")
	log.Println("Split Time: ", splitTime)
	log.Println("Extract Time: ", extractTime)
	log.Println("Encode Time: ", encodeTime)

	return nil
}

func SplitAndExtract()  error{
	processStart := time.Now()

	srcFile := "videos/tth.mp4"
	tempDir := "videos/temp/"
	segDir := tempDir + "segments/"
	videoDir := tempDir + "video/"
	audioDir := tempDir + "audio/"
	duration := "10"


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	sPam := ffmpeg.FFMPEGParam{
		InputVideo: 		srcFile,
		SegmentDuration: 	duration,
		SegmentExt: 		".mp4",
		SegmentName: 		"sample",
		OutputLocation: 	segDir,
		FrameRate: 			"30",
		Debug: 				false,
		SegmentList: 		segDir + "list.txt",
	}

	names, err := ffmpeg.Split(sPam)
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

		err = ffmpeg.ExtractAV(filename, videoDir + name + ".mp4" , audioDir + name + ".aac")
		if err != nil{
			return err
		}
		nameElement = nameElement.Next()
	}
	extractTime :=  time.Since(extractStart)


	processTime := time.Since(processStart)

	log.Println("PROCESSING TIME - ", processTime)
	log.Println("--------------------")
	log.Println("Split Time: ", splitTime)
	log.Println("Extract Time: ", extractTime)


	return nil

}

func SplitEncodeExtract()  error{
	//This will split upload file -> encode the segments -> extract seg to a/v
	//To compare processing time between this method and Split-Extract-Encode
	processStart := time.Now()

	srcFile := "videos/tth.mp4"
	tempDir := "videos/temp/"
	segDir := tempDir + "segments/"
	videoDir := tempDir + "video/"
	audioDir := tempDir + "audio/"
	encodeDir := tempDir + "encode/"
	duration := "10"


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	sPam := ffmpeg.FFMPEGParam{
		InputVideo: 		srcFile,
		SegmentDuration: 	duration,
		SegmentExt: 		".mp4",
		SegmentName: 		"sample",
		OutputLocation: 	segDir,
		FrameRate: 			"30",
		Debug: 				false,
		SegmentList: 		segDir + "list.txt",
	}

	names, err := ffmpeg.Split(sPam)
	splitTime := time.Since(splitStart)

	//Encode Video
	//Encode videos
	encodeStart := time.Now()
	tools.CreateDir(encodeDir)
	nameElement := names.Front()
	for (nameElement != nil){
		filename := segDir + nameElement.Value.(string)  // videos/temp/segments/sample-0000.mp4
		baseName := path.Base(filename)			 // sample-0000.mp4
		parts := strings.Split(baseName, ".")		 // [sample-0000, mp4]
		name := parts[0]				 // sample-0000

		//log.Println("File: ", filename)
		//log.Println("Name: ", name)

		err = ffmpeg.Transcode(filename, encodeDir , name +".mp4", "ultrafast", "mq")
		if err != nil{
			return err
		}
		nameElement = nameElement.Next()
	}
	encodeTime :=  time.Since(encodeStart)
	processTime := time.Since(processStart)



	//Extract Video
	extractStart := time.Now()
	nameElement = names.Front()
	tools.CreateDir(videoDir)
	tools.CreateDir(audioDir)
	for (nameElement != nil){
		filename := encodeDir + nameElement.Value.(string)  // videos/temp/segments/sample-0000.mp4
		baseName := path.Base(filename)			 // sample-0000.mp4
		parts := strings.Split(baseName, ".")		 // [sample-0000, mp4]
		name := parts[0]				 // sample-0000

		//log.Println("File: ", filename)
		//log.Println("Name: ", name)

		err = ffmpeg.ExtractAV(filename, videoDir + name + ".mp4" , audioDir + name + ".mp3")
		if err != nil{
			return err
		}
		nameElement = nameElement.Next()
	}
	extractTime :=  time.Since(extractStart)

	log.Println("PROCESSING TIME - ", processTime)
	log.Println("Encode the segment contains all audio")
	log.Println("--------------------")
	log.Println("Split Time: ", splitTime)
	log.Println("Encode Time: ", encodeTime)
	log.Println("Extract Time: ", extractTime)

	return nil

}

func Extract()  {
	input := "videos/sample.mp4"
	ffmpeg.ExtractAV(input, "noaudio.m4v", "novideo.mp3")
}

func ExtractAndMux()  {
	input := "videos/TTH.mp4"

	//Extract
	ffmpeg.ExtractAV(input, "video.mp4", "audio.aac")
	log.Println("extract complete")

	//Mux - not work with .h264 video
	//ffmpeg.Mux("noaudio.h264", "novideo.aac", "out.mp4")

	//work with .h264 video
	ffmpeg.MP4BoxMux("video.mp4", "audio.aac", "out.mp4")
}

func ExtractTranscodeMux()  {
	//This prove that transcode can be done on .h264
	input := "videos/TTH.mp4"

	//Extract
	ffmpeg.ExtractAV(input, "noaudio.h264", "novideo.aac")
	log.Println("extract complete")

	ffmpeg.Transcode("noaudio.h264", "", "noaudioTranscoded.h264", "ultrafast", "mq" )

	//ffmpeg does not work with .h264 video
	ffmpeg.MP4BoxMux("noaudioTranscoded.h264", "novideo.aac", "out.mp4")
}

func TestMux() {
	//video := "videos/noaudio.tmp"
	//audio := "videos/novideo.mp3"
	//output := "videos/muxed.mp4"

	video := "tmp/noaudio.tmp"
	audio := "tmp/novideo.mp3"
	output := "tmp/muxed.mp4"
	//Check existence of the file
	if _, err := os.Stat(video); os.IsNotExist(err) {
		log.Printf("-- file %s not exist", video)
	}else{
		log.Printf("-- file %s exist", video)
		ffmpeg.Mux(video, audio, output)
	}


}

func TestDemux() {
	video := "videos/muxed.mp4"
	audio := "videos/novideo2.mp3"
	output := "videos/muxed2.mp4"
	ffmpeg.Demux(video, audio, output)
}

func TestSplitAndMerge()  {
	srcFile := "videos/quangha.mp4"
	namesFile := "list.txt"
	tempDir := "videos/temp/"
	trcdDir := "videos/trcd/"
	outDir := "videos/out/"
	outFile := "final.mp4"
	duration := "10"
	//Split file
	log.Println("start splitting file")
	sPam := ffmpeg.FFMPEGParam{
		InputVideo: 		srcFile,
		SegmentDuration: 	duration,
		SegmentExt: 		".mp4",
		SegmentName: 		"sample",
		OutputLocation: 	tempDir,
		FrameRate: 			"30",
		Debug: 				false,
		SegmentList: 		tempDir + "list.txt",
	}

	names, err := ffmpeg.Split(sPam)
	chunks := names.Len()
	if(err == nil){
		//Make sure trcd folder exist
		if _, err := os.Stat(trcdDir); os.IsNotExist(err) {
			err := os.MkdirAll(trcdDir,0711)
			if err != nil {
				log.Println("Error creating trcd directory")
				return
			}
		}

		//Convert all chunks
		log.Println("start converting chunks")
		var waitGroup sync.WaitGroup
		waitGroup.Add(chunks) //wait until job done

		//Create routine for each chunks
		for i:=1 ; i <= chunks; i++{
			nameElement := names.Front()
			names.Remove(nameElement)
			chunkFile := tempDir + nameElement.Value.(string)
			go func(srcFile string, i int) {
				log.Printf("Worker %d started \n", i)
				ffmpeg.Transcode(chunkFile, trcdDir, tools.ZeroPad(i,3) + ".mp4", "ultrafast", "mq")
				waitGroup.Done()
				log.Printf("Worker %d exit \n", i)
			}(chunkFile, i)
		}

		waitGroup.Wait()
		log.Println("Transcoding Completed")

		//Merge file
		log.Println("start merging Trunks")
		ffmpeg.CreateMergeList(trcdDir, namesFile)
		ffmpeg.Merge(namesFile, outDir, outFile)
		ffmpeg.GenerateJPGThumbnail("videos/", "tth.mp4")
		//deleteDir(tempDir)
		//deleteDir(trcdDir)

		//Upload to the AWS S3
		//Upload(outFile)
		//deleteDir(outDir)
	}
}

func TestThumbnailGenerator(){
	ffmpeg.GenerateJPGThumbnail("videos/", "tth.mp4")
}

func TestSplitAndMergeDistributedWorkers()  {
	srcFile := "videos/sampl.mp4"
	mergeList := "list.txt"
	tempDir := "videos/temp/"
	trcdDir := "videos/trcd/"
	outDir := "videos/out/"
	outFile := "final.mp4"
	duration := "180"
	//Split file
	log.Println("start splitting file")
	sPam := ffmpeg.FFMPEGParam{
		InputVideo: 		srcFile,
		SegmentDuration: 	duration,
		SegmentExt: 		".mp4",
		SegmentName: 		"sample",
		OutputLocation: 	tempDir,
		FrameRate: 			"30",
		Debug: 				false,
		SegmentList: 		tempDir + "list.txt",
	}

	names, err := ffmpeg.Split(sPam)
	chunks := names.Len()
	if(err == nil){
		//Make sure trcd folder exist
		if _, err := os.Stat(trcdDir); os.IsNotExist(err) {
			err := os.MkdirAll(trcdDir,0711)
			if err != nil {
				log.Println("Error creating trcd directory")
				return
			}
		}

		//Convert all chunks
		log.Println("start converting chunks")
		var waitGroup *sync.WaitGroup
		waitGroup.Add(chunks) //wait until job done
		maxRouties := 20
		input := make(chan *worker.Args)
		//Make worker ready
		for i:= 1; i <= maxRouties; i++{
			wk := worker.Worker{
				Id: i,
				InputChan: input,
				Wg:waitGroup,
			}
			go wk.Run()
		}

		//Route works into channel
		for i:=1 ; i <= chunks; i++{
			nameElement := names.Front()
			names.Remove(nameElement)
			chunkFile := tempDir + nameElement.Value.(string)

			work := worker.Args{
				SrcFile: chunkFile,
				DestDir: trcdDir,
				Name:"trd_"+ tools.ZeroPad(i,3),
			}

			input <- &work
		}

		waitGroup.Wait()
		//go monitorWorker(&waitGroup, input)

		//Merge file
		log.Println("start merging Trunks")
		ffmpeg.CreateMergeList(trcdDir, mergeList)
		ffmpeg.Merge(mergeList, outDir, outFile)
		//deleteDir(tempDir)
		//deleteDir(trcdDir)

		//Upload to the AWS S3
		//Upload(outFile)
		//deleteDir(outDir)
	}
}

func monitorWorker(wg *sync.WaitGroup, jc chan *worker.Args ) {
	wg.Wait()
	close(jc)
}

func TestMP4() {
	var inputFile string
	var f *mp4.File

	flag.StringVar(&inputFile, "i", "", "-i videos/sample.mp4")
	flag.Parse()

	if inputFile == "" {
		flag.Usage()
		return
	}
	f, err := mp4.Open(inputFile)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer f.Close()

}