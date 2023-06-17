module ArkhamCookie/omglol

go 1.20

require (
	// ArkhamCookie/omglol v0.0.0
	omglol/profile v0.0.0
	// internal/http v0.0.0
	// omglol/cmd v0.0.0
)

replace (
	// ArkhamCookie/omglol => ../omg.lol
	omglol/profile => ./lib/profile
	// internal/http => ./internal/http
	// omglol/cmd => ./cmd
)

// require ArkhamCookie/omg.lol v1.0.0
