package homekit

import (
	haps "github.com/alpr777/homekit/hap-service"
	"github.com/brutella/hc/accessory"
)

//AccessoryFanV2Multifunc struct
type AccessoryFan2Speed struct {
	*accessory.Accessory
	Fan2 *haps.Fan2RS
}

//NewAccessoryFanV2Multifunc return AccessoryFanV2Multifunc (args... are not used)
func NewAccessoryFan2Speed(info accessory.Info, args ...interface{}) *AccessoryFan2Speed {
	acc := AccessoryFan2Speed{}
	acc.Accessory = accessory.New(info, accessory.TypeFan)
	acc.Fan2 = haps.NewFanV2Multifunc()
	acc.AddService(acc.Fan2.Service)
	return &acc
}

func (acc *AccessoryFan2Speed) OnValuesRemoteUpdates(fn func()) {
	acc.Fan2.Active.OnValueRemoteUpdate(func(int) { fn() })
	acc.Fan2.RotationSpeed.OnValueRemoteUpdate(func(float64) { fn() })
}
