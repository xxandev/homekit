package homekit

import (
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
