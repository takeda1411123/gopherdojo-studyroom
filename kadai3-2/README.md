# 課題3-1

分割ダウンローダを作ろう

* Rangeアクセスを用いる
* いくつかのゴルーチンでダウンロードしてマージする
* エラー処理を工夫する
  * golang.org/x/sync/errgourpパッケージなどを使ってみる
* キャンセルが発生した場合の実装を行う



## Setup
* init `$ go install`
## Usage
* `go run main.go [-i input url ] [-o output path] [-d division number]`

## sample file
**https://people.sc.fsu.edu/~jburkardt/data/csv/hw_200.csv**

メジャーリーグの野球選手：名前、チーム、位置、身長（インチ）、体重（ポンド）、年齢（年）。それぞれ6つの値を持つ1034レコード。最初のヘッダー行もあります。

