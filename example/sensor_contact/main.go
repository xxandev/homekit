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
	acc := homekit.NewAccessorySensorContact(accessory.Info{Name: "Contact", SerialNumber: "Ex-Contact", Model: "HAP-CNT", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		t := time.NewTicker(2 * time.Second)
		for {
			<-t.C
			acc.ContactSensor.ContactSensorState.SetValue(1)
			fmt.Printf("acc sensor contact update current state: %T - %v \n", acc.ContactSensor.ContactSensorState.GetValue(), acc.ContactSensor.ContactSensorState.GetValue())
			<-t.C
			acc.ContactSensor.ContactSensorState.SetValue(0)
			fmt.Printf("acc sensor contact update current state: %T - %v \n", acc.ContactSensor.ContactSensorState.GetValue(), acc.ContactSensor.ContactSensorState.GetValue())
		}
	}()
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
