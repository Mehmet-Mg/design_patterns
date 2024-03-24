package main

type StereoOffCommand struct {
	stereo Stereo
}

func NewStereoOffCommand(stereo Stereo) *StereoOffCommand {
	return &StereoOffCommand{
		stereo: stereo,
	}
}

func (s StereoOffCommand) Execute() {
	s.stereo.Off()
}

func (s StereoOffCommand) Undo() {
	s.stereo.On()
}
