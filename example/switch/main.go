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
	flag.StringVar(&config.Name, "n", "Switch", "homekit accessory name")
	flag.StringVar(&config.SN, "sn", "Ex-Switch", "homekit accessory serial number")
	flag.StringVar(&config.Host, "h", "", "homekit host, example: 192.168.1.xxx")
	flag.UintVar(&config.Port, "p", 10725, "homekit port, example: 10101, 10102...")
	flag.StringVar(&config.Pin, "pin", "19378246", "homekit pin, example: 82143697, 13974682")
	flag.Parse()

	homekit.OnLog(debug)
}

func main() {
	acc := homekit.NewAccessorySwitch(config.GetInfo("Ex-Switch"))
	log.SetPrefix(fmt.Sprintf("[%T] <%v> ", acc, acc.GetSN()))
	storage := hap.NewFsStore(fmt.Sprintf("./%s", acc.GetSN()))
	server, err := hap.NewServer(storage, acc.GetAccessory())
	if err != nil {
		log.Fatalf("error create hap server: %v\n", err)
	}
	log.Printf("hap server create successful.\n")

	go func() {
		for range time.Tick(time.Millisecond * 300000) {
			acc.Switch.On.SetValue(!acc.Switch.On.Value())
			log.Printf("update on: %[1]T - %[1]v\n", acc.Switch.On.Value())
		}
	}()
	acc.Switch.On.OnValueRemoteUpdate(func(v bool) {
		log.Printf("remote update on: %[1]T - %[1]v\n", v)
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
