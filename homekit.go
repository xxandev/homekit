package homekit

import (
	"fmt"
	"log"
	"sync"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

//https://github.com/homebridge/HAP-NodeJS/blob/master/src/lib/Accessory.ts >> export const enum Categories
const (
	AccessoryTypeUnknown            accessory.AccessoryType = 0
	AccessoryTypeOther              accessory.AccessoryType = 1
	AccessoryTypeBridge             accessory.AccessoryType = 2
	AccessoryTypeFan                accessory.AccessoryType = 3
	AccessoryTypeGarageDoorOpener   accessory.AccessoryType = 4
	AccessoryTypeLightbulb          accessory.AccessoryType = 5
	AccessoryTypeDoorLock           accessory.AccessoryType = 6
	AccessoryTypeOutlet             accessory.AccessoryType = 7
	AccessoryTypeSwitch             accessory.AccessoryType = 8
	AccessoryTypeThermostat         accessory.AccessoryType = 9
	AccessoryTypeSensor             accessory.AccessoryType = 10
	AccessoryTypeSecuritySystem     accessory.AccessoryType = 11
	AccessoryTypeDoor               accessory.AccessoryType = 12
	AccessoryTypeWindow             accessory.AccessoryType = 13
	AccessoryTypeWindowCovering     accessory.AccessoryType = 14
	AccessoryTypeProgrammableSwitch accessory.AccessoryType = 15
	AccessoryTypeRangeExtender      accessory.AccessoryType = 16
	AccessoryTypeIPCamera           accessory.AccessoryType = 17
	AccessoryTypeVideoDoorbell      accessory.AccessoryType = 18
	AccessoryTypeAirPurifier        accessory.AccessoryType = 19
	AccessoryTypeHeater             accessory.AccessoryType = 20
	AccessoryTypeAirConditioner     accessory.AccessoryType = 21
	AccessoryTypeHumidifier         accessory.AccessoryType = 22
	AccessoryTypeDehumidifier       accessory.AccessoryType = 23
	AccessoryTypeAppleTV            accessory.AccessoryType = 24
	AccessoryTypeSpeaker            accessory.AccessoryType = 26
	AccessoryTypeAirport            accessory.AccessoryType = 27
	AccessoryTypeSprinklers         accessory.AccessoryType = 28
	AccessoryTypeFaucets            accessory.AccessoryType = 29
	AccessoryTypeShowerSystems      accessory.AccessoryType = 30
	AccessoryTypeTelevision         accessory.AccessoryType = 31
	AccessoryTypeRemoteControl      accessory.AccessoryType = 32
	AccessoryTypeWiFiRouter         accessory.AccessoryType = 33
	AccessoryTypeAudioReceiver      accessory.AccessoryType = 34
	AccessoryTypeTVSetTopBox        accessory.AccessoryType = 35
	AccessoryTypeTVStick            accessory.AccessoryType = 36
)

const (
	Revision             string = "1.2.4"
	Manufacturer         string = "alpr777"
	MaxBridgeAccessories int    = 150
)

type HomeKit struct {
	mutex   sync.Mutex
	logger  *log.Logger
	started bool
	accs    []*accessory.Accessory
	apimap  map[uint64]AccessoryAPI
}

// NewContainer returns a container.
func New() *HomeKit {
	return &HomeKit{
		started: false,
		accs:    make([]*accessory.Accessory, 0),
		apimap:  make(map[uint64]AccessoryAPI),
	}
}

func (hk *HomeKit) SetLog(logger *log.Logger) {
	hk.mutex.Lock()
	hk.logger = logger
	hk.mutex.Unlock()
}

func (hk *HomeKit) AddAccessory(accs ...AccessoryAPI) error {
	hk.mutex.Lock()
	defer hk.mutex.Unlock()
	if hk.started {
		if hk.logger != nil {
			hk.logger.Printf("error add hap accessory, adding is possible only before run\n")
		}
		return fmt.Errorf("adding is possible only before run")
	}
	if accs == nil || len(accs) < 1 {
		if hk.logger != nil {
			hk.logger.Printf("error add hap accessory, no accessories selected to add\n")
		}
		return fmt.Errorf("no accessories selected to add ")
	}
	l := len(hk.accs)
	for n, acc := range accs {
		if acc == nil {
			if hk.logger != nil {
				hk.logger.Printf("error add hap accessory, max quantity \n")
			}
			continue
		}
		if l+n+1 > MaxBridgeAccessories {
			if hk.logger != nil {
				hk.logger.Printf("error add hap accessory, max quantity %v\n", MaxBridgeAccessories)
			}
			return fmt.Errorf("no accessories selected to add ")
		}
		/*
			+++
			check acc.GetID() > acc.GetID() == 0
		*/
		if hk.apimap[acc.GetID()] != nil {
			if hk.logger != nil {
				hk.logger.Printf("error add hap accessory %v / %v, id %v already taken\n", acc.GetSN(), acc.GetName(), acc.GetID())
			}
			return fmt.Errorf("id %v already taken", acc.GetID())
		}
		hk.apimap[acc.GetID()] = acc
		hk.accs = append(hk.accs, acc.GetAccessory())
	}
	return nil
}

func (hk *HomeKit) GetAccessory(id uint64) (AccessoryAPI, error) {
	hk.mutex.Lock()
	defer hk.mutex.Unlock()
	if hk.apimap[id] != nil {
		if hk.logger != nil {
			hk.logger.Printf("error get, accessory id %v not found\n", id)
		}
		return nil, fmt.Errorf("accessory id %v not found", id)
	}
	return hk.apimap[id], nil
}

//
func (hk *HomeKit) Run(BridgeInfo accessory.Info, TransportConfig hc.Config) error {
	bridge := accessory.NewBridge(BridgeInfo)
	transp, err := hc.NewIPTransport(TransportConfig, bridge.Accessory, hk.accs...)
	if err != nil {
		if hk.logger != nil {
			hk.logger.Printf("error create hap transport %v / %v: %v\n", bridge.Info.SerialNumber.GetValue(), bridge.Info.Name.GetValue(), err)
		}
		return err
	}
	if hk.logger != nil {
		hk.logger.Printf("hap transport %v / %v running.\n", bridge.Info.SerialNumber.GetValue(), bridge.Info.Name.GetValue())
		defer hk.logger.Printf("hap transport %v / %v stopped.\n", bridge.Info.SerialNumber.GetValue(), bridge.Info.Name.GetValue())
	}
	hk.started = true
	hc.OnTermination(func() { <-transp.Stop() })
	transp.Start()
	return nil
}
