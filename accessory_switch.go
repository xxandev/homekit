package homekit

import (
	"encoding/json"
	"fmt"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

type StatusAccessorySwitch struct {
	SN        string `json:"serial_numb"`
	TargetOn  bool   `json:"target_on"`
	CurrentOn bool   `json:"current_on"`
}

//AccessorySwitch struct
type AccessorySwitch struct {
	*accessory.Accessory
	Switch *service.Switch
}

// NewAccessorySwitch returns AccessorySwitch (args... are not used)
func NewAccessorySwitch(info accessory.Info, args ...interface{}) *AccessorySwitch {
	acc := AccessorySwitch{}
	acc.Accessory = accessory.New(info, accessory.TypeSwitch)
	acc.Switch = service.NewSwitch()
	acc.AddService(acc.Switch.Service)
	return &acc
}

func (acc *AccessorySwitch) OnValuesRemoteUpdates(fn func()) {
	acc.Switch.On.OnValueRemoteUpdate(func(_ bool) { fn() })
}

func (acc *AccessorySwitch) GetType() uint8                     { return uint8(acc.Accessory.Type) }
func (acc *AccessorySwitch) GetID() uint64                      { return acc.Accessory.ID }
func (acc *AccessorySwitch) GetSN() string                      { return acc.Accessory.Info.SerialNumber.GetValue() }
func (acc *AccessorySwitch) GetName() string                    { return acc.Accessory.Info.SerialNumber.GetValue() }
func (acc *AccessorySwitch) GetAccessory() *accessory.Accessory { return acc.Accessory }

func (acc *AccessorySwitch) GetJsonValues() string {
	type box struct {
		Status StatusAccessorySwitch `json:"status_accessory_switch"`
	}
	asb := box{StatusAccessorySwitch{
		SN:        acc.Accessory.Info.SerialNumber.GetValue(),
		CurrentOn: acc.Switch.On.GetValue(),
		TargetOn:  acc.Switch.On.GetValue(),
	}}
	res, _ := json.Marshal(asb)
	return string(res)
}

func (acc *AccessorySwitch) SetJsonValues(v string) error {
	type box struct {
		Status StatusAccessorySwitch `json:"status_accessory_switch"`
	}
	asb := box{}
	if err := json.Unmarshal([]byte(v), &asb); err != nil {
		return err
	}
	if asb.Status.SN != acc.Accessory.Info.SerialNumber.GetValue() {
		return fmt.Errorf("")
	}
	acc.Switch.On.SetValue(asb.Status.TargetOn)
	return nil
}

func (acc *AccessorySwitch) OnValuesRemoteJsonUpdate(fn func(v string)) {
	acc.Switch.On.OnValueRemoteUpdate(func(on bool) { fn(acc.GetJsonValues()) })
}
