package main

import (
	"time"

	"github.com/godbus/dbus/v5"
)

// metadata is a struct to easily define mpris metadata
// reference for fields: https://www.freedesktop.org/wiki/Specifications/mpris-spec/metadata
type metadata struct {
	TrackID string // unique track id, not really sure how this works
	Title string // music title
	Artist string // music artist
	Length time.Duration // length of music, when representing for dbus itll be in microseconds
	Genres []string
	BPM int
	ArtURL string // url to album art (file://path or http://url possibly?)
}

// will represent fields in metadata struct in a map, any fields that are
// empty/not set wont be present in the map
func (m *metadata) toMap() map[string]interface{} {
	var mprisMap = make(map[string]interface{})
	if m.TrackID != "" {
		mprisMap["mpris:trackid"] = dbus.ObjectPath(m.TrackID)
	}
	if m.Title != "" {
		mprisMap["xesam:title"] = m.Title
	}
	if m.Artist != "" {
		mprisMap["xesam:artist"] = m.Artist
	}
	if m.Length != 0 {
		mprisMap["mpris:length"] = m.Length.Microseconds()
	}
	if len(m.Genres) != 0 {
		mprisMap["xesam:genre"] = m.Genres
	}
	if m.BPM != 0 {
		mprisMap["xesam:audioBPM"] = m.BPM
	}
	if m.ArtURL != "" {
		mprisMap["mpris:artUrl"] = m.ArtURL
	}
	return mprisMap
}
