// this example implements a BLE accelerometer service.
// see https://lancaster-university.github.io/microbit-docs/resources/bluetooth/bluetooth_profile.html for the full spec.
package main

import (
	"time"

	"machine"

	"tinygo.org/x/bluetooth"
	"tinygo.org/x/drivers/lsm303agr"
)

var (
	adapter = bluetooth.DefaultAdapter

	tempMeasurement bluetooth.Characteristic
)

func main() {
	println("starting")

	machine.I2C0.Configure(machine.I2CConfig{})

	sensor := lsm303agr.New(machine.I2C0)
	err := sensor.Configure(lsm303agr.Configuration{}) //default settings
	if err != nil {
		for {
			println("Failed to configure", err.Error())
			time.Sleep(time.Second)
		}
	}

	// you can specify the following options to adjust accuracy, sensor range or save power.
	// see https://github.com/tinygo-org/drivers/blob/release/lsm303agr/registers.go for details:
	/*
		sensor.Configure(lsm303agr.Configuration{
			AccelPowerMode: lsm303agr.ACCEL_POWER_NORMAL,
			AccelRange:     lsm303agr.ACCEL_RANGE_2G,
			AccelDataRate:  lsm303agr.ACCEL_DATARATE_100HZ,
			MagPowerMode:   lsm303agr.MAG_POWER_NORMAL,
			MagSystemMode:  lsm303agr.MAG_SYSTEM_CONTINUOUS,
			MagDataRate:    lsm303agr.MAG_DATARATE_10HZ,
		})
	*/

	must("enable BLE stack", adapter.Enable())
	adv := adapter.DefaultAdvertisement()
	must("config adv", adv.Configure(bluetooth.AdvertisementOptions{
		LocalName:    "TinyGo",
		ServiceUUIDs: []bluetooth.UUID{bluetooth.ServiceUUIDMicrobitAccelerometer},
	}))
	must("start adv", adv.Start())

	must("add service", adapter.AddService(&bluetooth.Service{
		UUID: bluetooth.ServiceUUIDMicrobitTemperature,
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &tempMeasurement,
				UUID:   bluetooth.CharacteristicUUIDMicrobitTemperature,
				Value:  []byte{0},
				Flags:  bluetooth.CharacteristicNotifyPermission,
			},
		},
	}))

	buf := make([]byte, 1)

	var temp int32
	for {
		temp, _ = sensor.ReadTemperature()
		buf[0] = uint8(temp / 1000)
		println("Simplified TEMPERATURE", buf[0])
		time.Sleep(time.Second)

		tempMeasurement.Write(buf)
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
