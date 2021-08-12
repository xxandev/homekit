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
	acc := homekit.NewAccessorySensorAirQuality(accessory.Info{Name: "Air quality", SerialNumber: "Ex-Snr-AQ", Model: "HAP-SNR-QLT", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		t := time.NewTicker(10 * time.Second)
		for range t.C {
			if acc.AirQualitySensor.AirQuality.GetValue() >= 5 {
				acc.AirQualitySensor.AirQuality.SetValue(0)
			} else {
				acc.AirQualitySensor.AirQuality.SetValue(acc.AirQualitySensor.AirQuality.GetValue() + 1)
			}
			fmt.Printf("acc sensor air quality update current state: %T - %v \n", acc.AirQualitySensor.AirQuality.GetValue(), acc.AirQualitySensor.AirQuality.GetValue())
		}
	}()
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
