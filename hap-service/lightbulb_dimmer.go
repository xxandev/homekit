package haps

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//LightbulbDimmer (+On, Brightness)
type LightbulbDimmer struct {
	*service.Service
	On         *characteristic.On
	Brightness *characteristic.Brightness
}

//NewLightbulbDimmer return *LightbulbDimmer
func NewLightbulbDimmer() *LightbulbDimmer {
	svc := LightbulbDimmer{}
	svc.Service = service.New(service.TypeLightbulb)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	svc.Brightness = characteristic.NewBrightness()
	svc.AddCharacteristic(svc.Brightness.Characteristic)

	return &svc
}
