package address

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AddressExpirationData struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response struct {
		Message string `json:"message"`
		Expired bool   `json:"expired"`
	} `json:"response"`
}

type AddressAvaiableData struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response struct {
		Message      string `json:"message"`
		Address      string `json:"address"`
		Available    bool   `json:"available"`
		Availability string `json:"availability"`
	} `json:"response"`
}

// Check the availability of a given address.
func Available(address string) (bool, error) {
	// If no address was given, return an error (with a message).
	if address == "" {
		return false, errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/availability", address)
	resp, err := http.Get(target)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result AddressAvaiableData
	if err := json.Unmarshal(body, &result); err != nil {
		return false, errors.New("could not unmarshal JSON")
	}

	if result.Request.StatusCode != 200 {
		errorMsg := fmt.Sprintf("status code: %d", result.Request.StatusCode)
		return false, errors.New(errorMsg)
	}

	return result.Response.Available, nil
}

// Get the expiration date for an address
func Expiration(address string) (*AddressExpirationData, error) {
	// If no address was given, return an error (with a message).
	if address == "" {
		return nil, errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/expiration", address)
	resp, err := http.Get(target)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AddressExpirationData
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("could not unmarshal JSON")
	}

	if result.Request.StatusCode != 200 {
		errorMsg := fmt.Sprintf("status code: %d", result.Request.StatusCode)
		return nil, errors.New(errorMsg)
	}

	return &result, nil
}
