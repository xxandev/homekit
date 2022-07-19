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
	acc := homekit.NewAccessoryLightbulbDimmer(accessory.Info{Name: "Lightbulb", SerialNumber: "Ex-Lb-Dm", Model: "HAP-LB-DM", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		t := time.NewTicker(30 * time.Second)
		for range t.C {
			acc.LightbulbDimmer.On.SetValue(!acc.LightbulbDimmer.On.GetValue())
			fmt.Printf("acc lightbulb update on: %T - %v \n", acc.LightbulbDimmer.On.GetValue(), acc.LightbulbDimmer.On.GetValue())
		}
	}()
	go acc.LightbulbDimmer.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("acc lightbulb remote update on: %T - %v \n", v, v)
	})
	go acc.LightbulbDimmer.Brightness.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc lightbulb remote update brightness: %T - %v \n", v, v)
	})
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
