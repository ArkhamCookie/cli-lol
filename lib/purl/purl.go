package purl

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"internal/env"
)

type purlRetrieveData struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response struct {
		Message string `json:"message"`
		Purl    struct {
			Name    string `json:"name"`
			URL     string `json:"url"`
			Counter any    `json:"counter"`
			Listed  int    `json:"listed"`
		} `json:"purl"`
	} `json:"response"`
}

type purlListData struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response struct {
		Message string `json:"message"`
		Purls   []struct {
			Name    string `json:"name"`
			URL     string `json:"url"`
			Counter int    `json:"counter"`
		} `json:"purls"`
	} `json:"response"`
}

func Retrieve(address, purl string) (*purlRetrieveData, error) {
	if address == "" {
		return nil, errors.New("no address given")
	}
	if purl == "" {
		return nil, errors.New("no purl given")
	}
	apiKey, err := env.GetEnvValue("API_KEY", "")
	if err != nil {
		return nil, err
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/purl/%s", address, purl)
	var bearer = "Bearer " + apiKey

	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Error on response. \nError: %s", err)
		return nil, errors.New(errMsg)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result purlRetrieveData
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("could not unmarshal JSON")
	}

	return &result, nil
}

func List(address string) (*purlListData, error) {
	if address == "" {
		return nil, errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/purls", address)
	resp, err := http.Get(target)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result purlListData
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("could not unmarshal JSON")
	}

	if result.Request.StatusCode != 200 {
		errorMsg := fmt.Sprintf("status code: %d", result.Request.StatusCode)
		return nil, errors.New(errorMsg)
	}

	return &result, nil
}
