package homekit

import (
	"fmt"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySwitch struct
type AccessorySwitch struct {
	*accessory.A
	Switch *service.Switch
}

func (acc *AccessorySwitch) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySwitch) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySwitch) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySwitch) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySwitch) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySwitch) GetAccessory() *accessory.A {
	return acc.A
}

// NewAccessorySwitch returns AccessorySwitch (args... are not used)
func NewAccessorySwitch(info accessory.Info, args ...interface{}) *AccessorySwitch {
	acc := AccessorySwitch{}
	acc.A = accessory.New(info, accessory.TypeSwitch)
	acc.Switch = service.NewSwitch()
	acc.AddS(acc.Switch.S)
	return &acc
}

func (acc *AccessorySwitch) OnValuesRemoteUpdates(fn func()) {
	acc.Switch.On.OnValueRemoteUpdate(func(bool) { fn() })
}

func (acc *AccessorySwitch) OnExample() {
	go func() {
		for range time.Tick(30 * time.Second) {
			acc.Switch.On.SetValue(!acc.Switch.On.Value())
			fmt.Printf("[%T - %v - %v] acc switch update on: %T - %v\n",
				acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.Switch.On.Value(), acc.Switch.On.Value())
		}
	}()
	acc.Switch.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%T - %v - %v] remote update on: %T - %v \n",
			acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v, v)
	})
}
