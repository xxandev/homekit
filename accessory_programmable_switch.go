package homekit

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryProgrammableSwitch struct
type AccessoryProgrammableSwitch struct {
	*accessory.A
	ProgrammableSwitch *service.StatelessProgrammableSwitch
}

func (acc *AccessoryProgrammableSwitch) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryProgrammableSwitch) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryProgrammableSwitch) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryProgrammableSwitch) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryProgrammableSwitch) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryProgrammableSwitch) GetAccessory() *accessory.A {
	return acc.A
}

// NewAccessoryProgrammableSwitch returns *ProgrammableSwitch.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessoryProgrammableSwitch(info accessory.Info, args ...interface{}) *AccessoryProgrammableSwitch {
	acc := AccessoryProgrammableSwitch{}
	acc.A = accessory.New(info, accessory.TypeSwitch)
	acc.ProgrammableSwitch = service.NewStatelessProgrammableSwitch()
	acc.AddS(acc.ProgrammableSwitch.S)
	return &acc
}

// NewAccessoryProgrammableSwitch returns *ProgrammableSwitch.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccProgrammableSwitch(id uint64, info accessory.Info, args ...interface{}) *AccessoryProgrammableSwitch {
	acc := AccessoryProgrammableSwitch{}
	acc.A = accessory.New(info, accessory.TypeSwitch)
	acc.ProgrammableSwitch = service.NewStatelessProgrammableSwitch()
	acc.AddS(acc.ProgrammableSwitch.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessoryProgrammableSwitch) OnValuesRemoteUpdates(fn func()) {}
