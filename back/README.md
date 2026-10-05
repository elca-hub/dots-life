# dotslife/back

バックエンド

## 使用言語、主なモジュール

- 言語: Go(v1.25.0)
  - [echo](https://echo.labstack.com/ja/)
  - [GORM](https://gorm.io/ja_JP/)
  - [air](https://github.com/air-verse/air)(開発環境のみ)
  - [migrate](https://github.com/golang-migrate/migrate)

## 開発について

### migrateとGORM

本アプリでは **先にmigrateを実行し、テーブルの内容をGORMに適用する** スタイルをとっています。

そのため、新しくマイグレートを行うときには以下の手順で行なってください。

```bash
make migrate-create # マイグレーションファイルの作成。infra/migrationsにupとdownのファイルが作成されます
make migrate-up # 新しいマイグレーションの適用
make gorm-generate/build # 初回のみ。ビルドファイルの作成
make gorm-generate/run # SQLのテーブル内容からGORMモデルを作成
```

### `gorm-generate` について

`gorm-generate` は現在のデータベースのテーブル内容をGORMモデルに落とし込む際に使用します。
コードは `cmd/gorm-generate/` 内にあります。

コードの変更をした際は、 `make gorm-generate/build` を行なってから `make gorm-generate/run` を行なってください。

## アーキテクチャについて

本アプリでは私の学習目的のため、クリーンアーキテクチャを参考にしたアーキテクチャとなっています。

### 参考元

- [elca-hub/engineer-portfolio](https://github.com/elca-hub/engineer-portfolio)
- [GSabadini/go-clean-architecture](https://github.com/GSabadini/go-clean-architecture)
