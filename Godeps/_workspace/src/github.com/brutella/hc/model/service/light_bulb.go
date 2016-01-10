package service

import (
	"github.com/brutella/hc/model/characteristic"
)

// LightBulb is a service to represent a light bulb.
type LightBulb struct {
	*Service
	On         *characteristic.On
	Name       *characteristic.Name
	Brightness *characteristic.Brightness
	Saturation *characteristic.Saturation
	Hue        *characteristic.Hue
}

// NewLightBulb returns a light bulb service.
func NewLightBulb(name string, on bool) *LightBulb {
	onChar := characteristic.NewOn(on)
	nameChar := characteristic.NewName(name)
	brightness := characteristic.NewBrightness(100) // 100%
	saturation := characteristic.NewSaturation(0.0)
	hue := characteristic.NewHue(0.0)

	svc := New()
	svc.Type = typeLightbulb
	svc.AddCharacteristic(onChar.Characteristic)
	svc.AddCharacteristic(nameChar.Characteristic)
	svc.AddCharacteristic(brightness.Characteristic)
	svc.AddCharacteristic(saturation.Characteristic)
	svc.AddCharacteristic(hue.Characteristic)

	return &LightBulb{svc, onChar, nameChar, brightness, saturation, hue}
}
