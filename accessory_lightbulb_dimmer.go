package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

type serviceLightbulbDimmer struct {
	*service.Service
	On         *characteristic.On
	Brightness *characteristic.Brightness
}

func newServiceLightbulbDimmer() *serviceLightbulbDimmer {
	svc := serviceLightbulbDimmer{}
	svc.Service = service.New(service.TypeLightbulb)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	svc.Brightness = characteristic.NewBrightness()
	svc.AddCharacteristic(svc.Brightness.Characteristic)

	return &svc
}

//AccessoryLightbulbDimmer struct
type AccessoryLightbulbDimmer struct {
	*accessory.Accessory
	LightbulbDimmer *serviceLightbulbDimmer
}

//NewAccessoryLightbulbDimmer return AccessoryLightbulbDimmer (args... are not used)
func NewAccessoryLightbulbDimmer(info accessory.Info, args ...interface{}) *AccessoryLightbulbDimmer {
	acc := AccessoryLightbulbDimmer{}
	acc.Accessory = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbDimmer = newServiceLightbulbDimmer()
	acc.AddService(acc.LightbulbDimmer.Service)
	return &acc
}
