module ArkhamCookie/omglol

go 1.20

require (
	// internal/http v0.0.0 // indirect
	// ArkhamCookie/omglol v0.0.0
	// omglol/profile v0.0.0
	// omglol/address v0.0.0
	// omglol/cmd v0.0.0
)

replace (
	ArkhamCookie/omglol => ../omg.lol // main

	internal/http => ./internal/http

	// libs
	omglol/address => ./lib/address
	omglol/profile => ./lib/profile

	// commands
	omglol/cmd => ./cmd
)

// require ArkhamCookie/omg.lol v1.0.0
