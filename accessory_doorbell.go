package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryDoorbell provides RTP video streaming, Speaker and Mic controls
type AccessoryDoorbell struct {
	*accessory.Accessory
	Control          *service.Doorbell
	StreamManagement *service.CameraRTPStreamManagement
	Speaker          *service.Speaker
	Microphone       *service.Microphone
}

//NewAccessoryDoorbell returns a Video Doorbell accessory.
func NewAccessoryDoorbell(info accessory.Info, args ...interface{}) *AccessoryDoorbell {
	acc := AccessoryDoorbell{}
	acc.Accessory = accessory.New(info, accessory.TypeVideoDoorbell)
	acc.Control = service.NewDoorbell()
	acc.AddService(acc.Control.Service)

	acc.StreamManagement = service.NewCameraRTPStreamManagement()
	acc.AddService(acc.StreamManagement.Service)

	acc.Speaker = service.NewSpeaker()
	acc.AddService(acc.Speaker.Service)

	acc.Microphone = service.NewMicrophone()
	acc.AddService(acc.Microphone.Service)

	return &acc
}
