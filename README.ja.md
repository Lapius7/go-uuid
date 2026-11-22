<div align="center">

<p align="center"><a href="README.md">English</a> | <b>日本語</b></p>

<h1 align="center">UUID Server</h1>

<h3 align="center">HTTP リクエストのたびに UUID を生成する。</h3>

<p align="center">アクセスのたびに新しい UUID を返すシンプルな HTTP サーバーです。URL のパスで、RFC 9562 の UUID バージョンを選べます。</p>

<p align="center">
  <a href="https://github.com/Lapius7/go-uuid/releases"><img alt="Release" src="https://img.shields.io/github/v/release/Lapius7/go-uuid?label=release&style=flat-square&color=e8833a&logo=github"></a>
  <a href="https://www.npmjs.com/package/@lapius/go-uuid"><img alt="npm" src="https://img.shields.io/npm/v/@lapius/go-uuid?style=flat-square&color=cb3837&logo=npm"></a>
  <a href="go.mod"><img alt="Go" src="https://img.shields.io/github/go-mod/go-version/Lapius7/go-uuid?style=flat-square&color=00add8&logo=go&logoColor=white"></a>
  <a href="LICENSE"><img alt="License" src="https://img.shields.io/github/license/Lapius7/go-uuid?style=flat-square&color=97ca00"></a>
  <a href="https://github.com/Lapius7/go-uuid/stargazers"><img alt="Stars" src="https://img.shields.io/github/stars/Lapius7/go-uuid?style=flat-square&color=e3b341&logo=github"></a>
  <a href="https://github.com/Lapius7/go-uuid/commits/main"><img alt="Last commit" src="https://img.shields.io/github/last-commit/Lapius7/go-uuid?style=flat-square&color=3b82f6"></a>
</p>

<p align="center">
  <a href="#すぐ試す">すぐ試す</a> ·
  <a href="#インストール">インストール</a> ·
  <a href="#使い方">使い方</a> ·
  <a href="#エンドポイント">エンドポイント</a>
</p>

</div>

---

## すぐ試す

公開中のサーバーに接続します。

```bash
curl https://sandbox.lapius7.com/go-uuid/v4
```

## インストール

### pnpm（推奨）

> [!IMPORTANT]
> このパッケージは **pnpm 前提**です。npm と pnpm の両方に入れると同じコマンドが二重になるため、`pnpm add -g` だけを使ってください。

```bash
pnpm add -g @lapius/go-uuid
```

Linux（x64・arm64）、macOS（x64・arm64）、Windows（x64）のビルド済みバイナリが入ります。Go は不要です。Node.js 18 以降が必要です。更新も同じコマンドで行います。

### ソースから

```bash
git clone https://github.com/Lapius7/go-uuid.git && cd go-uuid
go run main.go
```

## 使い方

`go-uuid` を実行すると、ポート 7100 で HTTP サーバーが起動します。

```bash
go-uuid
```

別のターミナルから、パスで UUID のバージョンを指定して取得します。

```bash
curl http://localhost:7100/    # v7
curl http://localhost:7100/v4  # v4（ランダム）
curl "http://localhost:7100/v5?name=lapius7.com"
```

リクエストごとに、メソッド、パス、クエリ、接続元、User-Agent、生成した UUID をログに出します。パラメータが不正なときは `400 Bad Request` を返します。

## エンドポイント

| パス | バージョン | 説明 |
|---|---|---|
| `/` | v7 | デフォルト |
| `/v1` | v1 | タイムスタンプ + MAC アドレスベース |
| `/v2` | v2 | DCE Security。`?domain=person\|group\|org`（デフォルト: `person`）と `?id=<uint32>`（デフォルト: 実行ユーザーの UID）を指定できる |
| `/v3` | v3 | 名前ベース（MD5）。`?namespace=<UUID>`（デフォルト: DNS namespace）と `?name=<string>`（デフォルト: `example.com`）を指定できる |
| `/v4` | v4 | ランダム |
| `/v5` | v5 | 名前ベース（SHA-1）。パラメータは `/v3` と同じ |
| `/v6` | v6 | 時系列ソート可能なタイムスタンプベース |
| `/v7` | v7 | Unix ミリ秒タイムスタンプベース |

## 依存関係

- [github.com/google/uuid](https://github.com/google/uuid) v1.6.0

## ライセンス

MIT。[LICENSE](LICENSE) を参照してください。
