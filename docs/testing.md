テスト
====

meltでのコード例: [pkg/domain/service/user/service_test.go](https://github.com/karamaru-alpha/melt/blob/main/pkg/domain/service/user/service_test.go) 


## 方針

mapを用いたテーブルドリブンテストを用いる
- [Golangのテストはテーブルドリブンテストで！](https://qiita.com/takehanKosuke/items/cbfc88c4b7956adede79)
- [Goのテーブル駆動テストをわかりやすく書きたい](https://zenn.dev/kimuson13/articles/go_table_driven_test)

上位レイヤから注入された処理は適宜モックする
- [go/mock](https://github.com/golang/mock/tree/v1.5.0)
- [gomockを完全に理解する](https://zenn.dev/sanpo_shiho/articles/01da627ead98f5)
