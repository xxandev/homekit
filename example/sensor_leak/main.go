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
	acc := homekit.NewAccessorySensorLeak(accessory.Info{Name: "Leak", SerialNumber: "Ex-Leak", Model: "HAP-LEAK", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		t := time.NewTicker(2 * time.Second)
		for {
			<-t.C
			acc.LeakSensor.LeakDetected.SetValue(1)
			fmt.Printf("acc sensor leak update current state: %T - %v \n", acc.LeakSensor.LeakDetected.GetValue(), acc.LeakSensor.LeakDetected.GetValue())
			<-t.C
			acc.LeakSensor.LeakDetected.SetValue(0)
			fmt.Printf("acc sensor leak update current state: %T - %v \n", acc.LeakSensor.LeakDetected.GetValue(), acc.LeakSensor.LeakDetected.GetValue())
		}
	}()
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
