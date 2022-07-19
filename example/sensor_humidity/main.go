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
	acc := homekit.NewAccessorySensorHumidity(accessory.Info{Name: "Hum", SerialNumber: "Ex-Hum", Model: "HAP-HUM", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		t := time.NewTicker(2 * time.Second)
		for range t.C {
			acc.HumiditySensor.CurrentRelativeHumidity.SetValue(rand.Float64())
			fmt.Printf("acc sensor humidity update current state: %T - %v \n", acc.HumiditySensor.CurrentRelativeHumidity.GetValue(), acc.HumiditySensor.CurrentRelativeHumidity.GetValue())
		}
	}()
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
