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
	BRG_ADDRESS string = ":11102"
	BRG_PIN     string = "12344321"
	BRG_MODEL   string = "HAP-BRG"
)

type AccessoryInterface interface {
	GetAccessory() *accessory.A
	OnExample()
}

//do not use such names func!!!
func aadd(id uint64, a AccessoryInterface) *accessory.A {
	a.OnExample()
	acc := a.GetAccessory()
	acc.Id = id
	return acc
}

//do not use such names func!!!
func newinfo(name, sn, model string) accessory.Info {
	return accessory.Info{Name: name, SerialNumber: sn, Model: model, Manufacturer: homekit.Manufacturer, Firmware: homekit.Firmware}
}

func main() {
	homekit.OnLog(false)
	bridge := accessory.NewBridge(accessory.Info{Name: BRG_NAME, SerialNumber: BRG_SN, Model: BRG_MODEL, Manufacturer: homekit.Manufacturer, Firmware: homekit.Firmware})
	bridge.A.Id = 1
	llog := log.New(os.Stdout, fmt.Sprintf("[ %v / %v ] ", bridge.A.Info.SerialNumber.Value(), bridge.A.Info.Name.Value()), log.Ldate|log.Ltime|log.Lmsgprefix)
	storage := hap.NewFsStore(fmt.Sprintf("./%s", bridge.Info.SerialNumber.Value()))
	server, err := hap.NewServer(storage, bridge.A,
		aadd(2, homekit.NewAccessoryAirPurifier(newinfo("AirPurifier", "EX-AirPurifier", "HAP-AP"), 0, 0, 100, 1)),
		aadd(3, homekit.NewAccessoryDoor(newinfo("Door", "EX-dr", "HAP-DR"))),
		aadd(4, homekit.NewAccessoryFanRS(newinfo("Fan", "EX-fsp", "HAP-FRS"))),
		aadd(5, homekit.NewAccessoryFanSwitch(newinfo("Fan", "EX-fsw", "HAP-FS"))),
		aadd(6, homekit.NewAccessoryFan2RS(newinfo("Fan2", "EX-f2sp", "HAP-F2RS"))),
		aadd(7, homekit.NewAccessoryFan2Switch(newinfo("Fan2", "EX-f2sw", "HAP-F2S"))),
		aadd(8, homekit.NewAccessoryFaucet(newinfo("Faucet", "EX-fc", "HAP-FCT"))),
		aadd(9, homekit.NewAccessoryGate(newinfo("Gate", "EX-gt", "HAP-DT"))),
		aadd(10, homekit.NewAccessoryHumidifierDehumidifier(newinfo("HumDehum", "EX-hdm", "HAP-HMD"))),
		aadd(11, homekit.NewAccessoryIrrigation(newinfo("Irrigation", "EX-irg", "HAP-IRG"))),
		aadd(12, homekit.NewAccessoryLightbulbColored(newinfo("LightColor", "EX-lc", "HAP-LBC"))),
		aadd(13, homekit.NewAccessoryLightbulbDimmer(newinfo("LightDimm", "EX-ld", "HAP-LBD"))),
		aadd(14, homekit.NewAccessoryLightbulbSwitch(newinfo("LightSwitch", "EX-ls", "HAP-LBS"))),
		aadd(15, homekit.NewAccessoryLock(newinfo("Lock", "EX-lock", "HAP-LOCK"))),
		aadd(16, homekit.NewAccessoryOutlet(newinfo("Outlet", "EX-sw", "HAP-OTL"))),
		aadd(17, homekit.NewAccessorySecuritySystem(newinfo("Alarm", "EX-scr", "HAP-SS"))),
		aadd(18, homekit.NewAccessorySensorAirQuality(newinfo("AirQuality", "EX-saq", "HAP-SAQ"))),
		aadd(19, homekit.NewAccessorySensorContact(newinfo("Contact", "EX-sct", "HAP-SCNT"))),
		aadd(20, homekit.NewAccessorySensorDioxide(newinfo("Dioxide", "EX-sdx", "HAP-SDXD"))),
		aadd(21, homekit.NewAccessorySensorHumidity(newinfo("Humidity", "EX-shun", "HAP-SHUM"))),
		aadd(22, homekit.NewAccessorySensorLeak(newinfo("Leak", "EX-sleak", "HAP-SL"))),
		aadd(23, homekit.NewAccessorySensorLight(newinfo("Light", "EX-slt", "HAP-SLG"))),
		aadd(24, homekit.NewAccessorySensorMonoxide(newinfo("Monoxide", "EX-smx", "HAP-SMNX"))),
		aadd(25, homekit.NewAccessorySensorMotion(newinfo("Motion", "EX-smt", "HAP-SMNT"))),
		aadd(26, homekit.NewAccessorySensorSmoke(newinfo("Smoke", "EX-ssm", "HAP-SSMK"))),
		aadd(27, homekit.NewAccessorySensorTemperature(newinfo("Temperature", "EX-stemp", "HAP-STEMP"))),
		aadd(28, homekit.NewAccessorySwitch(newinfo("Switch", "EX-sw", "HAP-SW"))),
		aadd(29, homekit.NewAccessoryTelevision(newinfo("Television", "EX-tv", "HAP-TV"))),
		aadd(30, homekit.NewAccessoryThermostatAutomatic(newinfo("ThermAtm", "EX-trmatm", "HAP-TMA"))),
		aadd(31, homekit.NewAccessoryThermostat(newinfo("Thermostat", "EX-trm", "HAP-TM"))),
		aadd(32, homekit.NewAccessoryThermostat(newinfo("ThermostatHtn", "EX-trmhtn", "HAP-TM"), 0, 0, 1, 1)),
		aadd(33, homekit.NewAccessoryThermostat(newinfo("ThermostatUc", "EX-trmhuc", "HAP-TM"), 3, 3, 3, 0)),
		aadd(34, homekit.NewAccessoryWindow(newinfo("Window", "EX-wnd", "HAP-W"))),
		aadd(35, homekit.NewAccessoryWindowCovering(newinfo("WindowCovering", "EX-wndcvr", "HAP-WCR"))),
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
