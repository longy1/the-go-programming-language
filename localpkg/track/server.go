package track

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"time"
)

const templ = `<h1>Tracks</h1>
<table>
<tr style='text-align: left'>
<th><a href='?sorted=title'>Title</a></th>
<th><a href='?sorted=artist'>Artist</a></th>
<th><a href='?sorted=album'>Album</a></th>
<th><a href='?sorted=year'>Year</a></th>
<th><a href='?sorted=length'>Length</a></th>
</tr>
{{range .}}
<tr>
<td>{{.Title}}</td>
<td>{{.Artist}}</td>
<td>{{.Album}}</td>
<td>{{.Year}}</td>
<td>{{.Length}}</td>
</tr>
{{end}}
</table>
`

var testTracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, tLength("3m38s")},
	{"Go", "Moby", "Moby", 1992, tLength("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, tLength("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, tLength("4m24s")},
}

func tLength(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// change order
	q, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Print(err)
	}
	if len(q["sorted"]) > 0 {
		switch q["sorted"][0] {
		case "length":
			OrderedBy(length).Sort(testTracks)
		case "year":
			OrderedBy(year).Sort(testTracks)
		case "album":
			OrderedBy(album).Sort(testTracks)
		case "artist":
			OrderedBy(artist).Sort(testTracks)
		case "title":
			OrderedBy(title).Sort(testTracks)
		}
	}
	resp := template.Must(template.New("response").Parse(templ))
	if err := resp.Execute(w, testTracks); err != nil {
		log.Fatal(err)
	}
}

func StartServer() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe("127.0.0.1:1100", nil))
}
