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
	flag.StringVar(&config.Name, "n", "Door", "homekit accessory name")
	flag.StringVar(&config.SN, "sn", "Ex-Door", "homekit accessory serial number")
	flag.StringVar(&config.Host, "h", "", "homekit host, example: 192.168.1.xxx")
	flag.UintVar(&config.Port, "p", 10703, "homekit port, example: 10101, 10102...")
	flag.StringVar(&config.Pin, "pin", "19736428", "homekit pin, example: 82143697, 13974682")
	flag.Parse()

	homekit.OnLog(debug)
}

func main() {
	acc := homekit.NewAccessoryDoor(config.GetInfo("Ex-Door"))
	log.SetPrefix(fmt.Sprintf("[%T] <%v> ", acc, acc.GetSN()))
	storage := hap.NewFsStore(fmt.Sprintf("./%s", acc.GetSN()))
	server, err := hap.NewServer(storage, acc.GetAccessory())
	if err != nil {
		log.Fatalf("error create hap server: %v\n", err)
	}
	log.Printf("hap server create successful.\n")
	command := make(chan int, 5)
	go func() {
		callbackT := time.NewTimer(time.Millisecond * 10000)
		for {
			select {
			case cmd := <-command:
				callbackT.Stop()
				callbackT.Reset(time.Millisecond * 10000)
				acc.Door.TargetPosition.SetValue(cmd)
				acc.Door.CurrentPosition.SetValue(cmd)
				log.Printf("remote update target position, current: %[1]T - %[1]v, target: %[2]T - %[2]v\n", acc.Door.CurrentPosition.Value(), acc.Door.TargetPosition.Value())
			case <-callbackT.C:
				acc.Door.TargetPosition.SetValue(0)
				acc.Door.CurrentPosition.SetValue(0)
				log.Printf("update position, current: %[1]T - %[1]v, target: %[2]T - %[2]v\n", acc.Door.CurrentPosition.Value(), acc.Door.TargetPosition.Value())
			case <-time.Tick(time.Millisecond * 320000):
				callbackT.Stop()
				callbackT.Reset(time.Millisecond * 10000)
				acc.Door.TargetPosition.SetValue(100)
				acc.Door.CurrentPosition.SetValue(100)
				log.Printf("update position, current: %[1]T - %[1]v, target: %[2]T - %[2]v\n", acc.Door.CurrentPosition.Value(), acc.Door.TargetPosition.Value())
			}
		}
	}()
	acc.Door.TargetPosition.OnValueRemoteUpdate(func(v int) { command <- v })

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
