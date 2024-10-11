package backend

import (
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

type SoundPlayer struct {
	SoundFile *os.File
	Streamer  beep.StreamSeekCloser
	Err       error
}

func (s *SoundPlayer) Initialize(filename string) {
	var format beep.Format
	s.SoundFile, s.Err = os.Open(filename)
	if s.Err != nil {
		log.Fatal(s.Err)
		return
	}
	s.Streamer, format, s.Err = wav.Decode(s.SoundFile)
	if s.Err != nil {
		log.Fatal(s.Err)
		return
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
}

func (s *SoundPlayer) Play() {
	if s.Err != nil {
		return
	}
	done := make(chan bool)
	speaker.Play(beep.Seq(s.Streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
	s.Streamer.Seek(0)
}

func (s *SoundPlayer) ShutDown() {
	s.SoundFile.Close()
	s.Streamer.Close()
}
