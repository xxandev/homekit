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
	acc := homekit.NewAccessoryThermostat(accessory.Info{Name: "Thermostat", SerialNumber: "Ex-Therm-Non", Model: "HAP-TRM-NON", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision}, 3, 3, 3, 0)
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		tickerUpdateState := time.NewTicker(10 * time.Second)
		tickerUpdateTemp := time.NewTicker(2 * time.Second)
		for {
			select {
			case <-tickerUpdateState.C:
				state := acc.Thermostat.CurrentHeatingCoolingState.GetValue()
				if state >= 2 {
					acc.Thermostat.CurrentHeatingCoolingState.SetValue(0)
				} else {
					acc.Thermostat.CurrentHeatingCoolingState.SetValue(state + 1)
				}
				fmt.Printf("acc thermostat update current state: %T - %v \n", acc.Thermostat.CurrentHeatingCoolingState.GetValue(), acc.Thermostat.CurrentHeatingCoolingState.GetValue())
				continue
			case <-tickerUpdateTemp.C:
				acc.Thermostat.CurrentTemperature.SetValue(float64(time.Now().Second()-30) + float64(time.Now().Second()+40)/100)
				fmt.Printf("acc thermostat update current temp: %T - %v \n", acc.Thermostat.CurrentTemperature.GetValue(), acc.Thermostat.CurrentTemperature.GetValue())
				continue
			}
		}
	}()
	go acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc thermostat remote update target state: %T - %v \n", v, v)
	})
	go acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("acc thermostat remote update target temp: %T - %v \n", v, v)
	})
	fmt.Println("homekit accessory transport start [", "/", acc.Accessory.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
