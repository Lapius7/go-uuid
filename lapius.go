package main

import (
	"fmt"
	"os"
	"os/exec"
)

var version = "dev" // npm/build.mjs が -X main.version で埋め込む

// サーバーを起動する前に -h / -v だけ処理する
func init() {
	if len(os.Args) < 2 {
		return
	}
	switch os.Args[1] {
	case "-h", "--help", "help":
		fmt.Print(`go-uuid - アクセス時に UUID を返す HTTP サーバー（ポート 7100）

使い方:
  go-uuid            サーバーを起動
  go-uuid -v         バージョン

エンドポイント:
  /  /v7             v7（既定）
  /v1 〜 /v7         各バージョン（/v4 はランダム）
  /v5?name=<名前>    v5（名前ベース。/v3 も同様）

`)
	case "-v", "--version", "version":
		fmt.Println("go-uuid", version)
	default:
		return
	}
	fmt.Print(lapiusFooter())
	os.Exit(0)
}

// lapiusFooter は --help / --version の最後に出す作者表示と lapacks の案内。
func lapiusFooter() string {
	s := "作者: Lapius (https://github.com/Lapius7)\n"
	if _, err := exec.LookPath("lapacks"); err == nil {
		return s + "@lapius のツール: lapacks で一覧・インストール・更新\n"
	}
	return s + "@lapius のツール: npm i -g @lapius/lapacks で一覧・インストール・更新を管理\n"
}
