package ffmpeg

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"container/list"
	"gomultimedia/config"
	"strings"
	"gomultimedia/tools"
	"io/ioutil"
	"path"

)

func init() {
	//createDirectories()
	//tools.CreateDir("tmp/transcode-buffer/")
}

// MY TRANSCODER
func Transcode(oldFile string, destDir string,  name string, compressLevel string, quantizer string) error {
	//Compress Level Preset
	//1. ultrafast			4. faster			7. slow
	//2. superfast			5. fast				8. slower
	//3. veryfast			6. medium			9. veryslow

	//Constant Rate Factor (CRF) - The range of the quantizer scale is 0-51
	// A lower value is a higher quality and a subjectively sane range is 18-28.
	//  0 - lossless
	// 23 - default
	// 51 - worst
	// input quantizer : hq(18), lq(28), md(23)
	var qStr string
	switch quantizer {
	case "hq":
		qStr = " -crf 18"
	case "lq":
		qStr = " -crf 32"
	default:
		qStr = " -crf 23"
	}

	//comment it out if you do not want to see the debug info
	var debug string
	//debug = " -report "

	//-crf 22
	/* --------------------------COMMAND SECTION -------------------------------------------------------- */
	//Cut2it define, using this, work properly
	cmd := fmt.Sprintln(config.FFMPEG + " -i " + oldFile +
	" -y -c:v libx264 -preset " + compressLevel + qStr + debug + " -threads 0 -c:a aac -strict -2 " + destDir + name)


	//cmd := fmt.Sprintln("ffmpeg -i " + oldFile + " -c:v libx264 -preset ultrafast -threads 0 " +
	//                    "-c:a aac -strict -2 /tmp/transcode-buffer/" + newName + ".mp4")

	//To set the video bitrate of the output file to 64 kbit/s:
	//cmd := fmt.Sprintln("ffmpeg -i " + oldFile + " -c:v 64k -bufsize 64k " + newName)

	//To force the frame rate of the output file to 24 fps:
	//cmd := fmt.Sprintln("ffmpeg -i " + oldFile + " -r 24 " + newName)

//	log.Println(qStr)
//	cmd := fmt.Sprintln(config.FFMPEG + debug + " -i " + oldFile +
//	" -c:a libfdk_aac -ac 2 -ab 128k -c:v libx264 -x264opts 'keyint=24:min-keyint=24:no-scenecut' -b:v 1500k -maxrate 1500k -bufsize 1000k -vf 'scale=-1:720' " + destDir + name)
	/* -------------------------------------------------------------------------------------------------- */
	log.Println("Command :" + cmd)
	err := exec.Command("bash", "-c", cmd).Run()

	if err != nil {
		log.Println("Transcoding Err: " + err.Error())
		return err
	} else {
		log.Println("Transcoding Completed: " + oldFile)
		return nil
	}
}

func SplitVideo(input string, seconds string, outputDir string) (*list.List, error){
	//MP4Box -split 3 videos/duck.mp4 -out factory/segments/%s.mp4
	log.Println("-- Splitting video...", input)

	//Make sure temp folder exist
	err := tools.CreateDir(outputDir)
	if(err != nil){
		return list.New(), err
	}

	cmd := config.MP4BOX + " -split " + seconds + " " + input + " -out " + outputDir
	log.Println("cmd:" ,cmd)
	err = exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("--Split failure!", err.Error())
		return list.New(), err
	} else {
		//log.Println("-- Split file completed")
		names:= tools.GetFileNames(outputDir)
		return names, nil
	}

}

// SPLIT ONE MP4 INTO MULTIPLE MP4s WITH SAME LENGTH (SECONDS), n: seconds
func Split(input string, seconds int, outputDir string, videoId string, ext string) (*list.List, error){
	log.Println("-- Splitting video...", input)

	//Make sure temp folder exist
	err := tools.CreateDir(outputDir)
	if(err != nil){
		return list.New(), err
	}

	/* --------------------------COMMAND SECTION -------------------------------------------------------- */
	// OPT 2 - Split input file into equally files with segment in seconds, -vcodec will allow split AVI
//	cmd := config.FFMPEG + " -report -i " + input + " -vcodec copy -map 0 -segment_time " + strconv.Itoa(seconds) +
//			" -f segment -strict -2 " + outputDir + videoId + "-%04d" + ext

//	cmd := config.FFMPEG + " -report -i " + input + " -c copy -map 0 -segment_time " + strconv.Itoa(seconds) +
//	" -f segment " + outputDir +  "%04d" + ext

//	cmd := config.FFMPEG + " -report -i " + input + " -acodec copy -f segment -segment_time " + strconv.Itoa(seconds) +
//	" -vcodec copy -reset_timestamps 1 -map 0 -an " + outputDir +  "%04d" + ext

	//ffmpeg -i fff.avi -acodec copy -f segment -segment_time 10 -vcodec copy -reset_timestamps 1 -map 0 -an fff%d.avi

	//Using this, split without reencode
	//cmd := config.FFMPEG + " -i " + input + " -acodec copy -vcodec copy  -segment_time " + strconv.Itoa(seconds) +
	//	" -f segment " + outputDir +  "%04d" + ext

	//This split script produce the list.txt contains all the file name
	//ffmpeg -i videos/asia-copy.mp4 -acodec copy -vcodec copy -segment_list list.txt -force_key_frames  expr:gte(t,n_forced*3) -segment_time 3 -f segment factory/segments/%04d.mp4

	cmd := config.FFMPEG + " -i " + input + " -acodec copy -vcodec copy  -segment_time " + strconv.Itoa(seconds) +
	" -f segment -r 30 " + outputDir +  "%04d" + ext


	/* -------------------------------------------------------------------------------------------------- */
	log.Println("cmd:" ,cmd)
	err = exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("--Split failure!", err.Error())
		return list.New(), err
	} else {
		log.Println("-- Split file completed")
		names:= tools.GetFileNames(outputDir)
		return names, nil
	}
}


// MERGE MP4s FILES INTO ONE MP4
func Merge(namesFile string, outputFolder string, outputFile string) error{
	log.Println("-- merging segment started")
	//Make sure the output folder exist
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		err := os.MkdirAll(outputFolder,0711)
		if err != nil {
			log.Println("Error creating output directory")
			return err
		}
		log.Println("-- output folder is ready" )
	}

	//-y : overwrite the output
	fileName := outputFolder + outputFile
	//log.Println("-- output file name: " , fileName)

	//log.Println("file name: " + fileName)
	cmd := config.FFMPEG + " -f concat -i " + namesFile + " -c copy -y " + fileName
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("-- merging segment failure!")
		return err
	} else {
		log.Println("-- merging segment completed")
		return nil
	}
}

// GET SEGMENT DURATION - Return time in seconds
func GetSegmentDuration(segmentFilePath string) (string, error) {

	cmdStr := config.FFMPEG + " -i '" + segmentFilePath +
	"' 2>&1 | grep \"Duration\" | cut -d ' ' -f 4 | sed s/,// | awk '{ split($1, A, \":\"); print A[3]; }'"

	ffmpegCmd := exec.Command("sh", "-c", cmdStr)
	ffmpegCmdOutput, cmdErr := ffmpegCmd.CombinedOutput();

	if cmdErr != nil {
		return "", cmdErr
	}

	//Trim trailing \n character
	if(len(ffmpegCmdOutput) > 0){
		ffmpegCmdOutput = ffmpegCmdOutput[:len(ffmpegCmdOutput)-1]
		return string(ffmpegCmdOutput), nil
	}
	return "0", nil

}

// GET CODEC
func GetCodec(file string) (string, string, error) {
	cmd := fmt.Sprintln(config.FFPROBE + " -v error -show_entries stream=codec_name -of default=noprint_wrappers=1:nokey=1 " + file)

	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", "", err
	} else {
		cmdReturn := string(out[:len(out)-1])

		codecSlice := strings.Split(cmdReturn, "\n")
		var aCodec string
		vCodec := codecSlice[0]
		if len(codecSlice) == 2 {
			aCodec = codecSlice[1]
		}

		return vCodec, aCodec, nil
	}
}

// GET WIDTH & HEIGH
func GetResolution(file string) (int, int, error) {
	cmd := fmt.Sprintln(config.FFPROBE + " -v error -show_entries stream=width,height -of default=noprint_wrappers=1:nokey=1 " + file)

	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()

	size := len(out)
	log.Println("Get Width And Height")
	//spew.Dump(out)

	if err != nil && size == 10 {
		return 0, 0, err
	} else {

		widthStr := string(out[:4])
		log.Printf("Width Raw: %s \n", widthStr)
		width, _ := strconv.Atoi(widthStr)
		log.Println("Width String")


		heightStr := string(out[5:9])
		log.Printf("Height Raw : %s \n", heightStr)
		height, _ := strconv.Atoi(heightStr)
		log.Printf("Height : %d \n", height)


		if height == 0 || width == 0 {
			height = 1080
			width = 1960
		}

		return width, height, nil
	}
}

// GET SEGMENT DURATION v2 - Return time in hh:mm:ss.ssss
func GetDuration(file string) (string, error) {
	cmd := fmt.Sprintln(config.FFPROBE + " -v error -show_entries format=duration" +
								" -of default=noprint_wrappers=1:nokey=1 -sexagesimal " + file)

	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()

	if err != nil {
		return "", err
	} else {
		log.Printf("Duration Raw: %s", out)
		log.Println("Duration Parsed")
		temp := string(out[:len(out)-8])
		log.Printf("Duration: %s", temp)
		return string(out[:len(out)-8]), nil
	}

}

// GET SEGMENT DURATION v2 - Return time in hh:mm:ss.ssss
func GetAviDuration(file string) (string, error) {
	cmd := fmt.Sprintln(config.FFPROBE + " -v error -loglevel quiet -show_entries format=duration" +
	" -of default=noprint_wrappers=1:nokey=1 -sexagesimal " + file)

	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()

	if err != nil {
		log.Println(err.Error())
		return "", err
	} else {
		return string(out), nil
	}
}

func GetDurationInMillisecond(file string) (string) {
	defaultTime := "3000"
	cmd := fmt.Sprintln(config.FFPROBE + " -v error -loglevel quiet -show_entries format=duration" +
	" -of default=noprint_wrappers=1:nokey=1 -sexagesimal " + file)

	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()

	if err != nil {
		log.Println(err.Error())
		return defaultTime
	}

	duration := string(out)   //hh:mm:ss.mmmm
	durs := strings.Split(duration, ".")
	timepot := strings.Split(durs[0], ":")

	//invalid data
	if(len(timepot)<3){
		log.Println("len:", len(timepot))
		return defaultTime
	}

	hh := timepot[0]
	mm := timepot[1]
	ss := timepot[2]

	hv, _ := strconv.Atoi(hh)
	mv, _ := strconv.Atoi(mm)
	sv, _ := strconv.Atoi(ss)

	val := (hv*3600 + mv*60 + sv)*1000
	retval := strconv.Itoa(val)
	return retval
}

func SplitAvi(input string, seconds int, outputDir string, name string, ext string, debug bool)  (*list.List, error){
	//This slower than transcode & split
	log.Println("-- Splitting AVI video...", input)
	var report string
	if(debug){
		report = " -report "
	}else{
		report = ""
	}
	//Make sure temp folder exist
	err := tools.CreateDir(outputDir)
	if(err != nil){
		return list.New(), err
	}

	//Get duration
	names := list.New()
	segmentId := 1
	dur, err := GetAviDuration(input)					//hh:mm:ss.ms
	if(err != nil){
		names.PushBack(name + tools.ZeroPad(segmentId, 4) + ext)
		return names, err
	}
	hms := strings.Split(dur, ".")						//hh:mm:ss
	durInSeconds := tools.TimeStampToSeconds(hms[0])		//int, ex 35
	log.Println("Total seconds:", durInSeconds)
	log.Println("Split...")
	for i := 0; i < durInSeconds; i += seconds {
		ss := " -ss " + GetTimeStamp(i)
		tt := " -t " + strconv.Itoa(seconds)
		//Split file, this is fast seeker but a litter not exact as the slow seeker
		cmd := config.FFMPEG + report  + ss +  tt +  " -i " + input + " -c:v libx264 -c:a aac -y " + outputDir + name + tools.ZeroPad(segmentId, 4) + ext

		//Slow seeker, very slow, consume double time
		//cmd := config.FFMPEG + report  + ss + " " + tt +  " -i " + input + " -acodec copy -vcodec copy " + " -y " + outputDir + name + tools.ZeroPad(segmentId, 4) + ext
		log.Println("cmd:" ,cmd)
		err = exec.Command("bash", "-c", cmd).Run()
		if err != nil {
			log.Println("--Split failure!, Id: ", segmentId, err.Error())
		} else {
			names.PushBack(name + tools.ZeroPad(segmentId, 4) + ext)
			segmentId++
		}
	}

	return names, nil
}

func GetTimeStamp(durationInSeconds int) string  {
	hours	:= durationInSeconds/3600
	minutes := (durationInSeconds-hours*3600)/60
	seconds := (durationInSeconds-hours*3600 -minutes*60)

	hh:= tools.ZeroPad(hours, 2)
	mm:= tools.ZeroPad(minutes,2)
	ss:= tools.ZeroPad(seconds,2)

	return (hh + ":" + mm +":" + ss)
}

// CREATE MERGE LIST INTO A TEXT FILE, Using for Merging Video Files
func CreateMergeList(dir string, nameFile string) int{
	//Create file
	txtFile, err := os.Create(nameFile)
	if err != nil {
		log.Fatal(err)
	}
	defer txtFile.Close()

	//Read the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	//Write the file names to the file
	count := 0;
	for _, file := range files {
		//fmt.Println(file.Name())
		name:= "file '" + dir +  file.Name() + "'\n"
		txtFile.WriteString(name)
		count++
	}

	return count
}

// ENCODE VIDEO TO OGG FORMAT
func EncodeOGG(fileLoc string, fileName string) (fileNameOut string, err error) {
	log.Println("-- start encoding OGG...")
	fileSource := fileLoc + fileName
	// fileNameOut = m.returnBaseFilename(fileName) + ".ogg"
	fileNameOut = "encoded.ogg"
	fileDestination := fileLoc + fileNameOut

	maxWidth := "1280"
	maxHeight := "720"

	out, err := exec.Command(config.THEORA, fileSource, "-o", fileDestination, "--two pass", "--videobitrate",
		"900", "-x", maxWidth, "-y", maxHeight).Output()
	if err != nil {
		log.Println("-- encode to OGG failed:", err)
		return "", err
	}


	log.Printf("output:(%s), err(%v)\n", string(out), err)
	log.Println("-- complete encoding OGG...")
	return fileNameOut, nil
}

// ENCODE VIDEO TO WEBM FORMAT
func EncodeWEBM(fileLoc string, fileName string) (fileNameOut string, err error) {
	log.Println("-- start encoding WEBM...")
	fileSource := fileLoc + fileName
	fileNameOut = "encoded.webm"
	fileDestination := fileLoc + fileNameOut

	out, err := exec.Command(config.FFMPEG, "-i", fileSource, "-pass", "1", "-passlogfile", fileDestination, "-keyint_min",
		"0", "-g", "250", "-skip_threshold", "0", "-vcodec", "libvpx", "-b", "600k", "-s", "1280x720", "-aspect",
		"16:9", "-an", "-y", fileDestination).Output()


	if err != nil {
		log.Println("-- encode video to WEBM failed:", err)
		return "", err
	}

	out, err = exec.Command(config.FFMPEG, "-i", fileSource, "-pass", "2", "-passlogfile", fileDestination, "-keyint_min",
		"0", "-g", "250", "-skip_threshold", "0", "-vcodec", "libvpx", "-b", "600k", "-s", "1280x720", "-aspect",
		"16:9", "-acodec", "libvorbis", "-y", fileDestination).Output()

	if err != nil {
		log.Println("-- encode video to WEBM failed:", err)
		return "", err
	}


	log.Printf("output:(%s), err(%v)\n", string(out), err)
	log.Println("-- complete encoding WEBM...")
	return fileNameOut, nil
}

// GENERATE THUMBNAIL
func GenerateJPGThumbnail(fileLoc string, fileName string) (fileNameOut string, err error) {
	log.Println("-- start generating thumbnail...")
	fileSource := fileLoc + fileName
	fileNameOut = "encoded.jpg"
	fileDestination := fileLoc + fileNameOut

	//[]byte, err
	_, err = exec.Command("ffmpeg", "-i", fileSource, "-t", "0.001", "-ss", "7", "-vframes", "1", "-y", "-f",
		"mjpeg", fileDestination).Output()
	if err != nil {
		log.Println("-- create thumnail failed:", err)
		return "", err
	}

	log.Println("-- thumnail created.")
	return fileNameOut, nil
}

// MUX AUDIO & VIDEO INTO MP4
// Input is mp4 & .wav
func Mux(video string, audio string, newfile string)  error{
	//log.Println("video:", tools.IsExist(video))
	//log.Println("audio:", tools.IsExist(audio))
	dir, file := path.Split(video)
	log.Println(video)
	log.Println(audio)
	log.Println(dir)
	log.Println(file)

//	log.Println(tools.IsExist(video))
//	err := Transcode(video, dir, "videot.mp4", "ultrafast", "mq")
//	if err != nil {
//		log.Println("-- transcode failed, ", err)
//		return err
//	} else {
//		log.Println("-- transcode completed.")
//	}
//	log.Println(err)
//	srcVStream := dir + "videot" +  ".mp4"
//	log.Println(srcVStream)


	//Mux h264 & aac with FFMPEG is not successful, log said the h264 is not valid
//	cmd := fmt.Sprintln(config.FFMPEG + " -report  -i " + video + " -i " + audio +
//	" -vcodec copy -acodec copy -absf aac_adtstoasc  -y " + newfile)

	//This give similiar result as above, "noaudio.h264: Invalid data found when processing input"
//	cmd := fmt.Sprintln(config.FFMPEG + " -report  -i " + video + " -i " + audio +
//	" -c:v copy -c:a aac -strict experimental -y " + newfile)

	//Mux using MP4 Box, not working "cannot find H264 start code"
	//cmd := fmt.Sprintln(config.MP4BOX + " -fps 23.976 -add " + video + " -add " + audio + " " + newfile)

	//result: "noaudio.h264: Invalid data found when processing input"
	cmd := fmt.Sprintln(config.FFMPEG + " -framerate 25  -report  -i " + video + " -i " + audio +
	" -codec copy -y " + newfile)

	log.Println(cmd)
	errs := exec.Command("bash", "-c", cmd).Run()
	if errs != nil {
		log.Println("-- muxing was failed, ", errs)
		return errs
	} else {
		log.Println("-- muxing completed.")
		return nil
	}


}

// DEMUX AUDIO from VIDEO
// Replace sound track in the video with an input audio
func Demux(video string, audio string, newfile string)  error{
	cmd := fmt.Sprintln(config.FFMPEG + " -i " + video + " -i " + audio +
	" -c:v copy -c:a aac -strict experimental -map 0:v:0 -map 1:a:0 -y " + newfile)

	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("-- demux failed, ", err)
		return err
	} else {
		log.Println("-- demux completed.")
		return nil
	}
}

func ExtractAV(input string, vOuput string, aOutput string)  error{
	//extractStart := time.Now()
	//vCmd := fmt.Sprintln(config.FFMPEG + " -i " + input + " -c:v copy -an -y " + vOuput)
	vCmd := fmt.Sprintln(config.FFMPEG + " -i " + input + " -vcodec copy -an -bsf:v h264_mp4toannexb " + vOuput)
	//log.Println("vCmd:" , vCmd)
	//aCmd := fmt.Sprintln(config.FFMPEG + " -i " + input + " -c:v copy -vn -y " + aOutput)
	//aCmd := fmt.Sprintln(config.FFMPEG + " -i " + input + " -c copy -map 0:a:0 -y " + aOutput)

	aCmd := fmt.Sprintln(config.FFMPEG + " -i " + input + " -vn -acodec copy -bsf:a aac_adtstoasc " + aOutput)

	//For DTS Audio
	//aCmd := fmt.Sprintln(config.FFMPEG + " -i " + input + " -vn -c:a aac -b:a 128k " + aOutput)

	log.Println("aCmd:" , aCmd)
	err := exec.Command("bash", "-c", vCmd).Run()
	if err != nil {
		log.Println("-- extract video failed, ", err)
		return err
	}

	err = exec.Command("bash", "-c", aCmd).Run()
	if err != nil {
		log.Println("-- extract audio failed, ", err)
		return err
	}

	//extractTime :=  time.Since(extractStart)
	//log.Println("Processing time: ", extractTime)
	return nil
}

func MP4BoxExtractAudio(input string, vOuput string, aOutput string)  error{
	//extractStart := time.Now()

	//vCmd := fmt.Sprintln(config.FFMPEG + " -i " + input + " -c:v copy -an -y " + vOuput)
	//aCmd := fmt.Sprintln(config.FFMPEG + " -i " + input + " -c:v copy -vn -y " + aOutput)
	aCmd := fmt.Sprintln("mp4box -raw 2 " + input + " -out " + aOutput)



//	err := exec.Command("bash", "-c", vCmd).Run()
//	if err != nil {
//		log.Println("-- extract video failed, ", err)
//		return err
//	}

	err := exec.Command("bash", "-c", aCmd).Run()
	if err != nil {
		log.Println("-- extract audio failed, ", err)
		return err
	}

	//extractTime :=  time.Since(extractStart)
	//log.Println("Processing time: ", extractTime)
	return nil
}

func MP4BoxMux(video string, audio string, newfile string)error{
	//Mux using MP4 Box, not working "cannot find H264 start code"
	//cmd := fmt.Sprintln(config.MP4BOX + " -fps 23.976 -add " + video + " -add " + audio + " " + newfile)
	cmd := fmt.Sprintln("mp4box -add " + video + " -add " + audio + " " + newfile)

	//result: "noaudio.h264: Invalid data found when processing input"
	//cmd := fmt.Sprintln(config.FFMPEG + " -framerate 25  -report  -i " + video + " -i " + audio +
	//" -codec copy -y " + newfile)

	log.Println(cmd)
	errs := exec.Command("bash", "-c", cmd).Run()
	if errs != nil {
		log.Println("-- muxing was failed, ", errs)
		return errs
	} else {
		log.Println("-- muxing completed.")
		return nil
	}
}


func ExtractElementaryStream(input string, vOuput string) error{
	//Extract the raw video codec data as it is.
	//The extracted elementary streams are lacking the Video Object Layer (VOL) and the upper layers.
	// An extracted elementary stream by FFmpeg contains just sequence of Video Object Plane (VOP).
	vCmd := fmt.Sprintln(config.FFMPEG + " -i " + input + " -vcodec copy -an -y -f m4v " + vOuput)
	log.Println(vCmd)
	err := exec.Command("bash", "-c", vCmd).Run()
	if err != nil {
		log.Println("-- extract video failed, ", err)
		return err
	}
	return nil
}

func DashPackage(input string, output string)error{
	// -dash [DURATION]: enables MPEG-DASH segmentation, creating segments of the given duration (in milliseconds).
	// We advise you to set the duration to 2 seconds for Live and short VOD files, and 5 seconds for long VOD videos.
	// -rap -frag-rap: forces segments to begin with Random Access Points. Mandatory to have a working playback.
	// –profile [PROFILE]: MPEG-DASH profile. Set it to 'onDemand' for VOD videos, and 'live' for live streams.
	// -out [path/to/outpout.file]: output file location. This parameter is optional: by default, MP4box will create an
	// output.mpd file and the corresponding output.mp4 files in the current directory.
	// [path/to/input1.file]…: indicates where your input mp4 files are. They can be video or audio files.
	// -segment-name name
	// -moof-sn 2
	// profile dashavc264:onDemand or dashavc264:live
	cmd := fmt.Sprintln(config.MP4BOX + " -dash 3000 -rap -frag-rap  -profile dashavc264:onDemand -out " + output + " " + input)
	log.Println(cmd)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("-- dashed failed, ", err)
		return err
	}
	return nil
}

func SegmentDash(input string, output string)error{
	// -dash [DURATION]: enables MPEG-DASH segmentation, creating segments of the given duration (in milliseconds).
	// We advise you to set the duration to 2 seconds for Live and short VOD files, and 5 seconds for long VOD videos.
	// -rap -frag-rap: forces segments to begin with Random Access Points. Mandatory to have a working playback.
	// –profile [PROFILE]: MPEG-DASH profile. Set it to 'onDemand' for VOD videos, and 'live' for live streams.
	// -out [path/to/outpout.file]: output file location. This parameter is optional: by default, MP4box will create an
	// output.mpd file and the corresponding output.mp4 files in the current directory.
	// [path/to/input1.file]…: indicates where your input mp4 files are. They can be video or audio files.

	// Create dash segments from input file (input file is the long video)
	// Example: MP4Box -dash 10000 -frag 1000 -rap -segment-name myDash -subsegs-per-sidx 5 -url-template videos/TTH.mp4
	// will create myDash1-myDash25 + myDashInit + TTH_dash.mpd
	cmd := fmt.Sprintln(config.MP4BOX + " -dash 3000 " + " -segment-name " + output + " " + input)
	log.Println(cmd)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("-- dashed failed, ", err)
		return err
	}
	return nil
}

func DashPackageFFMPEG(input string, output string)error{
	// -dash [DURATION]: enables MPEG-DASH segmentation, creating segments of the given duration (in milliseconds).
	// We advise you to set the duration to 2 seconds for Live and short VOD files, and 5 seconds for long VOD videos.
	// -rap -frag-rap: forces segments to begin with Random Access Points. Mandatory to have a working playback.
	// –profile [PROFILE]: MPEG-DASH profile. Set it to 'onDemand' for VOD videos, and 'live' for live streams.
	// -out [path/to/outpout.file]: output file location. This parameter is optional: by default, MP4box will create an
	// output.mpd file and the corresponding output.mp4 files in the current directory.
	// [path/to/input1.file]…: indicates where your input mp4 files are. They can be video or audio files.
	cmd := fmt.Sprintln("ffmpeg -report -re -i "+ input + " -g 52 -acodec libvo_aacenc -ab 64k -vcodec libx264 -vb 448k -f mp4 -movflags frag_keyframe+empty_moov " + output)
	log.Println(cmd)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("-- dashed failed, ", err)
		return err
	}
	return nil
}

func DashPackageWithSequence(input string, output string, profile string, segmentId string)error{
	// -dash [DURATION]: enables MPEG-DASH segmentation, creating segments of the given duration (in milliseconds).
	// We advise you to set the duration to 2 seconds for Live and short VOD files, and 5 seconds for long VOD videos.
	// -rap -frag-rap: forces segments to begin with Random Access Points. Mandatory to have a working playback.
	// –profile [PROFILE]: MPEG-DASH profile. Set it to 'onDemand' for VOD videos, and 'live' for live streams.
	// -out [path/to/outpout.file]: output file location. This parameter is optional: by default, MP4box will create an
	// output.mpd file and the corresponding output.mp4 files in the current directory.
	// [path/to/input1.file]…: indicates where your input mp4 files are. They can be video or audio files.
	// -segment-name name
	// -moof-sn 2
	cmd := fmt.Sprintln(config.MP4BOX + " -dash 2000 -rap -frag-rap -moof-sn "+ segmentId + " -profile onDemand -out " + output + " " + input)
	//log.Println(cmd)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("-- dashed failed, ", err)
		return err
	}
	return nil
}

func RemoveMp4Moov(input string, output string) error {
	//Remove the moov & ftyp atom from the mp4 file
	//Limitation: only work with video frag file
	cmd := fmt.Sprintln(config.MP4EDIT + " --remove moov --remove ftyp --remove free --remove sidx " + input + " " + output)
	//log.Println(cmd)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("-- atom edit failed, ", err)
		return err
	}
	return nil
}

func MP4BoxConcat(dir string, outdir string, output string) error {
	//Read the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	//Create output folder
	err = tools.CreateDir(outdir)
	if err != nil {
		return err
	}

	//Write the file names to the file
	filenames := ""
	for _, file := range files {
		filenames = filenames + " -cat "+ dir + file.Name()
	}

	cmd := fmt.Sprintln(config.MP4BOX +  filenames + " " + outdir+ output)
	//log.Println(cmd)

	//log.Println(cmd)
	err = exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("-- Cat failed, ", err.Error())
		return err
	}
	return nil
}

func Aac2mp4(aac string, mp4 string)  error {
	cmd := fmt.Sprintln(config.MP4BOX +  " -add " + aac + " " + mp4)
	log.Println(cmd)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("-- Cat failed, ", err.Error())
		return err
	}
	return nil
}

// DASH-IF FROM VIDEO & AUDIO SEG - VOD
func ProduceDashif(duration string, mpdDir string, mpdName string, videoFile string, audioFile string )error{
	// STILL FOUND NO WAY TO PRODUCE DASH IF VIDEO FROM SEGMENTS, DASH IF FROM A MP4 IS OK
	// Profiles:
	// onDemand, main, simple, full, dashavc264:onDemand --> append all segments into init segment, seeker work
	// live , dashavc264:live--> create ms4 segments

	//---------------------AUDIO/VIDEO SEGS ARE SEPARATED----------------------------------------
	//Using profile main, URL in MPD is different from below - THIS APP DOES NOT PLAY ON DASH-IF
//	cmdV := fmt.Sprintln(config.MP4BOX +  " -dash " + strconv.Itoa(3000) +
//	" -mpd-refresh 3 -profile h264:live -rap -segment-ext mp4 -segment-name " + "v_" + " -dash-ctx " + mpdDir + "v-stream.txt"  +
//	" -out " + mpdDir + mpdName  +  " " + audioFile  + " " + videoFile)
//
//	cmdA := fmt.Sprintln(config.MP4BOX +  " -dash " + strconv.Itoa(3000) +
//	" -mpd-refresh 3 -profile h264:live -rap -segment-ext mp4 -segment-name " + "a_" + " -dash-ctx " + mpdDir + "a-stream.txt"  +
//	" -out " + mpdDir + mpdName  +  " " + videoFile  + " " + audioFile)
	//-----------------------------------------------------------------------------------------------------------------

	//-----------------------AUDIO/VIDEO SEGS APPEND TO AUDIO/VIDEO INIT SEG------------------------------------------------------------------------------------------
	//Using profile main, URL in MPD is different from below - THIS APP DOES NOT PLAY ON DASH-IF nor  OSMO4
	//cmdV := fmt.Sprintln(config.MP4BOX +  " -dash " + strconv.Itoa(3000) +
	//" -mpd-refresh 3 -profile full -rap -single-file -segment-name v_" + " -dash-ctx " + mpdDir + "v-stream.txt"  +
	//" -out " + mpdDir + mpdName  +  " " + audioFile  + " " + videoFile)

	//cmdA := fmt.Sprintln(config.MP4BOX +  " -dash " + strconv.Itoa(3000) +
	//" -mpd-refresh 3 -profile full -rap -single-file -segment-name a_" + " -dash-ctx " + mpdDir + "a-stream.txt"  +
	//" -out " + mpdDir + mpdName  +  " " + videoFile  + " " + audioFile)
	//-----------------------------------------------------------------------------------------------------------------

	//segment name change to -segment-name qh_ : PLAY VIDEO ONLY, need to make the audio and video segment has different name
	//cmd := fmt.Sprintln(config.MP4BOX +  " -dash " + strconv.Itoa(3000) +
	// " -mpd-refresh 3 -profile live -rap -segment-name qh_ -dash-ctx " + mpdDir + "stream.txt"  +
	//" -out " + mpdDir + mpdName  +  " " + audioFile  + " " + videoFile)

	//segment name change to -segment-name qh_ : PLAY AUDIO ONLY, need to make the audio and video segment has different name
	//cmd := fmt.Sprintln(config.MP4BOX +  " -dash " + strconv.Itoa(3000) +
	// " -mpd-refresh 3 -profile live -rap -segment-name qh_ -dash-ctx " + mpdDir + "stream.txt"  +
	//" -out " + mpdDir + mpdName  +  " " + videoFile  + " " + audioFile)

	//NON MULTIPLEX VIDEO, CAN PLAY BOTH OSMO4 & JW PLAYER. BUT NOT STILL THE END OF THE VIDEO
	MP4BoxAudioMux(videoFile, audioFile)
	//cmd := fmt.Sprintln(config.MP4BOX +  " -dash-ctx " + mpdDir + "stream.txt -dash " + strconv.Itoa(duration) +
	//" -dash-profile dashavc264:live -bs-switching no -rap -frag 3000 " + videoFile  + "#video "+
	//videoFile + "#audio -out " + mpdDir + mpdName)


	//Produce a list of URL instead of the URL template
	cmd := fmt.Sprintln(config.MP4BOX +  " -dash-ctx " + mpdDir + "stream.txt -dash " + duration +
	" -segment-timeline -profile live -bs-switching no -rap -frag " + duration + " " + videoFile  + "#video "+
	videoFile + "#audio -out " + mpdDir + mpdName)

	//minimumUpdatePeriod="PT0H0M2.000S"
	log.Println("cmd: ", cmd)
	//log.Println(cmd)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("Dash Err, " , err)
		return err
	} else {
		return nil
	}
}

// DASH-IF FROM MUX VIDEOSEG - VOD
func ProduceDashifFromMuxSeg(duration int, mpdDir string, mpdName string, videoFile string)error{
	//NON MULTIPLEX VIDEO, CAN PLAY BOTH OSMO4 & JW PLAYER STILL THE END OF THE VIDEO
	cmd := fmt.Sprintln(config.MP4BOX +  " -dash-ctx " + mpdDir + "stream.txt -dash " + strconv.Itoa(duration) +
	" -profile live -bs-switching no "  + " " + videoFile  + "#video " + videoFile + "#audio -out " + mpdDir + mpdName)

	//log.Println(cmd)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("Dash Err, " , err)
		return err
	} else {
		return nil
	}
}

// MULTIPLEX DASH FROM MUX SEGEMENT - VOD
func ProduceDash(duration int, mpdDir string, mpdName string, videoFile string, audioFile string, segmentName string)error{
	// Profiles: h264:live & h264 play the video to the end
	// onDemand, main, simple, full, dashavc264:onDemand --> append all segments into init segment, seeker work
	// live , dashavc264:live--> create ms4 segments

	//OPTION 1 - USING MUX SEGMENTS - MULTIPLEX OPTION I
	// RESULT: Dash segment are separated in files, play to the end of file  GRADE:*****
	MP4BoxAudioMux(videoFile, audioFile)
	cmd := fmt.Sprintln(config.MP4BOX +  " -dash-ctx " + mpdDir + segmentName + "stream.txt -dash " + strconv.Itoa(duration) +
	" -mpd-refresh 3 -profile h264:live -rap -frag 3000 -segment-name "+ segmentName  + " -out " + mpdDir + mpdName  +
	" " + videoFile)


	//OPTION 2 - USING MUX SEGMENTS - MULTIPLEX OPTION II
	// RESULT: Dash segments are created separately in files, NOT PLAY TO END
	//<<<<<<<<<<<<<<<<<Current OPT on MediaCluster>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//MP4BoxAudioMux(videoFile, audioFile)
	//cmd := fmt.Sprintln(config.MP4BOX +  " -dash-ctx " + mpdDir + "stream.txt -dash 3000  -mpd-refresh 3 -profile dashavc264:live -rap "  +
	//" -out " + mpdDir + mpdName  +  " -add " + videoFile  + "#video -add " + videoFile + "#audio -fps 30 seg" )

	//OPTION 3 - USING MUX SEGMENTS - MULTIPLEX OPTION III (Different profile from II cause the segment to merge)
	// RESULT: Dash segments are merged in the init segment, play better GRADE: *****
	//MP4BoxAudioMux(videoFile, audioFile)
	//cmd := fmt.Sprintln(config.MP4BOX +  " -dash-ctx " + mpdDir + "stream.txt -dash 3000  -mpd-refresh 3 -profile h264:live -rap "  +
	//" -out " + mpdDir + mpdName  +  " -add " + videoFile  + "#video -add " + videoFile + "#audio -fps 30 seg" )

	//OPTION 4 - USING NON-MUX SEGMENTS - MULTIPLEX OPTION IV
	// audio & video can be seperate (aac is accepted)
	// RESULT: Dash segments append into init segment - Seeker Support
	//cmd := fmt.Sprintln(config.MP4BOX +  " -dash-ctx " + mpdDir + "stream.txt -dash " + strconv.Itoa(duration) +
	//" -mpd-refresh 3 -profile h264:live -rap "  + " -out " + mpdDir + mpdName  +
	//" -add " + videoFile  + "#video -add " + audioFile + "#audio kd2"   )

	log.Println("Command :" , cmd)
	err := exec.Command("bash", "-c", cmd).Run()

	if err != nil {
		log.Println("Dash Err, " , err)
		return err
	} else {
		return nil
	}
}

// DASH-IF FROM VIDEO & AUDIO SEG - LIVE
func ProduceDashifLive(duration int, mpdDir string, mpdName string, videoFile string, audioFile string )error{
	// STILL FOUND NO WAY TO PRODUCE DASH IF VIDEO FROM SEGMENTS, DASH IF FROM A MP4 IS OK
	// Profiles:
	// onDemand, main, simple, full, dashavc264:onDemand --> append all segments into init segment, seeker work
	// live , dashavc264:live--> create ms4 segments

	//---------------------AUDIO/VIDEO SEGS ARE SEPARATED----------------------------------------
	//Using profile main, URL in MPD is different from below - THIS APP DOES NOT PLAY ON DASH-IF
	//	cmdV := fmt.Sprintln(config.MP4BOX +  " -dash " + strconv.Itoa(3000) +
	//	" -mpd-refresh 3 -profile h264:live -rap -segment-ext mp4 -segment-name " + "v_" + " -dash-ctx " + mpdDir + "v-stream.txt"  +
	//	" -out " + mpdDir + mpdName  +  " " + audioFile  + " " + videoFile)
	//
	//	cmdA := fmt.Sprintln(config.MP4BOX +  " -dash " + strconv.Itoa(3000) +
	//	" -mpd-refresh 3 -profile h264:live -rap -segment-ext mp4 -segment-name " + "a_" + " -dash-ctx " + mpdDir + "a-stream.txt"  +
	//	" -out " + mpdDir + mpdName  +  " " + videoFile  + " " + audioFile)
	//-----------------------------------------------------------------------------------------------------------------

	//-----------------------AUDIO/VIDEO SEGS APPEND TO AUDIO/VIDEO INIT SEG------------------------------------------------------------------------------------------
	//Using profile main, URL in MPD is different from below - THIS APP DOES NOT PLAY ON DASH-IF nor  OSMO4
	//cmdV := fmt.Sprintln(config.MP4BOX +  " -dash " + strconv.Itoa(3000) +
	//" -mpd-refresh 3 -profile full -rap -single-file -segment-name v_" + " -dash-ctx " + mpdDir + "v-stream.txt"  +
	//" -out " + mpdDir + mpdName  +  " " + audioFile  + " " + videoFile)

	//cmdA := fmt.Sprintln(config.MP4BOX +  " -dash " + strconv.Itoa(3000) +
	//" -mpd-refresh 3 -profile full -rap -single-file -segment-name a_" + " -dash-ctx " + mpdDir + "a-stream.txt"  +
	//" -out " + mpdDir + mpdName  +  " " + videoFile  + " " + audioFile)
	//-----------------------------------------------------------------------------------------------------------------

	//segment name change to -segment-name qh_ : PLAY VIDEO ONLY, need to make the audio and video segment has different name
	//cmd := fmt.Sprintln(config.MP4BOX +  " -dash " + strconv.Itoa(3000) +
	// " -mpd-refresh 3 -profile live -rap -segment-name qh_ -dash-ctx " + mpdDir + "stream.txt"  +
	//" -out " + mpdDir + mpdName  +  " " + audioFile  + " " + videoFile)

	//segment name change to -segment-name qh_ : PLAY AUDIO ONLY, need to make the audio and video segment has different name
	//cmd := fmt.Sprintln(config.MP4BOX +  " -dash " + strconv.Itoa(3000) +
	// " -mpd-refresh 3 -profile live -rap -segment-name qh_ -dash-ctx " + mpdDir + "stream.txt"  +
	//" -out " + mpdDir + mpdName  +  " " + videoFile  + " " + audioFile)

	//NON MULTIPLEX VIDEO, CAN PLAY BOTH OSMO4 & JW PLAYER STILL THE END OF THE VIDEO
	MP4BoxAudioMux(videoFile, audioFile)

	//cmd := fmt.Sprintln(config.MP4BOX +  " -dash-ctx " + mpdDir + "stream.txt -dash " + strconv.Itoa(duration) +
	//" -dash-profile dashavc264:live -bs-switching no "  + " " + videoFile  + "#video " + videoFile + "#audio -out " + mpdDir + mpdName)

	//USING THIS IN CASE ALL SEGMENTS HAVE SAME DURATION, ALL SEGMENTS KEPT
	//cmd := fmt.Sprintln(config.MP4BOX +  " -dash-ctx " + mpdDir + "stream.txt -dash " + strconv.Itoa(duration) +
	//" -dash-profile dashavc264:live -dynamic -time-shift -1 -fps 34 -mpd-refresh 5 -min-buffer 10000 -bs-switching inband  -rap -frag 3000 " +
	//videoFile  + "#video "+	videoFile + "#audio -out " + mpdDir + mpdName)

	//USING THIS SEGMENT TIMELINE FOR DEVIATION SEGMENTS, ALL SEGMENTS KEPT
	cmd := fmt.Sprintln(config.MP4BOX +  " -dash-ctx " + mpdDir + "stream.txt -dash " + strconv.Itoa(duration) +
	" -dash-profile dashavc264:live -dynamic -time-shift -1 -fps 34 -mpd-refresh 5 -min-buffer 3000 -bs-switching inband  -rap -frag 3000 " +
	videoFile  + "#video "+	videoFile + "#audio -out " + mpdDir + mpdName)

	//USING THIS IN CASE WANT TO HAVE SEGMENTS REPLACED
	//cmd := fmt.Sprintln(config.MP4BOX +  " -dash-ctx " + mpdDir + "stream.txt -dash " + strconv.Itoa(duration) +
	//" -dash-profile dashavc264:live -dynamic -fps 34 -mpd-refresh 1 -min-buffer 10000 -bs-switching inband  -rap -frag 3000 " +
	//videoFile  + "#video "+	videoFile + "#audio -out " + mpdDir + mpdName)

	log.Println("cmd: ", cmd)
	//log.Println(cmd)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("Dash Err, " , err)
		return err
	} else {
		return nil
	}
}

// MULTIPLEX DASH - LIVE
func ProduceDashLive(duration int, mpdDir string, mpdName string, videoFile string, audioFile string)error{
	// Profiles: h264:live & h264 play the video to the end
	// onDemand, main, simple, full, dashavc264:onDemand --> append all segments into init segment, seeker work
	// live , dashavc264:live--> create ms4 segments

	//previous segments are removed when new segment come
	MP4BoxAudioMux(videoFile, audioFile)
	cmd := fmt.Sprintln(config.MP4BOX +  " -dash-ctx " + mpdDir + "stream.txt -dash " + strconv.Itoa(duration) +
	" -dash-profile dashavc264:live -dynamic -time-shift -1 -fps 34 -mpd-refresh 1 -min-buffer 1000 -bs-switching inband -rap -frag " +
	strconv.Itoa(duration) + " -segment-name seg_ "  + " -out " + mpdDir + mpdName  +	" " + videoFile)

	//** -min-buffer durInMilSecs --> min buffer player need to fill before it can play. If data come fast, set this number
	//   smaller will let the player more responsive

	//New segment is appended to the init segment when it comes, cannot be done bc the init segment is opening by the player
//	MP4BoxAudioMux(videoFile, audioFile)
//	cmd := fmt.Sprintln(config.MP4BOX +  " -dash-ctx " + mpdDir + "stream.txt -dash " + strconv.Itoa(duration) +
//	" -mpd-refresh 3 -profile h264:live -rap -frag 1000 -dynamic -single-file -segment-name seg_ "  + " -out " + mpdDir + mpdName  +
//	" " + videoFile)
	log.Println("Command :" , cmd)
	err := exec.Command("bash", "-c", cmd).Run()

	if err != nil {
		log.Println("Dash Err, " , err)
		return err
	} else {
		return nil
	}
}

//Add audio track into video track
func MP4BoxAudioMux(video string, audio string)error{
	//Mux using MP4 Box, not working "cannot find H264 start code"
	//cmd := fmt.Sprintln(config.MP4BOX + " -fps 23.976 -add " + video + " -add " + audio + " " + newfile)
	cmd := fmt.Sprintln(config.MP4BOX+ " -add " + audio + " " + video)

	//result: "noaudio.h264: Invalid data found when processing input"
	//cmd := fmt.Sprintln(config.FFMPEG + " -framerate 25  -report  -i " + video + " -i " + audio +
	//" -codec copy -y " + newfile)

	log.Println(cmd)
	errs := exec.Command("bash", "-c", cmd).Run()
	if errs != nil {
		log.Println("-- muxing was failed, ", errs)
		return errs
	} else {
		//log.Println("-- muxing completed.")
		return nil
	}
}

// CONVERT AAC TO M4A
func ConvertAAC2M4A(aac string, m4a string) error {
	//ffmpeg -i muxtest/live/8.aac -c:a aac -b:a 128k muxtest/live/8.m4a
	cmd := fmt.Sprintln(config.FFMPEG + " -i " + aac + " -c:a aac -strict -2 -b:a 128k -y " + m4a)
	log.Println(cmd)
	errs := exec.Command("bash", "-c", cmd).Run()
	if errs != nil {
		log.Println("-- convert aac failed, ", errs)
		return errs
	} else {
		//log.Println("-- muxing completed.")
		return nil
	}

}


/*
usage: mp4fragment [options] <input> <output>
options are:
  --verbosity <n> sets the verbosity (details) level to <n> (between 0 and 3)
  --debug enable debugging information output
  --quiet don't print out notice messages
  --fragment-duration <milliseconds> (default = automatic)
  --timescale <n> (use 10000000 for Smooth Streaming compatibility)
  --track <track-id or type> only include media from one track (pass a track ID, 'audio', 'video' or 'subtitles')
  --index (re)create the segment index
  --trim trim excess media in longer tracks
  --no-tfdt don't add 'tfdt' boxes in the fragments (may be needed for legacy Smooth Streaming clients)
  --force-i-frame-sync <auto|all> treat all I-frames as sync samples (for open-gop sequences)
    'auto' only forces the flag if an open-gop source is detected, 'all' forces the flag in all cases
*/
func BentoFrag(input string, output string) error{
	cmd := fmt.Sprintln(config.BENTOFRAG + " " + input + " " + output)
	log.Println("cmd: ", cmd)
	//log.Println(cmd)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("Dash Err, ", err)
		return err
	} else {
		return nil
	}
}

func BentoDashIf(input BentoParam)error  {
	cmdStr := config.BENTODASH
	cmdStr += " " + input.InputVideo																	//input file

	if(input.MpdName != "") {cmdStr += " --mpd-name=" + input.MpdName}									//mpd file name
	if(input.Profile != "") {cmdStr += " --profile=" + input.Profile}   								//live, on-demand, hbbtv-1.5
	if(input.InitSegName != "") {cmdStr += " --init-segment=" + input.InitSegName		}				//init segment name
	if(input.MinBuffDuration != "") {cmdStr += " --min-buffer-time=" + input.MinBuffDuration}			//Minimum buffer time (in seconds)
	if(input.SmoothMpdName != "") {cmdStr += " --smooth-client-manifest-name=" + input.SmoothMpdName}	//Smooth Streaming Client Manifest file name
	if(input.NoMedia) {cmdStr += " --no-media"}															//no generate media files
	if(input.NoSplit) {cmdStr += " --no-split"}															//do not split into segment
	if(input.UseSegmentList) {cmdStr += " --use-segment-list"}											//Use segment lists instead of segment templates
	if(input.SmoothCompatible) {cmdStr += " --smooth"}													//Produce an output compatible with Smooth Streaming
	if(input.UseSegmentTimeLine) {cmdStr += " --use-segment-timeline"}									//Use segment timelines (necessary if segment durations	vary)
	if(input.Debug) {cmdStr += " --debug"}																//Use segment timelines (necessary if segment durations	vary)
	if(input.UseExistingDir) {cmdStr += " --force"}														//Allow output to an existing directory
	cmdStr += " --output-dir " + input.OutputDir																	//output directory

	log.Println("cmd:", cmdStr)
	err := exec.Command("bash", "-c", cmdStr).Run()
	if err != nil {
		log.Println("Dash Err, ", err)
		return err
	} else {
		return nil
	}
}

// NON-MULTIPLEX VIDEO FROM MUX VIDEOSEG - VOD
func CreateDashifFromMuxSeg(input MP4BoxParameter) error {
	//Options described on https://gpac.wp.mines-telecom.fr/mp4box/dash/
	cmdStr := config.MP4BOX
	if(input.DashCTX) {cmdStr += " -dash-ctx " + input.MpdDirectory + "stream.txt"}
	if(input.DashDuration != "") {cmdStr += " -dash " + input.DashDuration}
	if(input.UseSegmentTimeline){cmdStr += " -segment-timeline"}
	if(input.Profile != ""){cmdStr += " -profile " + input.Profile}
	if(input.BitstreamSwitch != ""){cmdStr += " -bs-switching " + input.BitstreamSwitch}
	if(input.RandomAccess){cmdStr += " -rap"}
	if(input.FragDuration != ""){cmdStr += " -frag " + input.FragDuration}
	cmdStr += " " + input.Video_Track1 + "#video"
	cmdStr += " " + input.Video_Track2 + "#video"
	cmdStr += " " + input.Video_Track3 + "#video"
	cmdStr += " " + input.Video_Track1 + "#audio"
	cmdStr += " -out " + input.MpdDirectory + input.MpdName

	log.Println(cmdStr)
	err := exec.Command("bash", "-c", cmdStr).Run()
	if err != nil {
		log.Println("Dash Err, ", err)
		return err
	} else {
		return nil
	}
}

type BentoParam struct {
	InputVideo			string
	OutputDir 			string
	OutputVideo			string
	Profile 			string
	MpdName 			string
	InitSegName 		string
	MinBuffDuration 	string
	SmoothMpdName 		string
	FramentDuration		string
	TimeScale			string
	TrackID				string
	NoMedia 			bool
	NoSplit 			bool
	UseSegmentList 		bool
	UseSegmentTimeLine 	bool
	SmoothCompatible 	bool
	Debug 				bool
	UseExistingDir  	bool
	IsQuiet             bool
	RecreateIndex		bool
	Trim 				bool
	Notfdt				bool
	ForceIframeSync		bool
}



type MP4BoxParameter struct{
	Video_Track1 		string		//specifies input file
	Video_Track2 		string		//multiple bitrate
	Video_Track3 		string		//multiple birate
	Audio_Track 		string
	DashDuration 		string		//enables DASH segmentation of input files with the given segment duration in ms
	FragDuration 		string		//specifies the duration of subsegments in ms
	MpdDirectory 		string		//specifies output directory for MPD.
	MpdName 			string		//specifies output file name for MPD.
	Profile 			string		//specifies the target DASH profile
	BitstreamSwitch 	string		//inband(default), merge, multi, no
	DashCTX 				bool		//stores and restore DASH timing from FILE
	UseSegmentTimeline 		bool		//uses SegmentTimeline when generating segments.
	RandomAccess 			bool		//forces segments to begin with random access points.
	FragmentRandomAccess	bool		//all fragments will begin with a random access points.
	Dynamic 				bool		//uses dynamic MPD type instead of static (always set for -dash-live)
	SegmentExt			string			//sets the segment extension. Default is m4s, null means no extension.
	SegmentName			string			//sets the segment name for generated segments
	TimeShift			string			//specifies MPD time shift buffer depth in seconds (default 0). Specify -1 to keep all files
	MinBuffer			string			//specifies MPD min buffer time in milliseconds.
	UseURLTemplate		bool			//uses SegmentTemplate instead of explicit sources in segments.
	SingleSegment		bool			//uses a single segment for each representation. Set by default for onDemand profile.
	SingleFile			bool			//uses a single file for each representation.
}

