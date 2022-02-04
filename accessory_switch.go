package homekit

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

//AccessorySwitch struct
type AccessorySwitch struct {
	*accessory.Accessory
	Switch *service.Switch
}

// NewAccessorySwitch returns AccessorySwitch (args... are not used)
func NewAccessorySwitch(info accessory.Info, args ...interface{}) *AccessorySwitch {
	acc := AccessorySwitch{}
	acc.Accessory = accessory.New(info, accessory.TypeSwitch)
	acc.Switch = service.NewSwitch()
	acc.AddService(acc.Switch.Service)
	return &acc
}

func (acc *AccessorySwitch) GetType() uint8                     { return uint8(acc.Accessory.Type) }
func (acc *AccessorySwitch) GetID() uint64                      { return acc.Accessory.ID }
func (acc *AccessorySwitch) GetSN() string                      { return acc.Accessory.Info.SerialNumber.GetValue() }
func (acc *AccessorySwitch) GetName() string                    { return acc.Accessory.Info.SerialNumber.GetValue() }
func (acc *AccessorySwitch) GetAccessory() *accessory.Accessory { return acc.Accessory }

func (acc *AccessorySwitch) OnValuesRemoteUpdates(fn func()) {
	acc.Switch.On.OnValueRemoteUpdate(func(bool) { fn() })
}
