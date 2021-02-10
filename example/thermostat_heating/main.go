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
	accessoryName string = "thermostat"
	accessorySn   string = "ExmplTRMHTN"
	accessoryPin  string = "11112222"
)

func invState(arg int) int {
	if arg > 0 {
		return 0
	}
	return 1
}

func main() {
	// runtime.GOMAXPROCS(4)
	// log.Debug.Enable()
	acc := homekit.NewAccessoryThermostat(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "EXAMPLE", Model: "ACC-TEST", FirmwareRevision: "1.2"}, 0, 0, 1, 1)
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: accessoryPin}, acc.Accessory)
	if err != nil {
		fmt.Println("accessory [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]", "error create transport:", err)
		os.Exit(1)
	}
	go func() {
		tickerUpdateState := time.NewTicker(10 * time.Second)
		tickerUpdateTemp := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-tickerUpdateState.C:
				acc.Thermostat.CurrentHeatingCoolingState.SetValue(invState(acc.Thermostat.CurrentHeatingCoolingState.GetValue()))
				fmt.Printf("acc update current state: %T - %v \n", acc.Thermostat.CurrentHeatingCoolingState.GetValue(), acc.Thermostat.CurrentHeatingCoolingState.GetValue())
				continue
			case <-tickerUpdateTemp.C:
				acc.Thermostat.CurrentTemperature.SetValue(float64(time.Now().Second()-30) + float64(time.Now().Second()+40)/100)
				fmt.Printf("acc update current temp: %T - %v \n", acc.Thermostat.CurrentTemperature.GetValue(), acc.Thermostat.CurrentTemperature.GetValue())
				continue
			}
		}
	}()
	go acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(state int) { fmt.Printf("acc remote update target state: %T - %v \n", state, state) })
	go acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(temp float64) { fmt.Printf("acc remote update target temp: %T - %v \n", temp, temp) })
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
