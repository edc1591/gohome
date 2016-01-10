package service

import (
	"github.com/brutella/hc/model"
	"testing"
)

func TestThermostat(t *testing.T) {
	thermostat := NewThermostat("Testthermostat", 10.5, -10, 100, 1)

	if is, want := thermostat.Type, typeThermostat; is != want {
		t.Fatalf("type: is=%v want=%v", is, want)
	}
	if is, want := thermostat.Name.GetValue(), "Testthermostat"; is != want {
		t.Fatalf("name: is=%v want=%v", is, want)
	}
	if is, want := thermostat.Temp.GetValue(), 10.5; is != want {
		t.Fatalf("temp: is=%v want=%v", is, want)
	}
	if is, want := thermostat.TargetTemp.GetValue(), 10.5; is != want {
		t.Fatalf("targettemp: is=%v want=%v", is, want)
	}
	// TODO(brutella): uint8 cast should not be required!
	if is, want := thermostat.Mode.GetValue(), uint8(model.HeatCoolModeOff); is != want {
		t.Fatalf("mode: is=%v want=%v", is, want)
	}
	// TODO(brutella): uint8 cast should not be required!
	if is, want := thermostat.TargetMode.GetValue(), uint8(model.HeatCoolModeOff); is != want {
		t.Fatalf("targetmode: is=%v want=%v", is, want)
	}
}

func TestThermometer(t *testing.T) {
	thermometer := NewThermometer("Thermometer", 10.5, -10, 100, 1)

	if is, want := thermometer.Type, typeThermostat; is != want {
		t.Fatalf("type: is=%v want=%v", is, want)
	}
	if is, want := thermometer.Name.GetValue(), "Thermometer"; is != want {
		t.Fatalf("name: is=%v want=%v", is, want)
	}
	if is, want := thermometer.Temp.GetValue(), 10.5; is != want {
		t.Fatalf("temp: is=%v want=%v", is, want)
	}
	if is, want := thermometer.TargetTemp.GetValue(), 10.5; is != want {
		t.Fatalf("targettemp: is=%v want=%v", is, want)
	}
	// TODO(brutella): cast should not be required!
	if is, want := thermometer.Mode.GetValue(), uint8(model.HeatCoolModeOff); is != want {
		t.Fatalf("mode: is=%v want=%v", is, want)
	}
	// TODO(brutella): cast should not be required!
	if is, want := thermometer.TargetMode.GetValue(), uint8(model.HeatCoolModeOff); is != want {
		t.Fatalf("targetmode: is=%v want=%v", is, want)
	}
}
