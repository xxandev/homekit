package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryDoorbell provides RTP video streaming, Speaker and Mic controls
type AccessoryDoorbell struct {
	*accessory.A
	Control          *service.Doorbell
	StreamManagement *service.CameraRTPStreamManagement
	Speaker          *service.Speaker
	Microphone       *service.Microphone
}

func (acc *AccessoryDoorbell) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryDoorbell) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryDoorbell) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryDoorbell) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryDoorbell) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryDoorbell) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryDoorbell returns a Video Doorbell accessory.
func NewAccessoryDoorbell(info accessory.Info, args ...interface{}) *AccessoryDoorbell {
	acc := AccessoryDoorbell{}
	acc.A = accessory.New(info, accessory.TypeVideoDoorbell)
	acc.Control = service.NewDoorbell()
	acc.AddS(acc.Control.S)

	acc.StreamManagement = service.NewCameraRTPStreamManagement()
	acc.AddS(acc.StreamManagement.S)

	acc.Speaker = service.NewSpeaker()
	acc.AddS(acc.Speaker.S)

	acc.Microphone = service.NewMicrophone()
	acc.AddS(acc.Microphone.S)

	return &acc
}

func (acc *AccessoryDoorbell) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessoryDoorbell) OnExample()                      {}
