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

// NewAccessoryProgrammableSwitch return AccessoryProgrammableSwitch (args... are not used)
func NewAccessoryProgrammableSwitch(info accessory.Info, args ...interface{}) *AccessoryProgrammableSwitch {
	acc := AccessoryProgrammableSwitch{}
	acc.A = accessory.New(info, accessory.TypeSwitch)
	acc.ProgrammableSwitch = service.NewStatelessProgrammableSwitch()
	acc.AddS(acc.ProgrammableSwitch.S)
	return &acc
}

func (acc *AccessoryProgrammableSwitch) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessoryProgrammableSwitch) OnExample()                      {}
