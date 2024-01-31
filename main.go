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

func main() {
	spaces := []Space{
		{name: "entropia", url: "https://club.entropia.de/spaceapi", status: "state.open"},
	}

	for _, space := range spaces {
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
