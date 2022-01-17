package main

import (
	"fmt"
	"time"

	"github.com/Pauloo27/go-mpris"
	"github.com/godbus/dbus/v5"
)

type player struct{
	playbackStatus mpris.PlaybackStatus
	position time.Duration
}

func (p *player) Get(iface, prop string) (dbus.Variant, *dbus.Error) {
	fmt.Println("Get", iface, prop)

	switch prop {
		case "Metadata":
			return dbus.MakeVariant(mdata.toMap()), nil
		case "Identity":
			return dbus.MakeVariant("Opera"), nil
		case "Position":
			return dbus.MakeVariant(int64(0)), nil
		case "PlaybackStatus":
			return dbus.MakeVariant(p.playbackStatus), nil
		default:
			return dbus.MakeVariant(""), nil
	}
}

func (p *player) GetAll(iface string) (map[string]dbus.Variant, *dbus.Error) {
	fmt.Println("GetAll", iface)

	return map[string]dbus.Variant{
		"PlaybackStatus": dbus.MakeVariant(p.playbackStatus),
		"LoopStatus": dbus.MakeVariant("None"),
		"Volume": dbus.MakeVariant(float64(1.0)),
		"Shuffle": dbus.MakeVariant(false),
		"Metadata": dbus.MakeVariant(mdata.toMap()),
	}, nil
}

func (p *player) setPlaybackStatus(status mpris.PlaybackStatus) {
	p.playbackStatus = status
}
