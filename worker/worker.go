package worker

import (
	"sync"
	"log"
	"goAV/transcode"
)

type Args struct{
	SrcFile string
	DestDir string
	Name string
}

type Worker struct{
	Id int
	InputChan chan *Args
	Wg sync.WaitGroup
}

func (w *Worker) Run(){
	//The worker once runs, will keep continue until the input channel is closed
	log.Printf("Worker %d started", w.Id)
	for input:= range w.InputChan{
		//call rpc method of the client to get Multiplication result
		transcoding.Transcode(input.SrcFile, input.DestDir, input.Name, ".mp4")
		w.Wg.Done()
	}
}