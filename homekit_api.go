package homekit

import (
	"github.com/brutella/hc/accessory"
)

type AccessoryAPI interface {
	GetType() uint8
	GetID() uint64
	GetSN() string
	GetName() string
	GetAccessory() *accessory.Accessory

	GetValues() (vals []interface{})
	GetValuesJson() string

	SetValues(val interface{}, vals ...interface{})
	SetValuesJson(vals string)

	OnValueRemoteUpdate(fn func(pointer uint8, val interface{}))
	OnValuesRemoteUpdate(fn func(vals ...interface{}))
	OnValuesRemoteUpdateJson(fn func(json string))
}
