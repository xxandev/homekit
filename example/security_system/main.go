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
	accessoryName string = "alarm"
	accessorySn   string = "ExmplSecSys"
	accessoryPin  string = "11112222"
)

func main() {
	// runtime.GOMAXPROCS(4)
	// log.Debug.Enable()
	acc := homekit.NewAccessorySecuritySystemSimple(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "EXAMPLE", Model: "ACC-TEST", FirmwareRevision: "1.2"})
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
				if acc.SecuritySystemSimple.SecuritySystemCurrentState.GetValue() >= 4 {
					acc.SecuritySystemSimple.SecuritySystemCurrentState.SetValue(0)
				} else {
					acc.SecuritySystemSimple.SecuritySystemCurrentState.SetValue(acc.SecuritySystemSimple.SecuritySystemCurrentState.GetValue() + 1)
				}
				fmt.Printf("acc update current state: %T - %v \n", acc.SecuritySystemSimple.SecuritySystemCurrentState.GetValue(), acc.SecuritySystemSimple.SecuritySystemCurrentState.GetValue())
				continue
			}
		}
	}()
	go acc.SecuritySystemSimple.SecuritySystemTargetState.OnValueRemoteUpdate(func(state int) { fmt.Printf("acc remote update target state: %T - %v \n", state, state) })
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
