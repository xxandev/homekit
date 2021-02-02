package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryLightbulbSwitch struct
type AccessoryLightbulbSwitch struct {
	*accessory.Accessory
	LightbulbSwitch *service.Lightbulb
}

//NewAccessoryLightbulbSwitch return AccessoryLightbulbSwitch (args... are not used)
func NewAccessoryLightbulbSwitch(info accessory.Info, args ...interface{}) *AccessoryLightbulbSwitch {
	acc := AccessoryLightbulbSwitch{}
	acc.Accessory = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbSwitch = service.NewLightbulb()
	acc.AddService(acc.LightbulbSwitch.Service)
	return &acc
}
