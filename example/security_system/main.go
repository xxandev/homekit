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
	acc := homekit.NewAccessorySecuritySystemSimple(accessory.Info{Name: "Alarm", SerialNumber: "Ex-Secur", Model: "HAP-SCR-SS", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	go func() {
		t := time.NewTicker(30 * time.Second)
		for range t.C {
			if acc.SecuritySystemSimple.SecuritySystemCurrentState.GetValue() >= 4 {
				acc.SecuritySystemSimple.SecuritySystemCurrentState.SetValue(0)
			} else {
				acc.SecuritySystemSimple.SecuritySystemCurrentState.SetValue(acc.SecuritySystemSimple.SecuritySystemCurrentState.GetValue() + 1)
			}
			fmt.Printf("acc security system update current state: %T - %v \n", acc.SecuritySystemSimple.SecuritySystemCurrentState.GetValue(), acc.SecuritySystemSimple.SecuritySystemCurrentState.GetValue())
		}
	}()
	go acc.SecuritySystemSimple.SecuritySystemTargetState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc security system remote update target state: %T - %v \n", v, v)
	})
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
