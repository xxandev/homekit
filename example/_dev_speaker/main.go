package main

import (
	"fmt"
	"os"

	"github.com/alpr777/homekit"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

const (
	accessoryName string = "speaker"
	accessorySn   string = "ExmplSPIK"
	accessoryPin  string = "11112222"
)

func main() {
	// runtime.GOMAXPROCS(4)
	// log.Debug.Enable()
	acc := homekit.NewAccessorySpeaker(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "EXAMPLE", Model: "ACC-TEST", FirmwareRevision: "1.2"})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: accessoryPin}, acc.Accessory)
	if err != nil {
		fmt.Println("accessory [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]", "error create transport:", err)
		os.Exit(1)
	}
	go acc.Speaker.Mute.OnValueRemoteUpdate(func(state bool) { fmt.Printf("acc remote update mute: %T - %v \n", state, state) })
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
