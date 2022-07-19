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
	acc := homekit.NewAccessoryWifiRouter(accessory.Info{Name: "Router", SerialNumber: "Ex-Rour", Model: "HAP-ROUTER", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	fmt.Println("homekit accessory transport start [", "/", acc.Accessory.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
