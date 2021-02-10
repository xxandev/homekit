package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//ServiceLightbulbDimmer -
type ServiceLightbulbDimmer struct {
	*service.Service
	On         *characteristic.On
	Brightness *characteristic.Brightness
}

//NewServiceLightbulbDimmer -
func NewServiceLightbulbDimmer() *ServiceLightbulbDimmer {
	svc := ServiceLightbulbDimmer{}
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
	LightbulbDimmer *ServiceLightbulbDimmer
}

//NewAccessoryLightbulbDimmer return AccessoryLightbulbDimmer (args... are not used)
func NewAccessoryLightbulbDimmer(info accessory.Info, args ...interface{}) *AccessoryLightbulbDimmer {
	acc := AccessoryLightbulbDimmer{}
	acc.Accessory = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbDimmer = NewServiceLightbulbDimmer()
	acc.AddService(acc.LightbulbDimmer.Service)
	return &acc
}
