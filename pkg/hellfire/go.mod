module github.com/kernelpanic77/hellfire/pkg/hellfire

go 1.22.5

require (
	github.com/kernelpanic77/hellfire/internal v0.0.0-00010101000000-000000000000
	github.com/panjf2000/ants v1.3.0
	github.com/schollz/progressbar/v3 v3.14.4
)

require (
	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/stretchr/testify v1.8.2 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/term v0.22.0 // indirect
)

replace github.com/kernelpanic77/hellfire/internal => ../../internal
