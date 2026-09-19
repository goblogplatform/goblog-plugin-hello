// Hello is the reference goblog WebAssembly plugin: it appends a
// configurable greeting to the footer of every page.
//
// Build:  GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -ldflags="-s -w" -o plugin.wasm .
// Every export takes JSON on stdin (pdk.Input) and returns JSON or HTML
// (pdk.Output). See goblog's README "WebAssembly plugins" for the contract.
package main

import (
	"encoding/json"
	"html"

	pdk "github.com/extism/go-pdk"
)

// hookInput is the ctx goblog passes to template hooks; only settings are
// needed here.
type hookInput struct {
	Settings map[string]string `json:"settings"`
}

//go:wasmexport identity
func identity() int32 {
	pdk.OutputString(`{"name":"hello","display_name":"Hello","version":"2.0.0"}`)
	return 0
}

//go:wasmexport settings
func settings() int32 {
	pdk.OutputString(`[` +
		`{"key":"enabled","type":"text","default":"true","label":"Enabled","description":"Set to 'true' to show the greeting"},` +
		`{"key":"message","type":"text","default":"Hello from a WebAssembly plugin","label":"Message","description":"Text shown at the bottom of every page"}` +
		`]`)
	return 0
}

//go:wasmexport template_footer
func templateFooter() int32 {
	var in hookInput
	if err := json.Unmarshal(pdk.Input(), &in); err != nil {
		pdk.SetErrorString("template_footer: " + err.Error())
		return 1
	}
	if in.Settings["enabled"] != "true" {
		return 0
	}
	// Always escape setting values: an admin typed them, but the browser
	// will trust whatever this returns.
	pdk.OutputString(`<p class="text-center text-muted">` + html.EscapeString(in.Settings["message"]) + `</p>`)
	return 0
}

func main() {}
