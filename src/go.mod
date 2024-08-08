module src

go 1.22.5

replace mathlib => ../mathlib

require (
	github.com/go-gl/gl v0.0.0-20231021071112-07e5d0ea2e71
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20240506104042-037f3cc74f2a
	github.com/go-gl/mathgl v1.1.0
	mathlib v0.0.0-00010101000000-000000000000
)

require golang.org/x/image v0.0.0-20210607152325-775e3b0c77b9 // indirect
