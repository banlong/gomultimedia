package config

const (



	//-----MEDIA CONFIGURATION -----//
	BUFFER_DIR 	string = "/transcode-buffer/"
	FFMPEG_PRO	string = "/opt/ffmpeg/ffmpeg"
	FFPROB_PRO  string = "/opt/ffmpeg/ffprobe"
	FFMPEG_LC	string = "ffmpeg"
	FFPROB_LC   string = "ffprobe"
	MP4BOX  	string = "MP4Box"
	MP4EDIT  	string = "mp4edit"
	BENTOFRAG   string = "mp4fragment"
	BENTODASH	string = "mp4dash"
	ROOT_DIR	string = "tmp/transcode-buffer/"
	STREAM_PUBLIC_URL	string	= "http://beta.cut2it.com"
	STREAM_LOCAL_URL	string = "http://localhost:9099"

	TEMP_DIR 	string = "temp/"
	HSL_DIR    string = "hsl/"
	MP4_DIR    string = "mp4/"
	OUT_DIR		string = "whole/"
	NAMES_FILE 	string = "list.txt"
	DURATION	int    = 3		//seconds

	//-------DATABASE CONFIGURATION ----- //
	LOCAL_DB_DIR = "c:/goDB/"
	PRO_DB_DIR = "/tmp/database/" //<-- Need to define
)

//----INTERNAL IDENTIFIER----//
var (
	MOIP string //Medulla-Oblongata source
	FFMPEG string = "ffmpeg"
	FFPROBE string = "ffprobe"
	THEORA string = "ffmpeg2theora"
	DB_DIR string
	STREAM_URL string
)
