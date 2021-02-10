package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

//ServiceSprinkler -
type ServiceSprinkler struct {
	*service.Service
	Active *characteristic.Active
}

//NewServiceSprinkler -
func NewServiceSprinkler() *ServiceSprinkler {
	svc := ServiceSprinkler{}
	svc.Service = service.New(service.TypeAirPurifier)

	svc.Active = characteristic.NewActive()
	svc.AddCharacteristic(svc.Active.Characteristic)

	return &svc
}

//AccessorySprinkler struct
type AccessorySprinkler struct {
	*accessory.Accessory
	Valve      *service.Valve
	Sprinklers *ServiceSprinkler
}

//NewAccessorySprinkler return AccessoryFaucet (args... are not used)
func NewAccessorySprinkler(info accessory.Info, args ...interface{}) *AccessorySprinkler {
	acc := AccessorySprinkler{}
	acc.Accessory = accessory.New(info, accessory.TypeSprinklers)
	acc.Sprinklers = NewServiceSprinkler()
	acc.AddService(acc.Sprinklers.Service)
	return &acc
}
