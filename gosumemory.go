package main

import "time"

// struct of data gosumemory sends over ws
type gosumemory struct {
	Settings struct {
		Folders struct {
			Songs string `json:"songs"` // path to songs folder
		} `json:"folders"`
	} `json:"settings"`
	Menu gosumemoryMenu `json:"menu"`
}

type gosumemoryMenu struct {
	Beatmap gosumemoryBeatmap `json:"bm"`
}

type gosumemoryBeatmap struct {
	ID int `json:"id"`
	SetID int `json:"set"`
	Metadata beatmapMetadata `json:"metadata"`
	Time gosumemoryBeatmapTime `json:"time"`
}

type gosumemoryBeatmapTime struct {
	Current int `json:"current"`
}

type beatmapMetadata struct {
	Artist string `json:"artist"`
	Title string `json:"title"`
}

func (g *gosumemory) GetSetID() int {
	return g.Menu.Beatmap.SetID
}

func (g *gosumemory) GetMusicTitle() string {
	return g.Menu.Beatmap.Metadata.Title
}

func (g *gosumemory) GetMusicArtist() string {
	return g.Menu.Beatmap.Metadata.Artist
}

func (g *gosumemory) GetPosition() time.Duration {
	// make duration from current time
	currentTime := time.Duration(g.Menu.Beatmap.Time.Current) * time.Millisecond
	return currentTime
}
