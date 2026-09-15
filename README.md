# goblog-plugin-hello

The example [goblog](https://github.com/goblogplatform/goblog) dynamic plugin. It appends a configurable greeting to the footer of every page — the smallest thing that proves a plugin is loaded.

## Install

```bash
curl -o plugins/dynamic/hello.go \
  https://raw.githubusercontent.com/goblogplatform/goblog-plugin-hello/v1.0.0/plugin.go
ENABLE_DYNAMIC_PLUGINS=true ./goblog
```

Every page now ends with a greeting. Requires goblog 0.2.6 or newer. With Docker, bind-mount `plugins/dynamic/` into the image as described in goblog's README.

## Settings

Under **Admin → Settings → Hello (example)**:

| Setting | Default | Meaning |
|---|---|---|
| `enabled` | `true` | Set to `false` to hide the greeting |
| `message` | `Hello from a dynamic plugin` | The text shown at the bottom of every page |

## Use it as a template

Copy this repository, rename the plugin (`Name()` must be unique — it keys the plugin's settings), and follow the [plugin directory contract](https://github.com/goblogplatform/plugins/blob/main/docs/CONTRACT.md) to publish it.

## License

Apache-2.0.
