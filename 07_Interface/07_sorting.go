package main

import (
	"The.Go.Programming.Language/localpkg/track"
	"fmt"
	"sort"
	"time"
)

var tracks = []*track.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {
	track.PrintTracks(tracks)
	fmt.Println()

	sort.Sort(track.ByArtist(tracks))
	track.PrintTracks(tracks)
	fmt.Println()

	sort.Sort(sort.Reverse(track.ByArtist(tracks)))
	track.PrintTracks(tracks)
	fmt.Println()

	custom := track.CustomSort{T: tracks}

	custom.LessFunc = func(x, y *track.Track) bool {
		return x.Year < y.Year
	}
	sort.Sort(custom)
	track.PrintTracks(tracks)
	fmt.Println()

	custom.LessFunc = func(x, y *track.Track) bool {
		return x.Artist < y.Artist
	}
	sort.Sort(custom)
	track.PrintTracks(tracks)
	fmt.Println()
}
