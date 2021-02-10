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
	acc := homekit.NewAccessoryDoorLock(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "EXAMPLE", Model: "ACC-TEST", FirmwareRevision: "1.2"})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: accessoryPin}, acc.Accessory)
	if err != nil {
		fmt.Println("accessory [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]", "error create transport:", err)
		os.Exit(1)
	}
	go acc.LockMechanism.LockTargetState.OnValueRemoteUpdate(func(state int) {
		fmt.Printf("acc remote update target position: %T - %v \n", state, state)
		acc.LockMechanism.LockCurrentState.SetValue(state)
		fmt.Printf("acc update current position: %T - %v \n", acc.LockMechanism.LockCurrentState.GetValue(), acc.LockMechanism.LockCurrentState.GetValue())
	})
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
