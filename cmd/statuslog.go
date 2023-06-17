package cmd

import (
	"internal/http"
)

func Statuslog() {
	http.Get("https://api.omg.lol/address/arkhamcookie/statuses/648d5e756cc10")
}
