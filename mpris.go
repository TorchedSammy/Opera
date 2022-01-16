package main

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

type player struct{
	metadata map[string]interface{}
}

func (p *player) updateData(m map[string]interface{}) {
	for k, v := range m {
		p.metadata[k] = v
	}
}

func (p *player) Get(iface, prop string) (dbus.Variant, *dbus.Error) {
	fmt.Println("Get", iface, prop)

	switch prop {
		case "Metadata":
			return dbus.MakeVariant(p.metadata), nil
		case "Identity":
			return dbus.MakeVariant("Opera"), nil
		case "Position":
			return dbus.MakeVariant(int64(0)), nil
		case "PlaybackStatus":
			return dbus.MakeVariant("Playing"), nil
		default:
			return dbus.MakeVariant(""), nil
	}
}

func (p *player) GetAll(iface string) (map[string]dbus.Variant, *dbus.Error) {
	fmt.Println("GetAll", iface)

	return map[string]dbus.Variant{
		"PlaybackStatus": dbus.MakeVariant("Playing"),
		"LoopStatus": dbus.MakeVariant("None"),
		"Volume": dbus.MakeVariant(float64(1.0)),
		"Shuffle": dbus.MakeVariant(false),
		"Metadata": dbus.MakeVariant(p.metadata),
	}, nil
}
