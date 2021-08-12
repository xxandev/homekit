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
	acc := homekit.NewAccessoryWindowCovering(accessory.Info{Name: "blind", SerialNumber: "Ex-Win-Cov", Model: "HAP-WIN-COV", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go acc.WindowCovering.TargetPosition.OnValueRemoteUpdate(func(state int) {
		fmt.Printf("acc window covering remote update target position: %T - %v \n", state, state)
		acc.WindowCovering.CurrentPosition.SetValue(state)
		fmt.Printf("acc window covering update current position: %T - %v \n", acc.WindowCovering.CurrentPosition.GetValue(), acc.WindowCovering.CurrentPosition.GetValue())
	})
	fmt.Println("homekit accessory transport start [", "/", acc.Accessory.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
