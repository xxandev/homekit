package homekit

import (
	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//AccessoryTelevision struct
type AccessoryTelevision struct {
	*accessory.Accessory
	Television *service.Television
	Speaker    *haps.TelevisionSpeaker
}

func (acc *AccessoryTelevision) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryTelevision) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryTelevision) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryTelevision) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryTelevision) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryTelevision returns AccessoryTelevision (args... are not used)
func NewAccessoryTelevision(info accessory.Info, args ...interface{}) *AccessoryTelevision {
	acc := AccessoryTelevision{}
	acc.Accessory = accessory.New(info, accessory.TypeTelevision)
	acc.Television = service.NewTelevision()
	acc.Speaker = haps.NewTelevisionSpeaker()

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

//ProcessInputSource
//  events[0](val string) - ConfiguredName
//  events[1](val int) - InputSourceType
//  events[2](val int) - IsConfigured
//  events[3](val int) - Identifier
//  events[4](val int) - TargetVisibilityState
//  events[5](val string) - Name
func (acc *AccessoryTelevision) ProcessInputSource(insource *service.InputSource, events ...func(interface{})) {
	n := len(events)
	if n > 0 {
		insource.ConfiguredName.OnValueRemoteUpdate(func(v string) { events[0](v) })
	}
	if n > 1 {
		insource.TargetVisibilityState.OnValueRemoteUpdate(func(v int) { events[1](v) })
	}
	if n > 2 {
		insource.InputSourceType.OnValueRemoteUpdate(func(v int) { events[2](v) })
	}
	if n > 3 {
		insource.IsConfigured.OnValueRemoteUpdate(func(v int) { events[3](v) })
	}
	if n > 4 {
		insource.Identifier.OnValueRemoteUpdate(func(v int) { events[4](v) })
	}
	if n > 5 {
		insource.Name.OnValueRemoteUpdate(func(v string) { events[5](v) })
	}
}

func (acc *AccessoryTelevision) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessoryTelevision) OnValuesRemoteUpdatesPrint()     {}
