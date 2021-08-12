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
	acc := homekit.NewAccessoryWindow(accessory.Info{Name: "Window", SerialNumber: "Ex-Window", Model: "HAP-WIN", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go acc.Window.TargetPosition.OnValueRemoteUpdate(func(state int) {
		fmt.Printf("acc window remote update target position: %T - %v \n", state, state)
		acc.Window.CurrentPosition.SetValue(state)
		fmt.Printf("acc window update current position: %T - %v \n", acc.Window.CurrentPosition.GetValue(), acc.Window.CurrentPosition.GetValue())
	})
	fmt.Println("homekit accessory transport start [", "/", acc.Accessory.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
