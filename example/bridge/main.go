package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/brutella/hap"
	"github.com/brutella/hap/accessory"
	"github.com/xxandev/homekit"
)

const MANUFACTURER string = ""

type Config struct{ homekit.AccessoryConfig }

var (
	debug  bool
	config Config
)

func init() {
	log.SetOutput(os.Stdout) // log.SetOutput(ioutil.Discard)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix)

	flag.BoolVar(&debug, "d", false, "hap debug log activate")
	flag.StringVar(&config.Name, "n", "Bridge", "homekit accessory name")
	flag.StringVar(&config.SN, "sn", "Ex-Bridge", "homekit accessory serial number")
	flag.StringVar(&config.Host, "h", "", "homekit host, example: 192.168.1.xxx")
	flag.UintVar(&config.Port, "p", 10701, "homekit port, example: 10101, 10102...")
	flag.StringVar(&config.Pin, "pin", "19378246", "homekit pin, example: 82143697, 13974682")
	flag.Parse()

	log.SetPrefix(fmt.Sprintf("[SHOWROOM] <%v> ", config.SN))
	homekit.OnLog(debug)
}

func main() {
	bridge := homekit.NewAccBridge(1, accessory.Info{Name: config.Name, SerialNumber: config.SN, Model: "HK-Showroom", Manufacturer: homekit.MANUFACTURER, Firmware: "1.0.0"})
	storage := hap.NewFsStore(fmt.Sprintf("./%s", bridge.GetSN()))
	server, err := hap.NewServer(storage, bridge.A,
		ExampleSwitch(10, "CY"),
		ExampleSwitch(11, "CY"),
		ExampleSwitch(12, "CY"),
		ExampleSensorMotion(13, "CY"),
		ExampleDoor(14, "CY"),
		ExampleGate(15, "CY"),
		ExampleIrrigation(16, "CY"),
		ExampleSwitch(20, "GT"),
		ExampleSwitch(21, "GT"),
		ExampleSwitch(22, "GT"),
		ExampleSwitch(23, "GT"),
		ExampleGate(24, "GT"),
		ExampleGate(25, "GT"),
		ExampleOutlet(26, "GT"),
		ExampleOutlet(27, "GT"),
		ExampleAirPurifier(30, "KT"),
		ExampleFaucet(31, "KT"),
		ExampleOutlet(32, "KT"),
		ExampleOutlet(33, "KT"),
		ExampleSensorLeak(34, "KT"),
		ExampleSensorHumidity(35, "KT"),
		ExampleSensorMotion(36, "KT"),
		ExampleSwitch(37, "KT"),
		ExampleSwitch(38, "KT"),
		ExampleLightbulbColored(39, "KT"),
		ExampleLightbulbDimmer(40, "KT"),
		ExampleWindow(41, "KT"),
		ExampleWindowCovering(42, "KT"),
		ExampleThermostatClimate(43, "KT"),
		ExampleOutlet(50, "BDR"),
		ExampleOutlet(51, "BDR"),
		ExampleSwitch(52, "BDR"),
		ExampleSwitch(53, "BDR"),
		ExampleLightbulbColored(54, "BDR"),
		ExampleLightbulbDimmer(55, "BDR"),
		ExampleSensorHumidity(56, "BDR"),
		ExampleSensorMotion(57, "BDR"),
		ExampleWindowCovering(58, "BDR"),
		ExampleThermostatClimate(59, "BDR"),
		ExampleOutlet(70, "LR"),
		ExampleOutlet(71, "LR"),
		ExampleSwitch(72, "LR"),
		ExampleSwitch(73, "LR"),
		ExampleLightbulbColored(74, "LR"),
		ExampleLightbulbDimmer(75, "LR"),
		ExampleSensorHumidity(76, "LR"),
		ExampleSensorMotion(77, "LR"),
		ExampleWindowCovering(78, "LR"),
		ExampleThermostatClimate(79, "LR"),
		ExampleThermostatHeating(80, "LR"),
		ExampleLightbulbColored(90, "WC"),
		ExampleLightbulbDimmer(91, "WC"),
		ExampleFanRS(92, "WC"),
		ExampleOutlet(93, "WC"),
		ExampleSensorLeak(94, "WC"),
		ExampleSensorMotion(95, "WC"),
		ExampleFaucet(96, "WC"),
		ExampleLightbulbColored(100, "BR"),
		ExampleLightbulbDimmer(101, "BR"),
		ExampleSensorHumidity(102, "BR"),
		ExampleFanRS(103, "BR"),
		ExampleOutlet(104, "BR"),
		ExampleSensorLeak(105, "BR"),
		ExampleSensorMotion(106, "BR"),
		ExampleSensorContact(107, "BR"),
		ExampleSensorContact(108, "BR"),
		ExampleWindowCovering(109, "BR"),
		ExampleFaucet(110, "BR"),
	)
	if err != nil {
		log.Fatalf("error create hap server: %v\n", err)
	}
	log.Printf("hap server create successful.\n")
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sig
		log.Printf("stop program signal.\n")
		signal.Stop(sig)
		cancel()
	}()
	homekit.SetServer(server, config.GetAddress(), config.GetPin())
	log.Printf("hap server starting set, address %v, pin %v.\n", server.Addr, server.Pin)
	log.Fatalf("hap server: %v\n", server.ListenAndServe(ctx))
}
