package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/characteristic"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryTelevision struct
type AccessoryTelevision struct {
	*accessory.A
	Television *haps.Television
	Speaker    *haps.TelevisionSpeaker
}

func (acc *AccessoryTelevision) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryTelevision) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryTelevision) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryTelevision) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryTelevision) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryTelevision) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryTelevision returns *Television.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessoryTelevision(info accessory.Info, args ...interface{}) *AccessoryTelevision {
	acc := AccessoryTelevision{}
	acc.A = accessory.New(info, accessory.TypeTelevision)
	acc.Television = haps.NewTelevision()
	acc.Speaker = haps.NewTelevisionSpeaker()
	acc.AddS(acc.Television.S)
	acc.AddS(acc.Speaker.S)
	return &acc
}

//NewAccessoryTelevision returns AccessoryTelevision (args... are not used)
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccTelevision(id uint64, info accessory.Info, args ...interface{}) *AccessoryTelevision {
	acc := AccessoryTelevision{}
	acc.A = accessory.New(info, accessory.TypeTelevision)
	acc.Television = haps.NewTelevision()
	acc.Speaker = haps.NewTelevisionSpeaker()
	acc.AddS(acc.Television.S)
	acc.AddS(acc.Speaker.S)
	acc.A.Id = id
	return &acc
}

//AddInputSource -
func (acc *AccessoryTelevision) AddInputSource(id int, name string, inputSourceType int) *haps.InputSource {
	inSource := haps.NewInputSource()

	inSource.Identifier.SetValue(id)
	inSource.ConfiguredName.SetValue(name)
	inSource.Name.SetValue(name)
	inSource.InputSourceType.SetValue(inputSourceType)
	inSource.IsConfigured.SetValue(characteristic.IsConfiguredConfigured)

	acc.AddS(inSource.S)
	acc.Television.S.AddS(inSource.S)
	return inSource
}

//ProcessInputSource
//  events[0](val string) - ConfiguredName
//  events[1](val int) - InputSourceType
//  events[2](val int) - IsConfigured
//  events[3](val int) - Identifier
//  events[4](val int) - TargetVisibilityState
//  events[5](val string) - Name
func (acc *AccessoryTelevision) ProcessInputSource(insource *haps.InputSource, events ...func(interface{})) {
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
