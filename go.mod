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
	github.com/alecthomas/chroma v0.10.0 // indirect
	github.com/aymanbagabas/go-osc52 v1.0.3 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/charmbracelet/glamour v0.6.0 // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/microcosm-cc/bluemonday v1.0.21 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.13.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/yuin/goldmark v1.5.2 // indirect
	github.com/yuin/goldmark-emoji v1.0.1 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
)

require (
	internal/env v0.0.0 // indirect
	omglol/address v0.0.0 // indirect
	omglol/purl v0.0.0 // indirect
	omglol/statuslog v0.0.0 // indirect
)
