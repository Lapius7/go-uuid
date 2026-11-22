<div align="center">

<p align="center"><b>English</b> | <a href="README.ja.md">日本語</a></p>

<h1 align="center">UUID Server</h1>

<h3 align="center">Generate a UUID on every HTTP request.</h3>

<p align="center">A small HTTP server that returns a new UUID each time you access it. Pick any RFC 9562 UUID version with the URL path.</p>

<p align="center">
  <a href="https://github.com/Lapius7/go-uuid/releases"><img alt="Release" src="https://img.shields.io/github/v/release/Lapius7/go-uuid?label=release&style=flat-square&color=e8833a&logo=github"></a>
  <a href="https://www.npmjs.com/package/@lapius/go-uuid"><img alt="npm" src="https://img.shields.io/npm/v/@lapius/go-uuid?style=flat-square&color=cb3837&logo=npm"></a>
  <a href="go.mod"><img alt="Go" src="https://img.shields.io/github/go-mod/go-version/Lapius7/go-uuid?style=flat-square&color=00add8&logo=go&logoColor=white"></a>
  <a href="LICENSE"><img alt="License" src="https://img.shields.io/github/license/Lapius7/go-uuid?style=flat-square&color=97ca00"></a>
  <a href="https://github.com/Lapius7/go-uuid/stargazers"><img alt="Stars" src="https://img.shields.io/github/stars/Lapius7/go-uuid?style=flat-square&color=e3b341&logo=github"></a>
  <a href="https://github.com/Lapius7/go-uuid/commits/main"><img alt="Last commit" src="https://img.shields.io/github/last-commit/Lapius7/go-uuid?style=flat-square&color=3b82f6"></a>
</p>

<p align="center">
  <a href="#try-it">Try it</a> ·
  <a href="#install">Install</a> ·
  <a href="#usage">Usage</a> ·
  <a href="#endpoints">Endpoints</a>
</p>

</div>

---

## Try it

Call the public server:

```bash
curl https://sandbox.lapius7.com/go-uuid/v4
```

## Install

### pnpm (recommended)

> [!IMPORTANT]
> This package assumes **pnpm**. Installing with both npm and pnpm leaves two copies of the same command, so use only `pnpm add -g`.

```bash
pnpm add -g @lapius/go-uuid
```

The package ships prebuilt binaries for Linux (x64, arm64), macOS (x64, arm64), and Windows (x64). You don't need Go. Node.js 18 or later is required. Run the same command to update.

### From source

```bash
git clone https://github.com/Lapius7/go-uuid.git && cd go-uuid
go run main.go
```

## Usage

Run `go-uuid` to start an HTTP server on port 7100.

```bash
go-uuid
```

From another terminal, choose the UUID version with the path.

```bash
curl http://localhost:7100/    # v7
curl http://localhost:7100/v4  # v4 (random)
curl "http://localhost:7100/v5?name=lapius7.com"
```

Each request is written to the log (method, path, query, remote address, user agent, and the generated UUID). An invalid parameter returns `400 Bad Request`.

## Endpoints

| Path | Version | Description |
|---|---|---|
| `/` | v7 | Default |
| `/v1` | v1 | Timestamp and MAC address based |
| `/v2` | v2 | DCE Security. Accepts `?domain=person\|group\|org` (default `person`) and `?id=<uint32>` (default: UID of the user running the server) |
| `/v3` | v3 | Name based (MD5). Accepts `?namespace=<UUID>` (default: DNS namespace) and `?name=<string>` (default `example.com`) |
| `/v4` | v4 | Random |
| `/v5` | v5 | Name based (SHA-1). Same parameters as `/v3` |
| `/v6` | v6 | Time-ordered, sortable timestamp |
| `/v7` | v7 | Unix millisecond timestamp |

## Dependencies

- [github.com/google/uuid](https://github.com/google/uuid) v1.6.0

## License

MIT. See [LICENSE](LICENSE).
