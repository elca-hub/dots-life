# dotslife

ドットで彩る、私の日々。

## アプリについて

dotslifeは主に多くの作業を行なっている人向けのアプリです。
例えば「アルバイト」「プログラミング」「課題」「研究」....
こういった様々なタスクを我々は普段行なっています。
そのタスクごとに色を与え、今日を鮮やかにする、そんなアプリを目指しています。

## 起動方法(docker)

### 初期設定

`back` ディレクトリに移動し、 `make` コマンドを使用してマイグレーションを行なってください。

```bash
docker compose up -d db # DBの起動
cd back
make migrate-up
docker compose down
```

### 実行

```bash
docker compose up
```

### 各種ヘルスチェック

```bash
curl http://localhost:5050/api/v1/ping # back
```

## 各種ドキュメント

本アプリではコーディングにも力を入れています。
下のリンク、または対応するディレクトリのREADME.mdから、コードに関する解説が見れます。

- [back](./back/README.md)
