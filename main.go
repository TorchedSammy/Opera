package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/godbus/dbus/v5"
	"github.com/gorilla/websocket"
)

var busName = "org.mpris.MediaPlayer2.opera"
var objectPath = dbus.ObjectPath("/org/mpris/MediaPlayer2")
var objectInterface = "org.mpris.MediaPlayer2.Player"
var wsUrl = "ws://127.0.0.1:24050/ws"
// uncommented is what gets updated
var mdata = map[string]interface{}{
//	"mpris:trackid": dbus.ObjectPath("/1"),
	"mpris:length": 0,
	"xesam:artist": []string{"Ashnikko"},
	"xesam:title": "Hi Its Me",
//	"xesam:genre": []string{"Pop"},
//	"xesam:audioBPM": 120,
//	"xesam:trackNumber": 1,
//	"xesam:discNumber": 1,
//	"xesam:url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
//	"mpris:artUrl": "https://i.ytimg.com/vi/dQw4w9WgXcQ/maxresdefault.jpg",
}
var currentSet int

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		panic(err)
	}

	c, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	initData := getData(c)
	currentSet = initData.GetSetID()

	opera := &player{
		metadata: mdata,
	}
	opera.updateData(map[string]interface{}{
		"xesam:artist": []string{initData.GetMusicArtist()},
		"xesam:title": initData.GetMusicTitle(),
	})

	conn.Export(opera, objectPath, "org.freedesktop.DBus.Properties")
	conn.RequestName(busName, dbus.NameFlagDoNotQueue)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	go func() {
		for {
			d := getData(c)
			// if set id is same as current set, do nothing
			if d.GetSetID() == currentSet {
				continue
			}
			currentSet = d.GetSetID()
			opera.updateData(map[string]interface{}{
				"xesam:artist": []string{d.GetMusicArtist()},
				"xesam:title": d.GetMusicTitle(),
			})

			fmt.Printf("%s - %s\n", d.GetMusicArtist(), d.GetMusicTitle())
			conn.Emit(objectPath, "org.freedesktop.DBus.Properties", "PropertiesChanged", map[string]dbus.Variant{
				"Metadata": dbus.MakeVariant(opera.metadata),
			})
		}
	}()

	go func() {
		<-sigchan
		conn.ReleaseName(busName)
		conn.Close()
		os.Exit(0)
	}()
	for {}
}

func getData(c *websocket.Conn) gosumemory {
	_, msg, err := c.ReadMessage()
	if err != nil {
		panic(err)
	}
	var data gosumemory
	err = json.Unmarshal(msg, &data)
	if err != nil {
		panic(err)
	}

	return data
}
