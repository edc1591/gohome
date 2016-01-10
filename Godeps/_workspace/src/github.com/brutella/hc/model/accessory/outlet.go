package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
	"net"
)

type outlet struct {
	*Accessory
	outlet *service.Outlet

	onChanged    func(bool)
	inUseChanged func(bool)
}

// NewOutlet returns an outlet which implements model.Outlet.
func NewOutlet(info model.Info) *outlet {
	accessory := New(info)
	s := service.NewOutlet(info.Name, false, true) // off

	accessory.AddService(s.Service)

	sw := outlet{accessory, s, nil, nil}

	s.On.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		if sw.onChanged != nil {
			sw.onChanged(s.On.On())
		}
	})

	s.InUse.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		if sw.inUseChanged != nil {
			sw.inUseChanged(s.InUse.InUse())
		}
	})

	return &sw
}

func (o *outlet) SetOn(on bool) {
	o.outlet.On.SetOn(on)
}

func (o *outlet) IsOn() bool {
	return o.outlet.On.On()
}

func (o *outlet) SetInUse(on bool) {
	o.outlet.InUse.SetInUse(on)
}

func (o *outlet) IsInUse() bool {
	return o.outlet.InUse.InUse()
}

func (o *outlet) OnStateChanged(fn func(bool)) {
	o.onChanged = fn
}

func (o *outlet) InUseStateChanged(fn func(bool)) {
	o.inUseChanged = fn
}
