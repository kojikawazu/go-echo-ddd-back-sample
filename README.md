# Go/EchoのDDD設計用サンプルコード

## Summary

Go/Echoを使用し、DDD設計を行う。

- 並列実行(ゴルーチン)
- Supabase接続(DB操作)
- 認証、認可

## Directory

以下ディレクトリ構成とする。

```bash
/project
  ├── cmd/                # エントリーポイント
  ├── config/             # 設定ファイル
  ├── internal/
  │   ├── domain/         # ドメイン層（エンティティ、リポジトリ、VO）
  │   ├── usecase/        # ユースケース層（アプリケーションサービス）
  │   ├── infrastructure/ # インフラ層（DB, API クライアント, リポジトリ実装）
  │   ├── interfaces/     # インターフェース層（HTTPハンドラ, gRPC, CLI）
  │   ├── middleware/     # Echoのミドルウェア
  │   ├── pkg/            # 汎用ユーティリティ
  │   ├── router/         # ルーティング
  │   ├── test/           # ユニット・統合テスト
  ├── Dockerfile
  ├── go.mod
  ├── go.sum
  └── README.md
```
