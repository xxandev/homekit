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

// NewAccessoryProgrammableSwitch return AccessoryProgrammableSwitch (args... are not used)
func NewAccessoryProgrammableSwitch(info accessory.Info, args ...interface{}) *AccessoryProgrammableSwitch {
	acc := AccessoryProgrammableSwitch{}
	acc.Accessory = accessory.New(info, accessory.TypeSwitch)
	acc.ProgrammableSwitch = service.NewStatelessProgrammableSwitch()
	acc.AddService(acc.ProgrammableSwitch.Service)
	return &acc
}
