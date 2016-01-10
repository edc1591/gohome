package controller

import (
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/container"
	"github.com/brutella/hc/netio/data"
	"github.com/brutella/log"

	"bytes"
	"encoding/json"

	"io"
	"io/ioutil"
	"net"
	"net/url"
)

// CharacteristicController implements the CharacteristicsHandler interface and provides
// read (GET) and write (POST) interfaces to the managed characteristics.
type CharacteristicController struct {
	container *container.Container
}

// NewCharacteristicController returns a new characteristic controller.
func NewCharacteristicController(m *container.Container) *CharacteristicController {
	return &CharacteristicController{container: m}
}

// HandleGetCharacteristics handles a get characteristic request.
func (ctr *CharacteristicController) HandleGetCharacteristics(form url.Values) (io.Reader, error) {
	var b bytes.Buffer
	aid, cid, err := ParseAccessoryAndCharacterID(form.Get("id"))
	containerChar := ctr.GetCharacteristic(aid, cid)
	if containerChar == nil {
		log.Printf("[WARN] No characteristic found with aid %d and iid %d\n", aid, cid)
		return &b, nil
	}

	chars := data.NewCharacteristics()
	char := data.Characteristic{AccessoryID: aid, ID: cid, Value: containerChar.GetValue(), Events: containerChar.EventsEnabled()}
	chars.AddCharacteristic(char)

	result, err := json.Marshal(chars)
	if err != nil {
		log.Println("[ERRO]", err)
	}

	b.Write(result)
	return &b, err
}

// HandleUpdateCharacteristics handles an update characteristic request. The bytes must represent
// a data.Characteristics json.
func (ctr *CharacteristicController) HandleUpdateCharacteristics(r io.Reader, conn net.Conn) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	var chars data.Characteristics
	err = json.Unmarshal(b, &chars)
	if err != nil {
		return err
	}

	log.Println("[VERB]", string(b))

	for _, c := range chars.Characteristics {
		characteristic := ctr.GetCharacteristic(c.AccessoryID, c.ID)
		if characteristic == nil {
			log.Printf("[ERRO] Could not find characteristic with aid %d and iid %d\n", c.AccessoryID, c.ID)
			continue
		}

		if c.Value != nil {
			characteristic.SetValueFromConnection(c.Value, conn)
		}

		if events, ok := c.Events.(bool); ok == true {
			characteristic.SetEventsEnabled(events)
		}
	}

	return err
}

// GetCharacteristic returns the characteristic with the specified accessory and characteristic id.
func (ctr *CharacteristicController) GetCharacteristic(accessoryID int64, characteristicID int64) *characteristic.Characteristic {
	for _, a := range ctr.container.Accessories {
		if a.GetID() == accessoryID {
			for _, s := range a.GetServices() {
				for _, c := range s.GetCharacteristics() {
					if c.GetID() == characteristicID {
						return c
					}
				}
			}
		}
	}
	return nil
}
