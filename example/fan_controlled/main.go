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
	acc := homekit.NewAccessoryFanControlled(accessory.Info{Name: "Fan", SerialNumber: "EX-Fan", Model: "HAP-FN-CTRL", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		t := time.NewTicker(30 * time.Second)
		for range t.C {
			acc.Fan.On.SetValue(!acc.Fan.On.GetValue())
			fmt.Printf("acc fan update on: %T - %v \n", acc.Fan.On.GetValue(), acc.Fan.On.GetValue())
		}
	}()
	go acc.Fan.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("acc fan remote update on: %T - %v \n", v, v)
	})
	go acc.Fan.RotationSpeed.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("acc fan remote update rotation speed: %T - %v \n", v, v)
	})
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
