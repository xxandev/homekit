package main

import (
	"fmt"
	"os"

	"github.com/alpr777/homekit"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

const (
	accessoryName string = "switch"
	accessorySn   string = "ExmplSW"
	accessoryPin  string = "11112222"
)

func main() {
	// runtime.GOMAXPROCS(4)
	// log.Debug.Enable()
	acc := homekit.NewAccessoryHumidifierDehumidifier(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "EXAMPLE", Model: "ACC-TEST", FirmwareRevision: "1.2"})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: accessoryPin}, acc.Accessory)
	if err != nil {
		fmt.Println("accessory [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]", "error create transport:", err)
		os.Exit(1)
	}
	go acc.HumidifierDehumidifier.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc humidifier-dehumidifier remote update active: %T - %v \n", v, v)
	})
	go acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc humidifier-dehumidifier remote update target state: %T - %v \n", v, v)
	})
	go acc.HumidifierDehumidifier.TargetRelativeHumidity.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("acc humidifier-dehumidifier remote update target relative humidity: %T - %v \n", v, v)
	})
	go acc.HumidifierDehumidifier.RelativeHumidityDehumidifierThreshold.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("acc humidifier-dehumidifier remote update relative humidity dehumidifier threshold: %T - %v \n", v, v)
	})
	go acc.HumidifierDehumidifier.RelativeHumidityHumidifierThreshold.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("acc humidifier-dehumidifier remote update relative humidity humidifier threshold: %T - %v \n", v, v)
	})
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
