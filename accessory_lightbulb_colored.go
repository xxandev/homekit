package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessoryLightbulbColored struct
type AccessoryLightbulbColored struct {
	*accessory.Accessory
	LightbulbColored *service.ColoredLightbulb
}

//NewAccessoryLightbulbColored return AccessoryLightbulbColored (args... are not used)
func NewAccessoryLightbulbColored(info accessory.Info, args ...interface{}) *AccessoryLightbulbColored {
	acc := AccessoryLightbulbColored{}
	acc.Accessory = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbColored = service.NewColoredLightbulb()
	acc.AddService(acc.LightbulbColored.Service)
	return &acc
}
