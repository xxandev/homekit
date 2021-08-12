package homekit

import (
	"encoding/json"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

type AccessoryThermostatStatuses struct {
	TargetState  int     `json:"target_state"`
	CurrentState int     `json:"current_state"`
	TargetTemp   float64 `json:"target_temp"`
	CurrentTemp  float64 `json:"current_temp"`
}

type AccessoryThermostatBox struct {
	Thermostat string                      `json:"thermostat"`
	Statuses   AccessoryThermostatStatuses `json:"statuses"`
}

func (box *AccessoryThermostatBox) Marshal() string {
	res, _ := json.Marshal(box)
	return string(res)
}

func (box *AccessoryThermostatBox) Unmarshal(s string) error {
	return json.Unmarshal([]byte(s), &box)
}

//AccessoryThermostat struct
type AccessoryThermostat struct {
	*accessory.Accessory
	Thermostat *service.Thermostat
}

//NewAccessoryThermostat returns AccessoryThermostat
//  args[0](int) - TargetHeatingCoolingState.SetValue(args[0]) default(0)
//  args[1](int) - TargetHeatingCoolingState.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetHeatingCoolingState.SetMaxValue(args[2]) default(3)
//  args[3](int) - TargetHeatingCoolingState.SetStepValue(args[3]) default(1)
//  args[4](float64) - TargetTemperature.SetValue(args[0]) default(25.0)
//  args[5](float64) - TargetTemperature.SetMinValue(args[1]) default(10.0)
//  args[6](float64) - TargetTemperature.SetMaxValue(args[2]) default(40.0)
//  args[7](float64) - TargetTemperature.SetStepValue(args[3]) default(1.0)
func NewAccessoryThermostat(info accessory.Info, args ...interface{}) *AccessoryThermostat {
	acc := AccessoryThermostat{}
	acc.Accessory = accessory.New(info, accessory.TypeThermostat)
	acc.Thermostat = service.NewThermostat()

	n := len(args)
	if n > 0 {
		acc.Thermostat.TargetHeatingCoolingState.SetValue(toInt(args[0], 0))
	}
	if n > 1 {
		acc.Thermostat.TargetHeatingCoolingState.SetMinValue(toInt(args[1], 0))
	}
	if n > 2 {
		acc.Thermostat.TargetHeatingCoolingState.SetMaxValue(toInt(args[2], 3))
	}
	if n > 3 {
		acc.Thermostat.TargetHeatingCoolingState.SetStepValue(toInt(args[3], 1))
	}

	if n > 4 {
		acc.Thermostat.TargetTemperature.SetValue(toFloat64(args[4], 25.0))
	} else {
		acc.Thermostat.TargetTemperature.SetValue(25.0)
	}
	if n > 5 {
		acc.Thermostat.TargetTemperature.SetMinValue(toFloat64(args[5], 10.0))
	} else {
		acc.Thermostat.TargetTemperature.SetMinValue(10.0)
	}
	if n > 6 {
		acc.Thermostat.TargetTemperature.SetMaxValue(toFloat64(args[6], 40.0))
	} else {
		acc.Thermostat.TargetTemperature.SetMaxValue(40.0)
	}
	if n > 7 {
		acc.Thermostat.TargetTemperature.SetStepValue(toFloat64(args[7], 1.0))
	} else {
		acc.Thermostat.TargetTemperature.SetStepValue(1.0)
	}

	acc.AddService(acc.Thermostat.Service)

	return &acc
}

func (acc *AccessoryThermostat) OnValueRemoteUpdateEmpty(fn func()) {
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(_ int) { fn() })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(_ float64) { fn() })
}

func (acc *AccessoryThermostat) GetType() uint8                     { return uint8(acc.Accessory.Type) }
func (acc *AccessoryThermostat) GetID() uint64                      { return acc.Accessory.ID }
func (acc *AccessoryThermostat) GetSN() string                      { return acc.Accessory.Info.SerialNumber.GetValue() }
func (acc *AccessoryThermostat) GetName() string                    { return acc.Accessory.Info.SerialNumber.GetValue() }
func (acc *AccessoryThermostat) GetAccessory() *accessory.Accessory { return acc.Accessory }

//GetValues - accessory get values
//  vals[0](int) - Thermostat.TargetHeatingCoolingState.GetValue()
//  vals[1](int) - CurrentHeatingCoolingState.GetValue()
//  vals[2](float64) - TargetTemperature.GetValue()
//  vals[3](float64) - CurrentTemperature.GetValue()
func (acc *AccessoryThermostat) GetValues() (vals []interface{}) {
	vals = append(vals,
		acc.Thermostat.TargetHeatingCoolingState.GetValue(),
		acc.Thermostat.CurrentHeatingCoolingState.GetValue(),
		acc.Thermostat.TargetTemperature.GetValue(),
		acc.Thermostat.CurrentTemperature.GetValue(),
	)
	return
}

func (acc *AccessoryThermostat) GetValuesJson() string {
	box := AccessoryThermostatBox{
		Thermostat: acc.Accessory.Info.SerialNumber.GetValue(),
		Statuses: AccessoryThermostatStatuses{
			TargetState:  acc.Thermostat.TargetHeatingCoolingState.GetValue(),
			CurrentState: acc.Thermostat.CurrentHeatingCoolingState.GetValue(),
			TargetTemp:   acc.Thermostat.TargetTemperature.GetValue(),
			CurrentTemp:  acc.Thermostat.CurrentTemperature.GetValue(),
		},
	}
	return box.Marshal()
}

//SetValues - accessory set values
//  val(int) - Thermostat.TargetHeatingCoolingState.SetValue(val)
//  vals[0](int) - Thermostat.CurrentHeatingCoolingState.SetValue(vals[0])
//  vals[1](float64) - Thermostat.TargetTemperature.SetValue(vals[1])
//  vals[2](float64) - acc.Thermostat.CurrentTemperature.SetValue(vals[2])
func (acc *AccessoryThermostat) SetValues(val interface{}, vals ...interface{}) {
	if v, ok := val.(int); ok {
		acc.Thermostat.TargetHeatingCoolingState.SetValue(v)
	}
	l := len(vals)
	if l > 0 {
		if v, ok := vals[0].(int); ok {
			acc.Thermostat.CurrentHeatingCoolingState.SetValue(v)
		}
	}
	if l > 1 {
		if v, ok := vals[0].(float64); ok {
			acc.Thermostat.TargetTemperature.SetValue(v)
		}
	}
	if l > 2 {
		if v, ok := vals[0].(float64); ok {
			acc.Thermostat.CurrentTemperature.SetValue(v)
		}
	}
}

func (acc *AccessoryThermostat) SetValuesJson(vals string) {
	box := AccessoryThermostatBox{}
	if err := box.Unmarshal(vals); err != nil || box.Thermostat != acc.Accessory.Info.SerialNumber.GetValue() {
		return
	}
	acc.Thermostat.TargetHeatingCoolingState.SetValue(box.Statuses.TargetState)
	acc.Thermostat.CurrentHeatingCoolingState.SetValue(box.Statuses.CurrentState)
	acc.Thermostat.TargetTemperature.SetValue(box.Statuses.TargetTemp)
	acc.Thermostat.CurrentTemperature.SetValue(box.Statuses.CurrentTemp)
}

//OnValueRemoteUpdate
//  0x1(int) - Thermostat.TargetHeatingCoolingState.GetValue()
//  0x2(float64) - Thermostat.TargetTemperature.GetValue()
func (acc *AccessoryThermostat) OnValueRemoteUpdate(fn func(charact uint8, val interface{})) {
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(state int) { fn(0x1, state) })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(temp float64) { fn(0x2, temp) })
}

//OnValuesRemoteUpdate
//  0x1(int) - Thermostat.TargetHeatingCoolingState.GetValue()
//  0x2(float64) - Thermostat.TargetTemperature.GetValue()
func (acc *AccessoryThermostat) OnValuesRemoteUpdate(fn func(vals ...interface{})) {
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(state int) { fn(state, acc.Thermostat.TargetTemperature.GetValue()) })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(temp float64) { fn(acc.Thermostat.TargetHeatingCoolingState.GetValue(), temp) })
}

func (acc *AccessoryThermostat) OnValuesRemoteUpdateJson(fn func(vals string)) {
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(state int) { fn(acc.GetValuesJson()) })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(temp float64) { fn(acc.GetValuesJson()) })
}

func (acc *AccessoryThermostat) OnValuesRemoteUpdateEmpty(fn func()) {
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(_ int) { fn() })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(_ float64) { fn() })
}
