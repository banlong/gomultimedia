package main
import (
	"github.com/gorilla/mux"
	"log"
	"time"
	"net/http"
	"html/template"
	"gomultimedia/tools"
	"bytes"
)

func main()  {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/mpdlive/{videoId}", liveManifestSegHandler)
	router.HandleFunc("/adt/{videoId}", adaptiveHandler)

	router.HandleFunc("/mpd/{videoId}", BugHandler)
	log.Println("Streaming service starting...")
	log.Fatal(http.ListenAndServe(":9011", router))
}

func adaptiveHandler(w http.ResponseWriter, r *http.Request){
	log.Println("Live Dash Manifest handler" , r.URL)
	//log.Println( r.Header)
	//Get video id from the url
	pars := mux.Vars(r)
	//video id is expected as "fdsfhkdshkfs.mp4"
	videoId := pars["videoId"]
	log.Println("VideoId:", videoId)

	//Get video bytes from db

	videoData, err := tools.GetBytes("adaptive/dash/mpd/"+ videoId)
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

func BugHandler(w http.ResponseWriter, r *http.Request){
	log.Println("Bud handler" , r.URL)
	pars := mux.Vars(r)
	videoId := pars["videoId"]
	log.Println("VideoId:", videoId)

	//Get video bytes from db

	videoData, err := tools.GetBytes("bug/mpd/"+ videoId)
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