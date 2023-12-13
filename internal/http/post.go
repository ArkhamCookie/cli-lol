package http

import (
	"errors"
	"io"
	"log"
	"net/http"
)

func Post(url string) (string, error) {
	// If no url was given, return an error (with a message).
	if url == "" {
		return "", errors.New("no address given")
	}

	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Fatalln("could not create GET request:", err)
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body), nil
}
