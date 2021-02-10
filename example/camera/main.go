package main

import (
	"fmt"
	"image"

	"github.com/alpr777/homekit"
	"github.com/alpr777/homekit/ffmpeg"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/log"
)

const (
	accessoryName string = "camera"
	accessorySn   string = "ExmplCam"
	accessoryPin  string = "19283746"
)

func main() {
	// runtime.GOMAXPROCS(4)
	// log.Debug.Enable()
	// ffmpeg.EnableVerboseLogging()
	acc := homekit.NewAccessoryCamera(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "EXAMPLE", Model: "ACC-TEST", FirmwareRevision: "1.2"})
	ffmpeg := acc.SetupFFMPEGStreaming(ffmpeg.Config{
		InputDevice:      "v4l2",
		InputFilename:    "/dev/video0",
		LoopbackFilename: "/dev/video1",
		H264Decoder:      "",
		H264Encoder:      "h264_omx",
		MinVideoBitrate:  0,
		MultiStream:      false,
	})
	cc := homekit.NewCameraControl()
	acc.Control.AddCharacteristic(cc.Assets.Characteristic)
	acc.Control.AddCharacteristic(cc.GetAsset.Characteristic)
	acc.Control.AddCharacteristic(cc.DeleteAssets.Characteristic)
	acc.Control.AddCharacteristic(cc.TakeSnapshot.Characteristic)

	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: accessoryPin, Port: "10201"}, acc.Accessory)
	if err != nil {
		log.Info.Panic(err)
	}

	transp.CameraSnapshotReq = func(width, height uint) (*image.Image, error) {
		return ffmpeg.Snapshot(width, height)
	}

	cc.SetupWithDir("./" + acc.Info.SerialNumber.GetValue())
	cc.CameraSnapshotReq = func(width, height uint) (*image.Image, error) {
		return ffmpeg.Snapshot(width, height)
	}
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
