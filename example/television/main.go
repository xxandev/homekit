package main

import (
	"fmt"
	"os"

	"github.com/alpr777/homekit"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
)

const (
	accessoryName string = "television"
	accessorySn   string = "ExmplTV"
	accessoryPin  string = "19283746"
)

func main() {
	// runtime.GOMAXPROCS(4)
	// log.Debug.Enable()
	acc := homekit.NewAccessoryTelevision(accessory.Info{Name: accessoryName, SerialNumber: accessorySn, Manufacturer: "EXAMPLE", Model: "ACC-TEST", FirmwareRevision: "1.2"})
	tvin1 := acc.AddInputSource(1, "HDMI 1", characteristic.InputSourceTypeHdmi)
	_ = acc.AddInputSource(2, "HDMI 2", characteristic.InputSourceTypeHdmi)
	_ = acc.AddInputSource(3, "YouTube", characteristic.InputSourceTypeApplication)
	_ = acc.AddInputSource(4, "Airplay", characteristic.InputSourceTypeAirplay)
	_ = acc.AddInputSource(5, "AndroidTV", characteristic.InputSourceTypeOther)
	_ = acc.AddInputSource(6, "AppleTV", characteristic.InputSourceTypeOther)
	_ = acc.AddInputSource(7, "Xbox", characteristic.InputSourceTypeOther)
	_ = acc.AddInputSource(8, "Paystation", characteristic.InputSourceTypeOther)
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + acc.Info.SerialNumber.GetValue(), Pin: accessoryPin}, acc.Accessory)
	if err != nil {
		fmt.Println("accessory [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]", "error create transport:", err)
		os.Exit(1)
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
	tvin1.Name.OnValueRemoteUpdate(func(v string) {
		fmt.Printf("input source %s remote update name: %T - %v \n", tvin1.Name.GetValue(), v, v)
	})

	go acc.Television.Active.OnValueRemoteUpdate(func(v int) { fmt.Printf("acc remote update active: %T - %v \n", v, v) })
	go acc.Television.ActiveIdentifier.OnValueRemoteUpdate(func(v int) { fmt.Printf("acc remote update active identifier: %T - %v \n", v, v) })
	go acc.Television.ConfiguredName.OnValueRemoteUpdate(func(v string) { fmt.Printf("acc remote update configured name: %T - %v \n", v, v) })
	go acc.Television.SleepDiscoveryMode.OnValueRemoteUpdate(func(v int) { fmt.Printf("acc remote update sleep discovery mode: %T - %v \n", v, v) })
	go acc.Television.Brightness.OnValueRemoteUpdate(func(v int) { fmt.Printf("acc remote update brightness: %T - %v \n", v, v) })
	go acc.Television.ClosedCaptions.OnValueRemoteUpdate(func(v int) { fmt.Printf("acc remote update closed captions: %T - %v \n", v, v) })
	go acc.Television.DisplayOrder.OnValueRemoteUpdate(func(v []byte) { fmt.Printf("acc remote update display order: %T - %v \n", v, v) })
	go acc.Television.CurrentMediaState.OnValueRemoteUpdate(func(v int) { fmt.Printf("acc remote update current media state: %T - %v \n", v, v) })
	go acc.Television.TargetMediaState.OnValueRemoteUpdate(func(v int) { fmt.Printf("acc remote update target media state: %T - %v \n", v, v) })
	go acc.Television.PowerModeSelection.OnValueRemoteUpdate(func(v int) { fmt.Printf("acc remote update power mode selection: %T - %v \n", v, v) })
	go acc.Television.PictureMode.OnValueRemoteUpdate(func(v int) { fmt.Printf("acc remote update picture mode: %T - %v \n", v, v) })
	go acc.Television.RemoteKey.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("acc remote update remote key: %T - %v >>", v, v)
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
	fmt.Println("homekit accessory transport start [", acc.Info.SerialNumber.GetValue(), "/", acc.Info.Name.GetValue(), "]")
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
