package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/brutella/hap"
	"github.com/xxandev/homekit"
)

type Config struct{ homekit.AccessoryConfig }

var (
	debug  bool
	config Config
)

func init() {
	log.SetOutput(os.Stdout) // log.SetOutput(ioutil.Discard)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmsgprefix)

	flag.BoolVar(&debug, "d", false, "hap debug log activate")
	flag.StringVar(&config.Name, "n", "Therm", "homekit accessory name")
	flag.StringVar(&config.SN, "sn", "Ex-Therm", "homekit accessory serial number")
	flag.StringVar(&config.Host, "h", "", "homekit host, example: 192.168.1.xxx")
	flag.UintVar(&config.Port, "p", 10727, "homekit port, example: 10101, 10102...")
	flag.StringVar(&config.Pin, "pin", "19378246", "homekit pin, example: 82143697, 13974682")
	flag.Parse()

	homekit.OnLog(debug)
}

func main() {
	acc := homekit.NewAccessoryThermostat(config.GetInfo("Ex-Therm"))
	log.SetPrefix(fmt.Sprintf("[%T] <%v> ", acc, acc.GetSN()))
	storage := hap.NewFsStore(fmt.Sprintf("./%s", acc.GetSN()))
	server, err := hap.NewServer(storage, acc.GetAccessory())
	if err != nil {
		log.Fatalf("error create hap server: %v\n", err)
	}
	log.Printf("hap server create successful.\n")

	go func() {
		for range time.Tick(time.Millisecond * 5000) {
			if acc.Thermostat.TargetHeatingCoolingState.Value() == 0 { //off
				acc.Thermostat.CurrentHeatingCoolingState.SetValue(acc.Thermostat.TargetHeatingCoolingState.Value())
				continue
			}
			if acc.Thermostat.TargetHeatingCoolingState.Value() == 1 { //heating
				if acc.Thermostat.CurrentTemperature.Value() > acc.Thermostat.TargetTemperature.Value() {
					acc.Thermostat.CurrentHeatingCoolingState.SetValue(0)
				} else {
					acc.Thermostat.CurrentHeatingCoolingState.SetValue(acc.Thermostat.TargetHeatingCoolingState.Value())
					acc.Thermostat.CurrentTemperature.SetValue(acc.Thermostat.CurrentTemperature.Value() + 0.25)
				}
			}
			if acc.Thermostat.TargetHeatingCoolingState.Value() == 2 { //cooling
				if acc.Thermostat.CurrentTemperature.Value() < acc.Thermostat.TargetTemperature.Value() {
					acc.Thermostat.CurrentHeatingCoolingState.SetValue(0)
				} else {
					acc.Thermostat.CurrentHeatingCoolingState.SetValue(acc.Thermostat.TargetHeatingCoolingState.Value())
					acc.Thermostat.CurrentTemperature.SetValue(acc.Thermostat.CurrentTemperature.Value() - 0.25)
				}
			}
			if acc.Thermostat.TargetHeatingCoolingState.Value() == 3 { //autom
				if acc.Thermostat.CurrentTemperature.Value() == acc.Thermostat.TargetTemperature.Value() {
					acc.Thermostat.CurrentHeatingCoolingState.SetValue(0)
				} else if acc.Thermostat.CurrentTemperature.Value() > acc.Thermostat.TargetTemperature.Value() {
					acc.Thermostat.TargetHeatingCoolingState.SetValue(2)
					acc.Thermostat.CurrentTemperature.SetValue(acc.Thermostat.CurrentTemperature.Value() - 0.25)
				} else if acc.Thermostat.CurrentTemperature.Value() < acc.Thermostat.TargetTemperature.Value() {
					acc.Thermostat.TargetHeatingCoolingState.SetValue(1)
					acc.Thermostat.CurrentTemperature.SetValue(acc.Thermostat.CurrentTemperature.Value() + 0.25)
				}
			}
			log.Printf("update thermostat, current state: %[1]T - %[1]v, target state: %[2]T - %[2]v, current temp: %[3]T - %[3]v, target temp: %[4]T - %[4]v, \n",
				acc.Thermostat.CurrentHeatingCoolingState.Value(),
				acc.Thermostat.TargetHeatingCoolingState.Value(),
				acc.Thermostat.CurrentTemperature.Value(),
				acc.Thermostat.TargetTemperature.Value(),
			)
		}
	}()
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(v int) {
		log.Printf("remote update target state: %[1]T - %[1]v\n", v)
	})
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(v float64) {
		log.Printf("remote update target temp: %[1]T - %[1]v\n", v)
	})

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sig
		log.Println("program stop.")
		signal.Stop(sig)
		cancel()
	}()
	homekit.SetServer(server, config.GetAddress(), config.GetPin())
	log.Printf("hap server starting set, address: %v, pin: %v.\n", server.Addr, server.Pin)
	log.Fatalf("hap server: %v\n", server.ListenAndServe(ctx))
}
