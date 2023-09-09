module ArkhamCookie/omglol

go 1.20

require (
	internal/http v0.0.0 // indirect
	// ArkhamCookie/omglol v0.0.0
	// omglol/profile v0.0.0
	omglol/address v0.0.0
	// omglol/cmd v0.0.0
)

replace (
	ArkhamCookie/omglol => ../omg.lol
	internal/http => ./internal/http
	omglol/address => ./lib/address
	omglol/cmd => ./cmd
	omglol/profile => ./lib/profile
)

// require ArkhamCookie/omg.lol v1.0.0
