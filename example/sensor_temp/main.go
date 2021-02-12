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
	accessoryName string = "temp"
	accessorySn   string = "ExmplST"
	accessoryPin  string = "11112222"
)

func main() {
	// runtime.GOMAXPROCS(4)
	// log.Debug.Enable()
	acc := homekit.NewAccessorySensorTemperature(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "alpr777", Model: "ACC-TEST", FirmwareRevision: "1.2"})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: accessoryPin}, acc.Accessory)
	if err != nil {
		fmt.Println("accessory [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]", "error create transport:", err)
		os.Exit(1)
	}
	go func() {
		tickerUpdateTemp := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-tickerUpdateTemp.C:
				acc.TempSensor.CurrentTemperature.SetValue(float64(time.Now().Second()-30) + float64(time.Now().Second()+40)/100)
				fmt.Printf("acc sensor temp update current state: %T - %v \n", acc.TempSensor.CurrentTemperature.GetValue(), acc.TempSensor.CurrentTemperature.GetValue())
				continue
			}
		}
	}()
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
