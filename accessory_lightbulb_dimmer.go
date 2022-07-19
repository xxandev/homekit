package homekit

import (
	"fmt"

	"github.com/brutella/hc/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessoryLightbulbDimmer struct
type AccessoryLightbulbDimmer struct {
	*accessory.Accessory
	LightbulbDimmer *haps.LightbulbDimmer
}

func (acc *AccessoryLightbulbDimmer) GetType() uint8 {
	return uint8(acc.Accessory.Type)
}

func (acc *AccessoryLightbulbDimmer) GetID() uint64 {
	return acc.Accessory.ID
}

func (acc *AccessoryLightbulbDimmer) GetSN() string {
	return acc.Accessory.Info.SerialNumber.GetValue()
}

func (acc *AccessoryLightbulbDimmer) GetName() string {
	return acc.Accessory.Info.Name.GetValue()
}

func (acc *AccessoryLightbulbDimmer) GetAccessory() *accessory.Accessory {
	return acc.Accessory
}

//NewAccessoryLightbulbDimmer return AccessoryLightbulbDimmer (args... are not used)
func NewAccessoryLightbulbDimmer(info accessory.Info, args ...interface{}) *AccessoryLightbulbDimmer {
	acc := AccessoryLightbulbDimmer{}
	acc.Accessory = accessory.New(info, accessory.TypeLightbulb)
	acc.LightbulbDimmer = haps.NewLightbulbDimmer()
	acc.AddService(acc.LightbulbDimmer.Service)
	return &acc
}

func (acc *AccessoryLightbulbDimmer) OnValuesRemoteUpdates(fn func()) {
	acc.LightbulbDimmer.On.OnValueRemoteUpdate(func(bool) { fn() })
	acc.LightbulbDimmer.Brightness.OnValueRemoteUpdate(func(int) { fn() })
}

func (acc *AccessoryLightbulbDimmer) OnExample() {
	acc.LightbulbDimmer.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%T - %s] remote update on: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
	acc.LightbulbDimmer.Brightness.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%T - %s] remote update brightness: %T - %v \n", acc, acc.Accessory.Info.SerialNumber.GetValue(), v, v)
	})
}
