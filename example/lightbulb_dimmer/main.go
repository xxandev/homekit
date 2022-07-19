package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/brutella/hap"
	"github.com/brutella/hap/accessory"
	"github.com/xxandev/homekit"
)

const (
	NAME    string = "Lightbulb"
	SN      string = "EX-Lightbulb"
	MODEL   string = "HAP-LBDM"
	ADDRESS string = ":11110"
	PIN     string = "12344321"
)

func main() {
	homekit.OnLog(false)
	acc := homekit.NewAccessoryLightbulbDimmer(accessory.Info{Name: "Lightbulb", SerialNumber: "Ex-Lb-Dm", Model: "HAP-LB-DM", Manufacturer: homekit.Manufacturer, Firmware: homekit.Firmware})
	llog := log.New(os.Stdout, fmt.Sprintf("[ %v / %v ] ", acc.A.Info.SerialNumber.Value(), acc.A.Info.Name.Value()), log.Ldate|log.Ltime|log.Lmsgprefix)
	storage := hap.NewFsStore(fmt.Sprintf("./%s", acc.Info.SerialNumber.Value()))
	server, err := hap.NewServer(storage, acc.A)
	if err != nil {
		llog.Fatalf("error create hap server: %v\n", err)
	}
	llog.Printf("hap server create successful.\n")

	go func() {
		for range time.Tick(30 * time.Second) {
			acc.LightbulbDimmer.On.SetValue(!acc.LightbulbDimmer.On.Value())
			llog.Printf("acc lightbulb update on: %T - %v \n", acc.LightbulbDimmer.On.Value(), acc.LightbulbDimmer.On.Value())
		}
	}()
	acc.LightbulbDimmer.On.OnValueRemoteUpdate(func(v bool) {
		llog.Printf("acc lightbulb remote update on: %T - %v \n", v, v)
	})
	acc.LightbulbDimmer.Brightness.OnValueRemoteUpdate(func(v int) {
		llog.Printf("acc lightbulb remote update brightness: %T - %v \n", v, v)
	})

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
