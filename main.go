//CONVERT A VIDEO
//Have to install the ffmpeg build version, set path to its \bin so that we can call ffmpeg

package main
import (
	"log"
	"gomultimedia/tools"
	"time"
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
	"bytes"
	"gomultimedia/transcode"
	"path"
	"strings"
)

func main(){


	//Start()
	//ProduceVODDash()
	//ProduceLiveDash()
	//ProduceDashIf()
}

func Start()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/mpdlive/{videoId}", liveManifestSegHandler)
	log.Println("Streaming service starting...")
	log.Fatal(http.ListenAndServe(":9011", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request)  {
	log.Println("home handler: " , r.URL)
	w.Header().Add("Content Type", "text/html")
	var tc *template.Template
	tc = template.Must(template.ParseFiles("index.html"))
	tc.Execute(w, nil)
}

func liveManifestSegHandler(w http.ResponseWriter, r *http.Request){
	log.Println("Live Dash Manifest handler" , r.URL)
	//log.Println( r.Header)
	//Get video id from the url
	pars := mux.Vars(r)
	//video id is expected as "fdsfhkdshkfs.mp4"
	videoId := pars["videoId"]
	log.Println("VideoId:", videoId)

	//Get video bytes from db

	videoData, err := tools.GetBytes("factory/mpd/"+ videoId)
	//log.Println(len(videoData))
	if(err != nil){
		log.Println("Get error: ", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-PINGOTHER, Range")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	vReader := bytes.NewReader(videoData)
	http.ServeContent(w, r, videoId, time.Now(), vReader)
}

func ProduceLiveDash() error{
	//Multiplex Dash, not play on DashIF

	processStart := time.Now()

	srcFile := "videos/kyduyen.mp4"
	tempDir := "factory/"
	segDir := tempDir + "segments/"
	videoDir := tempDir + "video/"
	audioDir := tempDir + "audio/"
	encodeDir := tempDir + "encode/"
	mpdDir := tempDir + "mpd/"
	audioExt := ".aac"
	duration := 10


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
		ffmpeg.ProduceDashLive(3000, mpdDir, "kd2.mpd",  videoFile, audioFile)
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
