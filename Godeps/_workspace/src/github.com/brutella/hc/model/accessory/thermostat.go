package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/service"
)

type thermostat struct {
	*Accessory

	thermostat *service.Thermostat
}

// NewThermometer returns a thermometer  which implements model.Thermometer.
func NewThermometer(info model.Info, temp, min, max, steps float64) *thermostat {
	accessory := New(info)
	t := service.NewThermometer(info.Name, temp, min, max, steps)

	accessory.AddService(t.Service)

	return &thermostat{accessory, t}
}

// NewThermostat returns a thermostat which implements model.Thermostat.
func NewThermostat(info model.Info, temp, min, max, steps float64) *thermostat {
	accessory := New(info)
	t := service.NewThermostat(info.Name, temp, min, max, steps)

	accessory.AddService(t.Service)

	return &thermostat{accessory, t}
}

func (t *thermostat) Temperature() float64 {
	return t.thermostat.Temp.Temperature()
}

func (t *thermostat) SetTemperature(value float64) {
	t.thermostat.Temp.SetTemperature(value)
}

func (t *thermostat) Unit() model.TempUnit {
	return t.thermostat.Unit.Unit()
}

func (t *thermostat) SetTargetTemperature(value float64) {
	t.thermostat.TargetTemp.SetTemperature(value)
}

func (t *thermostat) TargetTemperature() float64 {
	return t.thermostat.TargetTemp.Temperature()
}

func (t *thermostat) SetMode(value model.HeatCoolModeType) {
	if value != model.HeatCoolModeAuto {
		t.thermostat.Mode.SetHeatingCoolingMode(value)
	}
}

func (t *thermostat) Mode() model.HeatCoolModeType {
	return t.thermostat.Mode.HeatingCoolingMode()
}

func (t *thermostat) SetTargetMode(value model.HeatCoolModeType) {
	t.thermostat.TargetMode.SetHeatingCoolingMode(value)
}

func (t *thermostat) TargetMode() model.HeatCoolModeType {
	return t.thermostat.TargetMode.HeatingCoolingMode()
}
