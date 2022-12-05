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
	flag.StringVar(&config.Name, "n", "HumDehum", "homekit accessory name")
	flag.StringVar(&config.SN, "sn", "Ex-HumDehum", "homekit accessory serial number")
	flag.StringVar(&config.Host, "h", "", "homekit host, example: 192.168.1.xxx")
	flag.UintVar(&config.Port, "p", 10707, "homekit port, example: 10101, 10102...")
	flag.StringVar(&config.Pin, "pin", "19736428", "homekit pin, example: 82143697, 13974682")
	flag.Parse()

	homekit.OnLog(debug)
}

func main() {
	acc := homekit.NewAccessoryHumidifierDehumidifier(config.GetInfo("Ex-HumDehum"))
	log.SetPrefix(fmt.Sprintf("[%T] <%v> ", acc, acc.GetSN()))
	storage := hap.NewFsStore(fmt.Sprintf("./%s", acc.GetSN()))
	server, err := hap.NewServer(storage, acc.GetAccessory())
	if err != nil {
		log.Fatalf("error create hap server: %v\n", err)
	}
	log.Printf("hap server create successful.\n")

	acc.HumidifierDehumidifier.Active.OnValueRemoteUpdate(func(v int) {
		log.Printf("remote update active: %[1]T - %[1]v\n", v)
	})
	acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.OnValueRemoteUpdate(func(v int) {
		log.Printf("remote update target state: %[1]T - %[1]v\n", v)
	})
	acc.HumidifierDehumidifier.RelativeHumidityDehumidifierThreshold.OnValueRemoteUpdate(func(v float64) {
		log.Printf("remote update relative dehumidifier threshold: %[1]T - %[1]v\n", v)
	})
	acc.HumidifierDehumidifier.RelativeHumidityHumidifierThreshold.OnValueRemoteUpdate(func(v float64) {
		log.Printf("remote update relative humidifier threshold: %[1]T - %[1]v\n", v)
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
