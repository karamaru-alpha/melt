環境構築
====

TODO: dockerにのせる

## Go言語のセットアップ

### 1. Homebrewのインストール

[Homebrew](https://brew.sh/index_ja) はmacOSでも動作するパッケージ管理システム

```bash
$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

### 2. anyenvのインストール

[anyenv](https://github.com/anyenv/anyenv) は色んな言語の開発環境を簡単にセットアップできるツール

- install
```bash
$ brew install anyenv
$ anyenv install --init
```
- path通す(.zshrcとかに記述)
```.zshrc
export PATH="$HOME/.anyenv/bin:$PATH"
eval "$(anyenv init -)"
```

### 3. goenvのインストール / バージョン指定

[goenv](https://github.com/syndbg/goenv) はGoのバージョン管理をディレクトリごとにしてくれるツール

今回は`1.18.0`のGoを用いる

- install
```bash
$ anyenv install goenv
$ goenv install -l # installできるバージョン確認
$ goenv install ${go-version} # install
$ goenv versions # 使えるようになったバージョン確認
$ goenv local 1.18.0 # currentディレクトリで使用するバージョン指定
$ go version
  -> go version go1.18 darwin/arm64
```
- path通す(.zshrcとかに記述)
```.zshrc
export GOENV_ROOT="$HOME/.goenv"
export PATH="$GOENV_ROOT/bin:$PATH"
eval "$(goenv init -)"
```


### Goに慣れるために

- [Tour of Go](https://go-tour-jp.appspot.com/welcome/1): チュートリアルサイト
- [Go PlayGround](https://go.dev/play/): オンラインでGoが動かせるサイト

## 周辺パッケージのインストール

```bash
$ go mod tidy # `go.mod`を参考にアプリケーションコードで用いられているパッケージをインストール
$ make local-install # アプリの動作とは別に、localで用いるツール群の導入
```

以下、localで用いるツール群の説明

### google/wire
- [wire](https://github.com/google/wire) はDI(dependency injection)を簡潔にかけるようにするライブラリ
- 詳しくは[テストのドキュメント](./testing.md)を参照

### golang/mock

- [gomock](https://github.com/golang/mock) は、インターフェース定義からモックの生成を行うことができるライブラリ
- DIされた処理の振る舞いを仮決め（モック）することで、主にテスト対象のスコープを絞るために用いられる
- 詳しくは[テストのドキュメント](./testing.md)を参照


### golangci-lint
- [golangci-lint](https://github.com/golangci/golangci-lint) はGoの最もメジャーなlinter
- `golangci-lint run`でlinterが走る
