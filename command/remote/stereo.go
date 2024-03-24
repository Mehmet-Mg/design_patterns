package main

import "fmt"

type Stereo struct {
	location string
}

func NewStereo(location string) *Stereo {
	return &Stereo{
		location: location,
	}
}

func (s Stereo) On() {
	fmt.Println(s.location + " stereo is on")
}

func (s Stereo) Off() {
	fmt.Println(s.location + " stereo is off")
}

func (s Stereo) SetCD() {
	fmt.Println(s.location + " stereo is set for CD input")
}

func (s Stereo) SetDVD() {
	fmt.Println(s.location + " stereo is set for DVD input")
}

func (s Stereo) SetRadio() {
	fmt.Println(s.location + " stereo is set for Radio")
}

func (s Stereo) SetVolume(volume int) {
	// code to set the volume
	// valid range: 1-11 (after all 11 is better than 10, right?)
	fmt.Printf("%s stereo volume set to %d\n", s.location, volume)
}
