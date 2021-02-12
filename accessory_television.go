package homekit

import (
	"github.com/alpr777/homekit/hapservices"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//AccessoryTelevision struct
type AccessoryTelevision struct {
	*accessory.Accessory
	Television *service.Television
	Speaker    *hapservices.TelevisionSpeaker
}

// NewAccessoryTelevision returns AccessorySwitch (args... are not used)
func NewAccessoryTelevision(info accessory.Info, args ...interface{}) *AccessoryTelevision {
	acc := AccessoryTelevision{}
	acc.Accessory = accessory.New(info, accessory.TypeTelevision)
	acc.Television = service.NewTelevision()
	acc.Speaker = hapservices.NewTelevisionSpeaker()

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
