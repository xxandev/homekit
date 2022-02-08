package main

import (
	"fmt"
	"log"

	"github.com/alpr777/homekit"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

type AccessoryInterface interface {
	GetAccessory() *accessory.Accessory
	OnValuesRemoteUpdatesPrint()
}

//do not use such names func!!!
func aadd(a AccessoryInterface) *accessory.Accessory {
	go a.OnValuesRemoteUpdatesPrint()
	return a.GetAccessory()
}

//do not use such names func!!!
func newinfo(name, sn string) accessory.Info {
	return accessory.Info{ /* ID: uint64(n),*/ Name: name, SerialNumber: sn, Model: "HAP-ACC-TEST", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision}
}

func main() {
	// log.Debug.Enable()
	bridge := accessory.NewBridge(accessory.Info{ /* ID: 1,*/ Name: "Bridge", SerialNumber: "Ex-Brg", Model: "HAP-BRG", Manufacturer: homekit.Manufacturer, FirmwareRevision: homekit.Revision})
	transp, err := hc.NewIPTransport(hc.Config{StoragePath: "./" + bridge.Info.SerialNumber.GetValue(), Pin: "11223344"}, bridge.Accessory,
		aadd(homekit.NewAccessoryAirPurifier(newinfo("AirPurifier", "ex-ap"), 0, 0, 100, 1)),
		aadd(homekit.NewAccessoryDoor(newinfo("Door", "ex-dr"))),
		aadd(homekit.NewAccessoryFanSpeed(newinfo("Fan", "ex-fsp"))),
		aadd(homekit.NewAccessoryFanSwitch(newinfo("Fan", "ex-fsw"))),
		aadd(homekit.NewAccessoryFan2Speed(newinfo("Fan2", "ex-f2sp"))),
		aadd(homekit.NewAccessoryFan2Switch(newinfo("Fan2", "ex-f2sw"))),
		aadd(homekit.NewAccessoryFaucet(newinfo("Faucet", "ex-fc"))),
		aadd(homekit.NewAccessoryGate(newinfo("Gate", "ex-gt"))),
		aadd(homekit.NewAccessoryHumidifierDehumidifier(newinfo("HumDehum", "ex-hdm"))),
		aadd(homekit.NewAccessoryIrrigation(newinfo("Irrigation", "ex-irg"))),
		aadd(homekit.NewAccessoryLightbulbColored(newinfo("LightColor", "ex-lc"))),
		aadd(homekit.NewAccessoryLightbulbDimmer(newinfo("LightDimm", "ex-ld"))),
		aadd(homekit.NewAccessoryLightbulbSwitch(newinfo("LightSwitch", "ex-ls"))),
		aadd(homekit.NewAccessoryLock(newinfo("Lock", "ex-lock"))),
		aadd(homekit.NewAccessoryOutlet(newinfo("Outlet", "ex-sw"))),
		aadd(homekit.NewAccessorySecuritySystemSimple(newinfo("Alarm", "ex-scr"))),
		aadd(homekit.NewAccessorySensorAirQuality(newinfo("AirQuality", "ex-saq"))),
		aadd(homekit.NewAccessorySensorContact(newinfo("Contact", "ex-sct"))),
		aadd(homekit.NewAccessorySensorDioxide(newinfo("Dioxide", "ex-sdx"))),
		aadd(homekit.NewAccessorySensorHumidity(newinfo("Humidity", "ex-shun"))),
		aadd(homekit.NewAccessorySensorLeak(newinfo("Leak", "ex-sleak"))),
		aadd(homekit.NewAccessorySensorLight(newinfo("Light", "ex-slt"))),
		aadd(homekit.NewAccessorySensorMonoxide(newinfo("Monoxide", "ex-smx"))),
		aadd(homekit.NewAccessorySensorMotion(newinfo("Motion", "ex-smt"))),
		aadd(homekit.NewAccessorySensorSmoke(newinfo("Smoke", "ex-ssm"))),
		aadd(homekit.NewAccessorySensorTemperature(newinfo("Temperature", "ex-stemp"))),
		aadd(homekit.NewAccessorySwitch(newinfo("Switch", "ex-sw"))),
		aadd(homekit.NewAccessoryTelevision(newinfo("Television", "ex-tv"))),
		aadd(homekit.NewAccessoryThermostatAutomatic(newinfo("ThermAtm", "ex-trmatm"))),
		aadd(homekit.NewAccessoryThermostat(newinfo("Thermostat", "ex-trm"))),
		aadd(homekit.NewAccessoryThermostat(newinfo("ThermostatHtn", "ex-trmhtn"), 0, 0, 1, 1)),
		aadd(homekit.NewAccessoryThermostat(newinfo("ThermostatUc", "ex-trmhuc"), 3, 3, 3, 0)),
		aadd(homekit.NewAccessoryWindow(newinfo("Window", "ex-wnd"))),
		aadd(homekit.NewAccessoryWindowCovering(newinfo("WindowCovering", "ex-wndcvr"))),
	)
	if err != nil {
		log.Fatalf("[ %v / %v ] error create hap transport: %v\n", bridge.Info.SerialNumber.GetValue(), bridge.Info.Name.GetValue(), err)
	}
	fmt.Printf("[ %v / %v ] accessories transport start\n", bridge.Info.SerialNumber.GetValue(), bridge.Info.Name.GetValue())
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
}
