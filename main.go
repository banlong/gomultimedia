//CONVERT A VIDEO
//Have to install the ffmpeg build version, set path to its \bin so that
//we can call ffmpeg
//the wrapper does not working

package main
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
)

func main(){
	//It seems that Split-Extract-Encode have more advantage on the processing time. However this model have to
	//go with a method to stream video with different languages and resolution like HLS and MPEG DASH. Otherwise,
	//mux audio and video and streaming the whole mp4 will make these balance ie no improvement
	SplitEncodeExtract()
	SplitExtractEncode()
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
	duration := 10


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	names, err := ffmpeg.Split(srcFile, duration , segDir, "sample", ".mp4")
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

		err = ffmpeg.Extract(filename, videoDir + name + ".mp4" , audioDir + name + ".mp3")
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

		err = ffmpeg.Transcode(filename, encodeDir , name, ".mp4")
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
	duration := 10


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	names, err := ffmpeg.Split(srcFile, duration , segDir, "sample", ".mp4")
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

		err = ffmpeg.Extract(filename, videoDir + name + ".mp4" , audioDir + name + ".mp3")
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
	duration := 10


	//Split file
	splitStart := time.Now()
	log.Println("start splitting file")
	names, err := ffmpeg.Split(srcFile, duration , segDir, "sample", ".mp4")
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

		err = ffmpeg.Transcode(filename, encodeDir , name, ".mp4")
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

		err = ffmpeg.Extract(filename, videoDir + name + ".mp4" , audioDir + name + ".mp3")
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
	input := "videos/sample.mkv"
	ffmpeg.Extract(input, "noaudio.m4v", "novideo.mp3")
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
	srcFile := "videos/tth.mp4"
	namesFile := "list.txt"
	tempDir := "videos/temp/"
	trcdDir := "videos/trcd/"
	outDir := "videos/out/"
	outFile := "final.mp4"
	duration := 10
	//Split file
	log.Println("start splitting file")
	names, err := ffmpeg.Split(srcFile, duration , tempDir, "sample", ".mp4")
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
				ffmpeg.Transcode(chunkFile, trcdDir, tools.ZeroPad(i,3), ".mp4")
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
	duration := 180
	//Split file
	log.Println("start splitting file")
	names, err := ffmpeg.Split(srcFile, duration , tempDir, "sample", ".mp4")
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



