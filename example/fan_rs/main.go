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
	flag.StringVar(&config.Name, "n", "FanRS", "homekit accessory name")
	flag.StringVar(&config.SN, "sn", "Ex-FanRS", "homekit accessory serial number")
	flag.StringVar(&config.Host, "h", "", "homekit host, example: 192.168.1.xxx")
	flag.UintVar(&config.Port, "p", 10704, "homekit port, example: 10101, 10102...")
	flag.StringVar(&config.Pin, "pin", "19736428", "homekit pin, example: 82143697, 13974682")
	flag.Parse()

	homekit.OnLog(debug)
}

func main() {
	acc := homekit.NewAccessoryFanRS(config.GetInfo("Ex-FanRS"))
	log.SetPrefix(fmt.Sprintf("[%T] <%v> ", acc, acc.GetSN()))
	storage := hap.NewFsStore(fmt.Sprintf("./%s", acc.GetSN()))
	server, err := hap.NewServer(storage, acc.GetAccessory())
	if err != nil {
		log.Fatalf("error create hap server: %v\n", err)
	}
	log.Printf("hap server create successful.\n")

	go func() {
		for range time.Tick(time.Millisecond * 180000) {
			acc.Fan.On.SetValue(true)
			acc.Fan.RotationSpeed.SetValue(50)
			log.Printf("update, on: %[1]T - %[1]v, speed %[2]T - %[2]v\n", acc.Fan.On.Value(), acc.Fan.RotationSpeed.Value())
			time.Sleep(time.Millisecond * 30000)
			acc.Fan.RotationSpeed.SetValue(100)
			log.Printf("update, on: %[1]T - %[1]v, speed %[2]T - %[2]v\n", acc.Fan.On.Value(), acc.Fan.RotationSpeed.Value())
			time.Sleep(time.Millisecond * 60000)
			acc.Fan.On.SetValue(false)
			log.Printf("update, on: %[1]T - %[1]v, speed %[2]T - %[2]v\n", acc.Fan.On.Value(), acc.Fan.RotationSpeed.Value())
		}
	}()
	acc.Fan.On.OnValueRemoteUpdate(func(v bool) {
		log.Printf("remote update on: %[1]T - %[1]v\n", v)
	})
	acc.Fan.RotationSpeed.OnValueRemoteUpdate(func(v float64) {
		log.Printf("remote update rotation speed: %[1]T - %[1]v\n", v)
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
	log.Printf("hap server starting set, address %v, pin %v.\n", server.Addr, server.Pin)
	log.Fatalf("hap server: %v\n", server.ListenAndServe(ctx))
}
