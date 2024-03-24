package main

type StereoOnWithCDCommand struct {
	stereo Stereo
}

func NewStereoOnWithCDCommand(stereo Stereo) *StereoOnWithCDCommand {
	return &StereoOnWithCDCommand{
		stereo: stereo,
	}
}

func (s *StereoOnWithCDCommand) Execute() {
	s.stereo.On()
	s.stereo.SetCD()
	s.stereo.SetVolume(11)
}

func (s *StereoOnWithCDCommand) Undo() {
	s.stereo.Off()
}
