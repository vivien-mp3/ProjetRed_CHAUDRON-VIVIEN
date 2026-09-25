package common

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	initOnce  sync.Once
	initErr   error
	speakerSR beep.SampleRate
	lastsound string
)

func play(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Println("ouverture:", err)
		return
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		f.Close()
		log.Println("décodage:", err)
		return
	}

	// génère le son qu'une fois
	initOnce.Do(func() {
		speakerSR = format.SampleRate
		initErr = speaker.Init(speakerSR, speakerSR.N(time.Second/10))
	})
	if initErr != nil {
		streamer.Close()
		log.Println("speaker:", initErr)
		return
	}

	// rebeep le son si il est pas pareil ou jsp
	var s beep.Streamer = streamer
	if format.SampleRate != speakerSR {
		s = beep.Resample(4, format.SampleRate, speakerSR, streamer)
	}

	// ça joue le son -b (ceci est un pouce)
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		streamer.Close()
	})))
}

func PlaySound(s string) {
	if lastsound != s {
		StopAllSounds()
		go play("./assets/music/" + s + ".mp3")
	}
}

func PlaySFX(s string) {
	play("./assets/sfx/" + s + ".mp3")
}

func StopAllSounds() {
	speaker.Clear()
}