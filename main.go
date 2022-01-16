package main

import (
	"os"
	"os/signal"

	"github.com/godbus/dbus/v5"
)

var busName = "org.mpris.MediaPlayer2.opera"
var objectPath = dbus.ObjectPath("/org/mpris/MediaPlayer2")
var objectInterface = "org.mpris.MediaPlayer2.Player"

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	opera := &player{
		// test values, it being one of my favourite songs :)
		metadata: map[string]interface{}{
			"mpris:trackid": dbus.ObjectPath("/1"),
			"mpris:length": 100000,
			"xesam:artist": []string{"Ashnikko"},
			"xesam:title": "Hi Its Me",
			"xesam:album": "Hi It's Me",
			"xesam:albumArtist": []string{"Ashnikko"},
			"xesam:genre": []string{"Pop"},
			"xesam:audioBPM": 120,
			"xesam:trackNumber": 1,
			"xesam:discNumber": 1,
			"xesam:url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			"mpris:artUrl": "https://i.ytimg.com/vi/dQw4w9WgXcQ/maxresdefault.jpg",
		},
	}
	conn.RequestName(busName, dbus.NameFlagDoNotQueue)
	conn.Export(opera, objectPath, "org.freedesktop.DBus.Properties")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	go func() {
		<-sigchan
		conn.ReleaseName(busName)
		conn.Close()
		os.Exit(0)
	}()
	for {}
}
