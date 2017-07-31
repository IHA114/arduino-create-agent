// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "arduino-create-agent": Application Media Types
//
// Command:
// $ main

package app

import (
	"github.com/goadesign/goa"
)

// ArduinoAgentDiscover media type (default view)
//
// Identifier: application/vnd.arduino.agent.discover+json; view=default
type ArduinoAgentDiscover struct {
	// A list of devices connected through the network
	Network ArduinoAgentDiscoverNetworkCollection `form:"network" json:"network" xml:"network"`
	// A list of devices connected through the serial port
	Serial ArduinoAgentDiscoverSerialCollection `form:"serial" json:"serial" xml:"serial"`
}

// Validate validates the ArduinoAgentDiscover media type instance.
func (mt *ArduinoAgentDiscover) Validate() (err error) {
	if mt.Serial == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "serial"))
	}
	if mt.Network == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "network"))
	}
	if err2 := mt.Network.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if err2 := mt.Serial.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// ArduinoAgentDiscoverNetwork media type (default view)
//
// Identifier: application/vnd.arduino.agent.discover.network+json; view=default
type ArduinoAgentDiscoverNetwork struct {
	// IP Address
	Address string `form:"address" json:"address" xml:"address"`
	// Informations about the device
	Info string `form:"info" json:"info" xml:"info"`
	// The friendly name given to the device
	Name string `form:"name" json:"name" xml:"name"`
	// IP Port
	Port int `form:"port" json:"port" xml:"port"`
}

// Validate validates the ArduinoAgentDiscoverNetwork media type instance.
func (mt *ArduinoAgentDiscoverNetwork) Validate() (err error) {
	if mt.Address == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "address"))
	}

	if mt.Info == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "info"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// ArduinoAgentDiscoverNetworkCollection is the media type for an array of ArduinoAgentDiscoverNetwork (default view)
//
// Identifier: application/vnd.arduino.agent.discover.network+json; type=collection; view=default
type ArduinoAgentDiscoverNetworkCollection []*ArduinoAgentDiscoverNetwork

// Validate validates the ArduinoAgentDiscoverNetworkCollection media type instance.
func (mt ArduinoAgentDiscoverNetworkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ArduinoAgentDiscoverSerial media type (default view)
//
// Identifier: application/vnd.arduino.agent.discover.serial+json; view=default
type ArduinoAgentDiscoverSerial struct {
	// Vendor ID
	Pid string `form:"pid" json:"pid" xml:"pid"`
	// The port through which it's connected
	Port string `form:"port" json:"port" xml:"port"`
	// The Serial Number
	Serial *string `form:"serial,omitempty" json:"serial,omitempty" xml:"serial,omitempty"`
	// Vendor ID
	Vid string `form:"vid" json:"vid" xml:"vid"`
}

// Validate validates the ArduinoAgentDiscoverSerial media type instance.
func (mt *ArduinoAgentDiscoverSerial) Validate() (err error) {
	if mt.Vid == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "vid"))
	}
	if mt.Pid == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "pid"))
	}
	if mt.Port == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "port"))
	}
	return
}

// ArduinoAgentDiscoverSerialCollection is the media type for an array of ArduinoAgentDiscoverSerial (default view)
//
// Identifier: application/vnd.arduino.agent.discover.serial+json; type=collection; view=default
type ArduinoAgentDiscoverSerialCollection []*ArduinoAgentDiscoverSerial

// Validate validates the ArduinoAgentDiscoverSerialCollection media type instance.
func (mt ArduinoAgentDiscoverSerialCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
