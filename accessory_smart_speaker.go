package homekit

import (
	"github.com/brutella/hap/accessory"
	haps "github.com/xxandev/homekit/hap-service"
)

//AccessorySmartSpeaker struct
type AccessorySmartSpeaker struct {
	*accessory.A
	SmartSpeaker *haps.SmartSpeaker
}

func (acc *AccessorySmartSpeaker) GetType() byte {
	return acc.A.Type
}

func (acc *AccessorySmartSpeaker) GetID() uint64 {
	return acc.A.Id
}

func (acc *AccessorySmartSpeaker) SetID(id uint64) {
	acc.A.Id = id
}

func (acc *AccessorySmartSpeaker) GetSN() string {
	return acc.A.Info.SerialNumber.Value()
}

func (acc *AccessorySmartSpeaker) GetName() string {
	return acc.A.Info.Name.Value()
}

func (acc *AccessorySmartSpeaker) GetAccessory() *accessory.A {
	return acc.A
}

//NewAccessorySmartSpeaker returns *SmartSpeaker.
//  (COMPATIBILITY)  - left for compatibility, recommended NewAcc...(id, info, args..)
//
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccessorySmartSpeaker(info accessory.Info, args ...interface{}) *AccessorySmartSpeaker {
	acc := AccessorySmartSpeaker{}
	acc.A = accessory.New(info, AccessoryTypeSpeaker)
	acc.SmartSpeaker = haps.NewSmartSpeaker()
	acc.AddS(acc.SmartSpeaker.S)
	return &acc
}

//NewAccSmartSpeaker returns *SmartSpeaker.
//  HomeKit requires that every accessory has a unique id, which must not change between system restarts.
//  The best would be to specify the unique id for every accessory yourself.
//
//  id (uint64) - accessory aid
//  info (accessory.Info) - struct accessory.Info{Name, SerialNumber, Manufacturer, Model, Firmware string}
//  args... are not used
func NewAccSmartSpeaker(id uint64, info accessory.Info, args ...interface{}) *AccessorySmartSpeaker {
	acc := AccessorySmartSpeaker{}
	acc.A = accessory.New(info, AccessoryTypeSpeaker)
	acc.SmartSpeaker = haps.NewSmartSpeaker()
	acc.AddS(acc.SmartSpeaker.S)
	acc.A.Id = id
	return &acc
}

func (acc *AccessorySmartSpeaker) OnValuesRemoteUpdates(fn func()) {}
