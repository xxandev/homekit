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
	accessorySn   string = "ExmplAP"
	accessoryPin  string = "11112222"
)

func main() {
	// runtime.GOMAXPROCS(4)
	// log.Debug.Enable()
	acc := homekit.NewAccessoryAirPurifier(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "EXAMPLE", Model: "ACC-TEST", FirmwareRevision: "1.2"})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: accessoryPin}, acc.Accessory)
	if err != nil {
		fmt.Println("accessory [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]", "error create transport:", err)
		os.Exit(1)
	}
	go acc.AirPurifier.Active.OnValueRemoteUpdate(func(active int) { fmt.Printf("acc remote update active: %T - %v \n", active, active) })
	go acc.AirPurifier.TargetAirPurifierState.OnValueRemoteUpdate(func(state int) {
		fmt.Printf("acc remote update target state: %T - %v \n", state, state)
		acc.AirPurifier.CurrentAirPurifierState.SetValue(state)
		fmt.Printf("acc update current state: %T - %v \n", acc.AirPurifier.CurrentAirPurifierState.GetValue(), acc.AirPurifier.CurrentAirPurifierState.GetValue())
	})
	go acc.AirPurifier.RotationSpeed.OnValueRemoteUpdate(func(speed float64) { fmt.Printf("acc remote update rotation speed: %T - %v \n", speed, speed) })
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
