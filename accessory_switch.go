package homekit

import (
	"encoding/json"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

type AccessorySwitchStatuses struct {
	On bool `json:"on"`
}

type AccessorySwitchBox struct {
	Switch   string                  `json:"switch"`
	Statuses AccessorySwitchStatuses `json:"statuses"`
}

func (box *AccessorySwitchBox) Marshal() string {
	res, _ := json.Marshal(box)
	return string(res)
}

func (box *AccessorySwitchBox) Unmarshal(s string) error {
	return json.Unmarshal([]byte(s), &box)
}

//AccessorySwitch struct
type AccessorySwitch struct {
	// mutex sync.Mutex
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

func (acc *AccessorySwitch) GetType() uint8                     { return uint8(acc.Accessory.Type) }
func (acc *AccessorySwitch) GetID() uint64                      { return acc.Accessory.ID }
func (acc *AccessorySwitch) GetSN() string                      { return acc.Accessory.Info.SerialNumber.GetValue() }
func (acc *AccessorySwitch) GetName() string                    { return acc.Accessory.Info.SerialNumber.GetValue() }
func (acc *AccessorySwitch) GetAccessory() *accessory.Accessory { return acc.Accessory }

//GetValues - accessory get values
//  vals[0](bool) - accessory switch on value
func (acc *AccessorySwitch) GetValues() (vals []interface{}) {
	vals = append(vals,
		acc.Switch.On.GetValue(),
	)
	return
}

func (acc *AccessorySwitch) GetValuesJson() string {
	box := AccessorySwitchBox{
		Switch: acc.Accessory.Info.SerialNumber.GetValue(),
		Statuses: AccessorySwitchStatuses{
			On: acc.Switch.On.GetValue(),
		},
	}
	return box.Marshal()
}

//GetValues - accessory get values
//  val(bool) - accessory switch on value
func (acc *AccessorySwitch) SetValues(val interface{}, vals ...interface{}) {
	if v, ok := val.(bool); ok {
		acc.Switch.On.SetValue(v)
	}
}

func (acc *AccessorySwitch) SetValuesJson(vals string) {
	box := AccessorySwitchBox{}
	if err := box.Unmarshal(vals); err != nil || box.Switch != acc.Accessory.Info.SerialNumber.GetValue() {
		return
	}
	acc.Switch.On.SetValue(box.Statuses.On)
}

//OnValueRemoteUpdate
//  0x1(bool) - accessory switch on > on value remote update
func (acc *AccessorySwitch) OnValueRemoteUpdate(fn func(charact uint8, val interface{})) {
	acc.Switch.On.OnValueRemoteUpdate(func(on bool) { fn(0x1, on) })
}

//OnValuesRemoteUpdate
//  vals[0](bool) - accessory switch on > on value remote update
func (acc *AccessorySwitch) OnValuesRemoteUpdate(fn func(vals ...interface{})) {
	acc.Switch.On.OnValueRemoteUpdate(func(on bool) { fn(on) })
}

func (acc *AccessorySwitch) OnValuesRemoteUpdateJson(fn func(vals string)) {
	acc.Switch.On.OnValueRemoteUpdate(func(on bool) { fn(acc.GetValuesJson()) })
}

func (acc *AccessorySwitch) OnValuesRemoteUpdateEmpty(fn func()) {
	acc.Switch.On.OnValueRemoteUpdate(func(_ bool) { fn() })
}
