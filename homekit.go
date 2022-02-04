package homekit

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/log"
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
	Revision             string = "1.2.3"
	Manufacturer         string = "alpr777"
	MaxBridgeAccessories int    = 150
)

type ConfigAccessory struct {
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	SN   string `json:"sn" xml:"sn"`
	Port uint16 `json:"port" xml:"port"`
	Pin  string `json:"pin,omitempty" xml:"pin,omitempty"`
}

func (a *ConfigAccessory) GetInfo(manufacturer, model, revision string) accessory.Info {
	return accessory.Info{
		Name:             a.GetName(),
		SerialNumber:     a.GetSN(),
		Manufacturer:     manufacturer,
		Model:            model,
		FirmwareRevision: revision,
	}
}

func (a *ConfigAccessory) GetConfigHC(storagepath string) hc.Config {
	return hc.Config{
		StoragePath: storagepath,
		Pin:         a.GetPin(),
		Port:        a.GetPort(),
	}
}

//OnDebug
//  if use systemd, recommended flag 64, or 77 for full debug
func (a *ConfigAccessory) OnDebug(active bool) {
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
	return fmt.Sprint(a.Port)
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

func (b *ConfigBridge) GetInfo(manufacturer, model, revision string) accessory.Info {
	return accessory.Info{
		Name:             b.GetName(),
		SerialNumber:     b.GetSN(),
		Manufacturer:     manufacturer,
		Model:            model,
		FirmwareRevision: revision,
	}
}

func (b *ConfigBridge) GetConfigHC(storagepath string) hc.Config {
	return hc.Config{
		StoragePath: storagepath,
		Pin:         b.GetPin(),
		Port:        b.GetPort(),
	}
}

//OnDebug
//  if use systemd, recommended flag 64, or 77 for full debug
func (b *ConfigBridge) OnDebug(active bool) {
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

func (sa *ConfigSlaveAccessory) GetInfo(manufacturer, model, revision string) accessory.Info {
	return accessory.Info{
		ID:               sa.ID,
		Name:             sa.GetName(),
		SerialNumber:     sa.GetSN(),
		Manufacturer:     manufacturer,
		Model:            model,
		FirmwareRevision: revision,
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
