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
	acc := homekit.NewAccessorySensorMotion(accessory.Info{Name: "Contact", SerialNumber: "Ex-Contact", Model: "HAP-CNT", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		t := time.NewTicker(10 * time.Second)
		for range t.C {
			acc.MotionSensor.MotionDetected.SetValue(!acc.MotionSensor.MotionDetected.GetValue())
			fmt.Printf("acc sensor motion update current state: %T - %v \n", acc.MotionSensor.MotionDetected.GetValue(), acc.MotionSensor.MotionDetected.GetValue())
		}
	}()
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
