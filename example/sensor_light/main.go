package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/xxandev/homekit"
)

func main() {
	// log.Debug.Enable()
	acc := homekit.NewAccessorySensorLight(accessory.Info{Name: "Light", SerialNumber: "Ex-Light", Model: "HAP-LGHT", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		t := time.NewTicker(2 * time.Second)
		for range t.C {
			acc.LightSensor.CurrentAmbientLightLevel.SetValue(rand.Float64())
			fmt.Printf("acc sensor light update current state: %T - %v \n", acc.LightSensor.CurrentAmbientLightLevel.GetValue(), acc.LightSensor.CurrentAmbientLightLevel.GetValue())
		}
	}()
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
