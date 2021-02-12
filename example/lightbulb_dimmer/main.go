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
	accessoryName string = "lightbulb"
	accessorySn   string = "ExmplLBD"
	accessoryPin  string = "11112222"
)

func main() {
	// runtime.GOMAXPROCS(4)
	// log.Debug.Enable()
	acc := homekit.NewAccessoryLightbulbDimmer(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "alpr777", Model: "ACC-TEST", FirmwareRevision: "1.2"})
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
				acc.LightbulbDimmer.On.SetValue(!acc.LightbulbDimmer.On.GetValue())
				fmt.Printf("acc lightbulb update on: %T - %v \n", acc.LightbulbDimmer.On.GetValue(), acc.LightbulbDimmer.On.GetValue())
				continue
			}
		}
	}()
	go acc.LightbulbDimmer.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("acc lightbulb remote update on: %T - %v \n", v, v)
	})
	go acc.LightbulbDimmer.Brightness.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc lightbulb remote update brightness: %T - %v \n", v, v)
	})
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
