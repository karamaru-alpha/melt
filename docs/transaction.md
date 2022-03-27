トランザクション制御
==

- domain層にトランザクション制御/トランザクションオブジェクトの抽象が定義されている
    ```go
    package database
    
    import (
        "context"
    )
    
    // TxManager トランザクションマネージャー
    type TxManager interface {
        Transaction(ctx context.Context, f func(ctx context.Context, tx Tx) error) error
    }
    
    // Tx トランザクション
    type Tx interface {
        Commit() error
        Rollback() error
    }
    ```

- 処理内で一貫性を持たせたい場合、usecase層でトランザクションを呼び出す
    ```go
    package main
    
    import (
        "context"
        
        "github.com/karamaru-alpha/melt/pkg/domain/database"
        "github.com/karamaru-alpha/melt/pkg/domain/repository"
        "github.com/karamaru-alpha/melt/pkg/domain/entity"
        "github.com/karamaru-alpha/melt/pkg/merrors"
    )
    
    type interactor struct {
        userRepository repository.UserRepository
        txManager database.TxManager
    }
    
    func (i *interactor) Hoge(ctx context.Context) error {
        // トランザクション開始するよ！
        if err := i.txManager.Transaction(ctx, func(ctx context.Context, tx database.Tx) error {
            user := &entity.User{ID: "hoge", Name: "fuga"}
            // txを使ってトランザクション内でクエリ実行するよ！
            if err := userRepository.Insert(ctx, tx, user); err != nil {
                return merrors.Stack(err)
            }
  
            return nil
        }; err != nil {
            return merrors.Stack(err) 
        }
        
        return nil
    }
    ```

## 備考

- トランザクションの取り忘れを防ぐために、基本的にユーザーデータ関連のクエリにはtxを渡すように設計している。

- 他にもcontextにtxオブジェクトを詰めるやり方などがある
    - [Goとクリーンアーキテクチャとトランザクションと](https://qiita.com/miya-masa/items/316256924a1f0d7374bb)
    - [【Go】厳密なClean Architectureとその妥協案](https://qiita.com/ariku/items/659a11767912c2ec266d)
    
- トランザクション≠ロック
  - [「トランザクション張っておけば大丈夫」と思ってませんか？ バグの温床になる、よくある実装パターン](https://zenn.dev/tockn/articles/4268398c8ec9a9)
  - [データベースのロック(排他制御)とは？ロックの種類や仕組みを解説](https://www.youtube.com/watch?v=oV3VhDu9QHc&t=810s)

- 更新系のクエリが1つでもトランザクション制御は必要。

|     | プロセス1                                                           | プロセス2                                                        |
|-----|-----------------------------------------------------------------|--------------------------------------------------------------|
| 1   | `SELECT * FROM user WHERE name = 'hoge';` -> null               |                                                              |
| 2   |                                                                 | `SELECT * FROM user WHERE name = 'hoge';` -> null            |
| 3   | `INSERT INTO user (id, name) VALUES (null, 'hoge')` -> succeed! |                                  |hash}/docs/a/bb.md`|これは..ない|
| 4   |                | `INSERT INTO user (id, name) VALUES (null, 'hoge')` -> failed! |
