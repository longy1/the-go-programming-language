package track

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type ByArtist []*Track

func (t ByArtist) Len() int           { return len(t) }
func (t ByArtist) Less(i, j int) bool { return t[i].Artist < t[j].Artist }
func (t ByArtist) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

type CustomSort struct {
	T        []*Track
	LessFunc func(x, y *Track) bool
}

func (x CustomSort) Len() int           { return len(x.T) }
func (x CustomSort) Less(i, j int) bool { return x.LessFunc(x.T[i], x.T[j]) }
func (x CustomSort) Swap(i, j int)      { x.T[i], x.T[j] = x.T[j], x.T[i] }

type multipleSort struct {
	t        []*Track
	lessFunc []func(x, y *Track) bool
}

func (x *multipleSort) Sort(tracks []*Track) {
	x.t = tracks
	sort.Sort(x)
}

func (x *multipleSort) Len() int { return len(x.t) }

func (x *multipleSort) Less(i, j int) bool {
	for _, f := range x.lessFunc {
		if f(x.t[i], x.t[j]) {
			return true
		} else if f(x.t[j], x.t[i]) {
			return false
		}
	}
	return false
}

func (x *multipleSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func OrderedBy(less ...func(x, y *Track) bool) *multipleSort {
	return &multipleSort{lessFunc: less}
}

func title(x, y *Track) bool  { return x.Title < y.Title }
func artist(x, y *Track) bool { return x.Artist < y.Artist }
func album(x, y *Track) bool  { return x.Album < y.Album }
func year(x, y *Track) bool   { return x.Year < y.Year }
func length(x, y *Track) bool { return x.Length < y.Length }
