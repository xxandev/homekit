package main

import (
	"fmt"
	"log"

	"github.com/alpr777/homekit"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

func main() {
	// log.Debug.Enable()
	acc := homekit.NewAccessoryGate(accessory.Info{Name: "Gate", SerialNumber: "Ex-Gate", Model: "HAP-GT", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go acc.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(v int) {
		acc.GarageDoorOpener.CurrentDoorState.SetValue(v)
		fmt.Printf("acc gate update: target state %T - %v, current state %T - %v\n", v, v, acc.GarageDoorOpener.CurrentDoorState.GetValue(), acc.GarageDoorOpener.CurrentDoorState.GetValue())
	})
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
