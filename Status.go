package vlcremote

import (
	"encoding/xml"
)

type Status struct {
	XMLName       xml.Name `xml:"root"`
	FullScreen    string   `xml:"fullscreen"`
	AspectRation  string   `xml:"aspectratio"`
	AudioDelay    int      `xml:"audiodelay"`
	ApiVersion    int      `xml:"apiversion"`
	CurrentPlId   string   `xml:"currentplid"`
	Time          int      `xml:"time"`
	Volume        int      `xml:"volume"`
	Length        int      `xml:"length"`
	Random        bool     `xml:"random"`
	Rate          int      `xml:"rate"`
	State         string   `xml:"state"`
	Loop          bool     `xml:"loop"`
	Version       string   `xml:"version"`
	Position      float64  `xml:"position"`
	Repeat        bool     `xml:"repeat"`
	SubtitleDelay int      `xml:"subtitledelay"`
}
