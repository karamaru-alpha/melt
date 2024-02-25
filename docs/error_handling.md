エラーハンドリング
====

[xerrors](https://pkg.go.dev/golang.org/x/xerrors) を拡張したmerrors(melt-errorsの略)を用いる


### New/Newf

独自エラー(melt-error)を生成するメソッド

```go
package main

import (
    "github.com/karamaru-alpha/melt/pkg/merrors"
)

func F1() error {
    return merrors.Newf(cerrors.InvalidArgument, "Character is duplicated. ID: %q", character.ID)
}
```

### Stack
エラーログに経路を記録するためにスタックフレームを積むメソッド
エラーを伝搬する時のみ使用する

```go
package main

import (
    "github.com/karamaru-alpha/melt/pkg/merrors"
)

func F2() error {
    if err := F1(); err != nil {
        return merrors.Stack(err)
    }
	
    return nil
}
```

### Wrap/Wrapf
エラーを独自エラーにラップするメソッド
外部ライブラリからのエラーを変換するために使用する

```go
package main

import (
    "github.com/google/uuid"
	
    "github.com/karamaru-alpha/melt/pkg/merrors"
)

func F3() (uuid.UUID, error) {
    id, err := uuid.NewRandom()
    if err != nil {
        return uuid.Nil, merrors.Wrap(err, merrors.Internal, "fail to generate uuid")
    }
	
    return id, nil
}
```

### 備考

xerrorsはgo1.13でメンテが終了しているが、フレームを取るだけなので引き続き使用している。
