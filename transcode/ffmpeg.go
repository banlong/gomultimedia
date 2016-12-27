package ffmpeg
import (
	"hermes/config"
	"os/exec"
"log"
"gomultimedia/tools"
	"container/list"
)

type Ffmpeg struct {
	Command string
}

type FfmpegParam struct{
	//split params
	InputVideo string
	SegmentDuration string
	SegmentExt string
	SegmentName string
	OutputLocation string
	FrameRate string
	Debug bool
	SegmentList string
}

func NewFfmpeg() *Ffmpeg {
	fObj := Ffmpeg{
		Command: "",
	}
	return &fObj
}

func (obj *Ffmpeg) Split(input FfmpegParam) (*list.List, error) {
	obj.Command = config.FFMPEG
	obj.Command += " -i " + input.InputVideo
	obj.Command += " -acodec copy -vcodec copy"
	if(input.Debug){obj.Command += " -report"}
	if(input.SegmentDuration != "") {obj.Command += " -f segment -segment_time " + input.SegmentDuration}
	if(input.SegmentList != ""){ obj.Command += " -segment_list " + input.SegmentList}
	if(input.FrameRate !=""){obj.Command += " -r " + input.FrameRate}

	//Name
	obj.Command += " " + input.OutputLocation
	if(input.SegmentName != "") {obj.Command += input.SegmentName }
	obj.Command += "%04d" + input.SegmentExt

	//Make sure temp folder exist
	err := tools.CreateDir(input.OutputLocation)

	//log.Println(obj.Command)
	err = exec.Command("bash", "-c", obj.Command).Run()
	if err != nil {
		log.Println("Split Err, ", err)
		return list.New(), err
	} else {
		list := tools.ParseList(input.SegmentList)
		return list, nil
	}
}