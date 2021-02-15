package main

import (
	"fmt"
	"os"

	"github.com/alpr777/homekit"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

const (
	accessoryName string = "air purifier"
	accessorySn   string = "ExmplAirPur"
	accessoryPin  string = "11112222"
)

func main() {
	// runtime.GOMAXPROCS(4)
	// log.Debug.Enable()
	acc := homekit.NewAccessoryAirPurifier(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "alpr777", Model: "ACC-TEST", FirmwareRevision: "1.2"}, 0, 0, 100, 1)
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: accessoryPin}, acc.Accessory)
	if err != nil {
		fmt.Println("accessory [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]", "error create transport:", err)
		os.Exit(1)
	}
	go acc.AirPurifier.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc air purifier remote update active: %T - %v \n", v, v)
		if v > 0 {
			acc.AirPurifier.CurrentAirPurifierState.SetValue(2)
		} else {
			acc.AirPurifier.CurrentAirPurifierState.SetValue(0)
		}
		fmt.Printf("acc air purifier update current state: %T - %v \n", acc.AirPurifier.CurrentAirPurifierState.GetValue(), acc.AirPurifier.CurrentAirPurifierState.GetValue())
	})
	go acc.AirPurifier.RotationSpeed.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("acc air purifier remote update rotation speed: %T - %v \n", v, v)
	})
	go acc.AirPurifier.TargetAirPurifierState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc air purifier remote update target state: %T - %v \n", v, v)
	})
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
