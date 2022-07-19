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
	acc := homekit.NewAccessoryHumidifierDehumidifier(accessory.Info{Name: "Hum", SerialNumber: "Ex-Hum-Dehum", Model: "HAP-HM-DHM", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go acc.HumidifierDehumidifier.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc humidifier-dehumidifier remote update active: %T - %v \n", v, v)
	})
	go acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc humidifier-dehumidifier remote update target state: %T - %v \n", v, v)
	})
	go acc.HumidifierDehumidifier.RelativeHumidityDehumidifierThreshold.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("acc humidifier-dehumidifier remote update relative humidity dehumidifier threshold: %T - %v \n", v, v)
	})
	go acc.HumidifierDehumidifier.RelativeHumidityHumidifierThreshold.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("acc humidifier-dehumidifier remote update relative humidity humidifier threshold: %T - %v \n", v, v)
	})
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
