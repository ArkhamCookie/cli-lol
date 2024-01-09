module github.com/ArkhamCookie/omglol

go 1.20

require (
	omglol/cmd v0.0.0
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

require (
	internal/http v0.0.0 // indirect
	omglol/address v0.0.0 //indirect
	omglol/statuslog v0.0.0 //indirect
)

replace (
	internal/http => ./internal/http

	// libs
	omglol/address => ./lib/address
	omglol/profile => ./lib/profile
	omglol/statuslog => ./lib/statuslog

	// commands
	omglol/cmd => ./cmd
)
