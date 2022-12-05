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
	flag.StringVar(&config.Name, "n", "SHumidity", "homekit accessory name")
	flag.StringVar(&config.SN, "sn", "Ex-SHum", "homekit accessory serial number")
	flag.StringVar(&config.Host, "h", "", "homekit host, example: 192.168.1.xxx")
	flag.UintVar(&config.Port, "p", 10717, "homekit port, example: 10101, 10102...")
	flag.StringVar(&config.Pin, "pin", "19378246", "homekit pin, example: 82143697, 13974682")
	flag.Parse()

	homekit.OnLog(debug)
}

func main() {
	acc := homekit.NewAccessorySensorHumidity(config.GetInfo("Ex-SHum"))
	log.SetPrefix(fmt.Sprintf("[%T] <%v> ", acc, acc.GetSN()))
	storage := hap.NewFsStore(fmt.Sprintf("./%s", acc.GetSN()))
	server, err := hap.NewServer(storage, acc.GetAccessory())
	if err != nil {
		log.Fatalf("error create hap server: %v\n", err)
	}
	log.Printf("hap server create successful.\n")

	go func() {
		for range time.Tick(time.Millisecond * 3000) {
			if acc.HumiditySensor.CurrentRelativeHumidity.Value() >= acc.HumiditySensor.CurrentRelativeHumidity.MaxValue() {
				acc.HumiditySensor.CurrentRelativeHumidity.SetValue(acc.HumiditySensor.CurrentRelativeHumidity.MinValue())
			} else {
				acc.HumiditySensor.CurrentRelativeHumidity.SetValue(acc.HumiditySensor.CurrentRelativeHumidity.Value() + 0.25)
			}
			log.Printf("update current hum: %[1]T - %[1]v\n", acc.HumiditySensor.CurrentRelativeHumidity.Value())
		}
	}()

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
