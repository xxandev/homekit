package hapcam

//ConfigFFMPEG contains ffmpeg parameters
type ConfigFFMPEG struct {
	InputDevice      string
	InputFilename    string
	LoopbackFilename string
	H264Decoder      string
	H264Encoder      string
	MinVideoBitrate  int
	MultiStream      bool
}
