package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//ServiceTelevisionSpeaker -
type ServiceTelevisionSpeaker struct {
	*service.Service

	Mute              *characteristic.Mute
	Active            *characteristic.Active
	Volume            *characteristic.Volume
	VolumeControlType *characteristic.VolumeControlType
	VolumeSelector    *characteristic.VolumeSelector
}

//NewServiceTelevisionSpeaker -
func NewServiceTelevisionSpeaker() *ServiceTelevisionSpeaker {
	svc := ServiceTelevisionSpeaker{}
	svc.Service = service.New(service.TypeSpeaker)

	svc.Mute = characteristic.NewMute()
	svc.AddCharacteristic(svc.Mute.Characteristic)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	svc.Volume = characteristic.NewVolume()
	svc.AddCharacteristic(svc.Volume.Characteristic)

	svc.VolumeControlType = characteristic.NewVolumeControlType()
	svc.AddCharacteristic(svc.VolumeControlType.Characteristic)

	svc.VolumeSelector = characteristic.NewVolumeSelector()
	svc.AddCharacteristic(svc.VolumeSelector.Characteristic)

	return &svc
}

//AccessoryTelevision struct
type AccessoryTelevision struct {
	*accessory.Accessory
	Television        *service.Television
	TelevisionSpeaker *ServiceTelevisionSpeaker
}

// NewAccessoryTelevision returns AccessorySwitch (args... are not used)
func NewAccessoryTelevision(info accessory.Info, args ...interface{}) *AccessoryTelevision {
	acc := AccessoryTelevision{}
	acc.Accessory = accessory.New(info, accessory.TypeTelevision)
	acc.Television = service.NewTelevision()
	acc.TelevisionSpeaker = NewServiceTelevisionSpeaker()

	acc.AddService(acc.Television.Service)
	acc.AddService(acc.TelevisionSpeaker.Service)

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

//ProcessInputSource -
//ConfiguredName
//InputSourceType
//IsConfigured
//Identifier
//TargetVisibilityState
//Name
func (acc *AccessoryTelevision) ProcessInputSource(insource *service.InputSource, events ...func(interface{})) {
	amountEvents := len(events)
	if amountEvents > 0 {
		insource.ConfiguredName.OnValueRemoteUpdate(func(v string) { events[0](v) })
	}
	if amountEvents > 1 {
		insource.TargetVisibilityState.OnValueRemoteUpdate(func(v int) { events[1](v) })
	}
	if amountEvents > 2 {
		insource.InputSourceType.OnValueRemoteUpdate(func(v int) { events[2](v) })
	}
	if amountEvents > 3 {
		insource.IsConfigured.OnValueRemoteUpdate(func(v int) { events[3](v) })
	}
	if amountEvents > 4 {
		insource.Identifier.OnValueRemoteUpdate(func(v int) { events[4](v) })
	}
	if amountEvents > 5 {
		insource.Name.OnValueRemoteUpdate(func(v string) { events[5](v) })
	}
}
