package address

import (
	"errors"
	"fmt"
	"internal/http"
)

// Check the availability of a given address.
func Available(address string) error {
	// If no address was given, return an error (with a message).
	if address == "" {
		return errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/availability", address)
	http.Get(target)

	return nil
}
