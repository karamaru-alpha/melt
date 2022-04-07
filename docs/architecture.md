アーキテクチャ
==

レイヤードアーキテクチャがベース。**DDDではない**。

DIには [wire](https://github.com/google/wire) か [dig](https://github.com/uber-go/dig) を用いる



各層の依存は以下のようになっている。

- [Handler](#Handler)
- [UseCase](#UseCase)
- [Domain](#Domain)
- [Infra](#Infra)
- [例外](#例外)


## Handler
inputを整形してUseCaseを呼び出し、結果をoutputとして整形して返す。
HTTPサーバーであればrequestとresponseの整形。
UseCase,Domainレイヤーに依存する。

#### ControllerとPresenter
クリーンアーキテクチャでいうところのこの2つはHandlerレイヤーとして統合する。
req/resが主体なので1つに統合して問題ないと判断した。
streamingのようなPresenterをUseCase側から呼び出したくなったときに考える。

## UseCase
アプリケーション固有のビジネスルール。
複数のDomainServiceを呼び出しビジネスルールを作り出す。
仕様変更に大きく影響する。
Domainレイヤーに依存する。

#### InputPort,Interactor,OutputPort
クリーンアーキテクチャでいうところのこの3つはUseCaseレイヤーとして統合する。
HandlerのControllerとPresenterと同じ理由。

## Domain
[例外](#例外)を除き最下位レイヤー。どのレイヤーにも依存しない独立したレイヤー。

#### Entity
infraレイヤーに影響を及ぼす構造体群。構造体に関係する関数も存在する。

#### Repository
EntityオブジェクトのCRUD操作をするInterface。infraレイヤーがこれを実装する。
infraレイヤーで使う技術に依存してはならない。(sql.Dbやspanner.client等)

#### Service
UseCaseから受け取った値に応じてRepositoryを呼び出す。 1まとまりの処理はまとめてここに記述し各usecaseから呼び出す。 とても短く単純で、再利用性がない場合はusecaseでロジックを簡潔させても良い
。
**DDDのdomainServiceではない**

## Infra
特定の技術に特化したレイヤー。
ライブラリ等の初期化処理を行ったり、DomainRepositoryを実装したりする。
DI時のみ呼び出され、他のレイヤーからは依存しない。Domainレイヤーに依存する。

## 例外
どのレイヤーからも呼び出す必要があるもの。-> 最下位のDomainレイヤーからも呼び出す必要があるもの。

### 独自Errors
Infraレイヤー(外部ライブラリ)で発生したエラーをそのまま別レイヤーに伝播させないように必ずラップする。
技術に依存したエラーコードにしない(spannerFailureとか)。技術特有のエラーが発生した場合はその場でログとして出す。

### 独自Loggers
logger。ほぼ一番下位レイヤーなのでどこからでも呼び出せる。


### 独自Util
どのcampus-serverのパッケージにも依存しない独立した存在。一番下位レイヤーなのでどこからでも呼び出せる。
