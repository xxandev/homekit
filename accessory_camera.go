package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryCamera provides RTP video streaming.
type AccessoryCamera struct {
	*accessory.A
	Control           *service.CameraControl
	StreamManagement1 *service.CameraRTPStreamManagement
	StreamManagement2 *service.CameraRTPStreamManagement
}

func (acc *AccessoryCamera) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryCamera) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryCamera) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryCamera) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryCamera) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryCamera) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryCamera returns *Camera.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessoryCamera(info accessory.Info, args ...interface{}) *AccessoryCamera {
	acc := AccessoryCamera{}
	acc.A = accessory.New(info, accessory.TypeIPCamera)
	acc.Control = service.NewCameraControl()
	acc.AddS(acc.Control.S)

	// TODO (mah) a camera must support at least 2 rtp streams
	acc.StreamManagement1 = service.NewCameraRTPStreamManagement()
	acc.StreamManagement2 = service.NewCameraRTPStreamManagement()
	acc.AddS(acc.StreamManagement1.S)
	// acc.AddService(acc.StreamManagement2.Service)
	return &acc
}

//NewAccCamera returns *Camera.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccCamera(id uint64, info accessory.Info, args ...interface{}) *AccessoryCamera {
	acc := AccessoryCamera{}
	acc.A = accessory.New(info, accessory.TypeIPCamera)
	acc.Control = service.NewCameraControl()
	acc.AddS(acc.Control.S)

	// TODO (mah) a camera must support at least 2 rtp streams
	acc.StreamManagement1 = service.NewCameraRTPStreamManagement()
	acc.StreamManagement2 = service.NewCameraRTPStreamManagement()
	acc.AddS(acc.StreamManagement1.S)
	// acc.AddService(acc.StreamManagement2.Service)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryCamera) OnValuesRemoteUpdates(fn func()) {}
