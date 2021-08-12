package main

import (
	"fmt"
	"log"

	"github.com/alpr777/homekit"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
)

func main() {
	// log.Debug.Enable()
	acc := homekit.NewAccessoryTelevision(accessory.Info{Name: "Television", SerialNumber: "Ex-Tv", Model: "HAP-TV", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	tvin1 := acc.AddInputSource(1, "HDMI 1", characteristic.InputSourceTypeHdmi)
	_ = acc.AddInputSource(2, "HDMI 2", characteristic.InputSourceTypeHdmi)
	_ = acc.AddInputSource(3, "YouTube", characteristic.InputSourceTypeApplication)
	_ = acc.AddInputSource(4, "Airplay", characteristic.InputSourceTypeAirplay)
	_ = acc.AddInputSource(5, "AndroidTV", characteristic.InputSourceTypeOther)
	_ = acc.AddInputSource(6, "AppleTV", characteristic.InputSourceTypeOther)
	_ = acc.AddInputSource(7, "Xbox", characteristic.InputSourceTypeOther)
	_ = acc.AddInputSource(8, "Paystation", characteristic.InputSourceTypeOther)
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: "11223344"}, acc.Accessory)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue(), err)
	}
	tvin1.ConfiguredName.OnValueRemoteUpdate(func(v string) {
		fmt.Printf("input source %s remote update configured name: %T - %v \n", tvin1.Name.GetValue(), v, v)
	})
	tvin1.InputSourceType.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("input source %s remote update input source type: %T - %v \n", tvin1.Name.GetValue(), v, v)
	})
	tvin1.IsConfigured.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("input source %s remote update is configured: %T - %v \n", tvin1.Name.GetValue(), v, v)
	})
	tvin1.Identifier.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("input source %s remote update identifier: %T - %v \n", tvin1.Name.GetValue(), v, v)
	})
	tvin1.InputDeviceType.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("input source %s remote update input device type: %T - %v \n", tvin1.Name.GetValue(), v, v)
	})
	tvin1.TargetVisibilityState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("input source %s remote update target visibility state: %T - %v \n", tvin1.Name.GetValue(), v, v)
	})

	go acc.Speaker.Mute.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("acc speaker remote update mute: %T - %v \n", v, v)
	})
	go acc.Speaker.Volume.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc speaker remote update volume: %T - %v \n", v, v)
	})
	go acc.Speaker.VolumeControlType.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc speaker remote update volume control type: %T - %v \n", v, v)
	})
	go acc.Speaker.VolumeSelector.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc speaker remote update volume selector: %T - %v \n", v, v)
	})

	go acc.Television.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc television remote update active: %T - %v \n", v, v)
	})
	go acc.Television.ActiveIdentifier.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc television remote update active identifier: %T - %v \n", v, v)
	})
	go acc.Television.ConfiguredName.OnValueRemoteUpdate(func(v string) {
		fmt.Printf("acc television remote update configured name: %T - %v \n", v, v)
	})
	go acc.Television.SleepDiscoveryMode.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc television remote update sleep discovery mode: %T - %v \n", v, v)
	})
	go acc.Television.Brightness.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc television remote update brightness: %T - %v \n", v, v)
	})
	go acc.Television.ClosedCaptions.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc television remote update closed captions: %T - %v \n", v, v)
	})
	go acc.Television.DisplayOrder.OnValueRemoteUpdate(func(v []byte) {
		fmt.Printf("acc television remote update display order: %T - %v \n", v, v)
	})
	go acc.Television.CurrentMediaState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc television remote update current media state: %T - %v \n", v, v)
	})
	go acc.Television.TargetMediaState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc television remote update target media state: %T - %v \n", v, v)
	})
	go acc.Television.PowerModeSelection.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc television remote update power mode selection: %T - %v \n", v, v)
	})
	go acc.Television.PictureMode.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc television remote update picture mode: %T - %v \n", v, v)
	})
	go acc.Television.RemoteKey.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc television remote update remote key: %T - %v >>", v, v)
		switch v {
		case characteristic.RemoteKeyRewind:
			fmt.Println("Rewind")
		case characteristic.RemoteKeyFastForward:
			fmt.Println("Fast forward")
		case characteristic.RemoteKeyExit:
			fmt.Println("Exit")
		case characteristic.RemoteKeyPlayPause:
			fmt.Println("Play/Pause")
		case characteristic.RemoteKeyInfo:
			fmt.Println("Info")
		case characteristic.RemoteKeyNextTrack:
			fmt.Println("Next")
		case characteristic.RemoteKeyPrevTrack:
			fmt.Println("Prev")
		case characteristic.RemoteKeyArrowUp:
			fmt.Println("Up")
		case characteristic.RemoteKeyArrowDown:
			fmt.Println("Down")
		case characteristic.RemoteKeyArrowLeft:
			fmt.Println("Left")
		case characteristic.RemoteKeyArrowRight:
			fmt.Println("Right")
		case characteristic.RemoteKeySelect:
			fmt.Println("Select")
		case characteristic.RemoteKeyBack:
			fmt.Println("Back")
		}
	})
	fmt.Printf("[ %v / %v ] accessories transport start\n", acc.Accessory.Info.SerialNumber.GetValue(), acc.Accessory.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
