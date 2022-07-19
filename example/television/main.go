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
	"github.com/brutella/hap/characteristic"
	"github.com/xxandev/homekit"
)

const (
	NAME    string = "Television"
	SN      string = "EX-Tv"
	MODEL   string = "HAP-TV"
	ADDRESS string = ":11127"
	PIN     string = "12344321"
)

func main() {
	homekit.OnLog(false)
	acc := homekit.NewAccessoryTelevision(accessory.Info{Name: NAME, SerialNumber: SN, Model: MODEL, Manufacturer: homekit.Manufacturer, Firmware: homekit.Firmware})
	_ = acc.AddInputSource(1, "HDMI 1", characteristic.InputSourceTypeHdmi)
	_ = acc.AddInputSource(2, "HDMI 2", characteristic.InputSourceTypeHdmi)
	_ = acc.AddInputSource(3, "YouTube", characteristic.InputSourceTypeApplication)
	_ = acc.AddInputSource(4, "Airplay", characteristic.InputSourceTypeAirplay)
	_ = acc.AddInputSource(5, "AndroidTV", characteristic.InputSourceTypeOther)
	_ = acc.AddInputSource(6, "AppleTV", characteristic.InputSourceTypeOther)
	_ = acc.AddInputSource(7, "Xbox", characteristic.InputSourceTypeOther)
	_ = acc.AddInputSource(8, "Paystation", characteristic.InputSourceTypeOther)
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
