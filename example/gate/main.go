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
	flag.StringVar(&config.Name, "n", "Gate", "homekit accessory name")
	flag.StringVar(&config.SN, "sn", "Ex-Gate", "homekit accessory serial number")
	flag.StringVar(&config.Host, "h", "", "homekit host, example: 192.168.1.xxx")
	flag.UintVar(&config.Port, "p", 10706, "homekit port, example: 10101, 10102...")
	flag.StringVar(&config.Pin, "pin", "19736428", "homekit pin, example: 82143697, 13974682")
	flag.Parse()

	homekit.OnLog(debug)
}

func main() {
	acc := homekit.NewAccessoryGate(config.GetInfo("Ex-Gate"))
	log.SetPrefix(fmt.Sprintf("[%T] <%v> ", acc, acc.GetSN()))
	storage := hap.NewFsStore(fmt.Sprintf("./%s", acc.GetSN()))
	server, err := hap.NewServer(storage, acc.GetAccessory())
	if err != nil {
		log.Fatalf("error create hap server: %v\n", err)
	}
	log.Printf("hap server create successful.\n")

	acc.GarageDoorOpener.CurrentDoorState.SetValue(4) //unknown current state
	timer := time.NewTimer(time.Millisecond * 10000)
	go func() {
		for range timer.C {
			// closed --> Target(1), Current(1)
			// closing -> Target(1), Current(2)
			// opened --> Target(0), Current(0)
			// opening -> Target(0), Current(2)
			acc.GarageDoorOpener.CurrentDoorState.SetValue(acc.GarageDoorOpener.TargetDoorState.Value())
			log.Printf("update state, current: %[1]T - %[1]v, target: %[2]T - %[2]v\n", acc.GarageDoorOpener.CurrentDoorState.Value(), acc.GarageDoorOpener.TargetDoorState.Value())
		}
	}()
	acc.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(v int) {
		timer.Stop()
		timer.Reset(time.Millisecond * 10000)
		acc.GarageDoorOpener.CurrentDoorState.SetValue(2)
		log.Printf("remote update state, current: %[1]T - %[1]v, target: %[2]T - %[2]v\n", acc.GarageDoorOpener.CurrentDoorState.Value(), v)
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
