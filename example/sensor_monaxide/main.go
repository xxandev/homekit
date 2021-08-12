package main

import (
	"fmt"
	"log"
	"time"

	"github.com/alpr777/homekit"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

func main() {
	// log.Debug.Enable()
	acc := homekit.NewAccessorySensorMonoxide(accessory.Info{Name: "Monoxide", SerialNumber: "Ex-Monoxide", Model: "HAP-MXD", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		t := time.NewTicker(2 * time.Second)
		for {
			<-t.C
			acc.CarbonMonoxideSensor.CarbonMonoxideDetected.SetValue(1)
			fmt.Printf("acc sensor monoxide update current state: %T - %v \n", acc.CarbonMonoxideSensor.CarbonMonoxideDetected.GetValue(), acc.CarbonMonoxideSensor.CarbonMonoxideDetected.GetValue())
			<-t.C
			acc.CarbonMonoxideSensor.CarbonMonoxideDetected.SetValue(0)
			fmt.Printf("acc sensor monoxide update current state: %T - %v \n", acc.CarbonMonoxideSensor.CarbonMonoxideDetected.GetValue(), acc.CarbonMonoxideSensor.CarbonMonoxideDetected.GetValue())
		}
	}()
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
