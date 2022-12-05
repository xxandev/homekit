package homekit

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/brutella/hap"
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/log"
)

//https://github.com/homebridge/HAP-NodeJS/blob/master/src/lib/Accessory.ts >> export const enum Categories
const (
	AccessoryTypeUnknown            byte = 0
	AccessoryTypeOther              byte = 1
	AccessoryTypeBridge             byte = 2
	AccessoryTypeFan                byte = 3
	AccessoryTypeGarageDoorOpener   byte = 4
	AccessoryTypeLightbulb          byte = 5
	AccessoryTypeDoorLock           byte = 6
	AccessoryTypeOutlet             byte = 7
	AccessoryTypeSwitch             byte = 8
	AccessoryTypeThermostat         byte = 9
	AccessoryTypeSensor             byte = 10
	AccessoryTypeSecuritySystem     byte = 11
	AccessoryTypeDoor               byte = 12
	AccessoryTypeWindow             byte = 13
	AccessoryTypeWindowCovering     byte = 14
	AccessoryTypeProgrammableSwitch byte = 15
	AccessoryTypeRangeExtender      byte = 16
	AccessoryTypeIPCamera           byte = 17
	AccessoryTypeVideoDoorbell      byte = 18
	AccessoryTypeAirPurifier        byte = 19
	AccessoryTypeHeater             byte = 20
	AccessoryTypeAirConditioner     byte = 21
	AccessoryTypeHumidifier         byte = 22
	AccessoryTypeDehumidifier       byte = 23
	AccessoryTypeAppleTV            byte = 24
	AccessoryTypeSpeaker            byte = 26
	AccessoryTypeAirport            byte = 27
	AccessoryTypeSprinkler          byte = 28
	AccessoryTypeFaucet             byte = 29
	AccessoryTypeShowerSystems      byte = 30
	AccessoryTypeTelevision         byte = 31
	AccessoryTypeRemoteControl      byte = 32
	AccessoryTypeWiFiRouter         byte = 33
	AccessoryTypeAudioReceiver      byte = 34
	AccessoryTypeTVSetTopBox        byte = 35
	AccessoryTypeTVStick            byte = 36
)

const (
	FIRMWARE             string = "1.3"
	MANUFACTURER         string = "XXanDev/brutella"
	MaxBridgeAccessories int    = 150
)

func SetServer(s *hap.Server, address, pin string) error {
	// valid

	//set
	s.Addr = address
	s.Pin = pin
	return nil
}

//OnLog - on/off hap log
//  if use systemd, recommended flag 64, or 77 for full debug
func OnLog(active bool) {
	log.Debug.SetFlags(67) // if use systemd, recommended flag 64, or 77 for full debug
	log.Info.SetFlags(67)  // if use systemd, recommended flag 64, or 77 for full debug
	log.Debug.SetPrefix("[HAP_DBG]")
	log.Info.SetPrefix("[HAP_INFO]")
	log.Debug.Disable()
	log.Info.Disable()
	if active {
		log.Debug.Enable()
		log.Info.Enable()
	}
}

func LogOff() {
	log.Debug.SetFlags(67) // if use systemd, recommended flag 64, or 77 for full debug
	log.Info.SetFlags(67)  // if use systemd, recommended flag 64, or 77 for full debug
	log.Debug.SetPrefix("[HAP_DBG]")
	log.Info.SetPrefix("[HAP_INFO]")
	log.Debug.Disable()
	log.Info.Disable()
}

func LogOn() {
	log.Debug.SetFlags(67) // if use systemd, recommended flag 64, or 77 for full debug
	log.Info.SetFlags(67)  // if use systemd, recommended flag 64, or 77 for full debug
	log.Debug.SetPrefix("[HAP_DBG]")
	log.Info.SetPrefix("[HAP_INFO]")
	log.Debug.Enable()
	log.Info.Enable()
}

type AccessoryConfig struct {
	ID   uint64 `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	SN   string `json:"sn,omitempty" xml:"sn,omitempty"`
	Host string `json:"host,omitempty" xml:"host,omitempty"`
	Port uint   `json:"port,omitempty" xml:"port,omitempty"`
	Pin  string `json:"pin,omitempty" xml:"pin,omitempty"`
}

func (ac *AccessoryConfig) GetID() uint64 {
	return ac.ID
}

func (ac *AccessoryConfig) GetSN() string {
	if len(ac.SN) > 0 {
		return ac.SN
	}
	return fmt.Sprintf("hk-id-%d", ac.ID)
}

func (ac *AccessoryConfig) GetName() string {
	if len(ac.Name) < 1 {
		return fmt.Sprintf("acc-%s", ac.GetSN())
	}
	return ac.Name
}

func (ac *AccessoryConfig) GetPin() string {
	pin := regexp.MustCompile(`\D`).ReplaceAllString(ac.Pin, "")
	if len(pin) < 8 {
		return "19283746"
	}
	pin = string(pin[0:8])
	for i := 0; i <= 9; i++ {
		if strings.Count(pin, fmt.Sprint(i)) > 2 {
			return "19283746"
		}
	}
	switch pin {
	case "01234567", "76543210", "12345678", "87654321", "23456789", "98765432":
		return "19283746"
	}
	return pin
}

func (ac *AccessoryConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", ac.Host, ac.Port)
}

func (ac *AccessoryConfig) GetInfo(model string) accessory.Info {
	return accessory.Info{
		Name:         ac.GetName(),
		SerialNumber: ac.GetSN(),
		Manufacturer: MANUFACTURER,
		Model:        model,
		Firmware:     FIRMWARE,
	}
}
