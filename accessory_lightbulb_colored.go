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

func (acc *AccessoryLightbulbColored) OnValuesRemoteUpdates(fn func()) {
	acc.LightbulbColored.On.OnValueRemoteUpdate(func(bool) { fn() })
	acc.LightbulbColored.Brightness.OnValueRemoteUpdate(func(int) { fn() })
	acc.LightbulbColored.Saturation.OnValueRemoteUpdate(func(float64) { fn() })
	acc.LightbulbColored.Hue.OnValueRemoteUpdate(func(float64) { fn() })
}
