package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"io"
	"net/http"
)

type Space struct {
	name   string
	url    string
	status string
}

var Spaces []Space = []Space{
	{name: "entropia", url: "https://club.entropia.de/spaceapi", status: "state.open"},
	{name: "stratum0", url: "https://status.stratum0.org/status.json", status: "isOpen"},
}

func main() {
	for _, space := range Spaces {
		resp, err := http.Get(space.url)
		if err != nil {
			panic(err)
		}
		raw, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		open := gojsonq.New().FromString(string(raw)).Find(space.status).(bool)
		var status string
		if open {
			status = "open"
		} else {
			status = "closed"
		}

		fmt.Printf("%s\t%s\n", space.name, status)
	}
}
