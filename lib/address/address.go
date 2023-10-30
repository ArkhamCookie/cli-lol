package address

import (
	"errors"
	"fmt"
	"internal/http"
	"log"
)

var (
	err error
	response string
)

type addressData struct {
	Request  *myRequest  `json:"objectValue"`
	Response *myResponse `json:"objectValue"`
}

type myRequest struct {
	StatusCode    int  `json:"intValue"`
	SuccessStatus bool `json:"boolValue"`
}

type myResponse struct {
	Message       string `json:"stringValue"`
	TargetAddress string `json:"stringValue"`
	IsAvailable   bool   `json:"boolValue"`
	Availability  string `json:"stringValue"`
}

// Check the availability of a given address.
func Available(address string) (string, error) {
	// If no address was given, return an error (with a message).
	if address == "" {
		return "", errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/availability", address)
	response, err := http.Get(target)
	if err != nil {
		log.Fatal(err)
	}

	// ? Handle differently if encoding is required?
	// TODO: Return bool instead of response
	return string(response), nil
}

// Get the expiration date for an address
func Expiration(address string) (string, error) {
	// If no address was given, return an error (with a message).
	if address == "" {
		return "", errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/expiration", address)
	response, err := http.Get(target)
	if err != nil {
		log.Fatal(err)
	}

	return string(response), nil
}
