package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alpr777/homekit"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

const (
	accessoryName string = "sensor"
	accessorySn   string = "ExmplSAQ"
	accessoryPin  string = "11112222"
)

func main() {
	// runtime.GOMAXPROCS(4)
	// log.Debug.Enable()
	acc := homekit.NewAccessorySensorAirQuality(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "alpr777", Model: "ACC-TEST", FirmwareRevision: "1.2"})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: accessoryPin}, acc.Accessory)
	if err != nil {
		fmt.Println("accessory [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]", "error create transport:", err)
		os.Exit(1)
	}
	go func() {
		tickerUpdateTemp := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-tickerUpdateTemp.C:
				if acc.AirQualitySensor.AirQuality.GetValue() >= 5 {
					acc.AirQualitySensor.AirQuality.SetValue(0)
				} else {
					acc.AirQualitySensor.AirQuality.SetValue(acc.AirQualitySensor.AirQuality.GetValue() + 1)
				}
				fmt.Printf("acc sensor temp update current state: %T - %v \n", acc.AirQualitySensor.AirQuality.GetValue(), acc.AirQualitySensor.AirQuality.GetValue())
				continue
			}
		}
	}()
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
