package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/brutella/hap/accessory"
	"github.com/xxandev/homekit"
)

func init() { rand.Seed(time.Now().Unix()) }

func ExampleSwitch(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccSwitch(id,
		accessory.Info{
			Name:         fmt.Sprintf("Switch-%d", id),
			SerialNumber: fmt.Sprintf("%s-Sw-%d", zone, id),
			Model:        "Ex-Switch",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go func() {
		for range time.Tick(time.Millisecond * time.Duration(7200000+rand.Intn(1000000))) {
			acc.Switch.On.SetValue(!acc.Switch.On.Value())
			fmt.Printf("[%[1]T - %[2]v] update on: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.Switch.On.Value())
		}
	}()
	go acc.Switch.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%[1]T - %[2]v] remote update on: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	return acc.GetAccessory()
}

func ExampleAirPurifier(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccAirPurifier(id,
		accessory.Info{
			Name:         fmt.Sprintf("AirPurifier-%d", id),
			SerialNumber: fmt.Sprintf("%s-AirPur-%d", zone, id),
			Model:        "Ex-AirPrf",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		}, 0.0, 0.0, 100.0, 1.0)
	go acc.AirPurifier.Active.OnValueRemoteUpdate(func(v int) {
		if v > 0 {
			acc.AirPurifier.CurrentAirPurifierState.SetValue(2)
		} else {
			acc.AirPurifier.CurrentAirPurifierState.SetValue(0)
		}
		fmt.Printf("[%[1]T - %[2]v] remote update active: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	go acc.AirPurifier.TargetAirPurifierState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v] remote update target state: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	go acc.AirPurifier.RotationSpeed.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v] remote update rotation speed: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	return acc.GetAccessory()
}

func ExampleDoor(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccDoor(id,
		accessory.Info{
			Name:         fmt.Sprintf("Door-%d", id),
			SerialNumber: fmt.Sprintf("%s-Door-%d", zone, id),
			Model:        "Ex-Door",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		}, 0, 0, 100, 100)
	command := make(chan int, 5)
	go func() {
		callbackT := time.NewTimer(time.Millisecond * 10000)
		for {
			select {
			case cmd := <-command:
				acc.Door.TargetPosition.SetValue(cmd)
				acc.Door.CurrentPosition.SetValue(cmd)
				callbackT.Stop()
				callbackT.Reset(time.Millisecond * 10000)
				fmt.Printf("[%[1]T - %[2]v] update position, current: %[3]T - %[3]v, target: %[4]T - %[4]v\n",
					acc, acc.A.Info.SerialNumber.Value(), acc.Door.CurrentPosition.Value(), acc.Door.TargetPosition.Value())
			case <-callbackT.C:
				acc.Door.TargetPosition.SetValue(0)
				acc.Door.CurrentPosition.SetValue(0)
				fmt.Printf("[%[1]T - %[2]v] update position, current: %[3]T - %[3]v, target: %[4]T - %[4]v\n",
					acc, acc.A.Info.SerialNumber.Value(), acc.Door.CurrentPosition.Value(), acc.Door.TargetPosition.Value())
			case <-time.Tick(time.Millisecond * time.Duration(7200000+rand.Intn(1000000))):
				command <- 100
			}
		}
	}()
	go acc.Door.TargetPosition.OnValueRemoteUpdate(func(v int) {
		command <- v
		fmt.Printf("[%[1]T - %[2]v] remote update position: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})

	return acc.GetAccessory()
}

func ExampleFanRS(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccFanRS(id,
		accessory.Info{
			Name:         fmt.Sprintf("FanRS-%d", id),
			SerialNumber: fmt.Sprintf("%s-FanRS-%d", zone, id),
			Model:        "Ex-FanRS",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go func() {
		for range time.Tick(time.Millisecond * time.Duration(7200000+rand.Intn(1000000))) {
			acc.Fan.On.SetValue(true)
			acc.Fan.RotationSpeed.SetValue(50)
			fmt.Printf("[%[1]T - %[2]v] update on: %[3]T - %[3]v, speed %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.Fan.On.Value(), acc.Fan.RotationSpeed.Value())
			time.Sleep(time.Millisecond * 30000)
			acc.Fan.On.SetValue(true)
			acc.Fan.RotationSpeed.SetValue(100)
			fmt.Printf("[%[1]T - %[2]v] update on: %[3]T - %[3]v, speed %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.Fan.On.Value(), acc.Fan.RotationSpeed.Value())
			time.Sleep(time.Millisecond * 120000)
			acc.Fan.On.SetValue(false)
			fmt.Printf("[%[1]T - %[2]v] update on: %[3]T - %[3]v, speed %[4]T - %[4]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.Fan.On.Value(), acc.Fan.RotationSpeed.Value())
		}
	}()
	go acc.Fan.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%[1]T - %[2]v] remote update on: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	go acc.Fan.RotationSpeed.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v] remote update rotation speed: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	return acc.GetAccessory()
}

func ExampleFaucet(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccFaucet(id,
		accessory.Info{
			Name:         fmt.Sprintf("Faucet-%d", id),
			SerialNumber: fmt.Sprintf("%s-Faucet-%d", zone, id),
			Model:        "Ex-Faucet",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go func() {
		for range time.Tick(time.Millisecond * time.Duration(8200000+rand.Intn(1000000))) {
			acc.Valve.Active.SetValue(int(math.Pow(0, float64(acc.Valve.Active.Value()))))
			acc.Valve.InUse.SetValue(acc.Valve.Active.Value())
			fmt.Printf("[%[1]T - %[2]v] update active: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.Valve.Active.Value())
		}
	}()
	go acc.Valve.Active.OnValueRemoteUpdate(func(v int) {
		acc.Valve.InUse.SetValue(v)
		fmt.Printf("[%[1]T - %[2]v] remote update active: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	return acc.GetAccessory()
}

func ExampleGate(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccGate(id,
		accessory.Info{
			Name:         fmt.Sprintf("Gate-%d", id),
			SerialNumber: fmt.Sprintf("%s-Gate-%d", zone, id),
			Model:        "Ex-Gate",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	acc.GarageDoorOpener.CurrentDoorState.SetValue(4) //unknown current state
	acc.GarageDoorOpener.TargetDoorState.SetValue(1)
	timer := time.NewTimer(time.Millisecond * 10000)
	go func() {
		for range timer.C {
			// closed --> Target(1), Current(1)
			// closing -> Target(1), Current(2)
			// opened --> Target(0), Current(0)
			// opening -> Target(0), Current(2)
			acc.GarageDoorOpener.CurrentDoorState.SetValue(acc.GarageDoorOpener.TargetDoorState.Value())
			fmt.Printf("[%[1]T - %[2]v] update state, current: %[3]T - %[3]v, target: %[4]T - %[4]v\n",
				acc, acc.A.Info.SerialNumber.Value(), acc.GarageDoorOpener.CurrentDoorState.Value(), acc.GarageDoorOpener.TargetDoorState.Value())
		}
	}()
	go acc.GarageDoorOpener.TargetDoorState.OnValueRemoteUpdate(func(v int) {
		timer.Stop()
		timer.Reset(time.Millisecond * 10000)
		acc.GarageDoorOpener.CurrentDoorState.SetValue(2)
		fmt.Printf("[%[1]T - %[2]v] remote update state, current: %[3]T - %[3]v, target: %[4]T - %[4]v\n",
			acc, acc.A.Info.SerialNumber.Value(), acc.GarageDoorOpener.CurrentDoorState.Value(), v)
	})
	return acc.GetAccessory()
}

func ExampleHumidifierDehumidifier(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccHumidifierDehumidifier(id,
		accessory.Info{
			Name:         fmt.Sprintf("HumDehum-%d", id),
			SerialNumber: fmt.Sprintf("%s-HumDehum-%d", zone, id),
			Model:        "Ex-HumDehum",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go acc.HumidifierDehumidifier.Active.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v] remote update active: %[3]T - %[3]v\n",
			acc, acc.A.Info.SerialNumber.Value(), v)
	})
	go acc.HumidifierDehumidifier.TargetHumidifierDehumidifierState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v] remote update target state: %[3]T - %[3]v\n",
			acc, acc.A.Info.SerialNumber.Value(), v)
	})
	go acc.HumidifierDehumidifier.RelativeHumidityDehumidifierThreshold.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v] remote update relative dehumidifier threshold: %[3]T - %[3]v\n",
			acc, acc.A.Info.SerialNumber.Value(), v)
	})
	go acc.HumidifierDehumidifier.RelativeHumidityHumidifierThreshold.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v] remote update relative humidifier threshold: %[3]T - %[3]v\n",
			acc, acc.A.Info.SerialNumber.Value(), v)
	})
	return acc.GetAccessory()
}

func ExampleIrrigation(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccIrrigation(id,
		accessory.Info{
			Name:         fmt.Sprintf("Irrigation-%d", id),
			SerialNumber: fmt.Sprintf("%s-Irg-%d", zone, id),
			Model:        "Ex-Irg",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go func() {
		for range time.Tick(time.Millisecond * time.Duration(7200000+rand.Intn(1000000))) {
			acc.Valve.Active.SetValue(int(math.Pow(0, float64(acc.Valve.Active.Value()))))
			acc.Valve.InUse.SetValue(acc.Valve.Active.Value())
			fmt.Printf("[%[1]T - %[2]v] update active: %[3]T - %[3]v\n",
				acc, acc.A.Info.SerialNumber.Value(), acc.Valve.Active.Value())
		}
	}()
	go acc.Valve.Active.OnValueRemoteUpdate(func(v int) {
		acc.Valve.InUse.SetValue(v)
		fmt.Printf("[%[1]T - %[2]v] remote update active: %[3]T - %[3]v\n",
			acc, acc.A.Info.SerialNumber.Value(), v)
	})
	return acc.GetAccessory()
}

func ExampleLightbulbColored(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccLightbulbColored(id,
		accessory.Info{
			Name:         fmt.Sprintf("LbColor-%d", id),
			SerialNumber: fmt.Sprintf("%s-LbColor-%d", zone, id),
			Model:        "Ex-LbColor",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go func() {
		for range time.Tick(time.Millisecond * time.Duration(7200000+rand.Intn(1000000))) {
			acc.LightbulbColored.Brightness.SetValue(100)
			acc.LightbulbColored.On.SetValue(!acc.LightbulbColored.On.Value())
			fmt.Printf("[%[1]T - %[2]v] update on: %[3]T - %[3]v \n",
				acc, acc.A.Info.SerialNumber.Value(), acc.LightbulbColored.On.Value())
		}
	}()
	go acc.LightbulbColored.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%[1]T - %[2]v] remote update on: %[3]T - %[3]v\n",
			acc, acc.A.Info.SerialNumber.Value(), v)
	})
	go acc.LightbulbColored.Brightness.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v] remote update brightness: %[3]T - %[3]v\n",
			acc, acc.A.Info.SerialNumber.Value(), v)
	})
	go acc.LightbulbColored.Saturation.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v] remote update saturation: %[3]T - %[3]v\n",
			acc, acc.A.Info.SerialNumber.Value(), v)
	})
	go acc.LightbulbColored.Hue.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v] remote update hue: %[3]T - %[3]v\n",
			acc, acc.A.Info.SerialNumber.Value(), v)
	})
	return acc.GetAccessory()
}

func ExampleLightbulbDimmer(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccLightbulbDimmer(id,
		accessory.Info{
			Name:         fmt.Sprintf("LbDimm-%d", id),
			SerialNumber: fmt.Sprintf("%s-LbDimm-%d", zone, id),
			Model:        "Ex-LbDimm",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go func() {
		for range time.Tick(time.Millisecond * time.Duration(7200000+rand.Intn(1000000))) {
			acc.LightbulbDimmer.Brightness.SetValue(100)
			acc.LightbulbDimmer.On.SetValue(!acc.LightbulbDimmer.On.Value())
			fmt.Printf("[%[1]T - %[2]v] update on: %[3]T - %[3]v \n",
				acc, acc.A.Info.SerialNumber.Value(), acc.LightbulbDimmer.On.Value())
		}
	}()
	go acc.LightbulbDimmer.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%[1]T - %[2]v] remote update on: %[3]T - %[3]v\n",
			acc, acc.A.Info.SerialNumber.Value(), v)
	})
	go acc.LightbulbDimmer.Brightness.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v] remote update brightness: %[3]T - %[3]v\n",
			acc, acc.A.Info.SerialNumber.Value(), v)
	})
	return acc.GetAccessory()
}

func ExampleOutlet(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccOutlet(id,
		accessory.Info{
			Name:         fmt.Sprintf("Outlet-%d", id),
			SerialNumber: fmt.Sprintf("%s-Outlet-%d", zone, id),
			Model:        "Ex-Outlet",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go func() {
		for range time.Tick(time.Millisecond * time.Duration(7200000+rand.Intn(1000000))) {
			acc.Outlet.On.SetValue(!acc.Outlet.On.Value())
			fmt.Printf("[%[1]T - %[2]v] update on: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.Outlet.On.Value())
		}
	}()
	go acc.Outlet.On.OnValueRemoteUpdate(func(v bool) {
		fmt.Printf("[%[1]T - %[2]v] remote update on: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	return acc.GetAccessory()
}

func ExampleSensorContact(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccSensorContact(id,
		accessory.Info{
			Name:         fmt.Sprintf("SContact-%d", id),
			SerialNumber: fmt.Sprintf("%s-SContact-%d", zone, id),
			Model:        "Ex-SContact",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go func() {
		for range time.Tick(time.Millisecond * time.Duration(8200000+rand.Intn(1000000))) {
			acc.ContactSensor.ContactSensorState.SetValue(int(math.Pow(0, float64(acc.ContactSensor.ContactSensorState.Value()))))
			fmt.Printf("[%[1]T - %[2]v] update state: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.ContactSensor.ContactSensorState.Value())
		}
	}()
	return acc.GetAccessory()
}

func ExampleSensorHumidity(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccSensorHumidity(id,
		accessory.Info{
			Name:         fmt.Sprintf("SHum-%d", id),
			SerialNumber: fmt.Sprintf("%s-SHum-%d", zone, id),
			Model:        "Ex-SHum",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go func() {
		for range time.Tick(time.Millisecond * 1000) {
			acc.HumiditySensor.CurrentRelativeHumidity.SetValue(float64(rand.Intn(40) + time.Now().Second()/2))
			fmt.Printf("[%[1]T - %[2]v] update status: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.HumiditySensor.CurrentRelativeHumidity.Value())
		}
	}()
	return acc.GetAccessory()
}

func ExampleSensorLeak(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccSensorLeak(id,
		accessory.Info{
			Name:         fmt.Sprintf("SLeak-%d", id),
			SerialNumber: fmt.Sprintf("%s-SLeak-%d", zone, id),
			Model:        "Ex-SLeak",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go func() {
		for range time.Tick(time.Millisecond * time.Duration(8200000+rand.Intn(1000000))) {
			acc.LeakSensor.LeakDetected.SetValue(int(math.Pow(0, float64(acc.LeakSensor.LeakDetected.Value()))))
			fmt.Printf("[%[1]T - %[2]v] update state: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.LeakSensor.LeakDetected.Value())
		}
	}()
	return acc.GetAccessory()
}

func ExampleSensorLight(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccSensorLight(id,
		accessory.Info{
			Name:         fmt.Sprintf("SLight-%d", id),
			SerialNumber: fmt.Sprintf("%s-SLight-%d", zone, id),
			Model:        "Ex-SLight",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go func() {
		for range time.Tick(time.Millisecond * 3000) {
			if acc.LightSensor.CurrentAmbientLightLevel.Value() >= acc.LightSensor.CurrentAmbientLightLevel.MaxValue() {
				acc.LightSensor.CurrentAmbientLightLevel.SetValue(acc.LightSensor.CurrentAmbientLightLevel.MinValue())
			} else {
				acc.LightSensor.CurrentAmbientLightLevel.SetValue(acc.LightSensor.CurrentAmbientLightLevel.Value() + 0.25)
			}
			fmt.Printf("[%[1]T - %[2]v] update status: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.LightSensor.CurrentAmbientLightLevel.Value())
		}
	}()
	return acc.GetAccessory()
}

func ExampleSensorMotion(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccSensorMotion(id,
		accessory.Info{
			Name:         fmt.Sprintf("SMotion-%d", id),
			SerialNumber: fmt.Sprintf("%s-SMotion-%d", zone, id),
			Model:        "Ex-SMotion",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go func() {
		for range time.Tick(time.Millisecond * time.Duration(8200000+rand.Intn(1000000))) {
			acc.MotionSensor.MotionDetected.SetValue(!acc.MotionSensor.MotionDetected.Value())
			fmt.Printf("[%[1]T - %[2]v] update state: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.MotionSensor.MotionDetected.Value())
		}
	}()
	return acc.GetAccessory()
}

func ExampleSensorTemperature(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccSensorTemperature(id,
		accessory.Info{
			Name:         fmt.Sprintf("STemp-%d", id),
			SerialNumber: fmt.Sprintf("%s-STemp-%d", zone, id),
			Model:        "Ex-STemp",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	acc.TempSensor.CurrentTemperature.SetValue(25.00)
	go func() {
		for range time.Tick(time.Millisecond * 1000) {
			acc.TempSensor.CurrentTemperature.SetValue(float64(5 + rand.Intn(10) + time.Now().Second()/3))
			fmt.Printf("[%[1]T - %[2]v] update status: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), acc.TempSensor.CurrentTemperature.Value())
		}
	}()
	return acc.GetAccessory()
}

func ExampleThermostatClimate(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccThermostat(id,
		accessory.Info{
			Name:         fmt.Sprintf("ThrmClmt-%d", id),
			SerialNumber: fmt.Sprintf("%s-ThrmClmt-%d", zone, id),
			Model:        "Ex-ThrmClmt",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	acc.Thermostat.CurrentTemperature.SetValue(24.00)
	go func() {
		for range time.Tick(time.Millisecond * 30000) {
			if acc.Thermostat.TargetHeatingCoolingState.Value() == 0 { //off
				acc.Thermostat.CurrentHeatingCoolingState.SetValue(acc.Thermostat.TargetHeatingCoolingState.Value())
				continue
			}
			if acc.Thermostat.TargetHeatingCoolingState.Value() == 1 { //heating
				if acc.Thermostat.CurrentTemperature.Value() > acc.Thermostat.TargetTemperature.Value() {
					acc.Thermostat.CurrentHeatingCoolingState.SetValue(0)
				} else {
					acc.Thermostat.CurrentHeatingCoolingState.SetValue(acc.Thermostat.TargetHeatingCoolingState.Value())
					acc.Thermostat.CurrentTemperature.SetValue(acc.Thermostat.CurrentTemperature.Value() + 0.25)
				}
			}
			if acc.Thermostat.TargetHeatingCoolingState.Value() == 2 { //cooling
				if acc.Thermostat.CurrentTemperature.Value() < acc.Thermostat.TargetTemperature.Value() {
					acc.Thermostat.CurrentHeatingCoolingState.SetValue(0)
				} else {
					acc.Thermostat.CurrentHeatingCoolingState.SetValue(acc.Thermostat.TargetHeatingCoolingState.Value())
					acc.Thermostat.CurrentTemperature.SetValue(acc.Thermostat.CurrentTemperature.Value() - 0.25)
				}
			}
			if acc.Thermostat.TargetHeatingCoolingState.Value() == 3 { //automatic
				if acc.Thermostat.CurrentTemperature.Value() == acc.Thermostat.TargetTemperature.Value() {
					acc.Thermostat.CurrentHeatingCoolingState.SetValue(0)
				} else if acc.Thermostat.CurrentTemperature.Value() > acc.Thermostat.TargetTemperature.Value() {
					acc.Thermostat.TargetHeatingCoolingState.SetValue(2)
					acc.Thermostat.CurrentTemperature.SetValue(acc.Thermostat.CurrentTemperature.Value() - 0.25)
				} else if acc.Thermostat.CurrentTemperature.Value() < acc.Thermostat.TargetTemperature.Value() {
					acc.Thermostat.TargetHeatingCoolingState.SetValue(1)
					acc.Thermostat.CurrentTemperature.SetValue(acc.Thermostat.CurrentTemperature.Value() + 0.25)
				}
			}
			fmt.Printf("[%[1]T - %[2]v] update thermostat, current state: %[3]T - %[3]v, target state: %[4]T - %[4]v, current temp: %[5]T - %[5]v, target temp: %[6]T - %[6]v, \n",
				acc, acc.A.Info.SerialNumber.Value(),
				acc.Thermostat.CurrentHeatingCoolingState.Value(), acc.Thermostat.TargetHeatingCoolingState.Value(),
				acc.Thermostat.CurrentTemperature.Value(), acc.Thermostat.TargetTemperature.Value(),
			)
		}
	}()
	go acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v] remote update target state: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	go acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v] remote update target temp: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	return acc.GetAccessory()
}

func ExampleThermostatHeating(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccThermostat(id,
		accessory.Info{
			Name:         fmt.Sprintf("ThrmHtn-%d", id),
			SerialNumber: fmt.Sprintf("%s-ThrmHtn-%d", zone, id),
			Model:        "Ex-ThrmHtn",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		}, 0, 0, 1, 1)
	acc.Thermostat.CurrentTemperature.SetValue(24.00)
	go func() {
		for {
			select {
			case <-time.Tick(time.Millisecond * 5000):
				if acc.Thermostat.TargetHeatingCoolingState.Value() == 0 {
					if acc.Thermostat.TargetTemperature.Value()-10.00 > acc.Thermostat.CurrentTemperature.Value() {
						acc.Thermostat.CurrentTemperature.SetValue(acc.Thermostat.CurrentTemperature.Value() - 0.5)
					}
				}
			case <-time.Tick(time.Millisecond * 30000):
				if acc.Thermostat.TargetHeatingCoolingState.Value() == 0 { //off
					acc.Thermostat.CurrentHeatingCoolingState.SetValue(acc.Thermostat.TargetHeatingCoolingState.Value())
					continue
				}
				if acc.Thermostat.TargetHeatingCoolingState.Value() == 1 { //heating
					if acc.Thermostat.CurrentTemperature.Value() > acc.Thermostat.TargetTemperature.Value() {
						acc.Thermostat.CurrentHeatingCoolingState.SetValue(0)
					} else {
						acc.Thermostat.CurrentHeatingCoolingState.SetValue(acc.Thermostat.TargetHeatingCoolingState.Value())
						acc.Thermostat.CurrentTemperature.SetValue(acc.Thermostat.CurrentTemperature.Value() + 0.25)
					}
				}
				fmt.Printf("[%[1]T - %[2]v] update thermostat, current state: %[3]T - %[3]v, target state: %[4]T - %[4]v, current temp: %[5]T - %[5]v, target temp: %[6]T - %[6]v, \n",
					acc, acc.A.Info.SerialNumber.Value(),
					acc.Thermostat.CurrentHeatingCoolingState.Value(), acc.Thermostat.TargetHeatingCoolingState.Value(),
					acc.Thermostat.CurrentTemperature.Value(), acc.Thermostat.TargetTemperature.Value(),
				)
			}
		}
	}()
	go acc.Thermostat.TargetHeatingCoolingState.OnValueRemoteUpdate(func(v int) {
		fmt.Printf("[%[1]T - %[2]v] remote update target state: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	go acc.Thermostat.TargetTemperature.OnValueRemoteUpdate(func(v float64) {
		fmt.Printf("[%[1]T - %[2]v] remote update target temp: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	return acc.GetAccessory()
}

func ExampleWindow(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccWindow(id,
		accessory.Info{
			Name:         fmt.Sprintf("Window-%d", id),
			SerialNumber: fmt.Sprintf("%s-Window-%d", zone, id),
			Model:        "Ex-Window",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go acc.Window.TargetPosition.OnValueRemoteUpdate(func(v int) {
		acc.Window.CurrentPosition.SetValue(v)
		fmt.Printf("[%[1]T - %[2]v] remote update target position: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})

	go acc.Window.PositionState.OnValueRemoteUpdate(func(v int) {
		acc.Window.CurrentPosition.SetValue(acc.Window.TargetPosition.Value())
		fmt.Printf("[%[1]T - %[2]v] remote update position state: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	return acc.GetAccessory()
}

func ExampleWindowCovering(id uint64, zone string) *accessory.A {
	acc := homekit.NewAccWindowCovering(id,
		accessory.Info{
			Name:         fmt.Sprintf("Jalousie-%d", id),
			SerialNumber: fmt.Sprintf("%s-Jalousie-%d", zone, id),
			Model:        "Ex-Jalousie",
			Manufacturer: homekit.MANUFACTURER,
			Firmware:     homekit.FIRMWARE,
		})
	go acc.WindowCovering.TargetPosition.OnValueRemoteUpdate(func(v int) {
		acc.WindowCovering.CurrentPosition.SetValue(v)
		fmt.Printf("[%[1]T - %[2]v] remote update target position: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})

	go acc.WindowCovering.PositionState.OnValueRemoteUpdate(func(v int) {
		acc.WindowCovering.CurrentPosition.SetValue(acc.WindowCovering.TargetPosition.Value())
		fmt.Printf("[%[1]T - %[2]v] remote update position state: %[3]T - %[3]v\n", acc, acc.A.Info.SerialNumber.Value(), v)
	})
	return acc.GetAccessory()
}
