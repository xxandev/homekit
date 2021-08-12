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
	acc := homekit.NewAccessoryLock(accessory.Info{Name: "Lock", SerialNumber: "Ex-Lock", Model: "HAP-LCK", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go acc.LockMechanism.LockTargetState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc door lock remote update target position: %T - %v \n", v, v)
		acc.LockMechanism.LockCurrentState.SetValue(v)
		fmt.Printf("acc door lock update current position: %T - %v \n", acc.LockMechanism.LockCurrentState.GetValue(), acc.LockMechanism.LockCurrentState.GetValue())
	})
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
