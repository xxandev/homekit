package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryProgrammableSwitch struct
type AccessoryProgrammableSwitch struct {
	*accessory.Accessory
	ProgrammableSwitch *service.StatelessProgrammableSwitch
}

func (acc *AccessoryProgrammableSwitch) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryProgrammableSwitch) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryProgrammableSwitch) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryProgrammableSwitch) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryProgrammableSwitch) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

// NewAccessoryProgrammableSwitch return AccessoryProgrammableSwitch (args... are not used)
func NewAccessoryProgrammableSwitch(info accessory.Info, args ...interface{}) *AccessoryProgrammableSwitch {
	acc := AccessoryProgrammableSwitch{}
	acc.Accessory = accessory.New(info, accessory.TypeSwitch)
	acc.ProgrammableSwitch = service.NewStatelessProgrammableSwitch()
	acc.AddService(acc.ProgrammableSwitch.Service)
	return &acc
}

func (acc *AccessoryProgrammableSwitch) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessoryProgrammableSwitch) OnValuesRemoteUpdatesPrint()     {}
