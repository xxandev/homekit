package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryCamera provides RTP video streaming.
type AccessoryCamera struct {
	*accessory.Accessory
	Control           *service.CameraControl
	StreamManagement1 *service.CameraRTPStreamManagement
	StreamManagement2 *service.CameraRTPStreamManagement
}

func (acc *AccessoryCamera) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryCamera) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryCamera) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryCamera) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryCamera) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryCamera returns an IP camera accessory.
func NewAccessoryCamera(info accessory.Info, args ...interface{}) *AccessoryCamera {
	acc := AccessoryCamera{}
	acc.Accessory = accessory.New(info, accessory.TypeIPCamera)
	acc.Control = service.NewCameraControl()
	acc.AddService(acc.Control.Service)

	// TODO (mah) a camera must support at least 2 rtp streams
	acc.StreamManagement1 = service.NewCameraRTPStreamManagement()
	acc.StreamManagement2 = service.NewCameraRTPStreamManagement()
	acc.AddService(acc.StreamManagement1.Service)
	// acc.AddService(acc.StreamManagement2.Service)

	return &acc
}

func (acc *AccessoryCamera) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessoryCamera) OnValuesRemoteUpdatesPrint()     {}
