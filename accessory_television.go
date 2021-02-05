package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//AccessoryTelevision struct
type AccessoryTelevision struct {
	*accessory.Accessory
	Television *service.Television
	Speaker    *service.Speaker
}

// NewAccessoryTelevision returns AccessorySwitch (args... are not used)
func NewAccessoryTelevision(info accessory.Info, args ...interface{}) *AccessoryTelevision {
	acc := AccessoryTelevision{}
	acc.Accessory = accessory.New(info, accessory.TypeTelevision)
	acc.Television = service.NewTelevision()
	acc.Speaker = service.NewSpeaker()

	acc.Television.Brightness.SetValue(0)
	acc.Television.Brightness.SetMinValue(0)
	acc.Television.Brightness.SetMaxValue(100)
	acc.Television.Brightness.SetStepValue(1)

	acc.Television.TargetMediaState.SetValue(0)
	acc.Television.TargetMediaState.SetMinValue(0)
	acc.Television.TargetMediaState.SetMaxValue(2)
	acc.Television.TargetMediaState.SetStepValue(1)

	acc.Television.PictureMode.SetValue(0)
	acc.Television.PictureMode.SetMinValue(0)
	acc.Television.PictureMode.SetMaxValue(7)
	acc.Television.PictureMode.SetStepValue(1)

	acc.Television.RemoteKey.SetValue(0)
	acc.Television.RemoteKey.SetMinValue(0)
	acc.Television.RemoteKey.SetMaxValue(15)
	acc.Television.RemoteKey.SetStepValue(1)

	acc.AddService(acc.Television.Service)
	acc.AddService(acc.Speaker.Service)

	return &acc
}

//AddInputSource -
func (acc *AccessoryTelevision) AddInputSource(id int, name string, inputSourceType int) *service.InputSource {
	inSource := service.NewInputSource()

	inSource.Identifier.SetValue(id)
	inSource.ConfiguredName.SetValue(name)
	inSource.Name.SetValue(name)
	inSource.InputSourceType.SetValue(inputSourceType)
	inSource.IsConfigured.SetValue(characteristic.IsConfiguredConfigured)

	acc.AddService(inSource.Service)
	acc.Television.AddLinkedService(inSource.Service)
	return inSource
}
