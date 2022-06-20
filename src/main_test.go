package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ComparableData struct {
	titles, frags string
}

// { "th", "fr", "pi", "sh", "wu", "ar", "il", "ne", "se", "pl"},
func TestGetLevDistance(t *testing.T) {
	data := []struct {
		name, title, frag string
	}{
		{name: "frag-th", title: "Taming of the Shrew", frag: "th"},
		{name: "frag-Tameng", title: "Taming of the Shrew", frag: "Tameng"},
		{name: "frag-fr", title: "Taming of the Shrew", frag: "fr"},
		{name: "frag-sh", title: "Taming of the Shrew", frag: "sh"},
		{name: "frag-sa", title: "Taming of the Shrew", frag: "sa"},
		{name: "frag-of", title: "Taming of the Shrew", frag: "of"},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			dist := getLevDistance(tt.title, tt.frag)

			fmt.Printf("Tested %s vs %s and got a distance of %d.\n", tt.title, tt.frag, dist)
		})
	}

}

func TestURL(t *testing.T) {

}

type Tests struct {
	name  string
	url   string
	query string
}

func TestShowResult(t *testing.T) {

	loadList()

	tests := []Tests{
		{name: "frag-th", url: "http://localhost:9000/autocomplete", query: "term=th"},
		{name: "frag-sh", url: "http://localhost:9000/autocomplete", query: "term=sh"},
		{name: "frag-mark", url: "http://localhost:9000/autocomplete", query: "term=mark"},
		{name: "frag-fr", url: "http://localhost:9000/autocomplete", query: "term=fr"},
		{name: "frag-ss", url: "http://localhost:9000/autocomplete", query: "term=ss"},
		{name: "frag-we", url: "http://localhost:9000/autocomplete", query: "term=we"},
		{name: "frag-wind", url: "http://localhost:9000/autocomplete", query: "term=wind"},
		{name: "frag-hen", url: "http://localhost:9000/autocomplete", query: "term=hen"},
		{name: "frag-jum", url: "http://localhost:9000/autocomplete", query: "term=jum"},
		{name: "frag-jum", url: "http://localhost:9000/", query: "ter=jum"},
	}

	l, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	server := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				showResult(w, r)
			},
		),
	)
	defer server.Close()
	server.Listener.Close()
	server.Listener = l
	server.Start()

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			resp, err := http.Get(d.url + "?" + d.query)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			defer resp.Body.Close()
		})
	}
}
