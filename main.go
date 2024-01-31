package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strconv"
)

const API = "https://api.spaceapi.io/"

type Space struct {
	Data SpaceData `json:"data"`
}

type SpaceData struct {
	Name  string     `json:"space"`
	State SpaceState `json:"state"`
}

type SpaceState struct {
	Open any `json:"open"`
}

type Status map[string]bool

// orderKeys orders the keys of a space status map alphabetically
func (s Status) orderKeys() []string {
	keys := make([]string, 0)
	for key := range s {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}

// get performs an HTTP GET request, returning the result as a string
func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("expected 200 OK")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// parseAPI parses the Space API output into a map consisting of space/status pairs
func parseAPI(api []byte) (Status, error) {
	// parse the actual json
	spaces := make([]Space, 0)
	err := json.Unmarshal(api, &spaces)
	if err != nil {
		return nil, err
	}

	// iterate through all entries and populate the actual map from it
	status := make(Status)
	for _, space := range spaces {
		// skip spaces with missing name
		if space.Data.Name == "" {
			continue
		}

		open, err := strconv.ParseBool(fmt.Sprintf("%t", space.Data.State.Open))
		// skip spaces with invalid open status
		if err != nil {
			continue
		}
		status[space.Data.Name] = open
	}

	return status, nil
}

func main() {
	// perform an HTTP GET request
	api, err := get(API)
	if err != nil {
		panic(err)
	}

	// parse the HTTP GET request
	status, err := parseAPI(api)
	if err != nil {
		panic(err)
	}

	for _, space := range status.orderKeys() {
		fmt.Printf("%-49s %t\n", space, status[space])
	}
}
