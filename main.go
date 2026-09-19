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

// setting mirrors goblog's settings JSON shape.
type setting struct {
	Key         string `json:"key"`
	Type        string `json:"type"`
	Default     string `json:"default"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

func outputJSON(v any) int32 {
	if err := pdk.OutputJSON(v); err != nil {
		pdk.SetErrorString("encode output: " + err.Error())
		return 1
	}
	return 0
}

//go:wasmexport identity
func identity() int32 {
	return outputJSON(map[string]string{"name": "hello", "display_name": "Hello", "version": "2.0.0"})
}

//go:wasmexport settings
func settings() int32 {
	return outputJSON([]setting{
		{Key: "enabled", Type: "text", Default: "true", Label: "Enabled", Description: "Set to 'true' to show the greeting"},
		{Key: "message", Type: "text", Default: "Hello from a WebAssembly plugin", Label: "Message", Description: "Text shown at the bottom of every page"},
	})
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
