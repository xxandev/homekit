package main

import (
	"fmt"
	"log"
	"time"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/xxandev/homekit"
)

func main() {
	// log.Debug.Enable()
	acc := homekit.NewAccessorySensorDioxide(accessory.Info{Name: "Dioxide", SerialNumber: "Ex-Dioxide", Model: "HAP-DXD", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		t := time.NewTicker(2 * time.Second)
		for {
			<-t.C
			acc.CarbonDioxideSensor.CarbonDioxideDetected.SetValue(1)
			fmt.Printf("acc sensor dioxide update current state: %T - %v \n", acc.CarbonDioxideSensor.CarbonDioxideDetected.GetValue(), acc.CarbonDioxideSensor.CarbonDioxideDetected.GetValue())
			<-t.C
			acc.CarbonDioxideSensor.CarbonDioxideDetected.SetValue(0)
			fmt.Printf("acc sensor dioxide update current state: %T - %v \n", acc.CarbonDioxideSensor.CarbonDioxideDetected.GetValue(), acc.CarbonDioxideSensor.CarbonDioxideDetected.GetValue())
		}
	}()
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
