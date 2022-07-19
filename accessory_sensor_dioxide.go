package homekit

import (
	"fmt"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/service"
)

//AccessorySensorDioxide struct
type AccessorySensorDioxide struct {
	*accessory.A
	CarbonDioxideSensor *service.CarbonDioxideSensor
}

func (acc *AccessorySensorDioxide) GetType() uint8 {
	return acc.A.Type
}

func (acc *AccessorySensorDioxide) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySensorDioxide) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySensorDioxide) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySensorDioxide) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySensorDioxide) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySensorLight return AccessorySensorLight (args... are not used)
func NewAccessorySensorDioxide(info accessory.Info, args ...interface{}) *AccessorySensorDioxide {
	acc := AccessorySensorDioxide{}
	acc.A = accessory.New(info, accessory.TypeSensor)
	acc.CarbonDioxideSensor = service.NewCarbonDioxideSensor()
	acc.AddS(acc.CarbonDioxideSensor.S)
	return &acc
}

func (acc *AccessorySensorDioxide) OnValuesRemoteUpdates(fn func()) {}
func (acc *AccessorySensorDioxide) OnExample() {
	go func() {
		for range time.Tick(30 * time.Second) {
			if acc.CarbonDioxideSensor.CarbonDioxideDetected.Value() > 0 {
				acc.CarbonDioxideSensor.CarbonDioxideDetected.SetValue(0)
			} else {
				acc.CarbonDioxideSensor.CarbonDioxideDetected.SetValue(1)
			}
			fmt.Printf("[%[1]T - %[2]v - %[3]v] update current state: %[4]T - %[4]v \n", acc, acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value(), acc.CarbonDioxideSensor.CarbonDioxideDetected.Value())
		}
	}()
}
