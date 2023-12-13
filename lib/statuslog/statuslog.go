package statuslog

import (
	"errors"
	"fmt"
	"internal/http"
)

/** Types */
type Status struct {
	Id int `json:"id"`
	Address string `json:"address"`
	Created int `json:"created"`
	Emoji string `json:"emoji"`
	Content string `json:"content"`
}

type StatuslogData struct {
	http.RequestData

	Message string
	Statuses []Status
}

func Retrieve(address, status string) (*Status, error) {
	if address == "" {
		return nil, errors.New("no address given")
	}
	if status == "" {
		return nil, errors.New("no status given")
	}

	return nil, nil
}

func List(address string) (*Status, error) {
	if address == "" {
		return nil, errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/statuses/", address)
	fmt.Println(http.Get(target))

	return nil, nil
}
