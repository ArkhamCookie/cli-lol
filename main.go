package main

import (
	// "internal/http"
	// "omglol/cmd"
	"fmt"
	"omglol/profile"
)

func main() {
	// http.Get("https://api.omg.lol/address/arkhamcookie/statuses/648d5e756cc10")
	// cmd.Statuslog()
	temp := profile.Profile("email", "omg-lol@mail.arkhamcookie.com")
	fmt.Printf("%s\n", temp)
}
