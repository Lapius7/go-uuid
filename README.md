# go-uuid

アクセス時にUUIDを生成して返すシンプルなHTTPサーバー。RFC 9562のUUIDバージョンをURLパスで指定できる。

## 使い方

```bash
go run main.go
```

ポート7100で起動する。

```bash
curl http://localhost:7100/v4
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
