package vlcremote

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type VlcClient struct {
	host     string
	password string
}

func NewVlcClient(host string, password string) *VlcClient {
	c := new(VlcClient)
	c.host = host
	c.password = password
	return c
}

func (c VlcClient) get(path string, queryParams url.Values) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.host, path)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range queryParams {
		q.Set(k, v[0])
	}
	req.URL.RawQuery = q.Encode()

	authString := fmt.Sprintf(":%s", c.password)
	authString = base64.StdEncoding.EncodeToString([]byte(authString))
	authString = fmt.Sprintf("Basic %s", authString)

	req.Header.Set("Authorization", authString)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func unmarshallResponse(resp *http.Response, v any) error {
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	err := xml.Unmarshal(buf.Bytes(), v)
	return err
}

func (c VlcClient) GetPlaylist() (*PlaylistRoot, error) {
	path := "requests/playlist.xml"
	resp, err := c.get(path, make(url.Values))

	if err != nil {
		return nil, err
	}

	var plRoot PlaylistRoot
	if err = unmarshallResponse(resp, &plRoot); err != nil {
		return nil, err
	}

	return &plRoot, nil
}

func (c VlcClient) Command(params url.Values) (*Status, error) {
	path := "requests/status.xml"
	resp, err := c.get(path, params)
	if err != nil {
		return nil, err
	}

	var video Status
	if err = unmarshallResponse(resp, &video); err != nil {
		return nil, err
	}

	return &video, nil
}

func (c VlcClient) Play() (*Status, error) {
	params := make(url.Values)
	params.Set("command", "pl_play")
	return c.Command(params)
}

func (c VlcClient) Pause() (*Status, error) {
	params := make(url.Values)
	params.Set("command", "pl_pause")
	return c.Command(params)
}

func (c VlcClient) Stop() (*Status, error) {
	params := make(url.Values)
	params.Set("command", "pl_stop")
	return c.Command(params)
}

func (c VlcClient) Previous() (*Status, error) {
	params := make(url.Values)
	params.Set("command", "pl_previous")
	return c.Command(params)
}

func (c VlcClient) Next() (*Status, error) {
	params := make(url.Values)
	params.Set("command", "pl_next")
	return c.Command(params)
}

func (c VlcClient) SetVolume(val string) (*Status, error) {
	params := make(url.Values)
	params.Set("command", "volume")
	params.Set("val", val)
	return c.Command(params)
}

func (c VlcClient) FullScreen() (*Status, error) {
	params := make(url.Values)
	params.Set("command", "fullscreen")
	return c.Command(params)
}

func (c VlcClient) AddToPlaylistAndPlay(uri string) (*Status, error) {
	params := make(url.Values)
	params.Set("command", "in_play")
	params.Set("input", uri)
	return c.Command(params)
}

func (c VlcClient) AddToPlaylist(uri string) (*Status, error) {
	params := make(url.Values)
	params.Set("command", "in_enqueue")
	params.Set("input", uri)
	return c.Command(params)
}

func (c VlcClient) RemoveFromPlaylist(id int) (*Status, error) {
	params := make(url.Values)
	params.Set("command", "pl_delete")
	params.Set("id", strconv.Itoa(id))
	return c.Command(params)
}

func (c VlcClient) EmptyPlaylist() (*Status, error) {
	params := make(url.Values)
	params.Set("command", "pl_empty")
	return c.Command(params)
}
