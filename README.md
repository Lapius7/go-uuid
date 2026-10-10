# go-uuid

アクセス時にUUIDを生成して返すシンプルなHTTPサーバー。RFC 9562のUUIDバージョンをURLパスで指定できる。

## すぐ試す

公開中のサーバーに `curl` するだけで UUID が返る（インストール不要）。

```bash
curl https://sandbox.lapius7.com/go-uuid/v4
```

## インストール

### npm（推奨）

```bash
pnpm add -g @lapius/go-uuid
```

Linux / macOS（x64・arm64）/ Windows（x64）のビルド済みバイナリが入る（Go 不要、Node.js 18 以降）。更新も同じコマンドで行う。

### ソースから

```bash
git clone https://github.com/Lapius7/go-uuid.git && cd go-uuid
go run main.go
```

## 使い方

`go-uuid` を実行するとポート 7100 で HTTP サーバーが起動する。

```bash
go-uuid
```

別のターミナルから、パスで UUID のバージョンを指定して取得する。

```bash
curl http://localhost:7100/       # v7
curl http://localhost:7100/v4     # v4（ランダム）
curl "http://localhost:7100/v5?name=lapius7.com"
```

## エンドポイント

| パス | UUIDバージョン | 説明 |
|---|---|---|
| `/` | v7 | デフォルト |
| `/v1` | v1 | タイムスタンプ + MACアドレスベース |
| `/v2` | v2 | DCE Security。`?domain=person\|group\|org`（デフォルト: `person`）、`?id=<uint32>`（デフォルト: 実行ユーザーのUID）を指定可能 |
| `/v3` | v3 | 名前ベース（MD5）。`?namespace=<UUID>`（デフォルト: DNS namespace）、`?name=<string>`（デフォルト: `example.com`）を指定可能 |
| `/v4` | v4 | ランダム |
| `/v5` | v5 | 名前ベース（SHA-1）。パラメータは`/v3`と同様 |
| `/v6` | v6 | 時系列ソート可能なタイムスタンプベース |
| `/v7` | v7 | Unixミリ秒タイムスタンプベース |

## 依存関係

- [github.com/google/uuid](https://github.com/google/uuid) v1.6.0
