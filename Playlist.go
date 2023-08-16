package vlcremote

import "encoding/xml"

type PlaylistItem struct {
	XMLName  xml.Name `xml:"leaf"`
	Name     string   `xml:"name,attr"`
	Id       int      `xml:"id,attr"`
	Duration int      `xml:"duration,attr"`
	Uri      string   `xml:"uri,attr"`
	Current  string   `xml:"current,attr"`
}

type Playlist struct {
	XMLName xml.Name       `xml:"node"`
	Name    string         `xml:"name,attr"`
	Id      int            `xml:"id,attr"`
	Media   []PlaylistItem `xml:"leaf"`
}

type PlaylistRoot struct {
	XMLName   xml.Name   `xml:"node"`
	Name      string     `xml:"name,attr"`
	Id        int        `xml:"id,attr"`
	Playlists []Playlist `xml:"node"`
}
