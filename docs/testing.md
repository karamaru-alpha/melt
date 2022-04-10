テスト
====

**DI**によって注入された処理を**モック**して、**テーブル駆動**でテストする。

### DIとは

- DIとは、依存性の注入(dependency injection)の略
- 呼び出す処理を外側から「詰め込む/注入する」ことで、その処理の差し替えを可能にするデザインパターン

cf. [DIの仕組みをGoで実装して理解する](https://qiita.com/yoshinori_hisakawa/items/a944115eb77ed9247794)

### モックとは

- テストなどの目的で、任意の処理を別の処理に差し替える方法 （本番:mysqlに問い合わせ, テスト:オンメモリのダミーデータなど）

#### gomock
- [gomock](https://github.com/golang/mock) はinteraface定義からモックを作成するツール
- interafaceのあるfileに以下の記述をし、`go generate`すればモックができる
```go
//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
```

cf. [gomockを完全に理解する](https://zenn.dev/sanpo_shiho/articles/01da627ead98f5)


### テーブル駆動テストとは

- mapやslice、配列などでテストケースを列挙した後、それらをループさせてテストをいっぺんに行う方法
- 見やすさ的に、主にmapを用いたテーブルドリブンテストを採用する
- テストケースが1つの場合などは無理にテーブル駆動にする必要はない

cf. [Golangのテストはテーブルドリブンテストで！](https://qiita.com/takehanKosuke/items/cbfc88c4b7956adede79)
 / [Goのテーブル駆動テストをわかりやすく書きたい](https://zenn.dev/kimuson13/articles/go_table_driven_test)

### テスト項目について

- なるべく閾値をテストケースにする（e.g. 文字列の長さが10文字以下である場合、正常系は10文字、異常系は11文字で）
- 外部サービスからのエラーなど、伝搬するエラーは検証しない。(merrors.New(f)のみ検証する)

### 参考コード

```go
package user

import (
	"context"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/karamaru-alpha/melt/pkg/domain/database/mock_database"
	"github.com/karamaru-alpha/melt/pkg/domain/entity"
	"github.com/karamaru-alpha/melt/pkg/domain/repository/mock_repository"
	"github.com/karamaru-alpha/melt/pkg/merrors"
	testutil "github.com/karamaru-alpha/melt/pkg/test/util"
	"github.com/karamaru-alpha/melt/pkg/util/mock_util"
)

// モックする際に注入するハリボテのオブジェクト(mock)達
type mocks struct {
	userRepository *mock_repository.MockUserRepository
	ulidGenerator  *mock_util.MockULIDGenerator
	tx             *mock_database.MockTx
}

func newWithMocks(t *testing.T) (context.Context, *service, *mocks) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	// モックを生成する
	userRepository := mock_repository.NewMockUserRepository(ctrl)
	ulidGenerator := mock_util.NewMockULIDGenerator(ctrl)
	tx := mock_database.NewMockTx(ctrl)
	return ctx,
		New(userRepository, ulidGenerator).(*service), // テスト対象にモックを詰め込む
		&mocks{userRepository, ulidGenerator, tx} // モックオブジェクト一覧。後でモックオブジェクトの挙動を指定するときに用いる
}

func Test_Service(t *testing.T) {
	// テストケースをテーブル(map)にまとめたテスト
	for name, tt := range map[string]struct {
		name string
		err  error
		mock func(ctx context.Context, m *mocks)
	}{
		"正常系": {
			name: strings.Repeat("a", 10),
			mock: func(ctx context.Context, m *mocks) {
				name := strings.Repeat("a", 10)
				// EXPECT()でモックオブジェクトの動作を指定できる
				m.userRepository.EXPECT().SelectByName(ctx, m.tx, name).Return(nil, nil).Times(1)
				m.ulidGenerator.EXPECT().Generate().Return("id", nil).Times(1)
				m.userRepository.EXPECT().Insert(ctx, m.tx, &entity.User{
					ID:   "id",
					Name: name,
				}).Return(nil).Times(1)

			},
		},
		"異常系: 既に存在するname": {
			name: strings.Repeat("a", 10),
			mock: func(ctx context.Context, m *mocks) {
				name := strings.Repeat("a", 10)
				m.userRepository.EXPECT().SelectByName(ctx, m.tx, name).Return([]*entity.User{{}}, nil).Times(1)
			},
			err: merrors.Newf(merrors.InvalidArgument, "user is already exist. name: %s", strings.Repeat("a", 10)),
		},
		"異常系: nameが長すぎる": {
			name: strings.Repeat("a", 11),
			err:  merrors.Newf(merrors.InvalidArgument, "user name len should be %d~%d", 2, 10),
		},
		"異常系: nameが短すぎる": {
			name: strings.Repeat("a", 1),
			err:  merrors.Newf(merrors.InvalidArgument, "user name len should be %d~%d", 2, 10),
		},
	} {
		t.Run(name, func(t *testing.T) {
			ctx, s, m := newWithMocks(t)
			
			// s:テスト対象の実体, m:モックオブジェクト一覧
			
			// 全てのテストケースで共通するモック処理(EXPECT)があればここに記述する
			
			if tt.mock != nil {
				tt.mock(ctx, m)
			}

			err := s.Create(ctx, m.tx, tt.name) // テスト対象の呼び出し
			testutil.EqualMeltError(t, tt.err, err) // アサーション
		})
	}
}

```
