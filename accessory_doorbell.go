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

func (acc *AccessoryDoorbell) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryDoorbell) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryDoorbell) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryDoorbell) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryDoorbell) GetAccessory() *accessory.Accessory {
	return acc.Accessory
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

func (acc *AccessoryDoorbell) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessoryDoorbell) OnValuesRemoteUpdatesPrint()     {}
