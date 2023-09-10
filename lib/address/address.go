package address

import (
	"errors"
	"fmt"
	"internal/http"
)

// Address checks the availability of a given address.
func Address(address string) error {
	// If no address was given, return an error (with a message).
	if address == "" {
		return errors.New("no address given")
	}

	target := fmt.Sprintf("https://api.omg.lol/address/%s/availability", address)
	http.Get(target)

	return nil
}
