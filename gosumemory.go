package main

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
