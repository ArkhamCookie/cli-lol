package http

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// TODO: return http status codes
func Get(url string) (error) {
	// If no url was given, return an error (with a message).
	if url == "" {
		return errors.New("no address given")
	}

	resp, err := http.Get(url)

	// If an error was returned,
	// print it to console & exit program.
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	// If an error was returned,
	// print it to console & exit program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	return nil
}
