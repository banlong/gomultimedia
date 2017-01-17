package ffmpeg

type FFMPEGParam struct{
	InputVideo 	string
	SegmentDuration string
	SegmentExt 	string
	SegmentName 	string
	OutputLocation 	string
	FrameRate 	string
	Debug 		bool
	SegmentList 	string
	Quantifier 	string
	DisplayCMD      bool
}


type BentoParam struct {
	InputVideo		string
	OutputDir 		string
	OutputVideo		string
	Profile 		string
	MpdName 		string
	InitSegName 		string
	MinBuffDuration 	string
	SmoothMpdName 		string
	FramentDuration		string
	TimeScale		string
	TrackID			string
	NoMedia 		bool
	NoSplit 		bool
	UseSegmentList 		bool
	UseSegmentTimeLine 	bool
	SmoothCompatible 	bool
	Debug 			bool
	UseExistingDir  	bool
	IsQuiet             	bool
	RecreateIndex		bool
	Trim 			bool
	Notfdt			bool
	ForceIframeSync		bool
}

type MP4BoxParameter struct{
	Video_Track1 		string		//specifies input video
	Video_Track2 		string		//multiple bitrate
	Video_Track3 		string		//multiple bitrate
	Audio_Track 		string
	DashDuration 		string		//enables DASH segmentation of input files with the given segment duration in ms
	FragDuration 		string		//specifies the duration of sub-segments in ms
	MpdDirectory 		string		//specifies output directory for MPD.
	MpdName 		string		//specifies output file name for MPD.
	Profile 		string		//specifies the target DASH profile
	BitstreamSwitch 	string		//inband(default), merge, multi, no
	DashCTX 		bool		//stores and restore DASH timing from FILE
	UseSegmentTimeline 	bool		//uses SegmentTimeline when generating segments.
	RandomAccess 		bool		//forces segments to begin with random access points.
	FragmentRandomAccess	bool		//all fragments will begin with a random access points.
	Dynamic 		bool		//uses dynamic MPD type instead of static (always set for -dash-live)
	SegmentExt		string		//sets the segment extension. Default is m4s, null means no extension.
	SegmentName		string		//sets the segment name for generated segments
	TimeShift		string		//specifies MPD time shift buffer depth in seconds (default 0). Specify -1 to keep all files
	MinBuffer		string		//specifies MPD min buffer time in milliseconds.
	UseURLTemplate		bool		//uses SegmentTemplate instead of explicit sources in segments.
	SingleSegment		bool		//uses a single segment for each representation. Set by default for onDemand profile.
	SingleFile		bool		//uses a single file for each representation.
	DisplayCmdStr        bool   //display the command string
	CreateLog            bool   //generates log file for BIFS encoder and for LASeR encoder/decoder. The log is only useful to debug the scene codecs.
}

