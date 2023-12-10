package http

import (
	"errors"
	"io"
	"log"
	"net/http"
)

func Get(url string) (string, error) {
	// If no url was given, return an error (with a message).
	if url == "" {
		return "", errors.New("no address given")
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln("could not create GET request:", err)
	}

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		// TODO: respond differently depending on error
		return resp.Status, errors.New("request failed")
	}

	return string(body), nil
}
