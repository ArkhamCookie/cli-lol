package address

import (
	"errors"
	"fmt"
	"internal/http"
	"log"
)

// Check the availability of a given address.
func Available(address string) (string, error) {
	// If no address was given, return an error (with a message).
	if address == "" {
		return "", errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/availability", address)
	resp, err := http.Get(target)
	if err != nil {
		log.Fatal(err)
	}

	// ? Handle differently if encoding is required?
	// TODO: Return bool instead of response
	return string(resp), nil
}

// Get the expiration date for an address
func Expiration(address string) (string, error) {
	// If no address was given, return an error (with a message).
	if address == "" {
		return "", errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/expiration", address)
	resp, err := http.Get(target)
	if err != nil {
		log.Fatal(err)
	}

	return string(resp), nil
}
