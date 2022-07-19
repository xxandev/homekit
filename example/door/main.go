package main

import (
	"fmt"
	"log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/xxandev/homekit"
)

func main() {
	// log.Debug.Enable()
	acc := homekit.NewAccessoryDoor(accessory.Info{Name: "Door", SerialNumber: "Ex-Door", Model: "HAP-DR", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go acc.Door.TargetPosition.OnValueRemoteUpdate(func(v int) {
		acc.Door.CurrentPosition.SetValue(v)
		fmt.Printf("acc door update: target position %T - %v, current position %T - %v\n", v, v, acc.Door.CurrentPosition.GetValue(), acc.Door.CurrentPosition.GetValue())
	})
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
