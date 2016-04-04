//CONVERT A VIDEO
//Have to install the ffmpeg build version, set path to its \bin so that
//we can call ffmpeg
//the wrapper does not working

package main
import (
	"log"
	"os"
	"sync"
	"goAV/worker"
	"goAV/transcode"
	"goAV/tools"
)

func main(){
	TestMux()
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
		transcoding.Mux(video, audio, output)
	}


}

func TestDemux() {
	video := "videos/muxed.mp4"
	audio := "videos/novideo2.mp3"
	output := "videos/muxed2.mp4"
	transcoding.Demux(video, audio, output)
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
	names, err := transcoding.Split(srcFile, duration , tempDir, "sample", ".mp4")
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
				transcoding.Transcode(chunkFile, trcdDir, tools.ZeroPad(i,3), ".mp4")
				waitGroup.Done()
				log.Printf("Worker %d exit \n", i)
			}(chunkFile, i)
		}

		waitGroup.Wait()
		log.Println("Transcoding Completed")

		//Merge file
		log.Println("start merging Trunks")
		transcoding.CreateMergeList(trcdDir, namesFile)
		transcoding.Merge(namesFile, outDir, outFile)
		transcoding.GenerateJPGThumbnail("videos/", "tth.mp4")
		//deleteDir(tempDir)
		//deleteDir(trcdDir)

		//Upload to the AWS S3
		//Upload(outFile)
		//deleteDir(outDir)
	}
}

func TestThumbnailGenerator(){
	transcoding.GenerateJPGThumbnail("videos/", "tth.mp4")
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
	names, err := transcoding.Split(srcFile, duration , tempDir, "sample", ".mp4")
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
		transcoding.CreateMergeList(trcdDir, mergeList)
		transcoding.Merge(mergeList, outDir, outFile)
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



