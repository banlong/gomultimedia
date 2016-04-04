package transcoding

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"container/list"
	"goAV/config"
	"strings"
	"goAV/tools"
	"io/ioutil"
	"path"
)

func init() {
	//createDirectories()
	//tools.CreateDir("tmp/transcode-buffer/")
}

// MY TRANSCODER
func Transcode(oldFile string, destDir string,  name string, ext string) error {
	/* --------------------------COMMAND SECTION -------------------------------------------------------- */
	//Cut2it define
	cmd := fmt.Sprintln(config.FFMPEG + " -i " + oldFile +
	" -y -c:v libx264 -preset ultrafast -threads 0 -c:a aac -strict -2 " + destDir + name + ext )


	//cmd := fmt.Sprintln("ffmpeg -i " + oldFile + " -c:v libx264 -preset ultrafast -threads 0 " +
	//                    "-c:a aac -strict -2 /tmp/transcode-buffer/" + newName + ".mp4")

	//To set the video bitrate of the output file to 64 kbit/s:
	//cmd := fmt.Sprintln("ffmpeg -i " + oldFile + " -c:v 64k -bufsize 64k " + newName)

	//To force the frame rate of the output file to 24 fps:
	//cmd := fmt.Sprintln("ffmpeg -i " + oldFile + " -r 24 " + newName)
	/* -------------------------------------------------------------------------------------------------- */
	//log.Println("Command :" + cmd)
	err := exec.Command("bash", "-c", cmd).Run()

	if err != nil {
		return err
	} else {
		//log.Println("Transcoding Completed: " + oldFile)
		return nil
	}
}

// SPLIT ONE MP4 INTO MULTIPLE MP4s WITH SAME LENGTH (SECONDS), n: seconds
func Split(input string, seconds int, outputDir string, videoId string, ext string) (*list.List, error){
	log.Println("Start splitting origin video...")

	//Make sure temp folder exist
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir,0711)
		if err != nil {
			log.Println("Error creating temp directory")
			return list.New(),err
		}
	}

	/* --------------------------COMMAND SECTION -------------------------------------------------------- */
	// OPT 2 - Split input file into equally files with segment in seconds, -vcodec will allow split AVI
	cmd := config.FFMPEG + " -i " + input + " -vcodec copy -map 0 -segment_time " + strconv.Itoa(seconds) +
			" -f segment " + outputDir + videoId + "-%04d" + ext

//	cmd := config.FFMPEG + " -i " + input + " -c copy -map 0 -segment_time " + strconv.Itoa(seconds) +
//	" -f segment " + outputDir + videoId + "-%04d" + ext
	/* -------------------------------------------------------------------------------------------------- */
	//log.Println("cmd:" ,cmd)
	err := exec.Command("bash", "-c", cmd).Run()
	if err != nil {
		log.Println("Split failure!", err.Error())
		return list.New(), err
	} else {
		log.Println("Split file completed")
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

// GET SEGMENT DURATION
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

// GET SEGMENT DURATION v2
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

	log.Println(tools.IsExist(video))
	err := Transcode(video, dir, "videot", ".mp4")
	if err != nil {
		log.Println("-- transcode failed, ", err)
		return err
	} else {
		log.Println("-- transcode completed.")
	}
	log.Println(err)
	srcVStream := dir + "videot" +  ".mp4"
	log.Println(srcVStream)


	cmd := fmt.Sprintln(config.FFMPEG + " -i " + srcVStream + " -i " + audio +
	" -c:v copy -c:a aac -strict experimental  -y " + newfile)
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