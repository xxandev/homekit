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
	acc := homekit.NewAccessorySmartSpeaker(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "alpr777", Model: "ACC-TEST", FirmwareRevision: "1.2"})
	acc.SmartSpeaker.ConfiguredName.SetValue("Smart Speaker")
	acc.SmartSpeaker.Mute.SetValue(false)
	acc.SmartSpeaker.Volume.SetValue(100)
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: accessoryPin}, acc.Accessory)
	if err != nil {
		fmt.Println("accessory [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]", "error create transport:", err)
		os.Exit(1)
	}
	go acc.SmartSpeaker.TargetMediaState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc smart speaker remote update target media state: %T - %v \n", v, v)
		acc.SmartSpeaker.CurrentMediaState.SetValue(acc.SmartSpeaker.TargetMediaState.GetValue())
		fmt.Printf("acc smart speaker update current media state: %T - %v \n", acc.SmartSpeaker.CurrentMediaState.GetValue(), acc.SmartSpeaker.CurrentMediaState.GetValue())
	})
	go acc.SmartSpeaker.Name.OnValueRemoteUpdate(func(v string) {
		fmt.Printf("acc smart speaker remote update name: %T - %v \n", v, v)
	})
	go acc.SmartSpeaker.ConfiguredName.OnValueRemoteUpdate(func(v string) {
		fmt.Printf("acc smart speaker remote update configured name: %T - %v \n", v, v)
	})
	go acc.SmartSpeaker.Mute.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("acc smart speaker remote update mute: %T - %v \n", v, v)
	})
	go acc.SmartSpeaker.Volume.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc smart speaker remote update volume: %T - %v \n", v, v)
	})
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
