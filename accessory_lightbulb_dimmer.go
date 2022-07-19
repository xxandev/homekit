package homekit

import (
	"fmt"
	"time"

	"github.com/brutella/hap/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryLightbulbDimmer struct
type AccessoryLightbulbDimmer struct {
	*accessory.A
	LightbulbDimmer *haps.LightbulbDimmer
}

func (acc *AccessoryLightbulbDimmer) GetType() byte {
	return acc.A.Type
}

func (acc *AccessoryLightbulbDimmer) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessoryLightbulbDimmer) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessoryLightbulbDimmer) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessoryLightbulbDimmer) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessoryLightbulbDimmer) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessoryLightbulbDimmer return AccessoryLightbulbDimmer (args... are not used)
func NewAccessoryLightbulbDimmer(info accessory.Info, args ...interface{}) *AccessoryLightbulbDimmer {
	acc := AccessoryLightbulbDimmer{}
	acc.A = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbDimmer = haps.NewLightbulbDimmer()
	acc.AddS(acc.LightbulbDimmer.S)
	return &acc
}

func (acc *AccessoryLightbulbDimmer) OnValuesRemoteUpdates(fn func()) {
	acc.LightbulbDimmer.On.OnValueRemoteUpdate(func(bool) { fn() })
	acc.LightbulbDimmer.Brightness.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryLightbulbDimmer) OnExample() {
	go func() {
		for range time.Tick(30 * time.Second) {
			acc.LightbulbDimmer.On.SetValue(!acc.LightbulbDimmer.On.Value())
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update on: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.LightbulbDimmer.On.Value())
		}
	}()
	acc.LightbulbDimmer.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update on: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
	acc.LightbulbDimmer.Brightness.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v - %[3]v] remote update brightness: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), v)
	})
}
