module github.com/ArkhamCookie/cli-lol

go 1.20

replace (
	internal/env => ./internal/env
	omglol/address => ./lib/address
	omglol/cmd => ./cmd
	omglol/profile => ./lib/profile
	omglol/purl => ./lib/purl
	omglol/statuslog => ./lib/statuslog
)

require omglol/cmd v0.0.0

require (
	internal/env v0.0.0 // indirect
	omglol/address v0.0.0 // indirect
	omglol/purl v0.0.0 // indirect
	omglol/statuslog v0.0.0 // indirect
)
