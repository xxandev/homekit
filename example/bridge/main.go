package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/brutella/hap"
	"github.com/brutella/hap/accessory"
	"github.com/xxandev/homekit"
)

const (
	BRG_NAME    string = "Bridge"
	BRG_SN      string = "EX-Brg"
	BRG_MODEL   string = "HAP-BRG"
	BRG_ADDRESS string = ":11102"
	BRG_PIN     string = "12344321"
)

type AccessoryInterface interface {
	GetAccessory() *accessory.A
	OnExample()
}

//do not use such names func!!!
func aadd(a AccessoryInterface) *accessory.A {
	a.OnExample()
	return a.GetAccessory()
}

//do not use such names func!!!
func newinfo(name, sn, model string) accessory.Info {
	return accessory.Info{Name: name, SerialNumber: sn, Model: model, Manufacturer: homekit.Manufacturer, Firmware: homekit.Firmware}
}

func main() {
	homekit.OnLog(false)
	bridge := accessory.NewBridge(accessory.Info{Name: BRG_NAME, SerialNumber: BRG_SN, Model: BRG_MODEL, Manufacturer: homekit.Manufacturer, Firmware: homekit.Firmware})
	llog := log.New(os.Stdout, fmt.Sprintf("[ %v / %v ] ", bridge.A.Info.SerialNumber.Value(), bridge.A.Info.Name.Value()), log.Ldate|log.Ltime|log.Lmsgprefix)
	storage := hap.NewFsStore(fmt.Sprintf("./%s", bridge.Info.SerialNumber.Value()))
	server, err := hap.NewServer(storage, bridge.A,
		aadd(homekit.NewAccessoryAirPurifier(newinfo("AirPurifier", "EX-AirPurifier", "HAP-AP"), 0, 0, 100, 1)),
		aadd(homekit.NewAccessoryDoor(newinfo("Door", "EX-dr", "HAP-"))),
		aadd(homekit.NewAccessoryFanRS(newinfo("Fan", "EX-fsp", "HAP-"))),
		aadd(homekit.NewAccessoryFanSwitch(newinfo("Fan", "EX-fsw", "HAP-"))),
		aadd(homekit.NewAccessoryFan2RS(newinfo("Fan2", "EX-f2sp", "HAP-"))),
		aadd(homekit.NewAccessoryFan2Switch(newinfo("Fan2", "EX-f2sw", "HAP-"))),
		aadd(homekit.NewAccessoryFaucet(newinfo("Faucet", "EX-fc", "HAP-"))),
		aadd(homekit.NewAccessoryGate(newinfo("Gate", "EX-gt", "HAP-"))),
		aadd(homekit.NewAccessoryHumidifierDehumidifier(newinfo("HumDehum", "EX-hdm", "HAP-"))),
		aadd(homekit.NewAccessoryIrrigation(newinfo("Irrigation", "EX-irg", "HAP-"))),
		aadd(homekit.NewAccessoryLightbulbColored(newinfo("LightColor", "EX-lc", "HAP-"))),
		aadd(homekit.NewAccessoryLightbulbDimmer(newinfo("LightDimm", "EX-ld", "HAP-"))),
		aadd(homekit.NewAccessoryLightbulbSwitch(newinfo("LightSwitch", "EX-ls", "HAP-"))),
		aadd(homekit.NewAccessoryLock(newinfo("Lock", "EX-lock", "HAP-"))),
		aadd(homekit.NewAccessoryOutlet(newinfo("Outlet", "EX-sw", "HAP-"))),
		aadd(homekit.NewAccessorySecuritySystem(newinfo("Alarm", "EX-scr", "HAP-"))),
		aadd(homekit.NewAccessorySensorAirQuality(newinfo("AirQuality", "EX-saq", "HAP-"))),
		aadd(homekit.NewAccessorySensorContact(newinfo("Contact", "EX-sct", "HAP-"))),
		aadd(homekit.NewAccessorySensorDioxide(newinfo("Dioxide", "EX-sdx", "HAP-"))),
		aadd(homekit.NewAccessorySensorHumidity(newinfo("Humidity", "EX-shun", "HAP-"))),
		aadd(homekit.NewAccessorySensorLeak(newinfo("Leak", "EX-sleak", "HAP-"))),
		aadd(homekit.NewAccessorySensorLight(newinfo("Light", "EX-slt", "HAP-"))),
		aadd(homekit.NewAccessorySensorMonoxide(newinfo("Monoxide", "EX-smx", "HAP-"))),
		aadd(homekit.NewAccessorySensorMotion(newinfo("Motion", "EX-smt", "HAP-"))),
		aadd(homekit.NewAccessorySensorSmoke(newinfo("Smoke", "EX-ssm", "HAP-"))),
		aadd(homekit.NewAccessorySensorTemperature(newinfo("Temperature", "EX-stemp", "HAP-"))),
		aadd(homekit.NewAccessorySwitch(newinfo("Switch", "EX-sw", "HAP-"))),
		aadd(homekit.NewAccessoryTelevision(newinfo("Television", "EX-tv", "HAP-"))),
		aadd(homekit.NewAccessoryThermostatAutomatic(newinfo("ThermAtm", "EX-trmatm", "HAP-"))),
		aadd(homekit.NewAccessoryThermostat(newinfo("Thermostat", "EX-trm", "HAP-"))),
		aadd(homekit.NewAccessoryThermostat(newinfo("ThermostatHtn", "EX-trmhtn", "HAP-"), 0, 0, 1, 1)),
		aadd(homekit.NewAccessoryThermostat(newinfo("ThermostatUc", "EX-trmhuc", "HAP-"), 3, 3, 3, 0)),
		aadd(homekit.NewAccessoryWindow(newinfo("Window", "EX-wnd", "HAP-"))),
		aadd(homekit.NewAccessoryWindowCovering(newinfo("WindowCovering", "EX-wndcvr", "HAP-"))),
	)
	if err != nil {
		llog.Fatalf("error create hap server: %v\n", err)
	}
	llog.Printf("hap server create successful.\n")
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sig
		llog.Printf("stop program signal.\n")
		signal.Stop(sig)
		cancel()
	}()
	homekit.SetServer(server, BRG_ADDRESS, BRG_PIN)
	llog.Printf("hap server starting set, address %v, pin %v.\n", server.Addr, server.Pin)
	llog.Fatalf("hap server: %v\n", server.ListenAndServe(ctx))
}
