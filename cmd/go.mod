module omglol/cmd

go 1.20

require (
	internal/http v0.0.0
)

replace (
	internal/http => ../internal/http
)
