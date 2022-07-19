package haps

import (
	"github.com/brutella/hap/characteristic"
	"github.com/brutella/hap/service"
)

//LightbulbDimmer
//	◈ On
//	◇ Brightness
type LightbulbDimmer struct {
	*service.S
	On         *characteristic.On
	Brightness *characteristic.Brightness
}

//NewLightbulbDimmer return *LightbulbDimmer
func NewLightbulbDimmer() *LightbulbDimmer {
	svc := LightbulbDimmer{}
	svc.S = service.New(service.TypeLightbulb)

	svc.On = characteristic.NewOn()
	svc.AddC(svc.On.C)

	svc.Brightness = characteristic.NewBrightness()
	svc.AddC(svc.Brightness.C)

	return &svc
}
