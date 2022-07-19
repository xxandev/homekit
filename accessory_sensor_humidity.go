package homekit

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorHumidity struct
type AccessorySensorHumidity struct {
	*accessory.A
	HumiditySensor *service.HumiditySensor
}

func (acc *AccessorySensorHumidity) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySensorHumidity) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorHumidity) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorHumidity) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorHumidity) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorHumidity) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorHumidity returns AccessorySensorHumidity (args... are not used)
func NewAccessorySensorHumidity(info accessory.Info, args ...interface{}) *AccessorySensorHumidity {
	acc := AccessorySensorHumidity{}
	acc.A = accessory.New(info, accessory.TypeThermostat)
	acc.HumiditySensor = service.NewHumiditySensor()
	acc.AddS(acc.HumiditySensor.S)
	return &acc
}

func (acc *AccessorySensorHumidity) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorHumidity) OnExample() {
	go func() {
		for range time.Tick(10 * time.Second) {
			acc.HumiditySensor.CurrentRelativeHumidity.SetValue(rand.Float64())
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update current state: %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.HumiditySensor.CurrentRelativeHumidity.Value())
		}
	}()
}
