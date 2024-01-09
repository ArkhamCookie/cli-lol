package statuslog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

/** Types */
type StatuslogList struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response struct {
		Message  string `json:"message"`
		Statuses []struct {
			ID           string `json:"id"`
			Address      string `json:"address"`
			Created      string `json:"created"`
			RelativeTime string `json:"relative_time"`
			Emoji        string `json:"emoji"`
			Content      string `json:"content"`
			ExternalURL  string `json:"external_url"`
		} `json:"statuses"`
	} `json:"response"`
}

func Retrieve(address, status string) (*StatuslogData, error) {
	if address == "" {
		return nil, errors.New("no address given")
	}
	if status == "" {
		return nil, errors.New("no status given")
	}

	return nil, nil
}

func List(address string) (*StatuslogList, error) {
	if address == "" {
		return nil, errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/statuses/", address)
	resp, err := http.Get(target)
	if err != nil {
		log.Fatalln("could not create GET request:", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var result StatuslogList
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalln("could not unmarshal JSON")
	}

	if result.Request.StatusCode != 200 {
		errorMsg := fmt.Sprintf("status code: %d", result.Request.StatusCode)
		return nil, errors.New(errorMsg)
	}
	return &result, nil
}
