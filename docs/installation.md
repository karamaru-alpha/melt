環境構築
====


TODO: docker

## Homebrewのインストール

```bash
$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

## Golangのインストール

### [anyenv](https://github.com/anyenv/anyenv)

色んな言語の開発環境を簡単にセットアップできるツール

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

### [goenv](https://github.com/syndbg/goenv)

Goのバージョン管理をしてくれるツール

- install
```bash
$ anyenv install goenv
$ goenv install -l # installできるバージョン確認
$ goenv install ${go-version} # install
$ goenv versions # 使えるようになったバージョン確認
$ goenv local 1.18.0 # 特定ディレクトリで使用するバージョン指定
$ go version
```
- path通す(.zshrcとかに記述)
```.zshrc
export GOENV_ROOT="$HOME/.goenv"
export PATH="$GOENV_ROOT/bin:$PATH"
eval "$(goenv init -)"
```
