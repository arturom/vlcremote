package vlcremote

import (
	"fmt"
)

type HighLevelClient struct {
	vlc VlcClient
}

func NewHighLevelClient(vlc VlcClient) *HighLevelClient {
	c := new(HighLevelClient)
	c.vlc = vlc
	return c
}

func (c HighLevelClient) AddFileToPlaylist(filePath string) (*Status, error) {
	uri := fmt.Sprintf("file://%s", filePath)
	return c.vlc.AddToPlaylist(uri)
}

func (c HighLevelClient) AddFilesToPlaylist(filePaths []string) (*Status, error) {
	var status *Status
	for _, filePath := range filePaths {
		s, err := c.AddFileToPlaylist(filePath)
		if err != nil {
			return nil, err
		}
		status = s
	}
	return status, nil
}

func (c HighLevelClient) GetCurrentlyPlaying() (*PlaylistItem, error) {
	pl, err := c.vlc.GetPlaylist()
	if err != nil {
		return nil, err
	}

	for _, item := range pl.Playlists[0].Media {
		if item.Current == "current" {
			return &item, nil
		}
	}

	return nil, fmt.Errorf("no current item playing")
}
