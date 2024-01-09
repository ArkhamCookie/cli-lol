package statuslog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

type StatuslogRetrieve struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response struct {
		Message string `json:"message"`
		Status  struct {
			ID           string `json:"id"`
			Address      string `json:"address"`
			Created      string `json:"created"`
			RelativeTime string `json:"relative_time"`
			Emoji        string `json:"emoji"`
			Content      string `json:"content"`
			ExternalURL  string `json:"external_url"`
		} `json:"status"`
	} `json:"response"`
}

type StatuslogLatest struct {
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
			ExternalURL  any    `json:"external_url"`
		} `json:"statuses"`
	} `json:"response"`
}

func Retrieve(address, status string) (*StatuslogRetrieve, error) {
	if address == "" {
		return nil, errors.New("no address given")
	}
	if status == "" {
		return nil, errors.New("no status given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/statuses/%s", address, status)
	resp, err := http.Get(target)
	if err != nil {
		errorMsg := fmt.Sprintf("could not create GET request: %s", err)
		return nil, errors.New(errorMsg)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result StatuslogRetrieve
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("could not unmarshal JSON")
	}

	if result.Request.StatusCode != 200 {
		errorMsg := fmt.Sprintf("status code: %d", result.Request.StatusCode)
		return nil, errors.New(errorMsg)
	}

	return &result, nil
}

func List(address string) (*StatuslogList, error) {
	if address == "" {
		return nil, errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/statuses/", address)
	resp, err := http.Get(target)
	if err != nil {
		errorMsg := fmt.Sprintf("could not create GET request: %s", err)
		return nil, errors.New(errorMsg)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result StatuslogList
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("could not unmarshal JSON")
	}

	if result.Request.StatusCode != 200 {
		errorMsg := fmt.Sprintf("status code: %d", result.Request.StatusCode)
		return nil, errors.New(errorMsg)
	}
	return &result, nil
}

func Latest() (*StatuslogLatest, error) {
	resp, err := http.Get("https://api.omg.lol/statuslog/latest")
	if err != nil {
		errorMsg := fmt.Sprintf("could not create GET request: %s", err)
		return nil, errors.New(errorMsg)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result StatuslogLatest
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("could not unmarshal JSON")
	}

	if result.Request.StatusCode != 200 {
		errorMsg := fmt.Sprintf("status code: %d", result.Request.StatusCode)
		return nil, errors.New(errorMsg)
	}
	return &result, nil
}
