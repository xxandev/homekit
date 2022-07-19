package homekit

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

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
	Firmware             string = "1.2.3"
	Manufacturer         string = "alpr777"
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

type ConfigAccessory struct {
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	SN   string `json:"sn" xml:"sn"`
	Port uint16 `json:"port" xml:"port"`
	Pin  string `json:"pin,omitempty" xml:"pin,omitempty"`
}

func (a *ConfigAccessory) GetInfo(manufacturer, model, firmware string) accessory.Info {
	return accessory.Info{
		Name:         a.GetName(),
		SerialNumber: a.GetSN(),
		Manufacturer: manufacturer,
		Model:        model,
		Firmware:     firmware,
	}
}

func (a *ConfigAccessory) GetID() uint64 {
	return 1
}

func (a *ConfigAccessory) GetName() string {
	if len(a.Name) < 1 {
		return fmt.Sprintf("acc-%s", a.SN)
	}
	return a.Name
}

func (a *ConfigAccessory) GetSN() string {
	return a.SN
}

func (a *ConfigAccessory) GetPort() string {
	return fmt.Sprintf(":%v", a.Port)
}

func (a *ConfigAccessory) GetPin() string {
	a.Pin = strings.ReplaceAll(a.Pin, "-", "")
	if len(a.Pin) != 8 || !regexp.MustCompile(`^[0-9]+$`).MatchString(a.Pin) {
		return "13974268"
	}
	return a.Pin
}

func (a *ConfigAccessory) Valid() error {
	if len(a.SN) < 1 || !regexp.MustCompile(`^[a-zA-Z0-9\-]+$`).MatchString(a.SN) {
		return fmt.Errorf("invalid serial number, may contain at least 1 characters and consist of a-z, A-Z, 0-9, '-'")
	}
	if a.Port < 1 {
		return fmt.Errorf("invalid port")
	}
	if len(a.Pin) > 0 {
		a.Pin = strings.ReplaceAll(a.Pin, "-", "")
		if len(a.Pin) != 8 || !regexp.MustCompile(`^[0-9]+$`).MatchString(a.Pin) {
			return fmt.Errorf("invalid pin (does not consist of 8 numbers or not a number)")
		}
		for i := 0; i <= 9; i++ {
			if strings.Count(a.Pin, fmt.Sprint(i)) > 2 {
				return fmt.Errorf("invalid pin (numbers are repeated)")
			}
		}
		switch a.Pin {
		case "01234567", "76543210", "12345678", "87654321", "23456789", "98765432":
			return fmt.Errorf("invalid pin (too easy)")
		}
	} else {
		a.Pin = "13974268"
	}
	return nil
}

func (a *ConfigAccessory) Trial() {
	a.Name = fmt.Sprintf("acc-%v", rand.Intn(99)+time.Now().Second())
	a.SN = fmt.Sprintf("sn-%v", rand.Intn(99))
	a.Port = 10001
	a.Pin = "13974268"
}

type ConfigBridge struct {
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	SN   string `json:"sn" xml:"sn"`
	Port uint16 `json:"port" xml:"port"`
	Pin  string `json:"pin,omitempty" xml:"pin,omitempty"`
}

func (b *ConfigBridge) GetInfo(manufacturer, model, Firmware string) accessory.Info {
	return accessory.Info{
		Name:         b.GetName(),
		SerialNumber: b.GetSN(),
		Manufacturer: manufacturer,
		Model:        model,
		Firmware:     Firmware,
	}
}

func (b *ConfigBridge) GetID() uint64 {
	return 1
}

func (b *ConfigBridge) GetName() string {
	if len(b.Name) < 1 {
		b.Name = fmt.Sprintf("brg-%s", strings.ToLower(b.SN))
	}
	return b.Name
}

func (b *ConfigBridge) GetSN() string {
	return b.SN
}

func (b *ConfigBridge) GetPort() string {
	return fmt.Sprint(b.Port)
}

func (b *ConfigBridge) GetPin() string {
	b.Pin = strings.ReplaceAll(b.Pin, "-", "")
	if len(b.Pin) != 8 || !regexp.MustCompile(`^[0-9]+$`).MatchString(b.Pin) {
		b.Pin = "13974268"
	}
	return b.Pin
}

func (b *ConfigBridge) Valid() error {
	if len(b.SN) < 1 || !regexp.MustCompile(`^[a-zA-Z0-9\-]+$`).MatchString(b.SN) {
		return fmt.Errorf("invalid serial number, may contain at least 1 characters and consist of a-z, A-Z, 0-9, '-'")
	}
	if b.Port < 1 {
		return fmt.Errorf("invalid port")
	}
	if len(b.Pin) > 0 {
		b.Pin = strings.ReplaceAll(b.Pin, "-", "")
		if len(b.Pin) != 8 || !regexp.MustCompile(`^[0-9]+$`).MatchString(b.Pin) {
			return fmt.Errorf("invalid pin (does not consist of 8 numbers or not a number)")
		}
		for i := 0; i <= 9; i++ {
			if strings.Count(b.Pin, fmt.Sprint(i)) > 2 {
				return fmt.Errorf("invalid pin (numbers are repeated)")
			}
		}
		switch b.Pin {
		case "01234567", "76543210", "12345678", "87654321", "23456789", "98765432":
			return fmt.Errorf("invalid pin (too easy)")
		}
	} else {
		b.Pin = "13974268"
	}
	return nil
}

func (b *ConfigBridge) Trial() {
	b.Name = fmt.Sprintf("brg-%v", rand.Intn(time.Now().Second()))
	b.SN = fmt.Sprintf("sn-%v", rand.Intn(99))
	b.Port = 10001
	b.Pin = "13974268"
}

type ConfigSlaveAccessory struct {
	ID   uint64 `json:"id" xml:"id"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	SN   string `json:"sn,omitempty" xml:"sn,omitempty"`
}

func (sa *ConfigSlaveAccessory) GetInfo(manufacturer, model, firmware string) accessory.Info {
	return accessory.Info{
		Name:         sa.GetName(),
		SerialNumber: sa.GetSN(),
		Manufacturer: manufacturer,
		Model:        model,
		Firmware:     firmware,
	}
}

func (sa *ConfigSlaveAccessory) GetID() uint64 {
	return sa.ID
}

func (sa *ConfigSlaveAccessory) GetName() string {
	if len(sa.Name) > 1 {
		return sa.Name
	}
	sa.Name = fmt.Sprintf("acc-%d", sa.ID)
	return sa.Name
}

func (sa *ConfigSlaveAccessory) GetSN() string {
	if len(sa.SN) > 1 {
		return sa.SN
	}
	sa.SN = fmt.Sprintf("SL-ACC-%d", sa.ID)
	return sa.SN
}

func (sa *ConfigSlaveAccessory) Valid() error {
	if sa.ID < 2 {
		return fmt.Errorf("invalid id, cannot be less than 2 (id 1 bridge)")
	}
	return nil
}

func (sa *ConfigSlaveAccessory) Trial() {
	sa.ID = uint64(rand.Intn(150))
	sa.Name = fmt.Sprintf("acc-%d", sa.ID)
	sa.SN = fmt.Sprintf("SL-ACC-%d", sa.ID)
}
