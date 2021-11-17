package homekit

import (
	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessoryLightbulbDimmer struct
type AccessoryLightbulbDimmer struct {
	*accessory.Accessory
	LightbulbDimmer *haps.LightbulbDimmer
}

//NewAccessoryLightbulbDimmer return AccessoryLightbulbDimmer (args... are not used)
func NewAccessoryLightbulbDimmer(info accessory.Info, args ...interface{}) *AccessoryLightbulbDimmer {
	acc := AccessoryLightbulbDimmer{}
	acc.Accessory = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbDimmer = haps.NewLightbulbDimmer()
	acc.AddService(acc.LightbulbDimmer.Service)
	return &acc
}

func (acc *AccessoryLightbulbDimmer) OnValuesRemoteUpdates(fn func()) {
	acc.LightbulbDimmer.On.OnValueRemoteUpdate(func(_ bool) { fn() })
	acc.LightbulbDimmer.Brightness.OnValueRemoteUpdate(func(_ int) { fn() })
}
