package main

import (
	"fmt"
	"log"

	"github.com/alpr777/homekit"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

var conteiner = accessory.NewContainer()

func init() {

	for n := 1; n <= 25; n++ { // MAX 150 ACCESSORIES
		acc := homekit.NewAccessorySwitch(accessory.Info{ /* ID: uint64(n),*/ Name: fmt.Sprintf("Switch%v", n), SerialNumber: fmt.Sprintf("sw-hap-%v", n), Model: "HAP-SW", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
		go acc.Switch.On.OnValueRemoteUpdate(func(v bool) {
			fmt.Printf("acc switch remote update on: %T - %v \n", v, v)
		})
		conteiner.AddAccessory(acc.Accessory)
	}
}

func main() {
	// log.Debug.Enable()
	bridge := accessory.NewBridge(accessory.Info{ /* ID: uint64(n),*/ Name: "Bridge", SerialNumber: "Ex-Brg", Model: "HAP-BRG", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + bridge.Info.SerialNumber.GetValue(), Pin: "11223344"}, bridge.Accessory, conteiner.Accessories...)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", bridge.Info.SerialNumber.GetValue(), bridge.Info.Name.GetValue(), err)
	}
	fmt.Printf("[ %v / %v ] accessories transport start\n", bridge.Info.SerialNumber.GetValue(), bridge.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
