package main

import (
	"fmt"
	"log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/xxandev/homekit"
)

func main() {
	// log.Debug.Enable()
	acc := homekit.NewAccessorySmartSpeaker(accessory.Info{Name: "Speaker", SerialNumber: "Ex-Spk", Model: "HAP-SPK", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	acc.SmartSpeaker.ConfiguredName.SetValue("Smart Speaker")
	acc.SmartSpeaker.Mute.SetValue(false)
	acc.SmartSpeaker.Volume.SetValue(100)
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
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
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
