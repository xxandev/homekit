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
	accessoryName string = "light"
	accessorySn   string = "ExmplLB"
	accessoryPin  string = "11112222"
)

func main() {
	// runtime.GOMAXPROCS(4)
	// log.Debug.Enable()
	acc := homekit.NewAccessoryLightbulbColored(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "EXAMPLE", Model: "ACC-TEST", FirmwareRevision: "1.2"})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: accessoryPin}, acc.Accessory)
	if err != nil {
		fmt.Println("accessory [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]", "error create transport:", err)
		os.Exit(1)
	}
	go func() {
		tickerUpdateState := time.NewTicker(30 * time.Second)
		for {
			select {
			case <-tickerUpdateState.C:
				acc.LightbulbColored.On.SetValue(!acc.LightbulbColored.On.GetValue())
				fmt.Printf("acc lightbulb update on: %T - %v \n", acc.LightbulbColored.On.GetValue(), acc.LightbulbColored.On.GetValue())
				continue
			}
		}
	}()
	go acc.LightbulbColored.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("acc lightbulb remote update on: %T - %v \n", v, v)
	})
	go acc.LightbulbColored.Brightness.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc lightbulb remote update brightness: %T - %v \n", v, v)
	})
	go acc.LightbulbColored.Saturation.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("acc lightbulb remote update saturation: %T - %v \n", v, v)
	})
	go acc.LightbulbColored.Hue.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("acc lightbulb remote update hue: %T - %v \n", v, v)
	})
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
