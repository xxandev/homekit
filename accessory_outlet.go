package homekit

import (
	"fmt"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessoryOutlet struct
type AccessoryOutlet struct {
	*accessory.A
	Outlet *service.Outlet
}

func (acc *AccessoryOutlet) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryOutlet) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryOutlet) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryOutlet) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryOutlet) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryOutlet) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryOutlet return AccessoryOutlet (args... are not used)
func NewAccessoryOutlet(info accessory.Info, args ...interface{}) *AccessoryOutlet {
	acc := AccessoryOutlet{}
	acc.A = accessory.New(info, accessory.TypeOutlet)
	acc.Outlet = service.NewOutlet()
	acc.Outlet.OutletInUse.SetValue(true)
	acc.AddS(acc.Outlet.S)
	return &acc
}

func (acc *AccessoryOutlet) OnValuesRemoteUpdates(fn func()) {
	acc.Outlet.On.OnValueRemoteUpdate(func(bool) { fn() })
	// acc.Outlet.OutletInUse.OnValueRemoteUpdate(func(bool) { fn() })
}

func (acc *AccessoryOutlet) OnExample() {
	go func() {
		for range time.Tick(30 * time.Second) {
			acc.Outlet.On.SetValue(!acc.Outlet.On.Value())
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update on: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.Outlet.On.Value())
		}
	}()
	acc.Outlet.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update on: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
