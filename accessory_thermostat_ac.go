package homekit

/*


TEST


*/

import (
	"encoding/json"

	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

type AccessoryThermostatAcStatuses struct {
	TargetState  int     `json:"target_state"`
	CurrentState int     `json:"current_state"`
	TargetTemp   float64 `json:"target_temp"`
	CurrentTemp  float64 `json:"current_temp"`
}

type AccessoryThermostatAcBox struct {
	Thermostat string                      `json:"thermostat"`
	Statuses   AccessoryThermostatStatuses `json:"statuses"`
}

func (box *AccessoryThermostatAcBox) Marshal() string {
	res, _ := json.Marshal(box)
	return string(res)
}

func (box *AccessoryThermostatAcBox) Unmarshal(s string) error {
	return json.Unmarshal([]byte(s), &box)
}

//AccessoryThermostat struct
type AccessoryThermostatAC struct {
	*accessory.Accessory
	Thermostat struct {
		*service.Service
		CurrentHeatingCoolingState *characteristic.CurrentHeatingCoolingState
		TargetHeatingCoolingState  *characteristic.TargetHeatingCoolingState
		CurrentTemperature         *characteristic.CurrentTemperature
		TargetTemperature          *characteristic.TargetTemperature
		TemperatureDisplayUnits    *characteristic.TemperatureDisplayUnits
	}
	FanV2 struct {
		*service.Service
		Active          *characteristic.Active
		RotationSpeed   *characteristic.RotationSpeed
		TargetFanState  *characteristic.TargetFanState
		CurrentFanState *characteristic.CurrentFanState
	}
}

//NewAccessoryThermostatAC returns AccessoryThermostat
//  args[0](int) - TargetHeatingCoolingState.SetValue(args[0]) default(0)
//  args[1](int) - TargetHeatingCoolingState.SetMinValue(args[1]) default(0)
//  args[2](int) - TargetHeatingCoolingState.SetMaxValue(args[2]) default(3)
//  args[3](int) - TargetHeatingCoolingState.SetStepValue(args[3]) default(1)
//  args[4](float64) - TargetTemperature.SetValue(args[0]) default(25.0)
//  args[5](float64) - TargetTemperature.SetMinValue(args[1]) default(10.0)
//  args[6](float64) - TargetTemperature.SetMaxValue(args[2]) default(40.0)
//  args[7](float64) - TargetTemperature.SetStepValue(args[3]) default(1.0)
func NewAccessoryThermostatAC(info accessory.Info, args ...interface{}) *AccessoryThermostatAC {
	acc := AccessoryThermostatAC{}
	acc.Accessory = accessory.New(info, accessory.TypeThermostat)
	acc.Thermostat.Service = service.New(service.TypeThermostat)
	acc.FanV2.Service = service.New(service.TypeFanV2)

	acc.Thermostat.CurrentHeatingCoolingState = characteristic.NewCurrentHeatingCoolingState()
	acc.Thermostat.AddCharacteristic(acc.Thermostat.CurrentHeatingCoolingState.Characteristic)

	acc.Thermostat.TargetHeatingCoolingState = characteristic.NewTargetHeatingCoolingState()
	acc.Thermostat.AddCharacteristic(acc.Thermostat.TargetHeatingCoolingState.Characteristic)

	acc.Thermostat.CurrentTemperature = characteristic.NewCurrentTemperature()
	acc.Thermostat.AddCharacteristic(acc.Thermostat.CurrentTemperature.Characteristic)

	acc.Thermostat.TargetTemperature = characteristic.NewTargetTemperature()
	acc.Thermostat.AddCharacteristic(acc.Thermostat.TargetTemperature.Characteristic)

	acc.Thermostat.TemperatureDisplayUnits = characteristic.NewTemperatureDisplayUnits()
	acc.Thermostat.AddCharacteristic(acc.Thermostat.TemperatureDisplayUnits.Characteristic)

	acc.FanV2.Active = characteristic.NewActive()
	acc.FanV2.AddCharacteristic(acc.FanV2.Active.Characteristic)

	acc.FanV2.RotationSpeed = characteristic.NewRotationSpeed()
	acc.FanV2.AddCharacteristic(acc.FanV2.RotationSpeed.Characteristic)

	acc.FanV2.TargetFanState = characteristic.NewTargetFanState()
	acc.FanV2.AddCharacteristic(acc.FanV2.TargetFanState.Characteristic)

	acc.FanV2.CurrentFanState = characteristic.NewCurrentFanState()
	acc.FanV2.AddCharacteristic(acc.FanV2.CurrentFanState.Characteristic)

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
	acc.AddService(acc.FanV2.Service)
	return &acc
}

func (acc *AccessoryThermostatAC) GetType() uint8                     { return uint8(acc.Accessory.Type) }
func (acc *AccessoryThermostatAC) GetID() uint64                      { return acc.Accessory.ID }
func (acc *AccessoryThermostatAC) GetSN() string                      { return acc.Accessory.Info.SerialNumber.GetValue() }
func (acc *AccessoryThermostatAC) GetName() string                    { return acc.Accessory.Info.SerialNumber.GetValue() }
func (acc *AccessoryThermostatAC) GetAccessory() *accessory.Accessory { return acc.Accessory }

//GetValues - accessory get values
//  vals[0](int) - Thermostat.TargetHeatingCoolingState.GetValue()
//  vals[1](int) - CurrentHeatingCoolingState.GetValue()
//  vals[2](float64) - TargetTemperature.GetValue()
//  vals[3](float64) - CurrentTemperature.GetValue()
func (acc *AccessoryThermostatAC) GetValues() (vals []interface{}) {
	vals = append(vals,
		acc.Thermostat.TargetHeatingCoolingState.GetValue(),
		acc.Thermostat.CurrentHeatingCoolingState.GetValue(),
		acc.Thermostat.TargetTemperature.GetValue(),
		acc.Thermostat.CurrentTemperature.GetValue(),
	)
	return
}

func (acc *AccessoryThermostatAC) GetValuesJson() string {
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
func (acc *AccessoryThermostatAC) SetValues(val interface{}, vals ...interface{}) {
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

func (acc *AccessoryThermostatAC) SetValuesJson(vals string) {
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
func (acc *AccessoryThermostatAC) OnValueRemoteUpdate(fn func(charact uint8, val interface{})) {
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(state int) { fn(0x1, state) })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(temp float64) { fn(0x2, temp) })
}

//OnValuesRemoteUpdate
//  0x1(int) - Thermostat.TargetHeatingCoolingState.GetValue()
//  0x2(float64) - Thermostat.TargetTemperature.GetValue()
func (acc *AccessoryThermostatAC) OnValuesRemoteUpdate(fn func(vals ...interface{})) {
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(state int) { fn(state, acc.Thermostat.TargetTemperature.GetValue()) })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(temp float64) { fn(acc.Thermostat.TargetHeatingCoolingState.GetValue(), temp) })
}

func (acc *AccessoryThermostatAC) OnValuesRemoteUpdateJson(fn func(vals string)) {
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(state int) { fn(acc.GetValuesJson()) })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(temp float64) { fn(acc.GetValuesJson()) })
}

func (acc *AccessoryThermostatAC) OnValuesRemoteUpdateEmpty(fn func()) {
	acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(_ int) { fn() })
	acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(_ float64) { fn() })
}
