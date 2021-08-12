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
	acc := homekit.NewAccessoryLightbulbColored(accessory.Info{Name: "Lightbulb", SerialNumber: "Ex-Lb-Clr", Model: "HAP-LB-CLR", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		t := time.NewTicker(30 * time.Second)
		for range t.C {
			acc.LightbulbColored.On.SetValue(!acc.LightbulbColored.On.GetValue())
			fmt.Printf("acc lightbulb colored update on: %T - %v \n", acc.LightbulbColored.On.GetValue(), acc.LightbulbColored.On.GetValue())
		}
	}()
	go acc.LightbulbColored.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("acc lightbulb colored remote update on: %T - %v \n", v, v)
	})
	go acc.LightbulbColored.Brightness.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc lightbulb colored remote update brightness: %T - %v \n", v, v)
	})
	go acc.LightbulbColored.Saturation.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("acc lightbulb colored remote update saturation: %T - %v \n", v, v)
	})
	go acc.LightbulbColored.Hue.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("acc lightbulb colored remote update hue: %T - %v \n", v, v)
	})
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
