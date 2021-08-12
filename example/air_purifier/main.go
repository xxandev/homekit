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
	acc := homekit.NewAccessoryAirPurifier(accessory.Info{Name: "Air purifier", SerialNumber: "Ex-Air-Pru", Model: "HAP-AP", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision}, 0, 0, 100, 1)
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go acc.AirPurifier.Active.OnValueRemoteUpdate(func(v int) {
		if v > 0 {
			acc.AirPurifier.CurrentAirPurifierState.SetValue(2)
		} else {
			acc.AirPurifier.CurrentAirPurifierState.SetValue(0)
		}
		fmt.Printf("acc air purifier remote update: update active %T - %v, current state %T - %v \n", v, v, acc.AirPurifier.CurrentAirPurifierState.GetValue(), acc.AirPurifier.CurrentAirPurifierState.GetValue())
	})
	go acc.AirPurifier.RotationSpeed.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("acc air purifier remote update rotation speed: %T - %v \n", v, v)
	})
	go acc.AirPurifier.TargetAirPurifierState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc air purifier remote update target state: %T - %v \n", v, v)
	})
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
