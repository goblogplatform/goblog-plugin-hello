# goblog-plugin-hello

The reference [goblog](https://github.com/goblogplatform/goblog) WebAssembly plugin. It appends a configurable greeting to the footer of every page — the smallest thing that proves a plugin is loaded.

## Install

From your goblog's **Admin → Plugins → Browse**, search for *Hello* and click **Install** (goblog 0.2.9 or newer). Or download `plugin.wasm` from the [latest release](https://github.com/goblogplatform/goblog-plugin-hello/releases/latest) into `plugins/wasm/` and restart goblog.

## Settings

Under **Admin → Settings → Hello**:

| Setting | Default | Meaning |
|---|---|---|
| `enabled` | `true` | Set to `false` to hide the greeting |
| `message` | `Hello from a WebAssembly plugin` | The text shown at the bottom of every page |

## Build it yourself

```bash
GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -ldflags="-s -w" -o plugin.wasm .
```

Needs Go 1.25 or newer (`go.mod` pins the toolchain, so `GOTOOLCHAIN=auto` fetches it). The plugin talks to no network (`allowed_hosts` is empty), stores nothing, and only implements the `identity`, `settings` and `template_footer` exports.

## Use it as a template

Copy this repository, change the identity (`name` must be unique — it keys the plugin's settings) and follow the [plugin directory contract](https://github.com/goblogplatform/plugins/blob/main/docs/CONTRACT.md). Tag a release as `vX.Y.Z`; the workflow builds and uploads `plugin.wasm` for you.

## License

Apache-2.0.
