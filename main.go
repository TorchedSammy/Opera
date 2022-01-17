package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/godbus/dbus/v5"
	"github.com/gorilla/websocket"
	"github.com/Pauloo27/go-mpris"
)

var busName = "org.mpris.MediaPlayer2.opera"
var objectPath = dbus.ObjectPath("/org/mpris/MediaPlayer2")
var objectInterface = "org.mpris.MediaPlayer2.Player"
var wsUrl = "ws://127.0.0.1:24050/ws"
var mdata metadata
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
	mdata = metadata{
		Title: initData.GetMusicTitle(),
		Artist: initData.GetMusicArtist(),
	}
	currentSet = initData.GetSetID()

	opera := &player{
		playbackStatus: mpris.PlaybackPlaying,
	}

	conn.Export(opera, objectPath, "org.freedesktop.DBus.Properties")
	conn.RequestName(busName, dbus.NameFlagDoNotQueue)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	go func() {
		for {
			d := getData(c)
			prevPos := opera.position
			opera.position = d.GetPosition()
			// if set id is same as current set, check our pos
			if d.GetSetID() == currentSet {
				// this is kinda weird as sometimes the position is not updated while
				// its playing which causes a pause once and playing after
				// only fix seems to be changing gosumemory update times ...
				if prevPos == d.GetPosition() && opera.playbackStatus != mpris.PlaybackPaused {
					// set status to paused if the position is same as before,
					// but only when our status isnt paused to not spam dbus
					opera.setPlaybackStatus(mpris.PlaybackPaused)
					fmt.Println("Paused")
					conn.Emit(objectPath, "org.freedesktop.DBus.Properties.PropertiesChanged", objectInterface, map[string]dbus.Variant{
						"PlaybackStatus": dbus.MakeVariant(mpris.PlaybackPaused),
					})
				} else if prevPos != d.GetPosition() && opera.playbackStatus == mpris.PlaybackPaused {
					opera.setPlaybackStatus(mpris.PlaybackPlaying)
					fmt.Println("Playing")
					conn.Emit(objectPath, "org.freedesktop.DBus.Properties.PropertiesChanged", objectInterface, map[string]dbus.Variant{
						"PlaybackStatus": dbus.MakeVariant(mpris.PlaybackPlaying),
					})
				}
				continue
			}
			currentSet = d.GetSetID()
			mdata.Title = d.GetMusicTitle()
			mdata.Artist = d.GetMusicArtist()

			fmt.Printf("%s - %s\n", d.GetMusicArtist(), d.GetMusicTitle())
			conn.Emit(objectPath, "org.freedesktop.DBus.Properties", "PropertiesChanged", map[string]dbus.Variant{
				"Metadata": dbus.MakeVariant(mdata.toMap()),
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
