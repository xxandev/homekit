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
	acc := homekit.NewAccessorySwitch(accessory.Info{Name: "Switch", SerialNumber: "Ex-Switch", Model: "HAP-SW", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		t := time.NewTicker(30 * time.Second)
		for range t.C {
			acc.Switch.On.SetValue(!acc.Switch.On.GetValue())
			fmt.Printf("acc switch update on: %T - %v \n", acc.Switch.On.GetValue(), acc.Switch.On.GetValue())
		}
	}()
	go acc.Switch.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("acc switch remote update on: %T - %v \n", v, v)
	})
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
