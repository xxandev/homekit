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
	NAME    string = "Speaker"
	SN      string = "EX-Speaker"
	MODEL   string = "HAP-SPKR"
	ADDRESS string = ":11125"
	PIN     string = "12344321"
)

func main() {
	homekit.OnLog(false)
	acc := homekit.NewAccessorySmartSpeaker(accessory.Info{Name: NAME, SerialNumber: SN, Model: MODEL, Manufacturer: homekit.Manufacturer, Firmware: homekit.Firmware})
	acc.SmartSpeaker.ConfiguredName.SetValue("Smart Speaker")
	acc.SmartSpeaker.Mute.SetValue(false)
	acc.SmartSpeaker.Volume.SetValue(100)
	llog := log.New(os.Stdout, fmt.Sprintf("[ %v / %v ] ", acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value()), log.Ldate|log.Ltime|log.Lmsgprefix)
	storage := hap.NewFsStore(fmt.Sprintf("./%s", acc.Info.SerialNumber.Value()))
	server, err := hap.NewServer(storage, acc.A)
	if err != nil {
		llog.Fatalf("error create hap server: %v\n", err)
	}
	llog.Printf("hap server create successful.\n")
	acc.OnExample()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-sig
		llog.Printf("stop program signal.\n")
		signal.Stop(sig)
		cancel()
	}()
	homekit.SetServer(server, ADDRESS, PIN)
	llog.Printf("hap server starting set, address %v, pin %v.\n", server.Addr, server.Pin)
	llog.Fatalf("hap server: %v\n", server.ListenAndServe(ctx))
}
