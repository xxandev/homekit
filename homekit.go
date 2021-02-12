package homekit

import "github.com/brutella/hc/accessory"

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

//argToBool -
func argToBool(value interface{}, def bool) bool {
	if val, ok := value.(bool); ok == true {
		return val
	}
	return def
}

//argToInt - int signed, either 32 or 64 bits
func argToInt(value interface{}, def int) int {
	switch val := value.(type) {
	case int:
		return val
	case int8:
		return int(val)
	case int16:
		return int(val)
	case int32:
		return int(val)
	case int64:
		return int(val)
	}
	return def
}

//argToInt32 - int32 signed 32-bit integers (-2147483648 to 2147483647)
func argToInt32(value interface{}, def int32) int32 {
	switch val := value.(type) {
	case int32:
		return val
	case int:
		return int32(val)
	case int8:
		return int32(val)
	case int16:
		return int32(val)
	case int64:
		return int32(val)
	}
	return def
}

//argToInt64 - int64 signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
func argToInt64(value interface{}, def int64) int64 {
	switch val := value.(type) {
	case int64:
		return val
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	}
	return def
}

//argToFloat32 - float32 the set of all IEEE-754 32-bit floating-point numbers
func argToFloat32(value interface{}, def float32) float32 {
	switch val := value.(type) {
	case float32:
		return val
	case float64:
		return float32(val)
	case int:
		return float32(val)
	}
	return def
}

//argToFloat64 - float64 the set of all IEEE-754 64-bit floating-point numbers
func argToFloat64(value interface{}, def float64) float64 {
	switch val := value.(type) {
	case float64:
		return val
	case float32:
		return float64(val)
	case int:
		return float64(val)
	}
	return def
}
