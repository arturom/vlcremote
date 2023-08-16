package vlcremote

import (
	"fmt"
	"io/fs"
	"os"
	"testing"
	"time"
)

func printPlaylist(p PlaylistRoot) {
	fmt.Println("root", p.Id, p.Name)
	for _, p := range p.Playlists {
		fmt.Println("playlist", p.Id, p.Name)
		for _, i := range p.Media {
			fmt.Println("media", i.Id, i.Name, i.Duration, i.Uri, i.Current)
		}
	}
}

func printStatus(v Status) {
	fmt.Println("video", v.State)
}

func TestAll(t *testing.T) {
	port := 8080
	password := "password"

	fmt.Println(port, password)

	cmd, err := OpenVlc(port, password)
	if err != nil {
		t.Error(err)
	}

	time.Sleep(1000000000 * 1)

	host := fmt.Sprintf("http://localhost:%d", port)
	vlc := NewVlcClient(host, password)

	plRoot, err := vlc.GetPlaylist()
	if err != nil {
		t.Error(err)
	}
	printPlaylist(*plRoot)

	dir := "/Users/arturomejia/Documents/ts"
	fsys := os.DirFS(dir)
	files, err := fs.Glob(fsys, "*/*.ts")
	if err != nil {
		t.Error(err)
	}
	for _, f := range files {
		status, err := vlc.AddToPlaylist("file://" + dir + "/" + f)
		if err != nil {
			t.Error(err)
		}
		printStatus(*status)
		fmt.Println()
	}

	status, err := vlc.Stop()
	if err != nil {
		t.Error(err)
	}
	printStatus(*status)

	status, err = vlc.Play()
	if err != nil {
		t.Error(err)
	}
	printStatus(*status)

	time.Sleep(1000000000 * 1)

	status, err = vlc.Next()
	if err != nil {
		t.Error(err)
	}
	printStatus(*status)

	time.Sleep(1000000000 * 1)

	status, err = vlc.Next()
	if err != nil {
		t.Error(err)
	}
	printStatus(*status)

	time.Sleep(1000000000 * 1)

	plRoot, err = vlc.GetPlaylist()
	if err != nil {
		t.Error(err)
	}
	printPlaylist(*plRoot)
	cmd.Wait()

	if !cmd.ProcessState.Exited() {
		if err = cmd.Process.Kill(); err != nil {
			t.Error(err)
		}
	}

}
